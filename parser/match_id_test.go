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
	tc := dummyTC([]token.Token{
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match int
	tc = dummyTC([]token.Token{
		token.OfValue(token.INT_LITERAL, "123"),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// No match
	tc = dummyTC([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)

	// Invalid syntax
	tc = dummyTC([]token.Token{
		token.OfValue(token.INT_LITERAL, "abc"),
	})
	doTestMatch(t, tc, 0, true, doTest)
}

func TestMatchIdOrVoid_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdOrVoid(tc)
	}

	// Match ID
	tc := dummyTC([]token.Token{
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match VOID
	tc = dummyTC([]token.Token{
		token.OfKind(token.VOID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// No match
	tc = dummyTC([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)
}

func TestMatchIdArray_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdArray(tc)
	}

	// Match single
	tc := dummyTC([]token.Token{
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match multiple
	tc = dummyTC([]token.Token{
		token.OfKind(token.ID),
		token.OfKind(token.ID_DELIM),
		token.OfKind(token.ID),
		token.OfKind(token.ID_DELIM),
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 5, false, doTest)

	// No match
	tc = dummyTC([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)

	// Invalid syntax
	tc = dummyTC([]token.Token{
		token.OfKind(token.ID),
		token.OfKind(token.ID_DELIM),
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, true, doTest)
}

func TestMatchIdOrItem_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdOrItem(tc)
	}

	// Match ID
	tc := dummyTC([]token.Token{
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match item access
	tc = dummyTC([]token.Token{
		token.OfKind(token.ID),
		token.OfKind(token.OPEN_GUARD),
		token.OfValue(token.INT_LITERAL, "123"),
		token.OfKind(token.CLOSE_GUARD),
	})
	doTestMatch(t, tc, 4, false, doTest)

	// No match
	tc = dummyTC([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)
}
