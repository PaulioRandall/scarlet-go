package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseAssign parses the next statement as an assignment statement.
func parseAssign(tm *TokenMatcher) (ids []eval.Expr) {

	t, _ := readExpect(tm, token.ID, token.VOID)
	ids = append(ids, eval.NewForID(t))

	for 1 == tm.Match(token.DELIM) {
		_ = tm.Skip()

		t, _ = readExpect(tm, token.ID, token.VOID)
		ids = append(ids, eval.NewForID(t))
	}

	if 1 == tm.Match(token.FUNC) {
		_ = tm.Skip()
		_ = parseFuncDef(tm)
		//_ = parseFuncBody(tc)
	}

	// TODO: Other possibile expressions

	return
}

func parseFuncDef(tm *TokenMatcher) (f eval.Expr) {

	_, _ = readExpect(tm, token.OPEN_PAREN)

	//params := parseFuncParams(tm)
	//returns := parseFuncReturns(tm)

	_, _ = readExpect(tm, token.CLOSE_PAREN)

	// TODO: Body

	return
}

// parseFuncParams   := [ ID { "," ID } ].
func parseFuncParams(tm *TokenMatcher) (ps []eval.Expr) {

	if 0 == tm.Match(token.ID) {
		return
	}

	t, _ := tm.Read()
	ps = append(ps, eval.NewForID(t))

	for 1 == tm.Match(token.DELIM) {
		tm.Skip()

		t, _ = readExpect(tm, token.ID)
		ps = append(ps, eval.NewForID(t))
	}

	return
}

// parseFuncReturns := [ "->" ID { "," ID } ].
func parseFuncReturns(tm *TokenMatcher) (re []eval.Expr) {

	if 1 == tm.Match(token.RETURNS) {
		tm.Skip()

		t, _ := readExpect(tm, token.ID)
		re = append(re, eval.NewForID(t))

		for 1 == tm.Match(token.DELIM) {
			tm.Skip()

			t, _ = readExpect(tm, token.ID)
			re = append(re, eval.NewForID(t))
		}
	}

	return
}

func readExpect(tm *TokenMatcher, ks ...token.Kind) (_ token.Token, _ bool) {

	if 0 == tm.MatchAny(ks...) {
		kindNames := token.KindsToStrings(ks)
		s := "Expected" + strings.Join(kindNames, " or ") + " token"

		t, _ := tm.Peek()
		badToken(s, t)
	}

	return tm.Read()
}

func badToken(s string, t token.Token) {
	if t.IsZero() {
		panic(s)
	}

	panic(s + " [" + t.String() + "]")
}
