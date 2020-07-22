package lexeme

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"

	"github.com/stretchr/testify/require"
)

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
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)

	b.ShiftUp()
	fullEqual(t, b, nil, a, b)
	fullEqual(t, a, b, c, a)
	fullEqual(t, c, a, nil, c)

	c.ShiftUp()
	c.ShiftUp()
	fullEqual(t, c, nil, b, c)
	fullEqual(t, b, c, a, b)
	fullEqual(t, a, b, nil, a)
}

func Test_Lexeme_ShiftDown(t *testing.T) {

	a := tok("true", prop.PR_BOOL)
	b := tok("1", prop.PR_NUMBER)
	c := tok(`"abc"`, prop.PR_STRING)

	_ = feign(a, b, c)

	a.ShiftDown()
	fullEqual(t, b, nil, a, b)
	fullEqual(t, a, b, c, a)
	fullEqual(t, c, a, nil, c)

	a.ShiftDown()
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, a, c)
	fullEqual(t, a, c, nil, a)

	a.ShiftDown()
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, a, c)
	fullEqual(t, a, c, nil, a)

	b.ShiftDown()
	b.ShiftDown()
	fullEqual(t, c, nil, a, c)
	fullEqual(t, a, c, b, a)
	fullEqual(t, b, a, nil, b)
}

func Test_Lexeme_Prepend(t *testing.T) {

	a := tok("true", prop.PR_BOOL)
	b := tok("1", prop.PR_NUMBER)
	c := tok(`"abc"`, prop.PR_STRING)

	b.Prepend(a)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)

	c.Prepend(b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_Lexeme_Append(t *testing.T) {

	a := tok("true", prop.PR_BOOL)
	b := tok("1", prop.PR_NUMBER)
	c := tok(`"abc"`, prop.PR_STRING)

	b.Append(c)
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	a.Append(b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_Lexeme_Remove(t *testing.T) {

	a, b, c, _ := setupList()
	a.Remove()
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	a, b, c, _ = setupList()
	b.Remove()
	fullEqual(t, a, nil, c, a)
	fullEqual(t, c, a, nil, c)

	a, b, c, _ = setupList()
	c.Remove()
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)
}
