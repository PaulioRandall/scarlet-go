package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func expressions(p *parser) ([]Expression, error) {
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

func expression(p *parser) (Expression, error) {
	// pattern := identifier | literal

	switch {
	case p.match(IDENTIFIER), p.match(VOID):
		return identifier(p)

	case p.match(BOOL), p.match(NUMBER), p.match(STRING):
		return p.NewLiteral(p.any()), nil

	case p.accept(SUBTRACT):
		return negation(p)

	case p.match(FUNC):
		return function(p)
	}

	return nil, nil
}

func expectExpression(p *parser) (Expression, error) {
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

func identifier(p *parser) (Expression, error) {
	// pattern := IDENTIFIER [list_accessor]
	id := p.NewIdentifier(p.any())
	return maybeListAccessor(p, id)
}

func negation(p *parser) (Expression, error) {
	// pattern := MINUS expression

	expr, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	return p.NewNegation(expr), nil
}

func acceptDelimiter(p *parser, closingSignal Morpheme) bool {

	if p.accept(DELIMITER) {
		if p.accept(TERMINATOR) {
			return !p.match(closingSignal)
		}

		return true
	}

	return false
}

func maybeListAccessor(p *parser, maybeList Expression) (Expression, error) {
	// pattern := expression [GUARD_OPEN expression GUARD_CLOSE]

	if p.match(GUARD_OPEN) {
		return listAccessor(p, maybeList)
	}

	return maybeList, nil
}

func listAccessor(p *parser, left Expression) (Expression, error) {
	// pattern := GUARD_OPEN expression GUARD_CLOSE

	p.expect(GUARD_OPEN)

	index, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(GUARD_CLOSE)
	if e != nil {
		return nil, e
	}

	return p.NewListAccessor(left, index), nil
}

func function(p *parser) (Expression, error) {
	// pattern := FUNC function_parameters function_body

	key, e := p.expect(FUNC)
	if e != nil {
		return nil, e
	}

	params, e := functionParameters(p)
	if e != nil {
		return nil, e
	}

	body, e := functionBody(p)
	if e != nil {
		return nil, e
	}

	return p.NewFunction(key, params, body), nil
}

func functionParameters(p *parser) (Parameters, error) {
	// pattern := PAREN_OPEN [expression {DELIMITER expression}] PAREN_CLOSE

	NIL := Parameters{}

	open, e := p.expect(PAREN_OPEN)
	if e != nil {
		return NIL, e
	}

	inputs, outputs, e := parameterIdentifiers(p)
	if e != nil {
		return NIL, e
	}

	close, e := p.expect(PAREN_CLOSE)
	if e != nil {
		return NIL, e
	}

	return p.NewParameters(open, close, inputs, outputs), nil
}

func parameterIdentifiers(p *parser) (in []Token, out []Token, _ error) {

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

func functionParam(p *parser) (Token, bool, error) {
	output := p.accept(OUTPUT)
	id, e := p.expect(IDENTIFIER)
	return id, output, e
}

func functionBody(p *parser) (Block, error) {

	NIL := Block{}

	open, e := p.expect(BLOCK_OPEN)
	if e != nil {
		return NIL, e
	}

	p.accept(TERMINATOR)
	stats, e := functionStatements(p)

	p.accept(TERMINATOR)
	close, e := p.expect(BLOCK_CLOSE)
	if e != nil {
		return NIL, e
	}

	return p.NewBlock(open, close, stats), nil
}

func functionStatements(p *parser) ([]Statement, error) {

	var (
		st Statement
		e  error
		r  = []Statement{}
	)

	for loop := true; loop; {

		st, loop, e = functionStatement(p)
		if e != nil {
			return nil, e
		}

		if st != nil {
			r = append(r, st)
		}
	}

	return r, nil
}

func functionStatement(p *parser) (st Statement, more bool, e error) {

	st, e = statement(p)
	if e != nil {
		return nil, false, e
	}

	if st == nil {
		return nil, false, nil
	}

	return st, p.accept(TERMINATOR), nil
}

func operation(p *parser, left Expression) (Expression, error) {

	if !p.hasMore() {
		return left, nil
	}

	if Precedence(left) >= p.peek().Morpheme().Precedence() {
		return left, nil
	}

	op, e := p.expectAnyOf(OperatorTypes()...)
	if e != nil {
		return nil, e
	}

	right, e := operationRight(p)
	if e != nil {
		return nil, e
	}

	right, e = operation(p, right)
	if e != nil {
		return nil, e
	}

	left = p.NewOperation(op, left, right)
	return operation(p, left)
}

func operationRight(p *parser) (Expression, error) {
	switch {
	case p.match(IDENTIFIER):
		return identifier(p)

	case p.match(BOOL), p.match(NUMBER), p.match(STRING):
		return p.NewLiteral(p.any()), nil

	case p.accept(SUBTRACT):
		return negation(p)
	}

	return nil, err.New("Expected expression", err.At(p.any()))
}
