package lextest

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"

	"github.com/stretchr/testify/require"
)

func Feign(lexs ...*lexeme.Lexeme) *lexeme.Container {

	que := &lexeme.Container{}

	for _, l := range lexs {
		que.Put(l)
	}

	return que
}

func Lex(line, col int, raw string, tk lexeme.Token) *lexeme.Lexeme {
	return &lexeme.Lexeme{
		Tok:  tk,
		Raw:  raw,
		Line: line,
		Col:  col,
	}
}

func Tok(raw string, tk lexeme.Token) *lexeme.Lexeme {
	return &lexeme.Lexeme{
		Tok: tk,
		Raw: raw,
	}
}

func Equal(t *testing.T, exp, act *lexeme.Lexeme) {

	idx := 0
	for exp != nil || act != nil {

		if exp == nil && act != nil {
			require.Nil(t, act, "Want: EOF\nHave: %s", act.String())
		}

		if exp != nil && act == nil {
			require.NotNil(t, act, "Want: %s\nHave: nil", exp.String())
		}

		equalContent(t, exp, act, fmt.Sprintf(
			"Unexepected Lexeme[%d]\nWant: %s\nHave: %s",
			idx, exp.String(), act.String(),
		))

		equalContent(t, exp.Prev(), act.Prev(), fmt.Sprintf(
			"Unexepected Lexeme[%d].prev\nWant: %s\nHave: %s",
			idx, exp.String(), act.String(),
		))

		exp, act = exp.Next(), act.Next()
		idx++
	}
}

func equalContent(t *testing.T, exp, act *lexeme.Lexeme, msg string) {

	if exp == nil {
		require.Nil(t, act, msg)
		return
	}

	require.NotNil(t, act, msg)
	require.Equal(t, exp.Tok, act.Tok, msg)
	require.Equal(t, exp.Raw, act.Raw, msg)
	require.Equal(t, exp.Line, act.Line, msg)
	require.Equal(t, exp.Col, act.Col, msg)
}
