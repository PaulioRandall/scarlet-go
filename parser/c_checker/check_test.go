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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok(")", lexeme.R_PAREN),
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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(")", lexeme.R_PAREN),
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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok(")", lexeme.R_PAREN),
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
		lextest.Tok(")", lexeme.R_PAREN),
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
		lextest.Tok("(", lexeme.L_PAREN),
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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(")", lexeme.R_PAREN),
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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(")", lexeme.R_PAREN),
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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok("y", lexeme.IDENT),
		lextest.Tok(")", lexeme.R_PAREN),
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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok(")", lexeme.R_PAREN),
	)

	doErrorTest(t, in)
}

func Test3_1(t *testing.T) {

	// WHEN checking a simple assignment
	// THEN no errors should be returned
	// x:=1
	in := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
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
		lextest.Tok("a", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("b", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("c", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("d", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
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
		lextest.Tok("a", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("b", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test4_1(t *testing.T) {

	// WHEN checking a simple expression
	// THEN no errors should be returned
	// 1 + 2
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test4_2(t *testing.T) {

	// WHEN checking a complex expression
	// THEN no errors should be returned
	// 1 + 2 - 3 * 4 / 5 % 6
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("-", lexeme.SUB),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("4", lexeme.NUMBER),
		lextest.Tok("/", lexeme.DIV),
		lextest.Tok("5", lexeme.NUMBER),
		lextest.Tok("%", lexeme.REM),
		lextest.Tok("6", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test4_3(t *testing.T) {

	// WHEN checking an expression with duplicate operators
	// THEN an error should be returned
	// 1 + + 2
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test4_4(t *testing.T) {

	// WHEN checking an expression with duplicate operands
	// THEN an error should be returned
	// 1 1 + 2
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test4_5(t *testing.T) {

	// WHEN checking an expression with trailing operator
	// THEN an error should be returned
	// 1 + 2 -
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("-", lexeme.SUB),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doErrorTest(t, in)
}

func Test4_6(t *testing.T) {

	// WHEN checking a spell with a single simple expression argument
	// THEN no errors should be returned
	// @Println(1 + 2)
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test4_7(t *testing.T) {

	// WHEN checking an assignment with a single simple expression argument
	// THEN no errors should be returned
	// x := 1 + 2
	in := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}

func Test4_8(t *testing.T) {

	// WHEN checking a complex logical and comparison expression
	// THEN no errors should be returned
	// false || false && true || 1 < 2 && 3 >= 3 && 4 != 5
	in := lextest.Feign(
		lextest.Tok("false", lexeme.BOOL),
		lextest.Tok("||", lexeme.OR),
		lextest.Tok("false", lexeme.BOOL),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("||", lexeme.OR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("<", lexeme.LESS),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(">=", lexeme.MORE_EQUAL),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("4", lexeme.NUMBER),
		lextest.Tok("!=", lexeme.NOT_EQUAL),
		lextest.Tok("5", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in)
}
