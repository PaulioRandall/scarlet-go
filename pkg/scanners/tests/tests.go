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

	acts := f(in)

	check(t, exps, acts)
}

func A2_Whitespace(t *testing.T, f ScanFunc) {

	in := " \t\r\v\f"

	exps := []Token{
		Token{WHITESPACE, " \t\r\v\f", 0, 0},
	}

	acts := f(in)

	check(t, exps, acts)
}

func A3_Comments(t *testing.T, f ScanFunc) {

	in := "// This is a comment"

	exps := []Token{
		Token{COMMENT, "// This is a comment", 0, 0},
	}

	acts := f(in)

	check(t, exps, acts)
}

func A4_Match(t *testing.T, f ScanFunc) {

	in := "MATCH"

	exps := []Token{
		Token{MATCH, "MATCH", 0, 0},
	}

	acts := f(in)

	check(t, exps, acts)
}

func A5_Bool_False(t *testing.T, f ScanFunc) {

	in := "FALSE"

	exps := []Token{
		Token{BOOL, "FALSE", 0, 0},
	}

	acts := f(in)

	check(t, exps, acts)
}

func A6_Bool_True(t *testing.T, f ScanFunc) {

	in := "TRUE"

	exps := []Token{
		Token{BOOL, "TRUE", 0, 0},
	}

	acts := f(in)

	check(t, exps, acts)
}
