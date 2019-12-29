package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestMatchOperator(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchOperator(tc)
	}

	// Match
	tc := dummyTC([]token.Token{
		token.OfKind(token.OPERATOR),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// No match
	tc = dummyTC([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)
}
