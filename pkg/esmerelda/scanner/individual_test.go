package scanner

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func tok(ty TokenType, v string) Token {
	return NewToken(ty, v, 0, 0)
}

func Test_T1_1(t *testing.T) {
	okTest(t, "\n", []Token{tok(TK_NEWLINE, "\n")})
}

func Test_T1_2(t *testing.T) {
	okTest(t, "\r\n", []Token{tok(TK_NEWLINE, "\r\n")})
}

func Test_T2_1(t *testing.T) {
	okTest(t, " \t\r\v\f", []Token{tok(TK_WHITESPACE, " \t\r\v\f")})
}

func Test_T3_1(t *testing.T) {
	okTest(t, "// This is a comment", []Token{tok(TK_COMMENT, "// This is a comment")})
}

func Test_T4_1(t *testing.T) {
	okTest(t, "when", []Token{tok(TK_WHEN, "when")})
}

func Test_T5_1(t *testing.T) {
	okTest(t, "false", []Token{tok(TK_BOOL, "false")})
}

func Test_T5_2(t *testing.T) {
	okTest(t, "true", []Token{tok(TK_BOOL, "true")})
}

func Test_T7_1(t *testing.T) {
	okTest(t, "def", []Token{tok(TK_DEFINITION, "def")})
}

func Test_T8_1(t *testing.T) {
	okTest(t, "F", []Token{tok(TK_FUNCTION, "F")})
}

func Test_T9_1(t *testing.T) {
	okTest(t, "a", []Token{tok(TK_IDENTIFIER, "a")})
}

func Test_T9_2(t *testing.T) {
	okTest(t, "abc", []Token{tok(TK_IDENTIFIER, "abc")})
}

func Test_T9_3(t *testing.T) {
	okTest(t, "a_b", []Token{tok(TK_IDENTIFIER, "a_b")})
}

func Test_T9_4(t *testing.T) {
	okTest(t, "ab_", []Token{tok(TK_IDENTIFIER, "ab_")})
}

func Test_T9_5(t *testing.T) {
	okTest(t, "def_", []Token{tok(TK_IDENTIFIER, "def_")})
}

func Test_T9_6(t *testing.T) {
	okTest(t, "deff", []Token{tok(TK_IDENTIFIER, "deff")})
}

func Test_T9_7(t *testing.T) {
	okTest(t, "ddef", []Token{tok(TK_IDENTIFIER, "ddef")})
}

func Test_T10_1(t *testing.T) {
	okTest(t, ":=", []Token{tok(TK_ASSIGNMENT, ":=")})
}

func Test_T11_1(t *testing.T) {
	okTest(t, "^", []Token{tok(TK_OUTPUT, "^")})
}

func Test_T11_2(t *testing.T) {
	okTest(t, "->", []Token{tok(TK_OUTPUTS, "->")})
}

func Test_T12_1(t *testing.T) {
	okTest(t, "<=", []Token{tok(TK_LESS_THAN_OR_EQUAL, "<=")})
}

func Test_T13_1(t *testing.T) {
	okTest(t, ">=", []Token{tok(TK_MORE_THAN_OR_EQUAL, ">=")})
}

func Test_T14_1(t *testing.T) {
	okTest(t, "{", []Token{tok(TK_BLOCK_OPEN, "{")})
}

func Test_T15_1(t *testing.T) {
	okTest(t, "}", []Token{tok(TK_BLOCK_CLOSE, "}")})
}

func Test_T16_1(t *testing.T) {
	okTest(t, "(", []Token{tok(TK_PAREN_OPEN, "(")})
}

func Test_T17_1(t *testing.T) {
	okTest(t, ")", []Token{tok(TK_PAREN_CLOSE, ")")})
}

func Test_T18_1(t *testing.T) {
	okTest(t, "[", []Token{tok(TK_GUARD_OPEN, "[")})
}

func Test_T19_1(t *testing.T) {
	okTest(t, "]", []Token{tok(TK_GUARD_CLOSE, "]")})
}

func Test_T20_1(t *testing.T) {
	okTest(t, ",", []Token{tok(TK_DELIMITER, ",")})
}

func Test_T21_1(t *testing.T) {
	okTest(t, "_", []Token{tok(TK_VOID, "_")})
}

func Test_T22_1(t *testing.T) {
	okTest(t, ";", []Token{tok(TK_TERMINATOR, ";")})
}

func Test_T23_1(t *testing.T) {
	okTest(t, "@abc", []Token{tok(TK_SPELL, "@abc")})
}

func Test_T23_2(t *testing.T) {
	okTest(t, "@abc.xyz", []Token{tok(TK_SPELL, "@abc.xyz")})
}

func Test_T23_3(t *testing.T) {
	okTest(t, "@a.b.c.d", []Token{tok(TK_SPELL, "@a.b.c.d")})
}

func Test_T24_1(t *testing.T) {
	okTest(t, "+", []Token{tok(TK_PLUS, "+")})
}

func Test_T25_1(t *testing.T) {
	okTest(t, "-", []Token{tok(TK_MINUS, "-")})
}

func Test_T26_1(t *testing.T) {
	okTest(t, "*", []Token{tok(TK_MULTIPLY, "*")})
}

func Test_T27_1(t *testing.T) {
	okTest(t, "/", []Token{tok(TK_DIVIDE, "/")})
}

func Test_T28_1(t *testing.T) {
	okTest(t, "%", []Token{tok(TK_REMAINDER, "%")})
}

func Test_T29_1(t *testing.T) {
	okTest(t, "&", []Token{tok(TK_AND, "&")})
}

func Test_T30_1(t *testing.T) {
	okTest(t, "|", []Token{tok(TK_OR, "|")})
}

func Test_T31_1(t *testing.T) {
	okTest(t, "==", []Token{tok(TK_EQUAL, "==")})
}

func Test_T32_1(t *testing.T) {
	okTest(t, "!=", []Token{tok(TK_NOT_EQUAL, "!=")})
}

func Test_T33_1(t *testing.T) {
	okTest(t, "<", []Token{tok(TK_LESS_THAN, "<")})
}

func Test_T34_1(t *testing.T) {
	okTest(t, ">", []Token{tok(TK_MORE_THAN, ">")})
}

func Test_T35_1(t *testing.T) {
	okTest(t, `""`, []Token{tok(TK_STRING, `""`)})
}

func Test_T35_2(t *testing.T) {
	okTest(t, `"abc"`, []Token{tok(TK_STRING, `"abc"`)})
}

func Test_T36_1(t *testing.T) {
	okTest(t, "1", []Token{tok(TK_NUMBER, "1")})
}

func Test_T36_2(t *testing.T) {
	okTest(t, "123", []Token{tok(TK_NUMBER, "123")})
}

func Test_T36_3(t *testing.T) {
	okTest(t, "1.0", []Token{tok(TK_NUMBER, "1.0")})
}

func Test_T36_4(t *testing.T) {
	okTest(t, "123.456", []Token{tok(TK_NUMBER, "123.456")})
}

func Test_T37_1(t *testing.T) {
	okTest(t, "loop", []Token{tok(TK_LOOP, "loop")})
}

func Test_T41_1(t *testing.T) {
	okTest(t, "E", []Token{tok(TK_EXPR_FUNC, "E")})
}

func Test_T42_1(t *testing.T) {
	okTest(t, ":", []Token{tok(TK_THEN, ":")})
}

func Test_T43_1(t *testing.T) {
	okTest(t, "exit", []Token{tok(TK_EXIT, "exit")})
}

func Test_T44_1(t *testing.T) {
	okTest(t, "?", []Token{tok(TK_EXISTS, "?")})
}
