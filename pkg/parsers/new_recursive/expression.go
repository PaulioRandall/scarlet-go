package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func expressions(p *pipeline) ([]Expression, error) {
	// pattern := [expression {DELIM expression}]

	r := []Expression{}

	for first := true; p.hasMore(); first = false {

		var expr Expression
		var e error

		if first {
			expr, e = expression(p)
		} else {
			expr, e = expectExpression(p)
		}

		if e != nil {
			return nil, e
		}

		if expr == nil { // Only needed for the first expression
			return r, nil
		}

		r = append(r, expr)

		if !p.accept(DELIMITER) {
			return r, nil
		}
	}

	return nil, err.New("Expected expression", err.At(p.any()))
}

func expression(p *pipeline) (Expression, error) {
	// pattern := identifier | literal

	switch {
	case p.match(IDENTIFIER), p.match(VOID):
		return identifier(p)

	case p.match(BOOL), p.match(NUMBER), p.match(STRING):
		return literal(p)

	case p.match(SUBTRACT):
		return negation(p)

	case p.match(PAREN_OPEN):
		return group(p)
	}

	return nil, nil
}

func expectExpression(p *pipeline) (Expression, error) {
	// pattern := identifier | literal

	expr, e := expression(p)
	if e != nil {
		return nil, e
	}

	if expr == nil {
		return nil, err.New("Expected expression", err.At(p.any()))
	}

	return expr, nil
}

func identifier(p *pipeline) (Expression, error) {
	// pattern := IDENTIFIER [list_accessor]
	id := newIdentifier(p.any())
	return maybeListAccessor(p, id)
}

func literal(p *pipeline) (Expression, error) {
	// pattern := BOOL | NUMBER | STRING
	return newLiteral(p.any()), nil
}

func negation(p *pipeline) (Expression, error) {
	// pattern := MINUS expression

	_, e := p.expect(SUBTRACT)
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

	_, e := p.expect(PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	expr, e := expectOperation(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return expr, e
}

func acceptDelimiter(p *pipeline, closingSignal Morpheme) bool {

	if p.accept(DELIMITER) {
		if p.accept(TERMINATOR) {
			return !p.match(closingSignal)
		}

		return true
	}

	return false
}

func maybeListAccessor(p *pipeline, maybeList Expression) (Expression, error) {
	// pattern := expression [GUARD_OPEN expression GUARD_CLOSE]

	if p.match(GUARD_OPEN) {
		return listAccessor(p, maybeList)
	}

	return maybeList, nil
}

func listAccessor(p *pipeline, left Expression) (Expression, error) {
	// pattern := GUARD_OPEN expression GUARD_CLOSE

	p.expect(GUARD_OPEN)

	index, e := expectOperation(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(GUARD_CLOSE)
	if e != nil {
		return nil, e
	}

	return newListAccessor(left, index), nil
}

func operations(p *pipeline) ([]Expression, error) {
	// pattern := [operation {DELIM operation}]

	r := []Expression{}

	for first := true; p.hasMore(); first = false {

		var (
			expr Expression
			e    error
		)

		if first {
			expr, e = operation(p)
		} else {
			expr, e = expectOperation(p)
		}

		if e != nil {
			return nil, e
		}

		if expr == nil { // Only needed for the first expression
			return r, nil
		}

		r = append(r, expr)

		if !p.accept(DELIMITER) {
			return r, nil
		}
	}

	return nil, err.New("Expected expression", err.At(p.any()))
}

func operation(p *pipeline) (Expression, error) {

	left, e := operand(p)
	if e != nil {
		return nil, e
	}

	if left == nil {
		return nil, nil
	}

	return operationExpression(p, left, 0)
}

func operand(p *pipeline) (Expression, error) {

	switch {
	case p.match(IDENTIFIER), p.match(VOID):
		return identifier(p)

	case p.match(BOOL), p.match(NUMBER), p.match(STRING):
		return literal(p)

	case p.match(SUBTRACT):
		return negation(p)

	case p.match(PAREN_OPEN):
		return group(p)
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

func operationExpression(p *pipeline, left Expression, leftPriority int) (Expression, error) {

	if !p.hasMore() {
		return left, nil
	}

	rightPriority := p.peek().Morpheme().Precedence()
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

	right, e = operationExpression(p, right, rightPriority)
	if e != nil {
		return nil, e
	}

	left = newOperation(op, left, right)

	left, e = operationExpression(p, left, leftPriority)
	if e != nil {
		return nil, e
	}

	return left, nil
}

func expectOperation(p *pipeline) (Expression, error) {
	// pattern := operation

	expr, e := operation(p)
	if e != nil {
		return nil, e
	}

	if expr == nil {
		return nil, err.New("Expected expression", err.At(p.any()))
	}

	return expr, nil
}

func block(p *pipeline) (Block, error) {

	open, e := p.expect(BLOCK_OPEN)
	if e != nil {
		return nil, e
	}

	p.accept(TERMINATOR)
	stats, e := blockStatements(p)

	p.accept(TERMINATOR)
	close, e := p.expect(BLOCK_CLOSE)
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

	for loop := true; loop && p.hasMore(); {

		st, loop, e = blockStatement(p)
		if e != nil {
			return nil, e
		}

		if st != nil {
			r = append(r, st)
		}
	}

	return r, nil
}

func blockStatement(p *pipeline) (st Expression, more bool, e error) {

	st, e = statement(p)
	if e != nil {
		return nil, false, e
	}

	if st == nil {
		return nil, false, nil
	}

	return st, p.accept(TERMINATOR), nil
}

func function(p *pipeline) (Expression, error) {
	// pattern := FUNC function_parameters function_body

	key, e := p.expect(FUNC)
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
	// pattern := PAREN_OPEN [expression {DELIMITER expression}] PAREN_CLOSE

	open, e := p.expect(PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	inputs, outputs, e := parameterIdentifiers(p)
	if e != nil {
		return nil, e
	}

	close, e := p.expect(PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return newParameters(open, close, inputs, outputs), nil
}

func parameterIdentifiers(p *pipeline) (in []Token, out []Token, _ error) {

	in = []Token{}
	out = []Token{}

	if p.match(PAREN_CLOSE) {
		return in, out, nil
	}

	p.accept(TERMINATOR)
	for loop := true; loop; loop = acceptDelimiter(p, PAREN_CLOSE) {

		id, isOutput, e := functionParam(p)
		if e != nil {
			return nil, nil, e
		}

		if isOutput {
			out = append(out, id)
		} else {
			in = append(in, id)
		}
	}

	return in, out, nil
}

func functionParam(p *pipeline) (Token, bool, error) {
	output := p.accept(OUTPUT)
	id, e := p.expect(IDENTIFIER)
	return id, output, e
}

func expressionFunction(p *pipeline) (Expression, error) {
	// pattern := EXPR_FUNC exprFuncInputs expression

	key, e := p.expect(EXPR_FUNC)
	if e != nil {
		return nil, e
	}

	inputs, e := expressionFunctionParameters(p)
	if e != nil {
		return nil, e
	}

	expr, e := expectOperation(p)
	if e != nil {
		return nil, e
	}

	return newExpressionFunction(key, inputs, expr), nil
}

func expressionFunctionParameters(p *pipeline) ([]Token, error) {
	// pattern := PAREN_OPEN parameters PAREN_CLOSE

	_, e := p.expect(PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	if p.accept(PAREN_CLOSE) {
		return []Token{}, nil
	}

	in, e := expressionFunctionInputs(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return in, nil
}

func expressionFunctionInputs(p *pipeline) ([]Token, error) {
	// pattern := [identifier {DELIMITER identifier} [DELIMITER]]

	in := []Token{}

	for loop := true; loop; loop = acceptDelimiter(p, PAREN_CLOSE) {

		id, e := p.expect(IDENTIFIER)
		if e != nil {
			return nil, e
		}

		in = append(in, id)
	}

	return in, nil
}

func watch(p *pipeline) (Expression, error) {
	// pattern := WATCH BLOCK_OPEN {statements} BLOCK_CLOSE

	key, e := p.expect(WATCH)
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

	for first := true; first || p.accept(DELIMITER); first = false {

		id, e := p.expect(IDENTIFIER)
		if e != nil {
			return nil, e
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func guard(p *pipeline) (Expression, error) {

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

	open, e := p.expect(GUARD_OPEN)
	if e != nil {
		return nil, nil, e
	}

	condition, e := expectOperation(p)
	if e != nil {
		return nil, nil, e
	}

	_, e = p.expect(GUARD_CLOSE)
	if e != nil {
		return nil, nil, e
	}

	return open, condition, nil
}

func guardBody(p *pipeline) (Block, error) {

	if p.match(BLOCK_OPEN) {
		return block(p)
	}

	stat, e := expectStatement(p)
	if e != nil {
		return nil, e
	}

	stats := []Expression{stat}
	return newUnDelimiteredBlock(stats), nil
}

func match(p *pipeline) (Expression, error) {

	key, e := p.expect(MATCH)
	if e != nil {
		return nil, e
	}

	subject, e := expectOperation(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(BLOCK_OPEN)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TERMINATOR)
	if e != nil {
		return nil, e
	}

	cases, e := matchCases(p)
	if e != nil {
		return nil, e
	}

	close, e := p.expect(BLOCK_CLOSE)
	if e != nil {
		return nil, e
	}

	return newMatch(key, close, subject, cases), nil
}

func matchCases(p *pipeline) ([]MatchCase, error) {

	var (
		mc MatchCase
		e  error
	)

	cases := []MatchCase{}

	for !p.match(BLOCK_CLOSE) && p.hasMore() {

		if p.match(GUARD_OPEN) {
			mc, e = matchGuardCase(p)
		} else {
			mc, e = matchCase(p)
		}

		if e != nil {
			return nil, e
		}

		cases = append(cases, mc)

		_, e = p.expect(TERMINATOR)
		if e != nil {
			return nil, e
		}
	}

	return cases, nil
}

func matchGuardCase(p *pipeline) (MatchCase, error) {

	open, condition, e := guardCondition(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(DO)
	if e != nil {
		return nil, e
	}

	body, e := guardBody(p)
	if e != nil {
		return nil, e
	}

	return newGuard(open, condition, body), nil
}

func matchCase(p *pipeline) (MatchCase, error) {

	condition, e := expectOperation(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(DO)
	if e != nil {
		return nil, e
	}

	body, e := guardBody(p)
	if e != nil {
		return nil, e
	}

	return newMatchCase(condition, body), nil
}
