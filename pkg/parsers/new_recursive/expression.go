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

	case p.accept(LIST):
		return list(p)

		//case p.accept(FUNC)

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

func list(p *parser) (Expression, error) {
	// pattern := BLOCK_OPEN [TERMINATOR] listItems [TERMINATOR] BLOCK_CLOSE

	open, e := p.expect(BLOCK_OPEN)
	if e != nil {
		return nil, e
	}

	p.accept(TERMINATOR)
	items, e := listItems(p)
	if e != nil {
		return nil, e
	}

	close, e := p.expect(BLOCK_CLOSE)
	if e != nil {
		return nil, e
	}

	list := p.NewList(open, items, close)
	return maybeListAccessor(p, list)
}

func listItems(p *parser) ([]Expression, error) {
	// pattern := [expression {DELIMITER [TERMINATOR] expression}]

	items := []Expression{}

	for loop := true; loop; loop = acceptDelimiter(p, BLOCK_CLOSE) {

		expr, e := listItem(p)
		if e != nil {
			return nil, e
		}

		if expr == nil {
			return items, nil
		}

		items = append(items, expr)
	}

	return items, nil
}

func listItem(p *parser) (Expression, error) {
	// pattern := [expression]
	return expression(p)
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

	if p.accept(GUARD_OPEN) {

		index, e := expectExpression(p)
		if e != nil {
			return nil, e
		}

		_, e = p.expect(GUARD_CLOSE)
		if e != nil {
			return nil, e
		}

		return p.NewListAccessor(maybeList, index), nil
	}

	return maybeList, nil
}

func function(p *parser) (Expression, error) {
	// pattern := FUNC function_parameters function_body

	_, e := p.expect(FUNC)
	if e != nil {
		return nil, e
	}

	// parse parameters

	// parse body
	return nil, nil
}

func functionParameters(p *parser) (Expression, error) {
	// pattern := PAREN_OPEN [expression {DELIMITER expression}] PAREN_CLOSE

	open, e := p.expect(PAREN_OPEN)
	if e != nil {
		return nil, e
	}

	inputs := []Token{}
	outputs := []Token{}

	for loop := true; loop; loop = acceptDelimiter(p, PAREN_CLOSE) {

		id, isOutput, e := functionParam(p)
		if e != nil {
			return nil, e
		}

		if isOutput {
			outputs = append(outputs, id)
		} else {
			inputs = append(inputs, id)
		}
	}

	close, e := p.expect(PAREN_CLOSE)
	if e != nil {
		return nil, e
	}

	return p.NewParameters(open, close, inputs, outputs), nil
}

func functionParam(p *parser) (Token, bool, error) {
	output := p.accept(OUTPUT)
	id, e := p.expect(IDENTIFIER)
	return id, output, e
}
