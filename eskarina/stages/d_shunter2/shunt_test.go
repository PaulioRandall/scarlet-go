package shunter

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"
)

func doTest(t *testing.T, in, exp *lexeme.Container2) {
	act := ShuntAll(in)
	lextest.Equal2(t, exp.Head(), act.Head())
}

func Test1_1(t *testing.T) {

	// WHEN refixing a spell with no arguments
	// @Println()
	in := lextest.Feign2(
		lextest.Tok2("@Println", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign2(
		lextest.Tok2("", lexeme.CALLABLE),
		lextest.Tok2("@Println", lexeme.SPELL),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := lextest.Feign2(
		lextest.Tok2("@Println", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign2(
		lextest.Tok2("", lexeme.CALLABLE),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2("@Println", lexeme.SPELL),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := lextest.Feign2(
		lextest.Tok2("@Println", lexeme.SPELL),
		lextest.Tok2("(", lexeme.LEFT_PAREN),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2(",", lexeme.SEPARATOR),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2(")", lexeme.RIGHT_PAREN),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign2(
		lextest.Tok2("", lexeme.CALLABLE),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2("x", lexeme.IDENTIFIER),
		lextest.Tok2("@Println", lexeme.SPELL),
		lextest.Tok2("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}
