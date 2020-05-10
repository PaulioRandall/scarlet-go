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

func S4_MatchBlock(t *testing.T, f ScanFunc) {

	in := "MATCH {\n" +
		"\t[FALSE] x:=FALSE\n" +
		"\t[TRUE] x:=TRUE\n" +
		"}"

	exps := []Token{
		Token{MATCH, "MATCH", 0, 0}, // Line start
		Token{WHITESPACE, " ", 0, 5},
		Token{BLOCK_OPEN, "{", 0, 6},
		Token{NEWLINE, "\n", 0, 7}, // Line start
		Token{WHITESPACE, "\t", 1, 0},
		Token{GUARD_OPEN, "[", 1, 1},
		Token{BOOL, "FALSE", 1, 2},
		Token{GUARD_CLOSE, "]", 1, 7},
		Token{WHITESPACE, " ", 1, 8},
		Token{ID, "x", 1, 9},
		Token{ASSIGN, ":=", 1, 10},
		Token{BOOL, "FALSE", 1, 12},
		Token{NEWLINE, "\n", 1, 17},
		Token{WHITESPACE, "\t", 2, 0}, // Line start
		Token{GUARD_OPEN, "[", 2, 1},
		Token{BOOL, "TRUE", 2, 2},
		Token{GUARD_CLOSE, "]", 2, 6},
		Token{WHITESPACE, " ", 2, 7},
		Token{ID, "x", 2, 8},
		Token{ASSIGN, ":=", 2, 9},
		Token{BOOL, "TRUE", 2, 11},
		Token{NEWLINE, "\n", 2, 15},
		Token{BLOCK_CLOSE, "}", 3, 0}, // Line start
	}

	checkMany(t, exps, f(in))
}

func S5_FuncDef(t *testing.T, f ScanFunc) {

	in := "F(a,b -> c,d)"

	exps := []Token{
		Token{FUNC, "F", 0, 0},
		Token{PAREN_OPEN, "(", 0, 1},
		Token{ID, "a", 0, 2},
		Token{DELIM, ",", 0, 3},
		Token{ID, "b", 0, 4},
		Token{WHITESPACE, " ", 0, 5},
		Token{RETURNS, "->", 0, 6},
		Token{WHITESPACE, " ", 0, 8},
		Token{ID, "c", 0, 9},
		Token{DELIM, ",", 0, 10},
		Token{ID, "d", 0, 11},
		Token{PAREN_CLOSE, ")", 0, 12},
	}

	checkMany(t, exps, f(in))
}

func S6_FuncCall(t *testing.T, f ScanFunc) {

	in := "xyz(a,b)"

	exps := []Token{
		Token{ID, "xyz", 0, 0},
		Token{PAREN_OPEN, "(", 0, 3},
		Token{ID, "a", 0, 4},
		Token{DELIM, ",", 0, 5},
		Token{ID, "b", 0, 6},
		Token{PAREN_CLOSE, ")", 0, 7},
	}

	checkMany(t, exps, f(in))
}

func S7_Expression(t *testing.T, f ScanFunc) {

	in := "1+2-3*4/5%6"

	exps := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{ADD, "+", 0, 1},
		Token{NUMBER, "2", 0, 2},
		Token{SUBTRACT, "-", 0, 3},
		Token{NUMBER, "3", 0, 4},
		Token{MULTIPLY, "*", 0, 5},
		Token{NUMBER, "4", 0, 6},
		Token{DIVIDE, "/", 0, 7},
		Token{NUMBER, "5", 0, 8},
		Token{REMAINDER, "%", 0, 9},
		Token{NUMBER, "6", 0, 10},
	}

	checkMany(t, exps, f(in))
}
