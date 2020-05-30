package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func parseExpressions(p *parser) ([]Expression, error) {
	// pattern := [expression {DELIM expression}]

	exp, e := parseExpression(p)
	if e != nil {
		return nil, e
	}

	if exp != nil {
		return parseDelimExpressions(p, exp)
	}

	return nil, nil
}

func parseDelimExpressions(p *parser, first Expression) ([]Expression, error) {

	exps := []Expression{first}

	for p.accept(DELIMITER) {

		next, e := expectExpression(p)
		if e != nil {
			return nil, e
		}

		exps = append(exps, next)
	}

	return exps, nil
}

func parseExpression(p *parser) (expr Expression, e error) {
	// pattern := identifier | literal

	switch {
	case p.match(IDENTIFIER):
		expr = p.NewIdentifier(p.any())

	case isLiteral(p):
		expr = p.NewLiteral(p.any())
	}

	return
}

func expectExpression(p *parser) (Expression, error) {

	exp, e := parseExpression(p)
	if e != nil {
		return nil, e
	}

	if exp == nil {
		return nil, err.New("Expected expression", err.At(p.any()))
	}

	return exp, nil
}

func isLiteral(p *parser) bool {
	return p.match(VOID) ||
		p.match(BOOL) ||
		p.match(NUMBER) ||
		p.match(STRING)
}
