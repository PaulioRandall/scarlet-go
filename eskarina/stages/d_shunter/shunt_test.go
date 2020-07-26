package shunter

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"
)

func doTest(t *testing.T, in, exp *lexeme.Lexeme) {
	act := ShuntAll(in)
	lextest.Equal(t, exp, act)
}

func Test1_1(t *testing.T) {

	// WHEN refixing a spell with no arguments
	// @Println()
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok("(", lexeme.PR_PARENTHESIS, lexeme.PR_OPENER),
		lextest.Tok(")", lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.PR_CALLABLE),
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("@Print", lexeme.PR_SPELL),
		lextest.Tok("(", lexeme.PR_PARENTHESIS, lexeme.PR_OPENER),
		lextest.Tok("x", lexeme.PR_TERM),
		lextest.Tok(")", lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.PR_CALLABLE),
		lextest.Tok("x", lexeme.PR_TERM),
		lextest.Tok("@Print", lexeme.PR_SPELL),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok("(", lexeme.PR_PARENTHESIS, lexeme.PR_OPENER),
		lextest.Tok("x", lexeme.PR_TERM),
		lextest.Tok(",", lexeme.PR_SEPARATOR),
		lextest.Tok("y", lexeme.PR_TERM),
		lextest.Tok(",", lexeme.PR_SEPARATOR),
		lextest.Tok("z", lexeme.PR_TERM),
		lextest.Tok(")", lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.PR_CALLABLE),
		lextest.Tok("x", lexeme.PR_TERM),
		lextest.Tok("y", lexeme.PR_TERM),
		lextest.Tok("z", lexeme.PR_TERM),
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	doTest(t, in, exp)
}
