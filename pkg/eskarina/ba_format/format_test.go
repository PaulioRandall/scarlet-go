package format

import (
	"testing"

	//"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme/lextest"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func Test1_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	exp := lextest.Feign()

	act := trimLeadingSpace(given)
	lextest.Equal(t, exp, act)
}

func Test1_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	act := trimLeadingSpace(given)
	lextest.Equal(t, exp, act)
}

func Test2_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
	)

	exp := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	exp := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_4(t *testing.T) {

	given := lextest.Feign(
		lextest.Lex(2, 4, ",", prop.PR_SEPARATOR),
		lextest.Lex(2, 5, "1", prop.PR_LITERAL),
	)

	exp := lextest.Feign(
		lextest.Lex(2, 4, ",", prop.PR_SEPARATOR),
		lextest.Lex(2, 5, " ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Lex(2, 5, "1", prop.PR_LITERAL),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("(", prop.PR_OPENER),
	)

	exp := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_OPENER),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_4(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_OPENER),
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_OPENER),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_5(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(")", prop.PR_CLOSER),
	)

	exp := lextest.Feign(
		lextest.Tok(")", prop.PR_CLOSER),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test4_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("   ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	exp := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	act := reduceSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test4_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("\t", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	exp := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	act := reduceSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test5_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	exp := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	act := reduceEmptyLines(given)
	lextest.Equal(t, exp, act)
}

func Test5_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	exp := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	act := reduceEmptyLines(given)
	lextest.Equal(t, exp, act)
}

func Test6_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\r\n", prop.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	act := unifyLineEndings(given, "\n")
	lextest.Equal(t, exp, act)
}

func TestAll(t *testing.T) {

	// " @Println ( 1 , 1 , \n 1 ) \n "
	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("(", prop.PR_OPENER),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(")", prop.PR_CLOSER),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	// "@Println(1, 1,\n1)\n"
	exp := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_OPENER),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Lex(1, 0, "\t", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(")", prop.PR_CLOSER),
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	act := FormatAll(given, "\n")
	lextest.Equal(t, exp, act)
}
