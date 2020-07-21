package sanitiser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme/lextest"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
	//"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in, exp *lexeme.Lexeme) {
	act := SanitiseAll(in)
	lextest.Equal(t, exp, act)
}

func Test1_1(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(" ", prop.PR_REDUNDANT),
	)

	exp := (*lexeme.Lexeme)(nil)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	exp := (*lexeme.Lexeme)(nil)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(""),
		lextest.Tok("\n", prop.PR_TERMINATOR),
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	exp := lextest.Feign(
		lextest.Tok(""),
		lextest.Tok("\n", prop.PR_TERMINATOR),
	)

	doTest(t, in, exp)
}

func Test1_4(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("(", prop.PR_PARENTHESIS, prop.PR_OPENER),
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("(", prop.PR_PARENTHESIS, prop.PR_OPENER),
	)

	doTest(t, in, exp)
}

func Test1_5(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
	)

	doTest(t, in, exp)
}

func Test1_6(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(")", prop.PR_PARENTHESIS, prop.PR_CLOSER),
	)

	exp := lextest.Feign(
		lextest.Tok(")", prop.PR_PARENTHESIS, prop.PR_CLOSER),
	)

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("(", prop.PR_PARENTHESIS, prop.PR_OPENER),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok(",", prop.PR_DELIMITER, prop.PR_SEPARATOR),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok(",", prop.PR_DELIMITER, prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok(")", prop.PR_PARENTHESIS, prop.PR_CLOSER),
		lextest.Tok(" ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
	)

	// @Println(1,1)
	exp := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_PARENTHESIS, prop.PR_OPENER),
		lextest.Tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER),
		lextest.Tok(",", prop.PR_DELIMITER, prop.PR_SEPARATOR),
		lextest.Tok("1", prop.PR_TERM, prop.PR_LITERAL, prop.PR_NUMBER),
		lextest.Tok(")", prop.PR_PARENTHESIS, prop.PR_CLOSER),
		lextest.Tok("\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
	)

	doTest(t, in, exp)
}
