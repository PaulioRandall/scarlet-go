package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/gytha/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func isList(p *pipe) bool {
	return p.match(TK_LIST)
}

func parseList(p *pipe) Expression {
	// pattern := LIST LIST_OPEN {expression} LIST_CLOSE

	return List{
		Key:   p.expect(`parseList`, TK_LIST),
		Open:  p.expect(`parseList`, TK_BLOCK_OPEN),
		Exprs: parseExpressions(p),
		Close: p.expect(`parseList`, TK_BLOCK_CLOSE),
	}
}

func isListAccess(p *pipe) bool {
	return p.matchSequence(TK_IDENTIFIER, TK_GUARD_OPEN)
}

func parseListAccess(p *pipe) ListAccess {
	// pattern := ID GUARD_OPEN expression GUARD_CLOSE

	tk := p.expect(`parseListAccess`, TK_IDENTIFIER)
	id := Identifier{tk}

	p.expect(`parseListAccess`, TK_GUARD_OPEN)
	indexExp := parseListItemExpr(p)
	p.expect(`parseListAccess`, TK_GUARD_CLOSE)

	return ListAccess{id, indexExp}
}

func parseListItemExpr(p *pipe) Expression {

	var expr Expression

	if p.matchAny(TK_LIST_START, TK_LIST_END) {
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
