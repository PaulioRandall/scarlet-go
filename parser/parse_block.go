package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

func parseStatement(tc *TokenCollector) (_ eval.Expr) {

	if n := matchGlobalAssign(tc); n > 0 {
		return parseAssign(tc)
	}

	return
}

func parseAssign(tc *TokenCollector) (_ eval.Expr) {

	var assignToken token.Token

	// TODO: Count the IDs to make sure we have the same number of expressions

	for _, t := range tc.Take() {
		if t.Kind == token.GLOBAL {
			continue
		}

		if t.Kind == token.ID || t.Kind == token.VOID {
			// TODO: create expressions for these
		}

		if t.Kind == token.ASSIGN {
			assignToken = t
		}
	}

	println(assignToken.String()) // DELETE ME
	// TODO: Match the expressions then parse them

	return
}
