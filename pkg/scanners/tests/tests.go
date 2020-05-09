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

func A5_Bool_False(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BOOL, "FALSE", 0, 0}, f("FALSE"))
}

func A6_Bool_True(t *testing.T, f ScanFunc) {
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
