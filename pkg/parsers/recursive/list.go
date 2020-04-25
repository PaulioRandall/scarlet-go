package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isList(p *pipe) bool {
	return p.match(token.LIST)
}

// Expects the following token pattern:
// pattern := LIST LIST_OPEN {expression} LIST_CLOSE
func parseList(p *pipe) st.Expression {
	return st.List{
		Key:   p.expect(`parseList`, token.LIST),
		Open:  p.expect(`parseList`, token.BLOCK_OPEN),
		Exprs: parseExpressions(p),
		Close: p.expect(`parseList`, token.BLOCK_CLOSE),
	}
}

func isListAccess(p *pipe) bool {
	return p.matchSequence(token.ID, token.GUARD_OPEN)
}

// Expects the following token pattern:
// pattern := ID GUARD_OPEN expression GUARD_CLOSE
func parseListAccess(p *pipe) st.ListAccess {

	id := st.Identifier{
		Source: p.expect(`listAccess`, token.ID),
	}

	p.expect(`listAccess`, token.GUARD_OPEN)
	indexExp := parseExpression(p)

	if indexExp == nil {
		panic(err("listAccess", p.past(), `Expected an expression`))
	}

	p.expect(`listAccess`, token.GUARD_CLOSE)
	return st.ListAccess{id, indexExp}
}
