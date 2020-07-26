package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {

	lex := &Lexeme{}

	_ = PropToken(lex)
	_ = Snippet(lex)
	_ = Node(lex)
	var _ string = lex.String()
}

func Test_Lexeme_Is(t *testing.T) {

	lex := tok("1", PR_TERM, PR_LITERAL, PR_NUMBER)

	require.True(t, lex.Is(PR_TERM))
	require.True(t, lex.Is(PR_LITERAL))
	require.True(t, lex.Is(PR_NUMBER))

	require.False(t, lex.Is(PR_IDENTIFIER))
}

func Test_Lexeme_Has(t *testing.T) {

	lex := tok("1", PR_TERM, PR_LITERAL, PR_NUMBER)

	require.True(t, lex.Has(PR_TERM))
	require.True(t, lex.Has(PR_TERM, PR_LITERAL))
	require.True(t, lex.Has(PR_TERM, PR_LITERAL, PR_NUMBER))
	require.True(t, lex.Has())

	require.False(t, lex.Has(PR_IDENTIFIER))
	require.False(t, lex.Has(PR_TERM, PR_LITERAL, PR_BOOL))
}

func Test_Lexeme_Any(t *testing.T) {

	lex := tok("1", PR_TERM, PR_LITERAL, PR_NUMBER)

	require.True(t, lex.Any(PR_TERM))
	require.True(t, lex.Any(PR_LITERAL, PR_IDENTIFIER))
	require.True(t, lex.Any(PR_SPELL, PR_OPENER, PR_NUMBER))

	require.False(t, lex.Any())
	require.False(t, lex.Any(PR_IDENTIFIER))
	require.False(t, lex.Any(PR_SPELL, PR_OPENER, PR_CLOSER))
}

func Test_Lexeme_ShiftUp(t *testing.T) {

	a := tok("true", PR_BOOL)
	b := tok("1", PR_NUMBER)
	c := tok(`"abc"`, PR_STRING)

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

	a := tok("true", PR_BOOL)
	b := tok("1", PR_NUMBER)
	c := tok(`"abc"`, PR_STRING)

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

	a := tok("true", PR_BOOL)
	b := tok("1", PR_NUMBER)
	c := tok(`"abc"`, PR_STRING)

	b.Prepend(a)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)

	c.Prepend(b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_Lexeme_Append(t *testing.T) {

	a := tok("true", PR_BOOL)
	b := tok("1", PR_NUMBER)
	c := tok(`"abc"`, PR_STRING)

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
