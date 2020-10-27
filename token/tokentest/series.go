package tokentest

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"
	"github.com/PaulioRandall/scarlet-go/token/series"

	"github.com/stretchr/testify/require"
)

var empty lexeme.Lexeme

type LexemeIterator interface {
	More() bool
	Next() lexeme.Lexeme
	Get() lexeme.Lexeme
}

func RequireSeries(t *testing.T, exp, act LexemeIterator) {
	for i := 0; exp.More() || act.More(); i++ {

		if !exp.More() {
			require.True(t, false, errMsg(i, empty, act.Next()))
		}

		if !act.More() {
			require.True(t, false, errMsg(i, exp.Next(), empty))
		}

		require.Equal(t, exp.Next(), act.Next(), errMsg(i, exp.Get(), act.Get()))
	}
}

func FeignSeries(lexs ...lexeme.Lexeme) *series.Series {
	s := series.Make()
	for _, l := range lexs {
		s.Append(l)
	}
	return s
}

func errMsg(i int, exp, act lexeme.Lexeme) string {

	expStr := "Lexeme{}"
	if exp != empty {
		expStr = exp.String()
	}

	actStr := "Lexeme{}"
	if act != empty {
		actStr = act.String()
	}

	return fmt.Sprintf(
		"Unexpected lexeme at %d; have %s, want %s", i, actStr, expStr)
}
