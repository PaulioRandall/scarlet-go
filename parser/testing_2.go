package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

type matcher_2 func(tc *TokenCollector) int

func dummyTC_2(stream ...token.Token) *TokenCollector {
	st := lexor.DummyScanToken(stream)
	tr := NewTokenReader(st)
	return NewTokenCollector(tr)
}

func testMatcher_2(
	t *testing.T,
	exp int,
	err bool,
	f matcher_2,
	in ...token.Token,
) {

	tc := dummyTC(in)

	if err {
		require.Panics(t, func() { f(tc) }, "Expected a panic")
		return
	}

	require.Equal(t, exp, f(tc), "Expected %d matched tokens", exp)
}
