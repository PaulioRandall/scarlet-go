package sanitiser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"
)

func doTest(t *testing.T, in, exp *lexeme.Container2) {
	act := SanitiseAll(in)
	lextest.Equal2(t, exp.Head(), act.Head())
}

func Test1_1(t *testing.T) {

	in := lextest.Feign2(
		lextest.Tok2(" ", lexeme.WHITESPACE),
	)

	exp := lextest.Feign2()

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	in := lextest.Feign2(
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign2()

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	in := lextest.Feign2(
		lextest.Tok2("", lexeme.UNDEFINED),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign2(
		lextest.Tok2("", lexeme.UNDEFINED),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test1_4(t *testing.T) {

	in := lextest.Feign2(
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign2(
		lextest.Tok2("(", lexeme.LEFT_PAREN),
	)

	doTest(t, in, exp)
}

func Test1_5(t *testing.T) {

	in := lextest.Feign2(
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign2(
		lextest.Tok2(",", lexeme.SEPARATOR),
	)

	doTest(t, in, exp)
}

func Test1_6(t *testing.T) {

	in := lextest.Feign2(
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
	)

	exp := lextest.Feign2(
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
	)

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	in := lextest.Feign2(
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("@Println", lexeme.SPELL),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("1", lexeme.NUMBER),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("1", lexeme.NUMBER),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2("\n", lexeme.NEWLINE),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2(" ", lexeme.WHITESPACE),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	// @Println(1,1)
	exp := lextest.Feign2(
		lextest.Tok2("@Println", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("1", lexeme.NUMBER),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2("1", lexeme.NUMBER),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}
