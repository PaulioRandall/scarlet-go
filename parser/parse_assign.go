package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

// parseAssign parses the next statement as an assignment statement.
func parseAssign(tm *TokenMatcher) (_ Expr) {

	var ids []Expr
	var ex []Expr

	t, _ := readExpect(tm, token.ID, token.VOID)
	ids = append(ids, NewForID(t))

	for 1 == tm.Match(token.DELIM) {
		_ = tm.Skip()

		t, _ = readExpect(tm, token.ID, token.VOID)
		ids = append(ids, NewForID(t))
	}

	if 1 == tm.Match(token.FUNC) {
		_ = tm.Skip()
		ex = append(ex, parseFuncDef(tm))
		return NewForAssign(ids, ex)
	}

	// TODO: Other possibile expressions

	return
}

func parseFuncDef(tm *TokenMatcher) (f Expr) {

	t, _ := readExpect(tm, token.ID)
	id := NewForID(t)

	_, _ = readExpect(tm, token.OPEN_PAREN)

	params := parseFuncParams(tm)
	returns := parseFuncReturns(tm)

	_, _ = readExpect(tm, token.CLOSE_PAREN)

	var body []Expr
	if 1 == tm.Match(token.DO) {
		_, _ = readExpect(tm, token.DO)
		body = parseBlock(tm)
	} else {
		body = []Expr{parseStatement(tm)}
	}

	return NewForFuncDef(id, params, returns, body)
}

// parseFuncParams   := [ ID { "," ID } ].
func parseFuncParams(tm *TokenMatcher) (ps []Expr) {

	if 0 == tm.Match(token.ID) {
		return
	}

	t, _ := tm.Read()
	ps = append(ps, NewForID(t))

	for 1 == tm.Match(token.DELIM) {
		tm.Skip()

		t, _ = readExpect(tm, token.ID)
		ps = append(ps, NewForID(t))
	}

	return
}

// parseFuncReturns := [ "->" ID { "," ID } ].
func parseFuncReturns(tm *TokenMatcher) (re []Expr) {

	if 1 == tm.Match(token.RETURNS) {
		tm.Skip()

		t, _ := readExpect(tm, token.ID)
		re = append(re, NewForID(t))

		for 1 == tm.Match(token.DELIM) {
			tm.Skip()

			t, _ = readExpect(tm, token.ID)
			re = append(re, NewForID(t))
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
