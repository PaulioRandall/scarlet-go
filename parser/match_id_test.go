package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func doTestMatch(
	t *testing.T,
	tc *TokenCollector,
	exp int,
	err bool,
	f func(*TokenCollector) (interface{}, int),
) {

	if err {
		require.Panics(t, func() { f(tc) }, "Expected a panic")
		return
	}

	ev, n := f(tc)

	if exp > 0 {
		require.NotNil(t, ev, "Expected an Expr")
		require.Equal(t, exp, n, "Expected %d tokens used", exp)
	} else {
		require.Nil(t, ev, "Expected nil Expr")
		require.Empty(t, n, "Expected 0 tokens used")
	}
}

func TestMatchIdOrInt_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdOrInt(tc)
	}

	// Match ID
	tc := setupTokenCollector([]token.Token{
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match int
	tc = setupTokenCollector([]token.Token{
		token.OfValue(token.INT_LITERAL, "123"),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)

	// Invalid syntax
	tc = setupTokenCollector([]token.Token{
		token.OfValue(token.INT_LITERAL, "abc"),
	})
	doTestMatch(t, tc, 0, true, doTest)
}

func TestMatchIdOrVoid_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdOrVoid(tc)
	}

	// Match ID
	tc := setupTokenCollector([]token.Token{
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match VOID
	tc = setupTokenCollector([]token.Token{
		token.OfKind(token.VOID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)
}

func TestMatchIdArray_1(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchIdArray(tc)
	}

	// Match single
	tc := setupTokenCollector([]token.Token{
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 1, false, doTest)

	// Match multiple
	tc = setupTokenCollector([]token.Token{
		token.OfKind(token.ID),
		token.OfKind(token.ID_DELIM),
		token.OfKind(token.ID),
		token.OfKind(token.ID_DELIM),
		token.OfKind(token.ID),
	})
	doTestMatch(t, tc, 5, false, doTest)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, false, doTest)

	// Invalid syntax
	tc = setupTokenCollector([]token.Token{
		token.OfKind(token.ID),
		token.OfKind(token.ID_DELIM),
		token.OfKind(token.FUNC),
	})
	doTestMatch(t, tc, 0, true, doTest)
}
