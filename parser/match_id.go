package parser

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/parser/ctx"
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// ID               := LETTER { "\_" | LETTER } .
// INTEGER          := DIGIT { DIGIT } .
func matchIdOrInt(tc *TokenCollector) (_ eval.Expr, _ int) {

	t := tc.Read()

	if t.Kind == token.ID {
		return eval.NewForID(t), 1
	}

	if t.Kind != token.INT_LITERAL {
		tc.PutBack(1)
		return
	}

	i, e := strconv.Atoi(t.Value)
	if e != nil {
		panic(NewParseErr("Could not parse INT token value", nil, t))
	}

	v := ctx.NewValue(ctx.INT, i)
	return eval.NewForValue(v), 1
}

// ID_OR_VOID       := ID | "\_" .
func matchIdOrVoid(tc *TokenCollector) (_ eval.Expr, _ int) {

	t := tc.Read()

	if t.Kind != token.ID && t.Kind != token.VOID {
		tc.PutBack(1)
		return
	}

	return eval.NewForID(t), 1
}

// ID_ARRAY         := ID_OR_VOID { "," ID_OR_VOID } .
func matchIdArray(tc *TokenCollector) (_ []eval.Expr, _ int) {

	ev, n := matchIdOrVoid(tc)

	if ev == nil {
		return
	}

	ids := []eval.Expr{ev}
	n += matchMoreIds(ids, tc)

	return ids, n
}

// *ID_ARRAY        := ... { "," ID_OR_VOID } .
func matchMoreIds(ids []eval.Expr, tc *TokenCollector) (_ int) {

	n := 0

	for tc.Peek().Kind == token.DELIM {

		_ = tc.Read() // Skip the delimiter
		n++

		expr, count := matchIdOrVoid(tc)

		if expr == nil {
			panic(NewParseErr("Expected ID token", nil, tc.Peek()))
		}

		n += count
		ids = append(ids, expr)
	}

	return n
}

// ID_OR_ITEM       := ID [ ITEM_ACCESS ] .
func matchIdOrItem(tc *TokenCollector) (_ eval.Expr, _ int) {

	var (
		idExpr eval.Expr
		iExpr  eval.Expr
		i      int
	)

	t, n := tc.Read(), 1

	if t.Kind != token.ID {
		tc.PutBack(n)
		return
	}

	idExpr = eval.NewForID(t)
	iExpr, i = matchItemAccess(tc)

	if iExpr == nil {
		return idExpr, n
	}

	n += i
	return eval.NewForListAccess(idExpr, iExpr), n
}
