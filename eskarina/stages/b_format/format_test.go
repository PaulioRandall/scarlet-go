package format

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"
)

func Test1_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	exp := lextest.Feign()

	act := trimWhiteSpace(given)
	lextest.Equal(t, exp.Head(), act.Head())
}

func Test1_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	act := trimWhiteSpace(given)
	lextest.Equal(t, exp.Head(), act.Head())
}

/*
func Test2_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_4(t *testing.T) {

	given := lextest.Feign(
		lextest.Lex2(2, 4, ",", lexeme.SEPARATOR),
		lextest.Lex2(2, 5, "1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Lex2(2, 4, ",", lexeme.SEPARATOR),
		lextest.Lex2(2, 5, " ", lexeme.WHITESPACE),
		lextest.Lex2(2, 5, "1", lexeme.NUMBER),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("(", lexeme.LEFT_PAREN),
	)

	exp := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_4(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_5(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
	)

	exp := lextest.Feign(
		lextest.Tok(")", lexeme.RIGHT_PAREN),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test4_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	act := reduceSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test4_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("\t", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	act := reduceSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test5_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	act := reduceEmptyLines(given)
	lextest.Equal(t, exp, act)
}

func Test5_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	act := reduceEmptyLines(given)
	lextest.Equal(t, exp, act)
}

func Test6_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\r\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	act := unifyLineEndings(given, "\n")
	lextest.Equal(t, exp, act)
}

func Test7_1(t *testing.T) {

	// " @Println ( 1 , 1 , \n 1 ) \n "
	given := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	// "@Println(1, 1,\n1)\n"
	exp := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Lex2(1, 0, "\t", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	act := FormatAll(given, "\n")
	lextest.Equal(t, exp, act)
}
*/
