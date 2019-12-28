package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func dummyTC(stream []token.Token) *TokenCollector {
	st := lexor.DummyScanToken(stream)
	tr := NewTokenReader(st)
	return NewTokenCollector(tr)
}

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
