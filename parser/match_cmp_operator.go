package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// CMP_OPERATOR     := "=" | "#" | "<" | ">" | "<=" | ">=" .
func matchCmpOperator(tc *TokenCollector) (_ eval.Expr, _ int) {

	t := tc.Read()

	if t.Kind != token.CMP_OPERATOR {
		tc.PutBack(1)
		return
	}

	return eval.NewForOperator(t), 1
}
