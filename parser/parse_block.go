package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

func parseStatement(tc *TokenCollector) (_ eval.Expr) {

	//isGlobal = 1 == matchAny(tc, token.GLOBAL)
	tc.Take()

	if n := matchLeftSideOfAssign(tc); n > 0 {
		//ids, a := parseAssign(tc)
	}

	return
}

// parseAssign parses the next statement as an assignment statement. Assumes
// that the statement matches a valid assignment statement.
func parseAssign(tc *TokenCollector) (ids []eval.Expr, a token.Token) {

	for _, t := range tc.Take() {

		if t.Kind == token.ID || t.Kind == token.VOID {
			ids = append(ids, eval.NewForID(t))
			continue
		}

		if t.Kind == token.ASSIGN {
			a = t
		}
	}

	return
}

// parseExprArray parses an array of expressions.
func parseExprArray(tc *TokenCollector) (_ []eval.Expr) {

	return
}
