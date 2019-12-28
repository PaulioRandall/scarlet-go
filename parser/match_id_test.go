package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func doTestMatch(
	t *testing.T,
	tc *TokenCollector,
	exp, err bool,
	f func(*TokenCollector) interface{},
) {

	if err {
		require.Panics(t, func() { f(tc) }, "Expected a panic")
	} else if exp {
		require.NotNil(t, f(tc), "Expected an Expr")
	} else {
		require.Nil(t, f(tc), "Expected nil Expr")
	}
}

func TestMatchIdOrInt_1(t *testing.T) {

	doTest := func(tc *TokenCollector) interface{} {
		return matchIdOrInt(tc)
	}

	// Match ID
	tc := setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
	})
	doTestMatch(t, tc, true, false, doTest)

	// Match int
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.NUM_LITERAL, "123", 0, 0),
	})
	doTestMatch(t, tc, true, false, doTest)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTestMatch(t, tc, false, false, doTest)

	// Invalid syntax
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.NUM_LITERAL, "abc", 0, 0),
	})
	doTestMatch(t, tc, false, true, doTest)
}

func TestMatchIdOrVoid_1(t *testing.T) {

	doTest := func(tc *TokenCollector) interface{} {
		return matchIdOrVoid(tc)
	}

	// Match ID
	tc := setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
	})
	doTestMatch(t, tc, true, false, doTest)

	// Match VOID
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.VOID, "", 0, 0),
	})
	doTestMatch(t, tc, true, false, doTest)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTestMatch(t, tc, false, false, doTest)
}

func TestMatchIdArray_1(t *testing.T) {

	doTest := func(tc *TokenCollector) interface{} {
		return matchIdArray(tc)
	}

	// Match single
	tc := setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
	})
	doTestMatch(t, tc, true, false, doTest)

	// Match multiple
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
		token.NewToken(token.ID_DELIM, "", 0, 0),
		token.NewToken(token.ID, "", 0, 0),
		token.NewToken(token.ID_DELIM, "", 0, 0),
		token.NewToken(token.ID, "", 0, 0),
	})
	doTestMatch(t, tc, true, false, doTest)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTestMatch(t, tc, false, false, doTest)

	// Invalid syntax
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
		token.NewToken(token.ID_DELIM, "", 0, 0),
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTestMatch(t, tc, false, true, doTest)
}
