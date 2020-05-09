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

func A4_Key_Match(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MATCH, "MATCH", 0, 0}, f("MATCH"))
}

func A5_Key_False(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BOOL, "FALSE", 0, 0}, f("FALSE"))
}

func A6_Key_True(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BOOL, "TRUE", 0, 0}, f("TRUE"))
}

func A7_Key_List(t *testing.T, f ScanFunc) {
	checkOne(t, Token{LIST, "LIST", 0, 0}, f("LIST"))
}

func A8_Key_Fix(t *testing.T, f ScanFunc) {
	checkOne(t, Token{FIX, "FIX", 0, 0}, f("FIX"))
}

func A9_Key_Eof(t *testing.T, f ScanFunc) {
	check(t, []Token{}, f("EOF"))
}

func A10_Key_F(t *testing.T, f ScanFunc) {
	checkOne(t, Token{FUNC, "F", 0, 0}, f("F"))
}

func A11_Identifiers(t *testing.T, f ScanFunc) {
	checkOne(t, Token{ID, "a", 0, 0}, f("a"))
	checkOne(t, Token{ID, "abc", 0, 0}, f("abc"))
	checkOne(t, Token{ID, "a_c", 0, 0}, f("a_c"))
	checkOne(t, Token{ID, "ab_", 0, 0}, f("ab_"))
}

func A12_Sym_Assign(t *testing.T, f ScanFunc) {
	checkOne(t, Token{ASSIGN, ":=", 0, 0}, f(":="))
}

func A13_Sym_Returns(t *testing.T, f ScanFunc) {
	checkOne(t, Token{RETURNS, "->", 0, 0}, f("->"))
}

func A14_Sym_LessThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, Token{LESS_THAN_OR_EQUAL, "<=", 0, 0}, f("<="))
}

func A15_Sym_MoreThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MORE_THAN_OR_EQUAL, ">=", 0, 0}, f(">="))
}
