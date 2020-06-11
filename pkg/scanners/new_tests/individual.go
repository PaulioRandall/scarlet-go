package tests

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func tok(ty TokenType, v string) Token {
	return NewToken(ty, v, 0, 0)
}

func solo(ty TokenType, v string) (string, []Token) {
	return v, []Token{tok(ty, v)}
}

func T1_1() (in string, expects []Token) {
	return solo(TK_NEWLINE, "\n")
}

func T1_2() (in string, expects []Token) {
	return solo(TK_NEWLINE, "\r\n")
}

func T2_1() (in string, expects []Token) {
	return solo(TK_WHITESPACE, " \t\r\v\f")
}

func T3_1() (in string, expects []Token) {
	return solo(TK_COMMENT, "// This is a comment")
}

func T4_1() (in string, expects []Token) {
	return solo(TK_WHEN, "when")
}

func T5_1() (in string, expects []Token) {
	return solo(TK_BOOL, "false")
}

func T5_2() (in string, expects []Token) {
	return solo(TK_BOOL, "true")
}

func T7_1() (in string, expects []Token) {
	return solo(TK_DEFINITION, "def")
}

func T8_1() (in string, expects []Token) {
	return solo(TK_FUNCTION, "F")
}

func T9_1() (in string, expects []Token) {
	return solo(TK_IDENTIFIER, "a")
}

func T9_2() (in string, expects []Token) {
	return solo(TK_IDENTIFIER, "abc")
}

func T9_3() (in string, expects []Token) {
	return solo(TK_IDENTIFIER, "a_b")
}

func T9_4() (in string, expects []Token) {
	return solo(TK_IDENTIFIER, "ab_")
}

func T9_5() (in string, expects []Token) {
	return solo(TK_IDENTIFIER, "def_")
}

func T9_6() (in string, expects []Token) {
	return solo(TK_IDENTIFIER, "deff")
}

func T9_7() (in string, expects []Token) {
	return solo(TK_IDENTIFIER, "ddef")
}

func T10_1() (in string, expects []Token) {
	return solo(TK_ASSIGNMENT, ":=")
}

func T11_1() (in string, expects []Token) {
	return solo(TK_OUTPUT, "^")
}

func T12_1() (in string, expects []Token) {
	return solo(TK_LESS_THAN_OR_EQUAL, "<=")
}

func T13_1() (in string, expects []Token) {
	return solo(TK_MORE_THAN_OR_EQUAL, ">=")
}

func T14_1() (in string, expects []Token) {
	return solo(TK_BLOCK_OPEN, "{")
}

func T15_1() (in string, expects []Token) {
	return solo(TK_BLOCK_CLOSE, "}")
}

func T16_1() (in string, expects []Token) {
	return solo(TK_PAREN_OPEN, "(")
}

func T17_1() (in string, expects []Token) {
	return solo(TK_PAREN_CLOSE, ")")
}

func T18_1() (in string, expects []Token) {
	return solo(TK_GUARD_OPEN, "[")
}

func T19_1() (in string, expects []Token) {
	return solo(TK_GUARD_CLOSE, "]")
}

func T20_1() (in string, expects []Token) {
	return solo(TK_DELIMITER, ",")
}

func T21_1() (in string, expects []Token) {
	return solo(TK_VOID, "_")
}

func T22_1() (in string, expects []Token) {
	return solo(TK_TERMINATOR, ";")
}

func T23_1() (in string, expects []Token) {
	return solo(TK_SPELL, "@abc")
}

func T23_2() (in string, expects []Token) {
	return solo(TK_SPELL, "@abc.xyz")
}

func T23_3() (in string, expects []Token) {
	return solo(TK_SPELL, "@a.b.c.d")
}

func T24_1() (in string, expects []Token) {
	return solo(TK_PLUS, "+")
}

func T25_1() (in string, expects []Token) {
	return solo(TK_MINUS, "-")
}

func T26_1() (in string, expects []Token) {
	return solo(TK_MULTIPLY, "*")
}

func T27_1() (in string, expects []Token) {
	return solo(TK_DIVIDE, "/")
}

func T28_1() (in string, expects []Token) {
	return solo(TK_REMAINDER, "%")
}

func T29_1() (in string, expects []Token) {
	return solo(TK_AND, "&")
}

func T30_1() (in string, expects []Token) {
	return solo(TK_OR, "|")
}

func T31_1() (in string, expects []Token) {
	return solo(TK_EQUAL, "==")
}

func T32_1() (in string, expects []Token) {
	return solo(TK_NOT_EQUAL, "!=")
}

func T33_1() (in string, expects []Token) {
	return solo(TK_LESS_THAN, "<")
}

func T34_1() (in string, expects []Token) {
	return solo(TK_MORE_THAN, ">")
}

func T35_1() (in string, expects []Token) {
	return solo(TK_STRING, `""`)
}

func T35_2() (in string, expects []Token) {
	return solo(TK_STRING, `"abc"`)
}

func T36_1() (in string, expects []Token) {
	return solo(TK_NUMBER, "1")
}

func T36_2() (in string, expects []Token) {
	return solo(TK_NUMBER, "123")
}

func T36_3() (in string, expects []Token) {
	return solo(TK_NUMBER, "1.0")
}

func T36_4() (in string, expects []Token) {
	return solo(TK_NUMBER, "123.456")
}

func T37_1() (in string, expects []Token) {
	return solo(TK_LOOP, "loop")
}

func T41_1() (in string, expects []Token) {
	return solo(TK_EXPR_FUNC, "E")
}

func T42_1() (in string, expects []Token) {
	return solo(TK_THEN, "->")
}
