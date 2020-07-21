package scanner

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"

	"github.com/stretchr/testify/require"
)

func feign(lexs ...*lexeme.Lexeme) *lexeme.Lexeme {

	var lex *lexeme.Lexeme

	for _, l := range lexs {

		if lex == nil {
			lex = l
			continue
		}

		lex.Next = l
		l.Prev = lex
	}

	return lex
}

func lex(line, col int, raw string, props ...prop.Prop) *lexeme.Lexeme {
	return &lexeme.Lexeme{
		Props: props,
		Raw:   raw,
		Line:  line,
		Col:   col,
	}
}

func requireSlice(t *testing.T, exp, act *lexeme.Lexeme) {

	for a, b := exp, act; a != nil && b != nil; a, b = a.Next, b.Next {
		exp, act = a, b
		require.True(t, a != nil, "Want: %s\nHave: nil", a.String())
		require.True(t, b != nil, "Want: EOF\nHave: %s", b.String())
		requireLexeme(t, *a, *b)
	}

	for a, b := exp, act; a != nil && b != nil; a, b = a.Prev, b.Prev {
		requireLexeme(t, *a, *b)
	}
}

func requireLexeme(t *testing.T, exp, act lexeme.Lexeme) {

	msg := fmt.Sprintf("Expected: %s\nActual:  %s", exp.String(), act.String())

	require.Equal(t, exp.Props, act.Props, msg)
	require.Equal(t, exp.Raw, act.Raw, msg)
	require.Equal(t, exp.Line, act.Line, msg)
	require.Equal(t, exp.Col, act.Col, msg)
}

func Test1_1(t *testing.T) {

	exp := feign(
		lex(0, 0, "\n", prop.PR_REDUNDANT, prop.PR_NEWLINE),
	)

	act, e := ScanAll("\n")
	require.Nil(t, e, "%+v", e)

	requireSlice(t, exp, act)
}
