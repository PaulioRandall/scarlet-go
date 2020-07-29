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

func Test2_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
	)

	act := insertWhiteSpace(given)
	lextest.Equal(t, exp.Head(), act.Head())
}

func Test2_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	act := insertWhiteSpace(given)
	lextest.Equal(t, exp.Head(), act.Head())
}

func Test2_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Lex(2, 4, ",", lexeme.SEPARATOR),
		lextest.Lex(2, 5, "1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Lex(2, 4, ",", lexeme.SEPARATOR),
		lextest.Lex(2, 5, " ", lexeme.WHITESPACE),
		lextest.Lex(2, 5, "1", lexeme.NUMBER),
	)

	act := insertWhiteSpace(given)
	lextest.Equal(t, exp.Head(), act.Head())
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

	act := stripUselessLines(given)
	lextest.Equal(t, exp.Head(), act.Head())
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

	act := stripUselessLines(given)
	lextest.Equal(t, exp.Head(), act.Head())
}

func Test5_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	act := stripUselessLines(given)
	lextest.Equal(t, exp.Head(), act.Head())
}

func Test6_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\r\n", lexeme.NEWLINE),
		lextest.Tok("\r\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	act := unifyLineEndings(given)
	lextest.Equal(t, exp.Head(), act.Head())
}

func Test6_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Lex(2, 3, "\r\n", lexeme.NEWLINE),
		lextest.Lex(2, 5, "123", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Lex(2, 3, "\r\n", lexeme.NEWLINE),
		lextest.Lex(2, 5, "123", lexeme.NUMBER),
		lextest.Lex(2, 8, "\r\n", lexeme.NEWLINE),
	)

	act := unifyLineEndings(given)
	lextest.Equal(t, exp.Head(), act.Head())
}

/*
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
		lextest.Lex(1, 0, "\t", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Lex(0, 1, "\n", lexeme.NEWLINE),
	)

	act := FormatAll(given, "\n")
	lextest.Equal(t, exp.Head(), act.Head())
}
*/
