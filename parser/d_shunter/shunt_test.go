package shunter

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme/lextest"
)

func doTest(t *testing.T, in, exp *lexeme.Container) {
	act := ShuntAll(in)
	lextest.Equal(t, exp.Head(), act.Head())
}

func Test1_1(t *testing.T) {

	// WHEN refixing a spell with no arguments
	// @Println()
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.CALLABLE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.CALLABLE),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.CALLABLE),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN refixing an assignment
	// x := 1
	in := lextest.Feign(
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.CALLABLE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test2_2(t *testing.T) {

	// WHEN refixing a multi assignment
	// a, b, c := 1, 2, 3
	in := lextest.Feign(
		lextest.Tok("a", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("b", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("c", lexeme.IDENTIFIER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.CALLABLE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("c", lexeme.IDENTIFIER),
		lextest.Tok("b", lexeme.IDENTIFIER),
		lextest.Tok("a", lexeme.IDENTIFIER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}
