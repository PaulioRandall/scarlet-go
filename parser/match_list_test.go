package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestMatchListAccess_1(t *testing.T) {

	doTest := func(tc *TokenCollector) interface{} {
		return matchListAccess(tc)
	}

	// Match
	tc := setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
		token.NewToken(token.OPEN_GUARD, "", 0, 0),
		token.NewToken(token.INT_LITERAL, "123", 0, 0),
		token.NewToken(token.CLOSE_GUARD, "", 0, 0),
	})
	doTestMatch(t, tc, true, false, doTest)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
		token.NewToken(token.OPEN_GUARD, "", 0, 0),
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTestMatch(t, tc, false, false, doTest)
}

func TestMatchItemAccess_1(t *testing.T) {

	doTest := func(tc *TokenCollector) interface{} {
		return matchItemAccess(tc)
	}

	// Match
	tc := setupTokenCollector([]token.Token{
		token.NewToken(token.OPEN_GUARD, "", 0, 0),
		token.NewToken(token.INT_LITERAL, "123", 0, 0),
		token.NewToken(token.CLOSE_GUARD, "", 0, 0),
	})
	doTestMatch(t, tc, true, false, doTest)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.OPEN_GUARD, "", 0, 0),
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTestMatch(t, tc, false, false, doTest)
}
