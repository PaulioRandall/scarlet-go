package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func identifier(p *pipeline) (Expression, error) {
	// pattern := IDENTIFIER [list_accessor | function_call]

	tk, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	var id Expression = newIdentifier(tk)

	for more := true; more; {
		id, more, e = maybeMore(p, id)
		if e != nil {
			return nil, e
		}
	}

	return id, nil
}

func maybeMore(p *pipeline, expr Expression) (_ Expression, more bool, e error) {
	// pattern := expression list_accessor
	// pattern := expression function_call

	if p.match(TK_GUARD_OPEN) {

		expr, e = collectionAccessor(p, expr)
		if e != nil {
			return nil, false, e
		}

		return expr, true, nil
	}

	if p.match(TK_PAREN_OPEN) {

		expr, e = functionCall(p, expr)
		if e != nil {
			return nil, false, e
		}

		return expr, true, nil
	}

	return expr, false, nil
}

func literal(p *pipeline) (Expression, error) {
	// pattern := BOOL | NUMBER | STRING

	l, e := p.expectAnyOf(TK_BOOL, TK_NUMBER, TK_STRING)
	if e != nil {
		return nil, e
	}

	return newLiteral(l), nil
}

func negation(p *pipeline) (Expression, error) {
	// pattern := MINUS expression

	_, e := p.expect(TK_MINUS)
	if e != nil {
		return nil, e
	}

	expr, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	return newNegation(expr), nil
}

func group(p *pipeline) (Expression, error) {

	_, e := p.expect(TK_PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	expr, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return expr, e
}

func collectionAccessor(p *pipeline, left Expression) (Expression, error) {
	// pattern := GUARD_OPEN expression GUARD_CLOSE

	p.expect(TK_GUARD_OPEN)

	index, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_GUARD_CLOSE)
	if e != nil {
		return nil, e
	}

	return newCollectionAccessor(left, index), nil
}

func expressions(p *pipeline) ([]Expression, error) {
	// pattern := [operation {DELIM operation}]

	ops := []Expression{}

	for !p.match(TK_PAREN_CLOSE) && !p.match(TK_TERMINATOR) {

		expr, e := expectExpression(p)
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

func expression(p *pipeline) (Expression, error) {

	left, e := operand(p)
	if e != nil {
		return nil, e
	}

	if left == nil {
		return nil, nil
	}

	return operation(p, left, 0)
}

func expectExpression(p *pipeline) (Expression, error) {
	// pattern := operation

	expr, e := expression(p)
	if e != nil {
		return nil, e
	}

	if expr == nil {
		return nil, err.New("Expected expression", err.At(p.any()))
	}

	return expr, nil
}

func operand(p *pipeline) (Expression, error) {

	switch {
	case p.match(TK_VOID):
		return newVoid(p.any()), nil

	case p.match(TK_IDENTIFIER):
		return identifier(p)

	case p.match(TK_BOOL), p.match(TK_NUMBER), p.match(TK_STRING):
		return literal(p)

	case p.match(TK_MINUS):
		return negation(p)

	case p.match(TK_PAREN_OPEN):
		return group(p)

	case p.match(TK_SPELL):
		return spellCall(p)
	}

	return nil, nil
}

func expectOperand(p *pipeline) (Expression, error) {

	o, e := operand(p)
	if e != nil {
		return nil, e
	}

	if o == nil {
		return nil, err.New("Expected expression", err.At(p.any()))
	}

	return o, nil
}

func operation(p *pipeline, left Expression, leftPriority int) (Expression, error) {

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

	left = newOperation(op, left, right)

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

	stats, e := blockStatements(p)

	close, e := p.expect(TK_BLOCK_CLOSE)
	if e != nil {
		return nil, e
	}

	return newBlock(open, close, stats), nil
}

func blockStatements(p *pipeline) ([]Expression, error) {

	var (
		st Expression
		e  error
		r  = []Expression{}
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

func function(p *pipeline) (Expression, error) {
	// pattern := FUNC function_parameters function_body

	key, e := p.expect(TK_FUNCTION)
	if e != nil {
		return nil, e
	}

	params, e := functionParameters(p)
	if e != nil {
		return nil, e
	}

	body, e := block(p)
	if e != nil {
		return nil, e
	}

	return newFunction(key, params, body), nil
}

func functionParameters(p *pipeline) (Parameters, error) {
	// pattern := PAREN_OPEN inputs [THEN outputs] PAREN_CLOSE

	open, e := p.expect(TK_PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	inputs, e := parameterIdentifiers(p)
	if e != nil {
		return nil, e
	}

	outputs := []Token{}

	if p.accept(TK_OUTPUTS) {
		outputs, e = parameterIdentifiers(p)
		if e != nil {
			return nil, e
		}
	}

	close, e := p.expect(TK_PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return newParameters(open, close, inputs, outputs), nil
}

func parameterIdentifiers(p *pipeline) ([]Token, error) {
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

func functionCall(p *pipeline, f Expression) (Expression, error) {
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

	return newFunctionCall(close, f, args), nil
}

func expressionFunction(p *pipeline) (Expression, error) {
	// pattern := EXPR_FUNC exprFuncInputs expression

	key, e := p.expect(TK_EXPR_FUNC)
	if e != nil {
		return nil, e
	}

	inputs, e := expressionFunctionParameters(p)
	if e != nil {
		return nil, e
	}

	expr, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	return newExpressionFunction(key, inputs, expr), nil
}

func expressionFunctionParameters(p *pipeline) ([]Token, error) {
	// pattern := PAREN_OPEN parameters PAREN_CLOSE

	_, e := p.expect(TK_PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	if p.accept(TK_PAREN_CLOSE) {
		return []Token{}, nil
	}

	in, e := expressionFunctionInputs(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return in, nil
}

func expressionFunctionInputs(p *pipeline) ([]Token, error) {
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

func watch(p *pipeline) (Expression, error) {
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

	return newWatch(key, ids, body), nil
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

func guard(p *pipeline) (Expression, error) {
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

	return newGuard(open, condition, body), nil
}

func guardCondition(p *pipeline) (Token, Expression, error) {

	open, e := p.expect(TK_GUARD_OPEN)
	if e != nil {
		return nil, nil, e
	}

	condition, e := expectExpression(p)
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

	stat, e := expectStatement(p)
	if e != nil {
		return nil, e
	}

	stats := []Expression{stat}
	return newUnDelimiteredBlock(stats), nil
}

func when(p *pipeline) (Expression, error) {
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

	return newWhen(key, close, init, cases), nil
}

func whenInitialiser(p *pipeline) (Assignment, error) {

	id, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	target := newIdentifier(id)

	_, e = p.expect(TK_ASSIGNMENT)
	if e != nil {
		return nil, e
	}

	source, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	return newAssignment(target, source), nil
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

	return newGuard(open, condition, body), nil
}

func whenCase(p *pipeline) (WhenCase, error) {
	// pattern := object THEN guardBody

	object, e := expectExpression(p)
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

	return newWhenCase(object, body), nil
}

func loop(p *pipeline) (Expression, error) {
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

	return newLoop(key, init, g), nil
}

func loopInitialiser(p *pipeline) (Assignment, error) {
	// pattern := IDENTIFIER ASSIGN SOURCE

	id, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	target := newIdentifier(id)

	_, e = p.expect(TK_ASSIGNMENT)
	if e != nil {
		return nil, e
	}

	source, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	return newAssignment(target, source), nil
}

func spellCall(p *pipeline) (Expression, error) {
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

	var spell Expression = newSpellCall(id, close, args)

	for more := true; more; {
		spell, more, e = maybeMore(p, spell)
		if e != nil {
			return nil, e
		}
	}

	return spell, nil
}

func spellCallArguments(p *pipeline) ([]Expression, error) {
	// pattern  := [argument {DELIM argument}]
	// arugment := expression | block

	var expr Expression
	var e error

	args := []Expression{}

	for !p.match(TK_PAREN_CLOSE) && !p.match(TK_TERMINATOR) {

		if p.match(TK_BLOCK_OPEN) {
			expr, e = block(p)
		} else {
			expr, e = expectExpression(p)
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
