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

// @Deprecated
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

// @Deprecated
func setup() (a, b, c, d *Lexeme) {
	a = tok("true", BOOL)
	b = tok("1", NUMBER)
	c = tok(`"abc"`, STRING)
	d = tok("x", IDENTIFIER)
	return
}

// @Deprecated
func setupList() (a, b, c, d *Lexeme) {
	a, b, c, d = setup()
	_ = feign(a, b, c)
	return
}

// @Deprecated
func halfEqual(t *testing.T, exp, act *Lexeme) {

	if exp == nil {
		require.Nil(t, act)
		return
	}

	require.NotNil(t, act)
	require.Equal(t, exp.Tok, act.Tok)
	require.Equal(t, exp.Raw, act.Raw)
}

// @Deprecated
func fullEqual(t *testing.T, exp, prev, next, act *Lexeme) {

	require.NotNil(t, act)
	require.Equal(t, exp.Tok, act.Tok)
	require.Equal(t, exp.Raw, act.Raw)

	halfEqual(t, prev, act.Prev)
	halfEqual(t, next, act.Next)
}

func feign2(lexs ...*Lexeme) {

	var last *Lexeme

	for _, l := range lexs {

		if last != nil {
			last.append(l)
		}

		last = l
	}
}

func setup2() (a, b, c, d *Lexeme) {
	a = tok("1st", BOOL)
	b = tok("2nd", NUMBER)
	c = tok("3rd", STRING)
	d = tok("4th", IDENTIFIER)
	return
}

func setupContainer2() (_ *Container2, a, b, c, d *Lexeme) {

	a, b, c, d = setup2()

	a.prev, a.next = nil, b
	b.prev, b.next = a, c
	c.prev, c.next = b, nil

	con := &Container2{
		size: 3,
		head: a,
		tail: c,
	}

	return con, a, b, c, d
}

func fullEqual2(t *testing.T, exp, prev, next, act *Lexeme) {

	require.NotNil(t, act)
	require.Equal(t, exp.Tok, act.Tok)
	require.Equal(t, exp.Raw, act.Raw)

	halfEqual(t, prev, act.prev)
	halfEqual(t, next, act.next)
}
