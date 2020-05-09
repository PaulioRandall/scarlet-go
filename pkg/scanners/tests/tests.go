package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func A1_Newlines(t *testing.T, f ScanFunc) {
	checkOne(t, Token{NEWLINE, "\n", 0, 0}, f("\n"))
	checkOne(t, Token{NEWLINE, "\r\n", 0, 0}, f("\r\n"))
}

func A2_Whitespace(t *testing.T, f ScanFunc) {
	in := " \t\r\v\f"
	exp := Token{WHITESPACE, " \t\r\v\f", 0, 0}
	checkOne(t, exp, f(in))
}

func A3_Comments(t *testing.T, f ScanFunc) {
	in := "// This is a comment"
	exp := Token{COMMENT, "// This is a comment", 0, 0}
	checkOne(t, exp, f(in))
}

func A4_Match(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MATCH, "MATCH", 0, 0}, f("MATCH"))
}

func A5_False(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BOOL, "FALSE", 0, 0}, f("FALSE"))
}

func A6_True(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BOOL, "TRUE", 0, 0}, f("TRUE"))
}

func A7_List(t *testing.T, f ScanFunc) {
	checkOne(t, Token{LIST, "LIST", 0, 0}, f("LIST"))
}

func A8_Fix(t *testing.T, f ScanFunc) {
	checkOne(t, Token{FIX, "FIX", 0, 0}, f("FIX"))
}

func A9_Eof(t *testing.T, f ScanFunc) {
	check(t, []Token{}, f("EOF"))
}

func A10_F(t *testing.T, f ScanFunc) {
	checkOne(t, Token{FUNC, "F", 0, 0}, f("F"))
}

func A11_Identifiers(t *testing.T, f ScanFunc) {
	checkOne(t, Token{ID, "a", 0, 0}, f("a"))
	checkOne(t, Token{ID, "abc", 0, 0}, f("abc"))
	checkOne(t, Token{ID, "a_c", 0, 0}, f("a_c"))
	checkOne(t, Token{ID, "ab_", 0, 0}, f("ab_"))
}

func A12_Assign(t *testing.T, f ScanFunc) {
	checkOne(t, Token{ASSIGN, ":=", 0, 0}, f(":="))
}

func A13_Returns(t *testing.T, f ScanFunc) {
	checkOne(t, Token{RETURNS, "->", 0, 0}, f("->"))
}

func A14_LessThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, Token{LESS_THAN_OR_EQUAL, "<=", 0, 0}, f("<="))
}

func A15_MoreThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MORE_THAN_OR_EQUAL, ">=", 0, 0}, f(">="))
}

func A16_BlockOpen(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BLOCK_OPEN, "{", 0, 0}, f("{"))
}

func A17_BlockClose(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BLOCK_CLOSE, "}", 0, 0}, f("}"))
}

func A18_ParenOpen(t *testing.T, f ScanFunc) {
	checkOne(t, Token{PAREN_OPEN, "(", 0, 0}, f("("))
}

func A19_ParenClose(t *testing.T, f ScanFunc) {
	checkOne(t, Token{PAREN_CLOSE, ")", 0, 0}, f(")"))
}

func A20_GuardOpen(t *testing.T, f ScanFunc) {
	checkOne(t, Token{GUARD_OPEN, "[", 0, 0}, f("["))
}

func A21_GuardClose(t *testing.T, f ScanFunc) {
	checkOne(t, Token{GUARD_CLOSE, "]", 0, 0}, f("]"))
}

func A22_Delim(t *testing.T, f ScanFunc) {
	checkOne(t, Token{DELIM, ",", 0, 0}, f(","))
}

func A23_Void(t *testing.T, f ScanFunc) {
	checkOne(t, Token{VOID, "_", 0, 0}, f("_"))
}

func A24_Terminator(t *testing.T, f ScanFunc) {
	checkOne(t, Token{TERMINATOR, ";", 0, 0}, f(";"))
}

func A25_Spell(t *testing.T, f ScanFunc) {
	checkOne(t, Token{SPELL, "@", 0, 0}, f("@"))
}

func A26_Add(t *testing.T, f ScanFunc) {
	checkOne(t, Token{ADD, "+", 0, 0}, f("+"))
}

func A27_Subtract(t *testing.T, f ScanFunc) {
	checkOne(t, Token{SUBTRACT, "-", 0, 0}, f("-"))
}

func A28_Multiply(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MULTIPLY, "*", 0, 0}, f("*"))
}

func A29_Divide(t *testing.T, f ScanFunc) {
	checkOne(t, Token{DIVIDE, "/", 0, 0}, f("/"))
}

func A30_Remainder(t *testing.T, f ScanFunc) {
	checkOne(t, Token{REMAINDER, "%", 0, 0}, f("%"))
}
