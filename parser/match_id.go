package parser

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/parser/ctx"
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// ID               := LETTER { "\_" | LETTER } .
// INTEGER          := DIGIT { DIGIT } .
func matchIdOrInt(tc *TokenCollector) eval.Expr {

	t := tc.Read()

	if t.Kind == token.ID {
		return eval.NewForID(t)
	}

	if t.Kind != token.INT_LITERAL {
		tc.PutBack(1)
		return nil
	}

	i, e := strconv.Atoi(t.Value)
	if e != nil {
		panic(NewParseErr("Could not parse INT token value", nil, t))
	}

	v := ctx.NewValue(ctx.INT, i)
	return eval.NewForValue(v)
}

// ID_OR_VOID       := ID | "\_" .
func matchIdOrVoid(tc *TokenCollector) eval.Expr {

	t := tc.Read()

	if t.Kind != token.ID && t.Kind != token.VOID {
		tc.PutBack(1)
		return nil
	}

	return eval.NewForID(t)
}

// ID_ARRAY         := ID_OR_VOID { "," ID_OR_VOID } .
func matchIdArray(tc *TokenCollector) []eval.Expr {

	ev := matchIdOrVoid(tc)

	if ev == nil {
		return nil
	}

	ids := []eval.Expr{ev}
	matchMoreIds(ids, tc)

	return ids
}

// *ID_ARRAY        := ... { "," ID_OR_VOID } .
func matchMoreIds(ids []eval.Expr, tc *TokenCollector) {

	for tc.Peek().Kind == token.ID_DELIM {

		_ = tc.Read() // Skip the delimiter

		ev := matchIdOrVoid(tc)

		if ev == nil {
			panic(NewParseErr("Expected ID token", nil, tc.Peek()))
		}

		ids = append(ids, ev)
	}
}
