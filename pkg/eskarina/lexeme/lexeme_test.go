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

func check(t *testing.T, act *Lexeme, exps ...*Lexeme) {

	var last *Lexeme

	for _, exp := range exps {
		last = act
		halfEqual(t, exp, act)
		act = act.Next
	}

	require.Nil(t, act)
	act = last

	for i := len(exps) - 1; i >= 0; i-- {
		halfEqual(t, exps[i], act)
		act = act.Prev
	}
}

func Test_Lexeme_Has(t *testing.T) {

	lex := tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER)

	require.True(t, lex.Has(prop.PR_TERM))
	require.True(t, lex.Has(prop.PR_LITERAL))
	require.True(t, lex.Has(prop.PR_NUMBER))

	require.False(t, lex.Has(prop.PR_IDENTIFIER))
}

func Test_Lexeme_Is(t *testing.T) {

	lex := tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER)

	require.True(t, lex.Is(prop.PR_TERM))
	require.True(t, lex.Is(prop.PR_TERM, prop.PR_LITERAL))
	require.True(t, lex.Is(prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER))
	require.True(t, lex.Is())

	require.False(t, lex.Is(prop.PR_IDENTIFIER))
	require.False(t, lex.Is(prop.PR_TERM, prop.PR_LITERAL, prop.PR_BOOL))
}

func Test_Lexeme_Any(t *testing.T) {

	lex := tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER)

	require.True(t, lex.Any(prop.PR_TERM))
	require.True(t, lex.Any(prop.PR_LITERAL, prop.PR_IDENTIFIER))
	require.True(t, lex.Any(prop.PR_SPELL, prop.PR_OPENER, prop.PR_NUMBER))

	require.False(t, lex.Any())
	require.False(t, lex.Any(prop.PR_IDENTIFIER))
	require.False(t, lex.Any(prop.PR_SPELL, prop.PR_OPENER, prop.PR_CLOSER))
}

func Test_Lexeme_ShiftUp(t *testing.T) {

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

func Test_Lexeme_ShiftDown(t *testing.T) {

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

func Test_Lexeme_Prepend(t *testing.T) {

	a := tok("true", prop.PR_BOOL)
	b := tok("1", prop.PR_NUMBER)
	c := tok(`"abc"`, prop.PR_STRING)

	b.Prepend(a)
	check(t, a, a, b)

	c.Prepend(b)
	check(t, a, a, b, c)
}

func Test_Lexeme_Append(t *testing.T) {

	a := tok("true", prop.PR_BOOL)
	b := tok("1", prop.PR_NUMBER)
	c := tok(`"abc"`, prop.PR_STRING)

	b.Append(c)
	check(t, b, b, c)

	a.Append(b)
	check(t, a, a, b, c)
}

func Test_Lexeme_Remove(t *testing.T) {

	setup := func() (a, b, c *Lexeme) {
		a = tok("true", prop.PR_BOOL)
		b = tok("1", prop.PR_NUMBER)
		c = tok(`"abc"`, prop.PR_STRING)
		feign(a, b, c)
		return
	}

	a, b, c := setup()
	a.Remove()
	check(t, a, a)
	check(t, b, b, c)

	a, b, c = setup()
	b.Remove()
	check(t, a, a, c)

	a, b, c = setup()
	c.Remove()
	check(t, a, a, b)
}
