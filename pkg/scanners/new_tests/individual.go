package tests

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func tok(m Morpheme, v string) Token {
	return NewToken(m, v, 0, 0)
}

func solo(m Morpheme, v string) (string, []Token) {
	return v, []Token{tok(m, v)}
}

func T1_1() (in string, expects []Token) {
	return solo(NEWLINE, "\n")
}

func T1_2() (in string, expects []Token) {
	return solo(NEWLINE, "\r\n")
}

func T2_1() (in string, expects []Token) {
	return solo(WHITESPACE, " \t\r\v\f")
}

func T3_1() (in string, expects []Token) {
	return solo(COMMENT, "// This is a comment")
}

func T4_1() (in string, expects []Token) {
	return solo(WHEN, "when")
}

func T5_1() (in string, expects []Token) {
	return solo(BOOL, "false")
}

func T5_2() (in string, expects []Token) {
	return solo(BOOL, "true")
}

func T7_1() (in string, expects []Token) {
	return solo(DEF, "def")
}

func T8_1() (in string, expects []Token) {
	return solo(FUNC, "F")
}

func T9_1() (in string, expects []Token) {
	return solo(IDENTIFIER, "a")
}

func T9_2() (in string, expects []Token) {
	return solo(IDENTIFIER, "abc")
}

func T9_3() (in string, expects []Token) {
	return solo(IDENTIFIER, "a_b")
}

func T9_4() (in string, expects []Token) {
	return solo(IDENTIFIER, "ab_")
}

func T9_5() (in string, expects []Token) {
	return solo(IDENTIFIER, "def_")
}

func T9_6() (in string, expects []Token) {
	return solo(IDENTIFIER, "deff")
}

func T9_7() (in string, expects []Token) {
	return solo(IDENTIFIER, "ddef")
}

func T10_1() (in string, expects []Token) {
	return solo(ASSIGN, ":=")
}

func T11_1() (in string, expects []Token) {
	return solo(OUTPUT, "^")
}

func T12_1() (in string, expects []Token) {
	return solo(LESS_THAN_OR_EQUAL, "<=")
}

func T13_1() (in string, expects []Token) {
	return solo(MORE_THAN_OR_EQUAL, ">=")
}

func T14_1() (in string, expects []Token) {
	return solo(BLOCK_OPEN, "{")
}

func T15_1() (in string, expects []Token) {
	return solo(BLOCK_CLOSE, "}")
}

func T16_1() (in string, expects []Token) {
	return solo(PAREN_OPEN, "(")
}

func T17_1() (in string, expects []Token) {
	return solo(PAREN_CLOSE, ")")
}

func T18_1() (in string, expects []Token) {
	return solo(GUARD_OPEN, "[")
}

func T19_1() (in string, expects []Token) {
	return solo(GUARD_CLOSE, "]")
}

func T20_1() (in string, expects []Token) {
	return solo(DELIMITER, ",")
}

func T21_1() (in string, expects []Token) {
	return solo(VOID, "_")
}

func T22_1() (in string, expects []Token) {
	return solo(TERMINATOR, ";")
}

func T23_1() (in string, expects []Token) {
	return solo(SPELL, "@")
}

func T24_1() (in string, expects []Token) {
	return solo(ADD, "+")
}

func T25_1() (in string, expects []Token) {
	return solo(SUBTRACT, "-")
}

func T26_1() (in string, expects []Token) {
	return solo(MULTIPLY, "*")
}

func T27_1() (in string, expects []Token) {
	return solo(DIVIDE, "/")
}

func T28_1() (in string, expects []Token) {
	return solo(REMAINDER, "%")
}

func T29_1() (in string, expects []Token) {
	return solo(AND, "&")
}

func T30_1() (in string, expects []Token) {
	return solo(OR, "|")
}

func T31_1() (in string, expects []Token) {
	return solo(EQUAL, "==")
}

func T32_1() (in string, expects []Token) {
	return solo(NOT_EQUAL, "!=")
}

func T33_1() (in string, expects []Token) {
	return solo(LESS_THAN, "<")
}

func T34_1() (in string, expects []Token) {
	return solo(MORE_THAN, ">")
}

func T35_1() (in string, expects []Token) {
	return solo(STRING, `""`)
}

func T35_2() (in string, expects []Token) {
	return solo(STRING, `"abc"`)
}

func T36_1() (in string, expects []Token) {
	return solo(NUMBER, "1")
}

func T36_2() (in string, expects []Token) {
	return solo(NUMBER, "123")
}

func T36_3() (in string, expects []Token) {
	return solo(NUMBER, "1.0")
}

func T36_4() (in string, expects []Token) {
	return solo(NUMBER, "123.456")
}

func T37_1() (in string, expects []Token) {
	return solo(LOOP, "loop")
}

func T41_1() (in string, expects []Token) {
	return solo(EXPR_FUNC, "E")
}

func T42_1() (in string, expects []Token) {
	return solo(THEN, "->")
}
