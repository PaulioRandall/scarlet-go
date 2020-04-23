package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

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

	return exps
}

func parseExpression(p *pipe) st.Expression {

	var left st.Expression

	switch {
	case isFuncCall(p):
		p.expect(`rightSide`, token.ID)
		left := st.NewValueExpression(p.prior())
		return funcCall(p, left)

	case isListAccess(p):
		return listAccess(p)

	case literal(p):
		left = st.NewValueExpression(p.proceed())
		return operation(p, left, 0)

	case p.inspect(token.PAREN_OPEN):
		left = group(p)
		return operation(p, left, 0)

	case p.inspect(token.FUNC):
		return funcDef(p)

	case p.inspect(token.LIST):
		return list(p)
	}

	return nil
}
