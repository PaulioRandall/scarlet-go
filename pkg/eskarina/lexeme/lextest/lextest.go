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

	for exp != nil || act != nil {

		if exp == nil && act != nil {
			require.Nil(t, act, "Want: EOF\nHave: %s", act.String())
		}

		if exp != nil && act == nil {
			require.NotNil(t, act, "Want: %s\nHave: nil", exp.String())
		}

		equalContent(t, exp, act, fmt.Sprintf(
			"Unexepected Lexeme\nWant: %s\nHave: %s",
			exp.String(), act.String(),
		))

		equalContent(t, exp.Prev, act.Prev, fmt.Sprintf(
			"Unexepected Lexeme.Prev\nWant: %s\nHave: %s",
			exp.String(), act.String(),
		))

		exp, act = exp.Next, act.Next
	}
}

func equalContent(t *testing.T, exp, act *lexeme.Lexeme, msg string) {

	if exp == nil {
		require.Nil(t, act, msg)
		return
	}

	require.NotNil(t, act, msg)
	require.Equal(t, exp.Props, act.Props, msg)
	require.Equal(t, exp.Raw, act.Raw, msg)
	require.Equal(t, exp.Line, act.Line, msg)
	require.Equal(t, exp.Col, act.Col, msg)
}
