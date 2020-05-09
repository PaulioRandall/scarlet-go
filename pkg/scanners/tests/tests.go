package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func A1(t *testing.T, tag string, f ScanFunc) {

	in := "\n\r\n"

	exps := []Token{
		Token{NEWLINE, "\n", 0, 0},
		Token{NEWLINE, "\r\n", 1, 0},
	}

	acts := f(in)

	check(t, tag+".A1", exps, acts)
}

func A2(t *testing.T, tag string, f ScanFunc) {

	in := " \t\r\v\f"

	exps := []Token{
		Token{WHITESPACE, " \t\r\v\f", 0, 0},
	}

	acts := f(in)

	check(t, tag+".A2", exps, acts)
}
