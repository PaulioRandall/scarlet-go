package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// ITEM_ACCESS      := "[" ( ID | INTEGER ) "]" .
func matchItemAccess(tc *TokenCollector) (_ eval.Expr, _ int) {

	var (
		iExpr eval.Expr
		i     int
	)

	t, n := tc.Read(), 1

	if t.Kind != token.OPEN_GUARD {
		goto NO_MATCH
	}

	iExpr, i = matchIdOrInt(tc)

	if iExpr == nil {
		goto NO_MATCH
	}

	n += i
	t = tc.Read()
	n++

	if t.Kind != token.CLOSE_GUARD {
		goto NO_MATCH
	}

	return iExpr, n

NO_MATCH:
	tc.PutBack(n)
	return
}
