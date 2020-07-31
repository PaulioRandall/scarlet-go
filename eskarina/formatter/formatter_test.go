package formatter

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"
)

func doTest(t *testing.T, f func(*lexeme.Iterator), given, exp *lexeme.Container) {
	itr := given.Iterator()
	f(itr)
	lextest.Equal(t, exp.Head(), given.Head())
}

func Test1_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	exp := lextest.Feign()

	doTest(t, trimWhiteSpace, given, exp)
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

	doTest(t, trimWhiteSpace, given, exp)
}

func Test2_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
	)

	doTest(t, insertSeparatorSpaces, given, exp)
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

	doTest(t, insertSeparatorSpaces, given, exp)
}

func Test2_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	doTest(t, insertSeparatorSpaces, given, exp)
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

	doTest(t, stripUselessLines, given, exp)
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

	doTest(t, stripUselessLines, given, exp)
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

	doTest(t, stripUselessLines, given, exp)
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

	doTest(t, unifyLineEndings, given, exp)
}

func Test6_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\r\n", lexeme.NEWLINE),
		lextest.Tok("123", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok("\r\n", lexeme.NEWLINE),
		lextest.Tok("123", lexeme.NUMBER),
		lextest.Tok("\r\n", lexeme.NEWLINE),
	)

	doTest(t, unifyLineEndings, given, exp)
}

func Test7_1(t *testing.T) {

	// (
	//
	// (
	// 1)
	//
	// )
	given := lextest.Feign(
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
	)

	// (
	//
	//   (
	//     1)
	//
	// )
	exp := lextest.Feign(
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "\t", lexeme.WHITESPACE),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "\t\t", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
	)

	doTest(t, indentLines, given, exp)
}

func Test8_1(t *testing.T) {

	// @Println(
	// 	1,
	// 	1,
	// )
	given := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\t", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\t", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Lex(0, 0, "@Println", lexeme.SPELL),
		lextest.Lex(0, 8, "(", lexeme.LEFT_PAREN),
		lextest.Lex(0, 9, "\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "\t", lexeme.WHITESPACE),
		lextest.Lex(1, 1, "1", lexeme.NUMBER),
		lextest.Lex(1, 2, ",", lexeme.SEPARATOR),
		lextest.Lex(1, 3, "\n", lexeme.NEWLINE),
		lextest.Lex(2, 0, "\t", lexeme.WHITESPACE),
		lextest.Lex(2, 1, "1", lexeme.NUMBER),
		lextest.Lex(2, 2, ",", lexeme.SEPARATOR),
		lextest.Lex(2, 3, "\n", lexeme.NEWLINE),
		lextest.Lex(3, 0, ")", lexeme.RIGHT_PAREN),
		lextest.Lex(3, 1, "\n", lexeme.NEWLINE),
	)

	doTest(t, updatePositions, given, exp)
}

func Test10_1(t *testing.T) {

	// " @Println ( \n1 , 1 , \n 1 ) \n "
	given := lextest.Feign(
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("(", lexeme.LEFT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n\r", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.WHITESPACE),
		lextest.Tok(")", lexeme.RIGHT_PAREN),
		lextest.Tok(" ", lexeme.WHITESPACE),
	)

	// @Println(
	//   1, 1,
	//
	//   1)
	exp := lextest.Feign(
		lextest.Lex(0, 0, "@Println", lexeme.SPELL),
		lextest.Lex(0, 8, "(", lexeme.LEFT_PAREN),
		lextest.Lex(0, 9, "\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "\t", lexeme.WHITESPACE),
		lextest.Lex(1, 1, "1", lexeme.NUMBER),
		lextest.Lex(1, 2, ",", lexeme.SEPARATOR),
		lextest.Lex(1, 3, " ", lexeme.WHITESPACE),
		lextest.Lex(1, 4, "1", lexeme.NUMBER),
		lextest.Lex(1, 5, ",", lexeme.SEPARATOR),
		lextest.Lex(1, 6, "\n", lexeme.NEWLINE),
		lextest.Lex(2, 0, "\n", lexeme.NEWLINE),
		lextest.Lex(3, 0, "\t", lexeme.WHITESPACE),
		lextest.Lex(3, 1, "1", lexeme.NUMBER),
		lextest.Lex(3, 2, ")", lexeme.RIGHT_PAREN),
		lextest.Lex(3, 3, "\n", lexeme.NEWLINE),
	)

	format(given.Iterator())
	lextest.Equal(t, exp.Head(), given.Head())
}
