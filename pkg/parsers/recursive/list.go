package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isList(p *pipe) bool {
	return p.inspect(token.LIST)
}

// Expects the following token pattern:
// pattern := LIST LIST_OPEN {expression} LIST_CLOSE
func parseList(p *pipe) st.Expression {
	return st.List{
		Key:   p.expect(`parseList`, token.LIST),
		Open:  p.expect(`parseList`, token.LIST_OPEN),
		Exprs: parseExpressions(p),
		Close: p.expect(`parseList`, token.LIST_CLOSE),
	}
}

func isListAccess(p *pipe) bool {
	return p.isSequence(token.ID, token.GUARD_OPEN)
}

// Expects the following token pattern:
// pattern := ID GUARD_OPEN expression GUARD_CLOSE
func parseListAccess(p *pipe) st.ListAccess {

	id := st.Identifier{
		Fixed:  false,
		Source: p.expect(`listAccess`, token.ID),
	}

	p.expect(`listAccess`, token.GUARD_OPEN)
	indexExp := parseExpression(p)

	if indexExp == nil {
		panic(err("listAccess", p.prior(), `Expected an expression`))
	}

	p.expect(`listAccess`, token.GUARD_CLOSE)
	return st.ListAccess{id, indexExp}
}
