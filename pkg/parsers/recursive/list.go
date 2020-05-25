package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isList(p *pipe) bool {
	return p.match(LIST)
}

func parseList(p *pipe) Expression {
	// pattern := LIST LIST_OPEN {expression} LIST_CLOSE

	return List{
		Key:   p.expect(`parseList`, LIST),
		Open:  p.expect(`parseList`, BLOCK_OPEN),
		Exprs: parseExpressions(p),
		Close: p.expect(`parseList`, BLOCK_CLOSE),
	}
}

func isListAccess(p *pipe) bool {
	return p.matchSequence(IDENTIFIER, GUARD_OPEN)
}

func parseListAccess(p *pipe) ListAccess {
	// pattern := ID GUARD_OPEN expression GUARD_CLOSE

	tk := p.expect(`parseListAccess`, IDENTIFIER)
	id := Identifier{tk}

	p.expect(`parseListAccess`, GUARD_OPEN)
	indexExp := parseListItemExpr(p)
	p.expect(`parseListAccess`, GUARD_CLOSE)

	return ListAccess{id, indexExp}
}

func parseListItemExpr(p *pipe) Expression {

	var expr Expression

	if p.matchAny(LIST_START, LIST_END) {
		expr = ListItemRef{p.next()}
	} else {
		expr = parseExpression(p)
	}

	if expr == nil {
		err.Panic(
			errMsg("parseListItemExpr", `expression or position reference`, p.next()),
			err.At(p.next()),
		)
	}

	return expr
}
