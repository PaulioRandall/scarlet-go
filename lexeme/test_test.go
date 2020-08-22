package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func lex(line, col int, raw string, tk Token) *Lexeme {
	return &Lexeme{
		Tok:  tk,
		Raw:  raw,
		Line: line,
		Col:  col,
	}
}

func tok(raw string, tk Token) *Lexeme {
	return &Lexeme{
		Tok: tk,
		Raw: raw,
	}
}

func halfEqual(t *testing.T, exp, act *Lexeme) {

	if exp == nil {
		require.Nil(t, act)
		return
	}

	require.NotNil(t, act)
	require.Equal(t, exp.Tok, act.Tok)
	require.Equal(t, exp.Raw, act.Raw)
}

func fullEqual(t *testing.T, exp, prev, next, act *Lexeme) {

	require.NotNil(t, act)
	require.Equal(t, exp.Tok, act.Tok)
	require.Equal(t, exp.Raw, act.Raw)

	halfEqual(t, prev, act.prev)
	halfEqual(t, next, act.next)
}

func feign(lexs ...*Lexeme) *Lexeme {

	var first *Lexeme
	var last *Lexeme

	for _, l := range lexs {

		l.next = nil
		l.prev = nil

		if first == nil {
			first = l
		}

		if last != nil {
			append(last, l)
		}

		last = l
	}

	return first
}
