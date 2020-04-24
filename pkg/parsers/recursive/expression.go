package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// Expects the following token pattern:
// pattern := {expression} [TERMINATOR]
func parseExpressions(p *pipe) []st.Expression {

	var exps []st.Expression
	exp := parseExpression(p)

	for exp != nil {

		exps = append(exps, exp)

		if p.accept(token.DELIM) {
			exp = parseExpression(p)
		} else {
			exp = nil
		}
	}

	if p.inspect(token.TERMINATOR) {
		p.proceed()
	}

	return exps
}

// Expects the following token pattern:
// pattern := func_call | list_access | literal | group | func_def | list
func parseExpression(p *pipe) st.Expression {

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

	case isFuncDef(p):
		return parseFuncDef(p)

	case isList(p):
		return parseList(p)
	}

	return nil
}
