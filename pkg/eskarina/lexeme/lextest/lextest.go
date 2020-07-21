package lextest

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"

	"github.com/stretchr/testify/require"
)

func Feign(lexs ...*lexeme.Lexeme) *lexeme.Lexeme {

	var first *lexeme.Lexeme
	var last *lexeme.Lexeme

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

func Lex(line, col int, raw string, props ...prop.Prop) *lexeme.Lexeme {
	return &lexeme.Lexeme{
		Props: props,
		Raw:   raw,
		Line:  line,
		Col:   col,
	}
}

func Tok(raw string, props ...prop.Prop) *lexeme.Lexeme {
	return &lexeme.Lexeme{
		Props: props,
		Raw:   raw,
	}
}

func Equal(t *testing.T, exp, act *lexeme.Lexeme) {

	for a, b := exp, act; a != nil && b != nil; a, b = a.Next, b.Next {
		exp, act = a, b

		msg := fmt.Sprintf(
			"Unexepected Lexeme.Next\nExpected: %s\nActual:  %s",
			exp.String(), act.String(),
		)

		require.True(t, a != nil, "Want: %s\nHave: nil", a.String())
		require.True(t, b != nil, "Want: EOF\nHave: %s", b.String())
		equalContent(t, *a, *b, msg)
	}

	for a, b := exp, act; a != nil && b != nil; a, b = a.Prev, b.Prev {
		msg := fmt.Sprintf(
			"Unexepected Lexeme.Prev\nExpected: %s\nActual:  %s",
			exp.String(), act.String(),
		)
		equalContent(t, *a, *b, msg)
	}
}

func equalContent(t *testing.T, exp, act lexeme.Lexeme, msg string) {
	require.Equal(t, exp.Props, act.Props, msg)
	require.Equal(t, exp.Raw, act.Raw, msg)
	require.Equal(t, exp.Line, act.Line, msg)
	require.Equal(t, exp.Col, act.Col, msg)
}
