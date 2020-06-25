package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

func identifier(p *pipeline) (Expr, error) {
	// pattern := IDENTIFIER [list_accessor | function_call]

	tk, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	var id Expr = NewIdentifier(tk)

	for more := true; more; {

		id, more, e = maybeMore(p, id)
		if e != nil {
			return nil, e
		}
	}

	return id, nil
}

func maybeMore(p *pipeline, expr Expr) (_ Expr, more bool, e error) {
	// pattern := expression list_accessor
	// pattern := expression function_call

	if p.match(TK_GUARD_OPEN) {

		expr, e = containerItem(p, expr)
		if e != nil {
			return nil, false, e
		}

		return expr, true, nil
	}

	if p.match(TK_PAREN_OPEN) {

		expr, e = funcCall(p, expr)
		if e != nil {
			return nil, false, e
		}

		return expr, true, nil
	}

	return expr, false, nil
}

func literal(p *pipeline) (Expr, error) {
	// pattern := BOOL | NUMBER | STRING

	l, e := p.expectAnyOf(TK_BOOL, TK_NUMBER, TK_STRING)
	if e != nil {
		return nil, e
	}

	return NewLiteral(l), nil
}

func negation(p *pipeline) (Expr, error) {
	// pattern := MINUS expression

	_, e := p.expect(TK_MINUS)
	if e != nil {
		return nil, e
	}

	expr, e := expectExpr(p)
	if e != nil {
		return nil, e
	}

	return NewNegation(expr), nil
}

func exists(p *pipeline, subject Expr) (Expr, error) {
	// pattern := subject EXISTS

	close, e := p.expect(TK_EXISTS)
	if e != nil {
		return nil, e
	}

	return NewExists(close, subject), nil
}

func maybeExists(p *pipeline, subject Expr) (Expr, error) {
	if p.match(TK_EXISTS) {
		return exists(p, subject)
	}

	return subject, nil
}

func group(p *pipeline) (Expr, error) {

	_, e := p.expect(TK_PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	expr, e := expectExpr(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return expr, nil
}

func containerItem(p *pipeline, left Expr) (Expr, error) {
	// pattern := GUARD_OPEN expression GUARD_CLOSE

	p.expect(TK_GUARD_OPEN)

	index, e := expectExpr(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_GUARD_CLOSE)
	if e != nil {
		return nil, e
	}

	return NewContainerItem(left, index), nil
}

func expressions(p *pipeline) ([]Expr, error) {
	// pattern := [operation {DELIM operation}]

	ops := []Expr{}

	for !p.match(TK_PAREN_CLOSE) && !p.match(TK_TERMINATOR) {

		expr, e := expectExpr(p)
		if e != nil {
			return nil, e
		}

		ops = append(ops, expr)

		if !p.accept(TK_DELIMITER) {
			break
		}
	}

	return ops, nil
}

func expression(p *pipeline) (Expr, error) {

	left, e := operand(p)
	if e != nil {
		return nil, e
	}

	if left == nil {
		return nil, nil
	}

	return operation(p, left, 0)
}

func expectExpr(p *pipeline) (Expr, error) {
	// pattern := operation

	expr, e := expression(p)
	if e != nil {
		return nil, e
	}

	if expr == nil {
		return nil, err.NewBySnippet("Expected expression", p.any())
	}

	return expr, nil
}

func operand(p *pipeline) (expr Expr, e error) {

	switch {
	case p.match(TK_IDENTIFIER):
		expr, e = identifier(p)

	case p.match(TK_BOOL), p.match(TK_NUMBER), p.match(TK_STRING):
		expr, e = literal(p)

	case p.match(TK_MINUS):
		expr, e = negation(p)

	case p.match(TK_PAREN_OPEN):
		expr, e = group(p)

	case p.match(TK_SPELL):
		expr, e = spellCall(p)

	default:
		return nil, nil
	}

	if e != nil {
		return nil, e
	}

	return maybeExists(p, expr)
}

func expectOperand(p *pipeline) (Expr, error) {

	o, e := operand(p)
	if e != nil {
		return nil, e
	}

	if o == nil {
		return nil, err.NewBySnippet("Expected expression", p.any())
	}

	return o, nil
}

func operation(p *pipeline, left Expr, leftPriority int) (Expr, error) {

	if !p.hasMore() {
		return left, nil
	}

	rightPriority := p.peek().Type().Precedence()
	if leftPriority >= rightPriority {
		return left, nil
	}

	op, e := p.expectAnyOf(OperatorTypes()...)
	if e != nil {
		return nil, e
	}

	right, e := expectOperand(p)
	if e != nil {
		return nil, e
	}

	right, e = operation(p, right, rightPriority)
	if e != nil {
		return nil, e
	}

	left = NewOperation(op, left, right)

	left, e = operation(p, left, leftPriority)
	if e != nil {
		return nil, e
	}

	return left, nil
}

func block(p *pipeline) (Block, error) {

	open, e := p.expect(TK_BLOCK_OPEN)
	if e != nil {
		return nil, e
	}

	sts, e := blockStatements(p)

	close, e := p.expect(TK_BLOCK_CLOSE)
	if e != nil {
		return nil, e
	}

	return NewBlock(open, close, sts), nil
}

func blockStatements(p *pipeline) ([]Expr, error) {

	var (
		st Expr
		e  error
		r  = []Expr{}
	)

	for !p.match(TK_BLOCK_CLOSE) {

		st, e = expectStatement(p)
		if e != nil {
			return nil, e
		}

		r = append(r, st)

		if !p.accept(TK_TERMINATOR) {
			break
		}
	}

	return r, nil
}

func funcDef(p *pipeline) (Expr, error) {
	// pattern := FUNC function_parameters function_body

	key, e := p.expect(TK_FUNCTION)
	if e != nil {
		return nil, e
	}

	inputs, outputs, e := funcDefParams(p)
	if e != nil {
		return nil, e
	}

	body, e := funcBody(p)
	if e != nil {
		return nil, e
	}

	return NewFuncDef(key, inputs, outputs, body), nil
}

func funcDefParams(p *pipeline) ([]Token, []Token, error) {
	// pattern := PAREN_OPEN inputs [THEN outputs] PAREN_CLOSE

	_, e := p.expect(TK_PAREN_OPEN)
	if e != nil {
		return nil, nil, e
	}

	inputs, e := paramIdentifiers(p)
	if e != nil {
		return nil, nil, e
	}

	outputs := []Token{}

	if p.accept(TK_OUTPUTS) {
		outputs, e = paramIdentifiers(p)
		if e != nil {
			return nil, nil, e
		}
	}

	_, e = p.expect(TK_PAREN_CLOSE)
	if e != nil {
		return nil, nil, e
	}

	return inputs, outputs, nil
}

func paramIdentifiers(p *pipeline) ([]Token, error) {
	// pattern := [IDENTIFIER {DELIMITER IDENTIFIER}]

	ids := []Token{}

	for !p.match(TK_OUTPUTS) && !p.match(TK_PAREN_CLOSE) {

		id, e := p.expect(TK_IDENTIFIER)
		if e != nil {
			return nil, e
		}

		ids = append(ids, id)

		if !p.accept(TK_DELIMITER) {
			break
		}
	}

	return ids, nil
}

func funcBody(p *pipeline) (Expr, error) {

	switch {
	case p.match(TK_BLOCK_OPEN):
		return block(p)

	case p.match(TK_WATCH):
		return watch(p)

	case p.match(TK_WHEN):
		return when(p)

	case p.match(TK_GUARD_OPEN):

		open, condition, e := guardCondition(p)
		if e != nil {
			return nil, e
		}

		body, e := block(p)
		if e != nil {
			return nil, e
		}

		return NewGuard(open, condition, body), nil
	}

	return nil, err.NewBySnippet("Expected function body", p.any())
}

func funcCall(p *pipeline, f Expr) (Expr, error) {
	// pattern := PAREN_OPEN arguments PAREN_CLOSE

	_, e := p.expect(TK_PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	args, e := expressions(p)
	if e != nil {
		return nil, e
	}

	close, e := p.expect(TK_PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return NewFuncCall(close, f, args), nil
}

func exprFunc(p *pipeline) (Expr, error) {
	// pattern := EXPR_FUNC exprFuncInputs expression

	key, e := p.expect(TK_EXPR_FUNC)
	if e != nil {
		return nil, e
	}

	inputs, e := exprFuncParams(p)
	if e != nil {
		return nil, e
	}

	expr, e := expectExpr(p)
	if e != nil {
		return nil, e
	}

	return NewExprFunc(key, inputs, expr), nil
}

func exprFuncParams(p *pipeline) ([]Token, error) {
	// pattern := PAREN_OPEN parameters PAREN_CLOSE

	_, e := p.expect(TK_PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	if p.accept(TK_PAREN_CLOSE) {
		return []Token{}, nil
	}

	in, e := exprFuncInputs(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return in, nil
}

func exprFuncInputs(p *pipeline) ([]Token, error) {
	// pattern := [identifier {DELIMITER identifier} [DELIMITER]]

	in := []Token{}

	for !p.match(TK_PAREN_CLOSE) {

		id, e := p.expect(TK_IDENTIFIER)
		if e != nil {
			return nil, e
		}

		in = append(in, id)

		if !p.accept(TK_DELIMITER) {
			break
		}
	}

	return in, nil
}

func watch(p *pipeline) (Expr, error) {
	// pattern := WATCH BLOCK_OPEN {statements} BLOCK_CLOSE

	key, e := p.expect(TK_WATCH)
	if e != nil {
		return nil, e
	}

	ids, e := watchIdentifiers(p)
	if e != nil {
		return nil, e
	}

	body, e := block(p)
	if e != nil {
		return nil, e
	}

	return NewWatch(key, ids, body), nil
}

func watchIdentifiers(p *pipeline) ([]Token, error) {
	// pattern := IDENTIFIER {DELIM IDENTIFIER}

	ids := []Token{}

	for !p.match(TK_BLOCK_OPEN) {

		id, e := p.expect(TK_IDENTIFIER)
		if e != nil {
			return nil, e
		}

		ids = append(ids, id)

		if !p.accept(TK_DELIMITER) {
			break
		}
	}

	return ids, nil
}

func guard(p *pipeline) (Expr, error) {
	return guardExplicit(p)
}

func guardExplicit(p *pipeline) (Guard, error) {

	open, condition, e := guardCondition(p)
	if e != nil {
		return nil, e
	}

	body, e := guardBody(p)
	if e != nil {
		return nil, e
	}

	return NewGuard(open, condition, body), nil
}

func guardCondition(p *pipeline) (Token, Expr, error) {

	open, e := p.expect(TK_GUARD_OPEN)
	if e != nil {
		return nil, nil, e
	}

	condition, e := expectExpr(p)
	if e != nil {
		return nil, nil, e
	}

	_, e = p.expect(TK_GUARD_CLOSE)
	if e != nil {
		return nil, nil, e
	}

	return open, condition, nil
}

func guardBody(p *pipeline) (Block, error) {

	if p.match(TK_BLOCK_OPEN) {
		return block(p)
	}

	st, e := expectStatement(p)
	if e != nil {
		return nil, e
	}

	sts := []Expr{st}
	return NewUnDelimiteredBlock(sts), nil
}

func when(p *pipeline) (Expr, error) {
	// pattern := WHEN whenInitialiser BLOCK_OPEN whenCases BLOCK_CLOSE

	key, e := p.expect(TK_WHEN)
	if e != nil {
		return nil, e
	}

	init, e := whenInitialiser(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_BLOCK_OPEN)
	if e != nil {
		return nil, e
	}

	cases, e := whenCases(p)
	if e != nil {
		return nil, e
	}

	close, e := p.expect(TK_BLOCK_CLOSE)
	if e != nil {
		return nil, e
	}

	return NewWhen(key, close, init, cases), nil
}

func whenInitialiser(p *pipeline) (Assign, error) {

	id, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	target := NewIdentifier(id)

	_, e = p.expect(TK_ASSIGNMENT)
	if e != nil {
		return nil, e
	}

	source, e := expectExpr(p)
	if e != nil {
		return nil, e
	}

	return NewAssign(target, source), nil
}

func whenCases(p *pipeline) ([]WhenCase, error) {
	// pattern := {whenGuardCase | whenCase}

	var mc WhenCase
	var e error

	cases := []WhenCase{}

	for !p.match(TK_BLOCK_CLOSE) {

		if p.match(TK_GUARD_OPEN) {
			mc, e = whenGuardCase(p)
		} else {
			mc, e = whenCase(p)
		}

		if e != nil {
			return nil, e
		}

		cases = append(cases, mc)

		if !p.accept(TK_TERMINATOR) {
			break
		}
	}

	return cases, nil
}

func whenGuardCase(p *pipeline) (WhenCase, error) {
	// pattern := guardCondition THEN guardBody

	open, condition, e := guardCondition(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_THEN)
	if e != nil {
		return nil, e
	}

	body, e := guardBody(p)
	if e != nil {
		return nil, e
	}

	return NewGuard(open, condition, body), nil
}

func whenCase(p *pipeline) (WhenCase, error) {
	// pattern := object THEN guardBody

	object, e := expectExpr(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_THEN)
	if e != nil {
		return nil, e
	}

	body, e := guardBody(p)
	if e != nil {
		return nil, e
	}

	return NewWhenCase(object, body), nil
}

func loop(p *pipeline) (Expr, error) {
	// pattern := LOOP loopInitialiser guard

	key, e := p.expect(TK_LOOP)
	if e != nil {
		return nil, e
	}

	init, e := loopInitialiser(p)
	if e != nil {
		return nil, e
	}

	g, e := guardExplicit(p)
	if e != nil {
		return nil, e
	}

	return NewLoop(key, init, g), nil
}

func loopInitialiser(p *pipeline) (Assign, error) {
	// pattern := IDENTIFIER ASSIGN SOURCE

	id, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	target := NewIdentifier(id)

	_, e = p.expect(TK_ASSIGNMENT)
	if e != nil {
		return nil, e
	}

	source, e := expectExpr(p)
	if e != nil {
		return nil, e
	}

	return NewAssign(target, source), nil
}

func spellCall(p *pipeline) (Expr, error) {
	// pattern := SPELL PAREN_OPEN arguments PAREN_CLOSE

	id, e := p.expect(TK_SPELL)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	args, e := spellCallArguments(p)
	if e != nil {
		return nil, e
	}

	close, e := p.expect(TK_PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	var spell Expr = NewSpellCall(id, close, args)

	for more := true; more; {
		spell, more, e = maybeMore(p, spell)
		if e != nil {
			return nil, e
		}
	}

	return spell, nil
}

func spellCallArguments(p *pipeline) ([]Expr, error) {
	// pattern  := [argument {DELIM argument}]
	// arugment := expression | block

	var expr Expr
	var e error

	args := []Expr{}

	for !p.match(TK_PAREN_CLOSE) && !p.match(TK_TERMINATOR) {

		if p.match(TK_BLOCK_OPEN) {
			expr, e = block(p)
		} else {
			expr, e = expectExpr(p)
		}

		if e != nil {
			return nil, e
		}

		args = append(args, expr)

		if !p.accept(TK_DELIMITER) {
			break
		}
	}

	return args, nil
}
