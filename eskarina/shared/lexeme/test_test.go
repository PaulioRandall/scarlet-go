package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func tok(raw string, tk Token) *Lexeme {
	return &Lexeme{
		Tok: tk,
		Raw: raw,
	}
}

func feign(lexs ...*Lexeme) *Lexeme {

	var first *Lexeme
	var last *Lexeme

	for _, l := range lexs {

		if first == nil {
			first = l
			last = l
			continue
		}

		last.Append(l)
		last = l
	}

	return first
}

func setup() (a, b, c, d *Lexeme) {
	a = tok("true", BOOL)
	b = tok("1", NUMBER)
	c = tok(`"abc"`, STRING)
	d = tok("x", IDENTIFIER)
	return
}

func setupList() (a, b, c, d *Lexeme) {
	a, b, c, d = setup()
	_ = feign(a, b, c)
	return
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

	halfEqual(t, prev, act.Prev)
	halfEqual(t, next, act.Next)
}
