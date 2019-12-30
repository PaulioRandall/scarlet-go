package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

type matcher func(tc *TokenCollector) int

func testMatcher(
	t *testing.T,
	exp int,
	err bool,
	f matcher,
	in ...token.Token,
) {

	tc := dummyTC(in...)

	if err {
		require.Panics(t, func() { f(tc) }, "Expected a panic")
		return
	}

	require.Equal(t, exp, f(tc), "Expected %d matched tokens", exp)
}
