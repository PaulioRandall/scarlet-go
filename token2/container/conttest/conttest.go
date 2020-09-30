package conttest

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/container"
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"

	"github.com/stretchr/testify/require"
)

type LexItr interface {
	More() bool
	Next() lexeme.Lexeme
}

func Feign(lexs ...lexeme.Lexeme) *container.Container {
	con := container.New()
	for _, l := range lexs {
		con.Put(l)
	}
	return con
}

func RequireEqual(t *testing.T, exp, act LexItr) {
	for i := 0; exp.More() || act.More(); i++ {

		if !exp.More() {
			require.True(t, false,
				"Unexpected lexeme in iterator at %d, have %s", i, act.Next().String())
		}

		if !act.More() {
			require.True(t, false,
				"Unexpected iterator end at %d, want %s", i, exp.Next().String())
		}

		expLex, actLex := exp.Next(), act.Next()
		require.Equal(t, expLex, actLex,
			"Unexpected item at %d, have %s, want %s",
			i, expLex.String(), actLex.String)
	}
}
