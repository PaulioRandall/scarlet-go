package parser

import (
	"testing"

	//"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func TestMatchIdOrVoid_1(t *testing.T) {

	doTest := func(tc *TokenCollector, exp bool, err bool) {
		ev, e := matchIdOrVoid(tc)

		require.Equal(t, exp, ev != nil)
		require.Equal(t, err, e != nil)
	}

	// Match
	tc := setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
	})
	doTest(tc, true, false)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTest(tc, false, false)
}

func TestMatchIdArray_1(t *testing.T) {

	doTest := func(tc *TokenCollector, exp bool, err bool) {
		ev, e := matchIdArray(tc)

		require.Equal(t, exp, ev != nil)
		require.Equal(t, err, e != nil)
	}

	// Match single
	tc := setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
	})
	doTest(tc, true, false)

	// Match multiple
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
		token.NewToken(token.ID_DELIM, "", 0, 0),
		token.NewToken(token.ID, "", 0, 0),
		token.NewToken(token.ID_DELIM, "", 0, 0),
		token.NewToken(token.ID, "", 0, 0),
	})
	doTest(tc, true, false)

	// No match
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTest(tc, false, false)

	// Invalid syntax
	tc = setupTokenCollector([]token.Token{
		token.NewToken(token.ID, "", 0, 0),
		token.NewToken(token.ID_DELIM, "", 0, 0),
		token.NewToken(token.FUNC, "", 0, 0),
	})
	doTest(tc, false, true)
}
