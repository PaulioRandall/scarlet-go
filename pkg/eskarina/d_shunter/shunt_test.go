package shunter

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme/lextest"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
	//"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in, exp *lexeme.Lexeme) {
	act := ShuntAll(in)
	lextest.Equal(t, exp, act)
}

func Test1_1(t *testing.T) {

	// WHEN refixing a spell with no arguments
	// @Println()
	in := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_PARENTHESIS, prop.PR_OPENER),
		lextest.Tok(")", prop.PR_PARENTHESIS, prop.PR_CLOSER),
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	exp := lextest.Feign(
		lextest.Tok("", prop.PR_CALLABLE),
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("@Print", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_PARENTHESIS, prop.PR_OPENER),
		lextest.Tok("x", prop.PR_TERM),
		lextest.Tok(")", prop.PR_PARENTHESIS, prop.PR_CLOSER),
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	exp := lextest.Feign(
		lextest.Tok("", prop.PR_CALLABLE),
		lextest.Tok("x", prop.PR_TERM),
		lextest.Tok("@Print", prop.PR_SPELL),
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_PARENTHESIS, prop.PR_OPENER),
		lextest.Tok("x", prop.PR_TERM),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("y", prop.PR_TERM),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("z", prop.PR_TERM),
		lextest.Tok(")", prop.PR_PARENTHESIS, prop.PR_CLOSER),
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	exp := lextest.Feign(
		lextest.Tok("", prop.PR_CALLABLE),
		lextest.Tok("x", prop.PR_TERM),
		lextest.Tok("y", prop.PR_TERM),
		lextest.Tok("z", prop.PR_TERM),
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	doTest(t, in, exp)
}
