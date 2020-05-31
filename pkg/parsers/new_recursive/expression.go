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
