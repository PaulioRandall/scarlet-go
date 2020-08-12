package checker

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme/lextest"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in *lexeme.Container) {
	e := CheckAll(in)
	require.Nil(t, e, "unexpected error: %+v", e)
}

func doErrorTest(t *testing.T, in *lexeme.Container) {
	e := CheckAll(in)
	require.NotNil(t, e, "Expected error")
}

func Test1_1(t *testing.T) {

	// WHEN checking a spell with no arguments
	// THEN no errors should be returned
	// @Println()
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test1_2(t *testing.T) {

	// WHEN checking a spell with one argument
	// THEN no errors should be returned
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test1_3(t *testing.T) {

	// WHEN checking a spell with multiple arguments
	// THEN no errors should be returned
	// @Println(x, 1, true)
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test2_1(t *testing.T) {

	// WHEN checking a spell with missing opening parenthesis
	// THEN an error should be returned
	// @Println)
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_2(t *testing.T) {

	// WHEN checking a spell with missing closing parenthesis
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_3(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_4(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_5(t *testing.T) {

	// WHEN checking a spell with a missing value delimiter
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("y", lexeme.IDENTIFIER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test2_6(t *testing.T) {

	// WHEN checking a spell with a missing final terminator
	// THEN an error should be returned
	// @Println(
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
	)

	doErrorTest(t, in)
}

func Test3_1(t *testing.T) {

	// WHEN checking a simple assignment
	// THEN no errors should be returned
	// x:=1
	in := lextest.Feign(
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test3_2(t *testing.T) {

	// WHEN checking a multi assignment
	// THEN no errors should be returned
	// a,b,c,d:=x,true,1,"abc"
	in := lextest.Feign(
		lextest.Tok("a", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("b", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("c", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("d", lexeme.IDENTIFIER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(`"abc"`, lexeme.STRING),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test3_3(t *testing.T) {

	// WHEN checking a multi assignment with a missing separator
	// THEN an error should be returned
	// a,b:=true1
	in := lextest.Feign(
		lextest.Tok("a", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("b", lexeme.IDENTIFIER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}
