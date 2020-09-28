package conttest

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/container"
	"github.com/PaulioRandall/scarlet-go/token/lexeme"

	"github.com/stretchr/testify/require"
)

type LexItr interface {
	HasNext() bool
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
	for i := 0; exp.HasNext() || act.HasNext(); i++ {

		if !exp.HasNext() {
			require.True(t, false,
				"Unexpected lexeme in iterator at %d, have %s", i, act.Next().String())
		}

		if !act.HasNext() {
			require.True(t, false,
				"Unexpected iterator end at %d, want %s", i, exp.Next().String())
		}

		expLex, actLex := exp.Next(), act.Next()
		require.Equal(t, expLex, actLex,
			"Unexpected item at %d, have %s, want %s",
			i, expLex.String(), actLex.String)
	}
}
