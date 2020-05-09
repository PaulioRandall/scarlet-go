package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func A1_Newlines(t *testing.T, f ScanFunc) {

	in := "\n\r\n"

	exps := []Token{
		Token{NEWLINE, "\n", 0, 0},
		Token{NEWLINE, "\r\n", 1, 0},
	}

	check(t, exps, f(in))
}

func A2_Whitespace(t *testing.T, f ScanFunc) {

	in := " \t\r\v\f"

	exps := []Token{
		Token{WHITESPACE, " \t\r\v\f", 0, 0},
	}

	check(t, exps, f(in))
}

func A3_Comments(t *testing.T, f ScanFunc) {

	in := "// This is a comment"

	exps := []Token{
		Token{COMMENT, "// This is a comment", 0, 0},
	}

	check(t, exps, f(in))
}

func A4_Match(t *testing.T, f ScanFunc) {

	in := "MATCH"

	exps := []Token{
		Token{MATCH, "MATCH", 0, 0},
	}

	check(t, exps, f(in))
}

func A5_Bool_False(t *testing.T, f ScanFunc) {

	in := "FALSE"

	exps := []Token{
		Token{BOOL, "FALSE", 0, 0},
	}

	check(t, exps, f(in))
}

func A6_Bool_True(t *testing.T, f ScanFunc) {

	in := "TRUE"

	exps := []Token{
		Token{BOOL, "TRUE", 0, 0},
	}

	check(t, exps, f(in))
}

func A7_List(t *testing.T, f ScanFunc) {

	in := "LIST"

	exps := []Token{
		Token{LIST, "LIST", 0, 0},
	}

	check(t, exps, f(in))
}

func A8_Fix(t *testing.T, f ScanFunc) {

	in := "FIX"

	exps := []Token{
		Token{FIX, "FIX", 0, 0},
	}

	check(t, exps, f(in))
}

func A9_Eof(t *testing.T, f ScanFunc) {

	in := "EOF"

	exps := []Token{}

	check(t, exps, f(in))
}

func A10_F(t *testing.T, f ScanFunc) {

	in := "F"

	exps := []Token{
		Token{FUNC, "F", 0, 0},
	}

	check(t, exps, f(in))
}
