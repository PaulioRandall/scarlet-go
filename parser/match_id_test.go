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
		require.Panics(t, func() { f(tc) })
	} else if exp {
		require.NotNil(t, f(tc))
	} else {
		require.Nil(t, f(tc))
	}
}

func TestMatchIdOrVoid_1(t *testing.T) {

	doTest := func(tc *TokenCollector) interface{} {
		return matchIdOrVoid(tc)
	}

	// Match
	tc := setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
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
