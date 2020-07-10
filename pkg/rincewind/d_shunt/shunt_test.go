package shunt

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in, exps []token.Token) {

	acts, e := ShuntAll(in)
	if e != nil {
		require.Nil(t, fmt.Sprintf("%+v", e))
	}

	testutils.RequireTokenSlice(t, exps, acts)
}

func doErrorTest(t *testing.T, in []token.Token) {
	_, e := ShuntAll(in)
	require.NotNil(t, e, "Expected an error")
}

func tok(gen GenType, sub SubType, raw string) token.Tok {
	return token.Tok{
		Gen:    gen,
		Sub:    sub,
		RawStr: raw,
		ColEnd: len(raw),
	}
}

func Test1_1(t *testing.T) {

	// WHEN refixing a spell with no arguments
	// @Println()
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND magic token indicating parameter start inserted before spell
	exp := []token.Token{
		tok(GEN_PARAMS, SUB_UNDEFINED, "("),
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND the argument is placed before the spell
	// AND magic token indicating parameter start inserted before the argument
	exp := []token.Token{
		tok(GEN_PARAMS, SUB_UNDEFINED, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "y"),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "z"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND the delimiters are removed
	// AND the arguments are placed before the spell
	// AND magic token indicating parameter start inserted before the arguments
	exp := []token.Token{
		tok(GEN_PARAMS, SUB_UNDEFINED, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "y"),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "z"),
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}
