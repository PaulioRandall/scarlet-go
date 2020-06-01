package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func expressions(p *parser) ([]Expression, error) {
	// pattern := [expression {DELIM expression}]

	exp, e := expression(p)
	if e != nil {
		return nil, e
	}

	if exp != nil {
		return delimExpressions(p, exp)
	}

	return nil, nil
}

func delimExpressions(p *parser, left Expression) ([]Expression, error) {
	// pattern := expression {DELIMITER expression}

	exps := []Expression{left}

	for p.accept(DELIMITER) {

		next, e := expectExpression(p)
		if e != nil {
			return nil, e
		}

		exps = append(exps, next)
	}

	return exps, nil
}

func expression(p *parser) (Expression, error) {
	// pattern := identifier | literal

	switch {
	case p.match(IDENTIFIER), p.match(VOID):
		return p.NewIdentifier(p.any()), nil

	case p.match(BOOL), p.match(NUMBER), p.match(STRING):
		return p.NewLiteral(p.any()), nil

	case p.accept(SUBTRACT):
		return negation(p)

	case p.accept(LIST):
		return list(p)
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

	return p.NewList(open, items, close), nil
}

func listItems(p *parser) ([]Expression, error) {
	// pattern := [expression {DELIMITER [TERMINATOR] expression}]
	return listItemDelimited(p, []Expression{}, true)
}

func listItemDelimited(p *parser, items []Expression, first bool) ([]Expression, error) {

	expr, e := listItem(p, first)
	if e != nil {
		return nil, e
	}

	if expr == nil {
		return items, nil
	}

	items = append(items, expr)

	if p.accept(DELIMITER) {
		if p.accept(TERMINATOR) && p.match(BLOCK_CLOSE) {
			return items, nil
		}

		return listItemDelimited(p, items, false)
	}

	return items, nil
}

func listItem(p *parser, first bool) (Expression, error) {
	if first {
		return expression(p)
	}

	return expectExpression(p)
}
