package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

type matcher func(tc *TokenCollector) (interface{}, int)

func dummyTC(stream []token.Token) *TokenCollector {
	st := lexor.DummyScanToken(stream)
	tr := NewTokenReader(st)
	return NewTokenCollector(tr)
}

func testMatcher(
	t *testing.T,
	exp int,
	err bool,
	f matcher,
	in ...token.Token,
) {
	doTestMatch(t, dummyTC(in), exp, err, f)
}

func doTestMatch(
	t *testing.T,
	tc *TokenCollector,
	exp int,
	err bool,
	f matcher,
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
