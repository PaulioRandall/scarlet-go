package checker

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in *lexeme.Lexeme) {
	e := CheckAll(in)
	require.Nil(t, e, "unexpected error: %+v", e)
}

func doErrorTest(t *testing.T, in *lexeme.Lexeme) {
	e := CheckAll(in)
	require.NotNil(t, e, "Expected error")
}

func Test1_1(t *testing.T) {

	// WHEN checking a spell with no arguments
	// THEN no errors should be returned
	// @Println()
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test1_2(t *testing.T) {

	// WHEN checking a spell with one argument
	// THEN no errors should be returned
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test1_3(t *testing.T) {

	// WHEN checking a spell with multiple arguments
	// THEN no errors should be returned
	// @Println(x, 1, true)
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2("1", lexeme.NUMBER),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2("true", lexeme.BOOL),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test2_1(t *testing.T) {

	// WHEN checking a spell with missing opening parenthesis
	// THEN an error should be returned
	// @Println)
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_2(t *testing.T) {

	// WHEN checking a spell with missing closing parenthesis
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_3(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_4(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_5(t *testing.T) {

	// WHEN checking a spell with a missing value delimiter
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2("y", lexeme.IDENTIFIER),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_6(t *testing.T) {

	// WHEN checking a spell with a missing final terminator
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok2("@Print", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
	)

	doErrorTest(t, in)
}
