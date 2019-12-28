package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestMatchItemAccess_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchItemAccess(tc)
	}

	// Match
	tc := dummyTC([]token.Token{
		token.OfKind(token.OPEN_GUARD),
		token.OfValue(token.INT_LITERAL, "123"),
		token.OfKind(token.CLOSE_GUARD),
	})
	doTestMatch(t, tc, 3, false, doTest)

	// No match
	tc = dummyTC([]token.Token{
		token.OfKind(token.OPEN_GUARD),
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)
}
