package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestMatchIdOrInt_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdOrInt(tc)
	}

	// Match ID
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.ID),
	)

	// Match int
	testMatcher(t, 1, false, doTest,
		token.OfValue(token.INT_LITERAL, "123"),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)

	// Invalid syntax
	testMatcher(t, 0, true, doTest,
		token.OfValue(token.INT_LITERAL, "abc"),
	)
}

func TestMatchIdOrVoid_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdOrVoid(tc)
	}

	// Match ID
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.ID),
	)

	// Match VOID
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.VOID),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)
}

func TestMatchIdArray_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdArray(tc)
	}

	// Match single
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.ID),
	)

	// Match multiple
	testMatcher(t, 5, false, doTest,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.ID),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)

	// Invalid syntax
	testMatcher(t, 0, true, doTest,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.FUNC),
	)
}

func TestMatchIdOrItem_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdOrItem(tc)
	}

	// Match ID
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.ID),
	)

	// Match item access
	testMatcher(t, 4, false, doTest,
		token.OfKind(token.ID),
		token.OfKind(token.OPEN_GUARD),
		token.OfValue(token.INT_LITERAL, "123"),
		token.OfKind(token.CLOSE_GUARD),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)
}
