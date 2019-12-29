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
	testMatcher(t, 3, false, doTest,
		token.OfKind(token.OPEN_GUARD),
		token.OfValue(token.INT_LITERAL, "123"),
		token.OfKind(token.CLOSE_GUARD),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.OPEN_GUARD),
		token.OfKind(token.FUNC),
	)
}
