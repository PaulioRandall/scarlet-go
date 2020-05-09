package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func S1_Assignment(t *testing.T, f ScanFunc) {

	in := "x := 1"

	exps := []Token{
		Token{ID, "x", 0, 0},
		Token{WHITESPACE, " ", 0, 1},
		Token{ASSIGN, ":=", 0, 2},
		Token{WHITESPACE, " ", 0, 4},
		Token{NUMBER, "1", 0, 5},
	}

	checkMany(t, exps, f(in))
}

func S2_MultiAssignment(t *testing.T, f ScanFunc) {

	in := "x,y:=1,TRUE"

	exps := []Token{
		Token{ID, "x", 0, 0},
		Token{DELIM, ",", 0, 1},
		Token{ID, "y", 0, 2},
		Token{ASSIGN, ":=", 0, 3},
		Token{NUMBER, "1", 0, 5},
		Token{DELIM, ",", 0, 6},
		Token{BOOL, "TRUE", 0, 7},
	}

	checkMany(t, exps, f(in))
}

func S3_GuardBlock(t *testing.T, f ScanFunc) {

	in := "[1<2] x:=TRUE"

	exps := []Token{
		Token{GUARD_OPEN, "[", 0, 0},
		Token{NUMBER, "1", 0, 1},
		Token{LESS_THAN, "<", 0, 2},
		Token{NUMBER, "2", 0, 3},
		Token{GUARD_CLOSE, "]", 0, 4},
		Token{WHITESPACE, " ", 0, 5},
		Token{ID, "x", 0, 6},
		Token{ASSIGN, ":=", 0, 7},
		Token{BOOL, "TRUE", 0, 9},
	}

	checkMany(t, exps, f(in))
}
