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
	testMatcher(t, 1, false, doTest,
		token.OfValue(token.BOOL_LITERAL, "TRUE"),
	)

	// Match bool `false`
	testMatcher(t, 1, false, doTest,
		token.OfValue(token.BOOL_LITERAL, "FALSE"),
	)

	// Match int
	testMatcher(t, 1, false, doTest,
		token.OfValue(token.INT_LITERAL, "123"),
	)

	// Match real
	testMatcher(t, 1, false, doTest,
		token.OfValue(token.REAL_LITERAL, "123.456"),
	)

	// Match string
	testMatcher(t, 1, false, doTest,
		token.OfValue(token.STR_LITERAL, "wololo"),
	)

	// Match template
	testMatcher(t, 1, false, doTest,
		token.OfValue(token.STR_TEMPLATE, "wololo"),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)

	// Match invalid bool
	testMatcher(t, 0, true, doTest,
		token.OfValue(token.BOOL_LITERAL, "?"),
	)

	// Match invalid real
	testMatcher(t, 0, true, doTest,
		token.OfValue(token.REAL_LITERAL, "?"),
	)
}
