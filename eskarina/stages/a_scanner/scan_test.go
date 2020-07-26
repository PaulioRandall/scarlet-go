package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in string, exp *lexeme.Lexeme) {
	act, e := ScanStr(in)
	require.Nil(t, e, "%+v", e)
	lextest.Equal(t, exp, act)
}

func doErrTest(t *testing.T, in string) {
	_, e := ScanStr(in)
	require.NotNil(t, e, "Expected an error")
}

func Test1_1(t *testing.T) {
	doTest(t, "\n", lextest.Feign(
		lextest.Lex(0, 0, "\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
	))
}

func Test1_2(t *testing.T) {
	doTest(t, "\r\n", lextest.Feign(
		lextest.Lex(0, 0, "\r\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
	))
}

func Test2_1(t *testing.T) {
	doTest(t, "# Comment", lextest.Feign(
		lextest.Lex(0, 0, "# Comment", lexeme.PR_REDUNDANT, lexeme.PR_COMMENT),
	))
}

func Test3_1(t *testing.T) {
	doTest(t, "   \t\v\f", lextest.Feign(
		lextest.Lex(0, 0, "   \t\v\f", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
	))
}

func Test4_1(t *testing.T) {
	doTest(t, "false", lextest.Feign(
		lextest.Lex(0, 0, "false", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_BOOL),
	))
}

func Test4_2(t *testing.T) {
	doTest(t, "true", lextest.Feign(
		lextest.Lex(0, 0, "true", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_BOOL),
	))
}

func Test4_3(t *testing.T) {
	doTest(t, "abc", lextest.Feign(
		lextest.Lex(0, 0, "abc", lexeme.PR_TERM, lexeme.PR_ASSIGNEE, lexeme.PR_IDENTIFIER),
	))
}

func Test4_4(t *testing.T) {
	doTest(t, "ab_c", lextest.Feign(
		lextest.Lex(0, 0, "ab_c", lexeme.PR_TERM, lexeme.PR_ASSIGNEE, lexeme.PR_IDENTIFIER),
	))
}

func Test5_1(t *testing.T) {
	doTest(t, "@abc", lextest.Feign(
		lextest.Lex(0, 0, "@abc", lexeme.PR_SPELL),
	))
}

func Test5_2(t *testing.T) {
	doTest(t, "@abc.xyz", lextest.Feign(
		lextest.Lex(0, 0, "@abc.xyz", lexeme.PR_SPELL),
	))
}

func Test5_3(t *testing.T) {
	doTest(t, "@a.b.c", lextest.Feign(
		lextest.Lex(0, 0, "@a.b.c", lexeme.PR_SPELL),
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
		lextest.Lex(0, 0, `""`, lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_STRING),
	))
}

func Test6_2(t *testing.T) {
	doTest(t, `"abc"`, lextest.Feign(
		lextest.Lex(0, 0, `"abc"`, lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_STRING),
	))
}

func Test6_3(t *testing.T) {
	doTest(t, `"\"abc\""`, lextest.Feign(
		lextest.Lex(0, 0, `"\"abc\""`, lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_STRING),
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
		lextest.Lex(0, 0, "1", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
	))
}

func Test7_2(t *testing.T) {
	doTest(t, "123", lextest.Feign(
		lextest.Lex(0, 0, "123", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
	))
}

func Test7_3(t *testing.T) {
	doTest(t, "123.456", lextest.Feign(
		lextest.Lex(0, 0, "123.456", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
	))
}

func Test7_4(t *testing.T) {
	doErrTest(t, "123.")
}

func Test7_5(t *testing.T) {
	doErrTest(t, "123.a")
}