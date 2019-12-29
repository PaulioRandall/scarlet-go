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
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.OPERATOR),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)
}

func TestMatchParam(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchParam(tc)
	}

	// Match id
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.ID),
	)

	// Match literal
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.STR_LITERAL),
	)

	// Match void
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.VOID),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)
}
