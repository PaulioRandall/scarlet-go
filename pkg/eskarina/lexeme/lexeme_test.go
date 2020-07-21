package lexeme

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"

	"github.com/stretchr/testify/require"
)

func tok(raw string, props ...prop.Prop) *Lexeme {
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

func check(t *testing.T, act *Lexeme, exps ...*Lexeme) {

	req := func(exp, act *Lexeme) {
		require.NotNil(t, act)
		require.Equal(t, exp.Props, act.Props)
		require.Equal(t, exp.Raw, act.Raw)
	}

	var last *Lexeme

	for _, exp := range exps {
		last = act
		req(exp, act)
		act = act.Next
	}

	require.Nil(t, act)
	act = last

	for i := len(exps) - 1; i >= 0; i-- {
		req(exps[i], act)
		act = act.Prev
	}
}

func Test_Has(t *testing.T) {

	lex := tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER)

	require.True(t, lex.Has(prop.PR_TERM))
	require.True(t, lex.Has(prop.PR_LITERAL))
	require.True(t, lex.Has(prop.PR_NUMBER))

	require.False(t, lex.Has(prop.PR_IDENTIFIER))
}

func Test_Is(t *testing.T) {

	lex := tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER)

	require.True(t, lex.Is(prop.PR_TERM))
	require.True(t, lex.Is(prop.PR_TERM, prop.PR_LITERAL))
	require.True(t, lex.Is(prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER))
	require.True(t, lex.Is())

	require.False(t, lex.Is(prop.PR_IDENTIFIER))
	require.False(t, lex.Is(prop.PR_TERM, prop.PR_LITERAL, prop.PR_BOOL))
}

func Test_Any(t *testing.T) {

	lex := tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER)

	require.True(t, lex.Any(prop.PR_TERM))
	require.True(t, lex.Any(prop.PR_LITERAL, prop.PR_IDENTIFIER))
	require.True(t, lex.Any(prop.PR_SPELL, prop.PR_OPENER, prop.PR_NUMBER))

	require.False(t, lex.Any())
	require.False(t, lex.Any(prop.PR_IDENTIFIER))
	require.False(t, lex.Any(prop.PR_SPELL, prop.PR_OPENER, prop.PR_CLOSER))
}

func Test_ShiftUp(t *testing.T) {

	a := tok("true", prop.PR_BOOL)
	b := tok("1", prop.PR_NUMBER)
	c := tok(`"abc"`, prop.PR_STRING)

	_ = feign(a, b, c)

	a.ShiftUp()
	check(t, a, a, b, c)

	b.ShiftUp()
	check(t, b, b, a, c)

	c.ShiftUp()
	c.ShiftUp()
	check(t, c, c, b, a)
}

func Test_ShiftDown(t *testing.T) {

	a := tok("true", prop.PR_BOOL)
	b := tok("1", prop.PR_NUMBER)
	c := tok(`"abc"`, prop.PR_STRING)

	_ = feign(a, b, c)

	a.ShiftDown()
	check(t, b, b, a, c)

	a.ShiftDown()
	check(t, b, b, c, a)

	a.ShiftDown()
	check(t, b, b, c, a)

	b.ShiftDown()
	b.ShiftDown()
	check(t, c, c, a, b)
}
