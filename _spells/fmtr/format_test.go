package fmtr

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/lexeme/lextest"
)

func doTest(t *testing.T, f func(*lexeme.Iterator), given, exp *lexeme.Container) {
	f(given.Iterator())
	lextest.Equal(t, exp.Head(), given.Head())
}

func Test1_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", lexeme.SPACE),
	)

	exp := lextest.Feign()

	doTest(t, trimSpaces, given, exp)
}

func Test1_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.SPACE),
	)

	exp := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, trimSpaces, given, exp)
}

func Test2_1(t *testing.T) {

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

func Test2_2(t *testing.T) {

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

func Test2_3(t *testing.T) {

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

func Test3_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_4(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("}", lexeme.R_CURLY),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok("}", lexeme.R_CURLY),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_5(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("]", lexeme.R_SQUARE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	exp := lextest.Feign(
		lextest.Tok("]", lexeme.R_SQUARE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_6(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("{", lexeme.L_CURLY),
	)

	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("{", lexeme.L_CURLY),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_7(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("# Comment", lexeme.COMMENT),
	)

	exp := lextest.Feign(
		lextest.Tok("# Comment", lexeme.COMMENT),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_8(t *testing.T) {

	given := lextest.Feign(
		lextest.Lex(0, 0, "1", lexeme.NUMBER),
		lextest.Lex(0, 1, "# Comment", lexeme.COMMENT),
	)

	exp := lextest.Feign(
		lextest.Lex(0, 0, "1", lexeme.NUMBER),
		lextest.Lex(0, 0, " ", lexeme.SPACE),
		lextest.Lex(0, 1, "# Comment", lexeme.COMMENT),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_9(t *testing.T) {

	// x:=1
	given := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
	)

	// x := 1 + 2
	exp := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_10(t *testing.T) {

	// x :=1
	given := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
	)

	// x := 1
	exp := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_11(t *testing.T) {

	// 1+2*3
	given := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("3", lexeme.NUMBER),
	)

	// 1 + 2 * 3
	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("3", lexeme.NUMBER),
	)

	doTest(t, insertSpaces, given, exp)
}

func Test3_12(t *testing.T) {

	// 1+ 2 *3
	given := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("3", lexeme.NUMBER),
	)

	// 1 + 2 * 3
	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("3", lexeme.NUMBER),
	)

	doTest(t, insertSpaces, given, exp)
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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(")", lexeme.R_PAREN),
	)

	// (
	//
	//   (
	//     1)
	//
	// )
	exp := lextest.Feign(
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "\t", lexeme.SPACE),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "\t\t", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(")", lexeme.R_PAREN),
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
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\t", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\t", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Lex(0, 0, "@Println", lexeme.SPELL),
		lextest.Lex(0, 8, "(", lexeme.L_PAREN),
		lextest.Lex(0, 9, "\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "\t", lexeme.SPACE),
		lextest.Lex(1, 1, "1", lexeme.NUMBER),
		lextest.Lex(1, 2, ",", lexeme.DELIM),
		lextest.Lex(1, 3, "\n", lexeme.NEWLINE),
		lextest.Lex(2, 0, "\t", lexeme.SPACE),
		lextest.Lex(2, 1, "1", lexeme.NUMBER),
		lextest.Lex(2, 2, ",", lexeme.DELIM),
		lextest.Lex(2, 3, "\n", lexeme.NEWLINE),
		lextest.Lex(3, 0, ")", lexeme.R_PAREN),
		lextest.Lex(3, 1, "\n", lexeme.NEWLINE),
	)

	doTest(t, updatePositions, given, exp)
}

func Test9_1(t *testing.T) {

	// 1# First
	// 2 # Second
	// 3  # Third
	given := lextest.Feign(
		lextest.Lex(0, 0, "1", lexeme.NUMBER),
		lextest.Lex(0, 1, "# First", lexeme.COMMENT),
		lextest.Lex(0, 8, "\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "2", lexeme.NUMBER),
		lextest.Lex(1, 1, " ", lexeme.SPACE),
		lextest.Lex(1, 2, "# Second", lexeme.COMMENT),
		lextest.Lex(1, 10, "\n", lexeme.NEWLINE),
		lextest.Lex(2, 0, "3", lexeme.NUMBER),
		lextest.Lex(2, 1, "  ", lexeme.SPACE),
		lextest.Lex(2, 3, "# Third", lexeme.COMMENT),
		lextest.Lex(2, 10, "\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Lex(0, 0, "1", lexeme.NUMBER),
		lextest.Lex(0, 0, "  ", lexeme.SPACE),
		lextest.Lex(0, 1, "# First", lexeme.COMMENT),
		lextest.Lex(0, 8, "\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "2", lexeme.NUMBER),
		lextest.Lex(1, 1, " ", lexeme.SPACE),
		lextest.Lex(0, 0, " ", lexeme.SPACE),
		lextest.Lex(1, 2, "# Second", lexeme.COMMENT),
		lextest.Lex(1, 10, "\n", lexeme.NEWLINE),
		lextest.Lex(2, 0, "3", lexeme.NUMBER),
		lextest.Lex(2, 1, "  ", lexeme.SPACE),
		lextest.Lex(2, 3, "# Third", lexeme.COMMENT),
		lextest.Lex(2, 10, "\n", lexeme.NEWLINE),
	)

	doTest(t, alignComments, given, exp)
}

func Test9_2(t *testing.T) {

	// 1# First
	// 2 # Second
	// 3
	// 4  # Fourth
	// 5# Fifth
	given := lextest.Feign(
		lextest.Lex(0, 0, "1", lexeme.NUMBER),
		lextest.Lex(0, 1, "# First", lexeme.COMMENT),
		lextest.Lex(0, 8, "\n", lexeme.NEWLINE),

		lextest.Lex(1, 0, "2", lexeme.NUMBER),
		lextest.Lex(1, 1, " ", lexeme.SPACE),
		lextest.Lex(1, 2, "# Second", lexeme.COMMENT),
		lextest.Lex(1, 10, "\n", lexeme.NEWLINE),

		lextest.Lex(2, 0, "3", lexeme.NUMBER),
		lextest.Lex(2, 1, "\n", lexeme.NEWLINE),

		lextest.Lex(3, 0, "4", lexeme.NUMBER),
		lextest.Lex(3, 1, "  ", lexeme.SPACE),
		lextest.Lex(3, 3, "# Fourth", lexeme.COMMENT),
		lextest.Lex(3, 11, "\n", lexeme.NEWLINE),

		lextest.Lex(4, 0, "5", lexeme.NUMBER),
		lextest.Lex(4, 1, "# Fifth", lexeme.COMMENT),
		lextest.Lex(4, 11, "\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Lex(0, 0, "1", lexeme.NUMBER),
		lextest.Lex(0, 0, " ", lexeme.SPACE),
		lextest.Lex(0, 1, "# First", lexeme.COMMENT),
		lextest.Lex(0, 8, "\n", lexeme.NEWLINE),

		lextest.Lex(1, 0, "2", lexeme.NUMBER),
		lextest.Lex(1, 1, " ", lexeme.SPACE),
		lextest.Lex(1, 2, "# Second", lexeme.COMMENT),
		lextest.Lex(1, 10, "\n", lexeme.NEWLINE),

		lextest.Lex(2, 0, "3", lexeme.NUMBER),
		lextest.Lex(2, 1, "\n", lexeme.NEWLINE),

		lextest.Lex(3, 0, "4", lexeme.NUMBER),
		lextest.Lex(3, 1, "  ", lexeme.SPACE),
		lextest.Lex(3, 3, "# Fourth", lexeme.COMMENT),
		lextest.Lex(3, 11, "\n", lexeme.NEWLINE),

		lextest.Lex(4, 0, "5", lexeme.NUMBER),
		lextest.Lex(0, 0, "  ", lexeme.SPACE),
		lextest.Lex(4, 1, "# Fifth", lexeme.COMMENT),
		lextest.Lex(4, 11, "\n", lexeme.NEWLINE),
	)

	doTest(t, alignComments, given, exp)
}

func Test99_1(t *testing.T) {

	// " @Println ( \n1 , 1+2 , \n 1 ) \n "
	given := lextest.Feign(
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n\r", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok(" ", lexeme.SPACE),
	)

	// @Println(
	//   1, 1,
	//
	//   1)
	exp := lextest.Feign(
		lextest.Lex(0, 0, "@Println", lexeme.SPELL),
		lextest.Lex(0, 8, "(", lexeme.L_PAREN),
		lextest.Lex(0, 9, "\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, "\t", lexeme.SPACE),
		lextest.Lex(1, 1, "1", lexeme.NUMBER),
		lextest.Lex(1, 2, " ", lexeme.SPACE),
		lextest.Lex(1, 3, "+", lexeme.ADD),
		lextest.Lex(1, 4, " ", lexeme.SPACE),
		lextest.Lex(1, 5, "2", lexeme.NUMBER),
		lextest.Lex(1, 6, ",", lexeme.DELIM),
		lextest.Lex(1, 7, " ", lexeme.SPACE),
		lextest.Lex(1, 8, "1", lexeme.NUMBER),
		lextest.Lex(1, 9, ",", lexeme.DELIM),
		lextest.Lex(1, 10, "\n", lexeme.NEWLINE),
		lextest.Lex(2, 0, "\n", lexeme.NEWLINE),
		lextest.Lex(3, 0, "\t", lexeme.SPACE),
		lextest.Lex(3, 1, "1", lexeme.NUMBER),
		lextest.Lex(3, 2, ")", lexeme.R_PAREN),
		lextest.Lex(3, 3, "\n", lexeme.NEWLINE),
	)

	format(given.Iterator())
	lextest.Equal(t, exp.Head(), given.Head())
}
