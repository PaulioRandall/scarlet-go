package parser

/*
import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

type matcher func(tm *TokenMatcher) int

func testMatcher(
	t *testing.T,
	exp int,
	err bool,
	f matcher,
	in ...token.Token,
) {

	tm := dummyTM(in...)

	if err {
		require.Panics(t, func() { f(tm) }, "Expected a panic")
		return
	}

	require.Equal(t, exp, f(tm), "Expected %d matched tokens", exp)
}
*/
