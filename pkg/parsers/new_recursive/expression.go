package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func parseExpressions(p *pipe) ([]Expression, error) {
	// pattern := [expression {DELIM expression}]

	exp := parseExpression(p)

	if exp == nil {
		return nil, nil
	}

	return parseDelimitedExpressions(p, exp)
}

func parseDelimitedExpressions(p *pipe, first Expression) ([]Expression, error) {

	exps := []Expression{first}

	for p.accept(DELIMITER) {
		exp, e := expectExpression(p)

		if e != nil {
			return nil, e
		}

		exps = append(exps, exp)
	}

	return exps, nil
}

func parseExpression(p *pipe) Expression {

	// pattern := identifier | literal

	switch {
	case p.match(IDENTIFIER):
		return parseIdentifier(p)

	case isLiteral(p):
		return parseLiteral(p)
	}

	return nil
}

func expectExpression(p *pipe) (Expression, error) {
	exp := parseExpression(p)

	if exp == nil {
		return nil, err.New("Expected expression", err.At(p.any()))
	}

	return exp, nil
}

func parseIdentifier(p *pipe) Expression {
	// pattern := IDENTIFIER

	return Identifier{
		TK: p.any(),
	}
}

func isLiteral(p *pipe) bool {
	return p.match(VOID) ||
		p.match(BOOL) ||
		p.match(NUMBER) ||
		p.match(STRING)
}

func parseLiteral(p *pipe) Expression {
	// pattern := IDENTIFIER

	return Literal{
		TK: p.any(),
	}
}
