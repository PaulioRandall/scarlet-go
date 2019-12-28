package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// LIST_ACCESS      := ID [ ITEM_ACCESS ] .
func matchListAccess(tc *TokenCollector) eval.Expr {

	t := tc.Read()

	if t.Kind != token.ID {
		tc.PutBack(1)
		return nil
	}

	indexEv := matchItemAccess(tc)

	if indexEv == nil {
		tc.PutBack(1)
		return nil
	}

	idEv := eval.NewForID(t)
	return eval.NewForListAccess(idEv, indexEv)
}

// ITEM_ACCESS      := "[" ( ID | INTEGER ) "]" .
func matchItemAccess(tc *TokenCollector) eval.Expr {

	t := tc.Read()

	if t.Kind != token.OPEN_GUARD {
		tc.PutBack(1)
		return nil
	}

	indexExpr := matchIdOrInt(tc)

	if indexExpr == nil {
		tc.PutBack(1)
		return nil
	}

	t = tc.Read()

	if t.Kind != token.CLOSE_GUARD {
		tc.PutBack(3)
		return nil
	}

	return indexExpr
}
