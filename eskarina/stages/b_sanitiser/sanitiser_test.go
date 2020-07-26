package sanitiser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"
)

func doTest(t *testing.T, in, exp *lexeme.Lexeme) {
	act := SanitiseAll(in)
	lextest.Equal(t, exp, act)
}

func Test1_1(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(" ", lexeme.PR_REDUNDANT),
	)

	exp := (*lexeme.Lexeme)(nil)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	exp := (*lexeme.Lexeme)(nil)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(""),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	exp := lextest.Feign(
		lextest.Tok(""),
		lextest.Tok("\n", lexeme.PR_TERMINATOR),
	)

	doTest(t, in, exp)
}

func Test1_4(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("(", lexeme.PR_PARENTHESIS, lexeme.PR_OPENER),
		lextest.Tok("\n", lexeme.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("(", lexeme.PR_PARENTHESIS, lexeme.PR_OPENER),
	)

	doTest(t, in, exp)
}

func Test1_5(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(",", lexeme.PR_SEPARATOR),
		lextest.Tok("\n", lexeme.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.PR_SEPARATOR),
	)

	doTest(t, in, exp)
}

func Test1_6(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(",", lexeme.PR_SEPARATOR),
		lextest.Tok(")", lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER),
	)

	exp := lextest.Feign(
		lextest.Tok(")", lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER),
	)

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("(", lexeme.PR_PARENTHESIS, lexeme.PR_OPENER),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("1", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok(",", lexeme.PR_DELIMITER, lexeme.PR_SEPARATOR),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("1", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok(",", lexeme.PR_DELIMITER, lexeme.PR_SEPARATOR),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok(")", lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER),
		lextest.Tok(" ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
	)

	// @Println(1,1)
	exp := lextest.Feign(
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok("(", lexeme.PR_PARENTHESIS, lexeme.PR_OPENER),
		lextest.Tok("1", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Tok(",", lexeme.PR_DELIMITER, lexeme.PR_SEPARATOR),
		lextest.Tok("1", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Tok(")", lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER),
		lextest.Tok("\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
	)

	doTest(t, in, exp)
}
