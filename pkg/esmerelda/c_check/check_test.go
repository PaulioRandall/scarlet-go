package check

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in []token.Token) {

	acts, e := CheckAll(in)
	if e != nil {
		require.Nil(t, fmt.Sprintf("%+v", e))
	}

	testutils.RequireTokenSlice(t, in, acts)
}

func doErrorTest(t *testing.T, in []token.Token) {
	_, e := CheckAll(in)
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

	// WHEN checking a spell with no arguments
	// THEN no errors should be returned
	// @Println()
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test1_2(t *testing.T) {

	// WHEN checking a spell with one argument
	// THEN no errors should be returned
	// @Println(x)
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test1_3(t *testing.T) {

	// WHEN checking a spell with multiple arguments
	// THEN no errors should be returned
	// @Println(x, 1, true)
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_LITERAL, SUB_NUMBER, "1"),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_LITERAL, SUB_BOOL, "true"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test2_1(t *testing.T) {

	// WHEN checking a spell with missing opening parenthesis
	// THEN an error should be returned
	// @Println)
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_2(t *testing.T) {

	// WHEN checking a spell with missing closing parenthesis
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_3(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_4(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_5(t *testing.T) {

	// WHEN checking a spell with a missing value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "y"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_6(t *testing.T) {

	// WHEN checking a spell with a missing final terminator
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	doErrorTest(t, in)
}
