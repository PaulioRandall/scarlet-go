package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isList(p *pipe) bool {
	return p.match(token.LIST)
}

func parseList(p *pipe) st.Expression {
	// pattern := LIST LIST_OPEN {expression} LIST_CLOSE

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

func parseListAccess(p *pipe) st.ListAccess {
	// pattern := ID GUARD_OPEN expression GUARD_CLOSE

	tk := p.expect(`listAccess`, token.ID)
	id := st.Identifier(tk)

	p.expect(`listAccess`, token.GUARD_OPEN)
	indexExp := parseExpression(p)

	if indexExp == nil {
		panic(err("listAccess", p.past(), `Expected an expression`))
	}

	p.expect(`listAccess`, token.GUARD_CLOSE)
	return st.ListAccess{id, indexExp}
}
