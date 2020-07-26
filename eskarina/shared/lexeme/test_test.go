package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func tok(raw string, props ...Prop) *Lexeme {
	return &Lexeme{
		Props: props,
		Raw:   raw,
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
	a = tok("true", PR_LITERAL, PR_BOOL)
	b = tok("1", PR_LITERAL, PR_NUMBER)
	c = tok(`"abc"`, PR_LITERAL, PR_STRING)
	d = tok("x", PR_IDENTIFIER)
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
	require.Equal(t, exp.Props, act.Props)
	require.Equal(t, exp.Raw, act.Raw)
}

func fullEqual(t *testing.T, exp, prev, next, act *Lexeme) {

	require.NotNil(t, act)
	require.Equal(t, exp.Props, act.Props)
	require.Equal(t, exp.Raw, act.Raw)

	halfEqual(t, prev, act.Prev)
	halfEqual(t, next, act.Next)
}
