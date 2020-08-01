package sanitiser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme/lextest"
)

func doTest(t *testing.T, in, exp *lexeme.Container) {
	SanitiseAll(in)
	lextest.Equal(t, exp.Head(), in.Head())
}

func Test1_1(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	exp := lextest.Feign()

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign()

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("", lexeme.UNDEFINED),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.UNDEFINED),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test1_4(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("(", lexeme.LEFT_PAREN),
	)

	doTest(t, in, exp)
}

func Test1_5(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
	)

	doTest(t, in, exp)
}

func Test1_6(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
	)

	exp := lextest.Feign(
		lextest.Tok(")", lexeme.RIGHT_PAREN),
	)

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// @Println(1,1)
	exp := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}
