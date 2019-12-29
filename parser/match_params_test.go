package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

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

func TestMatchParamList(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchParamList(tc)
	}

	// Match single
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.ID),
	)

	// Match multiple
	testMatcher(t, 5, false, doTest,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.STR_LITERAL),
		token.OfKind(token.DELIM),
		token.OfKind(token.VOID),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.OPERATOR),
	)

	// Error
	testMatcher(t, 0, true, doTest,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
	)
}
