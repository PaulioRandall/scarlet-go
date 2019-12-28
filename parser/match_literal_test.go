package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestMatchLiteral_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchLiteral(tc)
	}

	// Match bool `true`
	tc := dummyTC([]token.Token{
		token.OfValue(token.TRUE, "TRUE"),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match bool `false`
	tc = dummyTC([]token.Token{
		token.OfValue(token.FALSE, "FALSE"),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match int
	tc = dummyTC([]token.Token{
		token.OfValue(token.INT_LITERAL, "123"),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match real
	tc = dummyTC([]token.Token{
		token.OfValue(token.REAL_LITERAL, "123.456"),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match string
	tc = dummyTC([]token.Token{
		token.OfValue(token.STR_LITERAL, "wololo"),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match template
	tc = dummyTC([]token.Token{
		token.OfValue(token.STR_TEMPLATE, "wololo"),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// No match
	tc = dummyTC([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)

	// Match invalid bool
	tc = dummyTC([]token.Token{
		token.OfValue(token.TRUE, "?"),
	})
	doTestMatch(t, tc, 0, true, doTest)

	// Match invalid real
	tc = dummyTC([]token.Token{
		token.OfValue(token.REAL_LITERAL, "?"),
	})
	doTestMatch(t, tc, 0, true, doTest)
}
