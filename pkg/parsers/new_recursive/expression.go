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

func expression(p *parser) (expr Expression, e error) {
	// pattern := identifier | literal

	switch {
	case p.match(IDENTIFIER), p.match(VOID):
		expr = p.NewIdentifier(p.any())

	case p.match(BOOL), p.match(NUMBER), p.match(STRING):
		expr = p.NewLiteral(p.any())
	}

	return
}

func expectExpression(p *parser) (Expression, error) {
	// pattern := identifier | literal

	exp, e := expression(p)
	if e != nil {
		return nil, e
	}

	if exp == nil {
		return nil, err.New("Expected expression", err.At(p.any()))
	}

	return exp, nil
}
