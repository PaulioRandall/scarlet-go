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

func halfEqual(t *testing.T, exp, act *Lexeme) {

	if exp == nil {
		require.Nil(t, act)
		return
	}

	require.NotNil(t, act)
	require.Equal(t, exp.Tok, act.Tok)
	require.Equal(t, exp.Raw, act.Raw)
}

func feign(lexs ...*Lexeme) {

	var last *Lexeme

	for _, l := range lexs {

		if last != nil {
			append(last, l)
		}

		last = l
	}
}

func setup() (a, b, c, d *Lexeme) {
	a = tok("1st", BOOL)
	b = tok("2nd", NUMBER)
	c = tok("3rd", STRING)
	d = tok("4th", IDENTIFIER)
	return
}

func setupContainer() (_ *Container, a, b, c, d *Lexeme) {

	a, b, c, d = setup()

	a.prev, a.next = nil, b
	b.prev, b.next = a, c
	c.prev, c.next = b, nil

	con := &Container{
		size: 3,
		head: a,
		tail: c,
	}

	return con, a, b, c, d
}

func fullEqual(t *testing.T, exp, prev, next, act *Lexeme) {

	require.NotNil(t, act)
	require.Equal(t, exp.Tok, act.Tok)
	require.Equal(t, exp.Raw, act.Raw)

	halfEqual(t, prev, act.prev)
	halfEqual(t, next, act.next)
}
