package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func parseExpressions(p *pipe) []st.Expression {
	// pattern := [expression {DELIM expression}] [TERMINATOR]

	var exps []st.Expression
	exp := parseExpression(p)

	for exp != nil {
		exps = append(exps, exp)

		if p.accept(token.DELIM) {
			exp = parseExpression(p)

			if exp == nil {
				panic(err("parseExpressions", p.past(), 1, `Expected expression`))
			}

		} else {
			exp = nil
		}
	}

	p.accept(token.TERMINATOR) // In some cases
	return exps
}

func parseExpression(p *pipe) st.Expression {
	// pattern := func_call | list_access | literal | group | list

	var left st.Expression

	switch {
	case isFuncCall(p):
		return parseFuncCall(p)

	case isListAccess(p):
		return parseListAccess(p)

	case isLiteral(p):
		left = parseLiteral(p)
		return parseOperation(p, left, 0)

	case isGroup(p):
		left = parseGroup(p)
		return parseOperation(p, left, 0)

	case isList(p):
		return parseList(p)
	}

	return nil
}
