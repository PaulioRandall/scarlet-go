package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme/lextest"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in string, exp *lexeme.Container) {
	act, e := ScanStr(in)
	require.Nil(t, e, "%+v", e)
	lextest.Equal(t, exp.Head(), act.Head())
}

func doErrTest(t *testing.T, in string) {
	_, e := ScanStr(in)
	require.NotNil(t, e, "Expected an error")
}

func Test1_1(t *testing.T) {
	doTest(t, "\n", lextest.Feign(
		lextest.Lex(0, 0, "\n", lexeme.NEWLINE),
	))
}

func Test1_2(t *testing.T) {
	doTest(t, "\r\n", lextest.Feign(
		lextest.Lex(0, 0, "\r\n", lexeme.NEWLINE),
	))
}

func Test2_1(t *testing.T) {
	doTest(t, "# Comment", lextest.Feign(
		lextest.Lex(0, 0, "# Comment", lexeme.COMMENT),
	))
}

func Test3_1(t *testing.T) {
	doTest(t, "   \t\v\f", lextest.Feign(
		lextest.Lex(0, 0, "   \t\v\f", lexeme.WHITESPACE),
	))
}

func Test4_1(t *testing.T) {
	doTest(t, "false", lextest.Feign(
		lextest.Lex(0, 0, "false", lexeme.BOOL),
	))
}

func Test4_2(t *testing.T) {
	doTest(t, "true", lextest.Feign(
		lextest.Lex(0, 0, "true", lexeme.BOOL),
	))
}

func Test4_3(t *testing.T) {
	doTest(t, "abc", lextest.Feign(
		lextest.Lex(0, 0, "abc", lexeme.IDENT),
	))
}

func Test4_4(t *testing.T) {
	doTest(t, "ab_c", lextest.Feign(
		lextest.Lex(0, 0, "ab_c", lexeme.IDENT),
	))
}

func Test5_1(t *testing.T) {
	doTest(t, "@abc", lextest.Feign(
		lextest.Lex(0, 0, "@abc", lexeme.SPELL),
	))
}

func Test5_2(t *testing.T) {
	doTest(t, "@abc.xyz", lextest.Feign(
		lextest.Lex(0, 0, "@abc.xyz", lexeme.SPELL),
	))
}

func Test5_3(t *testing.T) {
	doTest(t, "@a.b.c", lextest.Feign(
		lextest.Lex(0, 0, "@a.b.c", lexeme.SPELL),
	))
}

func Test5_4(t *testing.T) {
	doErrTest(t, "@abc.")
}

func Test5_5(t *testing.T) {
	doErrTest(t, "@abc._")
}

func Test6_1(t *testing.T) {
	doTest(t, `""`, lextest.Feign(
		lextest.Lex(0, 0, `""`, lexeme.STRING),
	))
}

func Test6_2(t *testing.T) {
	doTest(t, `"abc"`, lextest.Feign(
		lextest.Lex(0, 0, `"abc"`, lexeme.STRING),
	))
}

func Test6_3(t *testing.T) {
	doTest(t, `"\"abc\""`, lextest.Feign(
		lextest.Lex(0, 0, `"\"abc\""`, lexeme.STRING),
	))
}

func Test6_4(t *testing.T) {
	doErrTest(t, `"`)
}

func Test6_5(t *testing.T) {
	doErrTest(t, `"\"`)
}

func Test6_6(t *testing.T) {
	doErrTest(t, `"\"abc`)
}

func Test7_1(t *testing.T) {
	doTest(t, "1", lextest.Feign(
		lextest.Lex(0, 0, "1", lexeme.NUMBER),
	))
}

func Test7_2(t *testing.T) {
	doTest(t, "123", lextest.Feign(
		lextest.Lex(0, 0, "123", lexeme.NUMBER),
	))
}

func Test7_3(t *testing.T) {
	doTest(t, "123.456", lextest.Feign(
		lextest.Lex(0, 0, "123.456", lexeme.NUMBER),
	))
}

func Test7_4(t *testing.T) {
	doErrTest(t, "123.")
}

func Test7_5(t *testing.T) {
	doErrTest(t, "123.a")
}

func Test8_1(t *testing.T) {
	doTest(t, ":=", lextest.Feign(
		lextest.Lex(0, 0, ":=", lexeme.ASSIGN),
	))
}

func Test8_2(t *testing.T) {
	doErrTest(t, ":")
}

func Test9_1(t *testing.T) {
	doTest(t, "+", lextest.Feign(
		lextest.Lex(0, 0, "+", lexeme.ADD),
	))
}

func Test9_2(t *testing.T) {
	doTest(t, "-", lextest.Feign(
		lextest.Lex(0, 0, "-", lexeme.SUB),
	))
}

func Test9_3(t *testing.T) {
	doTest(t, "*", lextest.Feign(
		lextest.Lex(0, 0, "*", lexeme.MUL),
	))
}

func Test9_4(t *testing.T) {
	doTest(t, "/", lextest.Feign(
		lextest.Lex(0, 0, "/", lexeme.DIV),
	))
}

func Test9_5(t *testing.T) {
	doTest(t, "%", lextest.Feign(
		lextest.Lex(0, 0, "%", lexeme.REM),
	))
}

func Test10_1(t *testing.T) {
	doTest(t, "&&", lextest.Feign(
		lextest.Lex(0, 0, "&&", lexeme.AND),
	))
}

func Test10_2(t *testing.T) {
	doTest(t, "||", lextest.Feign(
		lextest.Lex(0, 0, "||", lexeme.OR),
	))
}

func Test10_3(t *testing.T) {
	doErrTest(t, "&")
}

func Test10_4(t *testing.T) {
	doErrTest(t, "|")
}

func Test11_1(t *testing.T) {
	doTest(t, "<", lextest.Feign(
		lextest.Lex(0, 0, "<", lexeme.LESS),
	))
}

func Test11_2(t *testing.T) {
	doTest(t, ">", lextest.Feign(
		lextest.Lex(0, 0, ">", lexeme.MORE),
	))
}

func Test11_3(t *testing.T) {
	doTest(t, "<=", lextest.Feign(
		lextest.Lex(0, 0, "<=", lexeme.LESS_EQUAL),
	))
}

func Test11_4(t *testing.T) {
	doTest(t, ">=", lextest.Feign(
		lextest.Lex(0, 0, ">=", lexeme.MORE_EQUAL),
	))
}

func Test11_5(t *testing.T) {
	doTest(t, "==", lextest.Feign(
		lextest.Lex(0, 0, "==", lexeme.EQUAL),
	))
}

func Test11_6(t *testing.T) {
	doTest(t, "!=", lextest.Feign(
		lextest.Lex(0, 0, "!=", lexeme.NOT_EQUAL),
	))
}

func Test11_7(t *testing.T) {
	doErrTest(t, "=")
}

func Test11_8(t *testing.T) {
	doErrTest(t, "!")
}

func Test99_0(t *testing.T) {

	given := "@Println(1,\n true,\n \"heir\")\n"

	doTest(t, given, lextest.Feign(
		lextest.Lex(0, 0, "@Println", lexeme.SPELL),
		lextest.Lex(0, 8, "(", lexeme.L_PAREN),
		lextest.Lex(0, 9, "1", lexeme.NUMBER),
		lextest.Lex(0, 10, ",", lexeme.DELIM),
		lextest.Lex(0, 11, "\n", lexeme.NEWLINE),
		lextest.Lex(1, 0, " ", lexeme.WHITESPACE),
		lextest.Lex(1, 1, "true", lexeme.BOOL),
		lextest.Lex(1, 5, ",", lexeme.DELIM),
		lextest.Lex(1, 6, "\n", lexeme.NEWLINE),
		lextest.Lex(2, 0, " ", lexeme.WHITESPACE),
		lextest.Lex(2, 1, `"heir"`, lexeme.STRING),
		lextest.Lex(2, 7, ")", lexeme.R_PAREN),
		lextest.Lex(2, 8, "\n", lexeme.NEWLINE),
	))
}
