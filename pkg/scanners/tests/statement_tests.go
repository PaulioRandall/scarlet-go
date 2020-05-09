package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func S1_Assignment(t *testing.T, f ScanFunc) {

	exps := []Token{
		Token{ID, "x", 0, 0},
		Token{WHITESPACE, " ", 0, 1},
		Token{ASSIGN, ":=", 0, 2},
		Token{WHITESPACE, " ", 0, 4},
		Token{NUMBER, "1", 0, 5},
	}

	checkMany(t, exps, f("x := 1"))
}
