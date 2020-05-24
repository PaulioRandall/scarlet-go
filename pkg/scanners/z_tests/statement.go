package z_tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func S1_Assignment(t *testing.T, f ScanFunc) {

	in := "x : 1"

	exps := []Token{
		tok{IDENTIFIER, "x", 0, 0},
		tok{WHITESPACE, " ", 0, 1},
		tok{ASSIGN, ":", 0, 2},
		tok{WHITESPACE, " ", 0, 3},
		tok{NUMBER, "1", 0, 4},
	}

	checkMany(t, exps, f(in))
}

func S2_MultiAssignment(t *testing.T, f ScanFunc) {

	in := "x,y:1,TRUE"

	exps := []Token{
		tok{IDENTIFIER, "x", 0, 0},
		tok{DELIMITER, ",", 0, 1},
		tok{IDENTIFIER, "y", 0, 2},
		tok{ASSIGN, ":", 0, 3},
		tok{NUMBER, "1", 0, 4},
		tok{DELIMITER, ",", 0, 5},
		tok{BOOL, "TRUE", 0, 6},
	}

	checkMany(t, exps, f(in))
}

func S3_GuardBlock(t *testing.T, f ScanFunc) {

	in := "[1<2] x:TRUE"

	exps := []Token{
		tok{GUARD_OPEN, "[", 0, 0},
		tok{NUMBER, "1", 0, 1},
		tok{LESS_THAN, "<", 0, 2},
		tok{NUMBER, "2", 0, 3},
		tok{GUARD_CLOSE, "]", 0, 4},
		tok{WHITESPACE, " ", 0, 5},
		tok{IDENTIFIER, "x", 0, 6},
		tok{ASSIGN, ":", 0, 7},
		tok{BOOL, "TRUE", 0, 8},
	}

	checkMany(t, exps, f(in))
}

func S4_MatchBlock(t *testing.T, f ScanFunc) {

	in := "MATCH {\n" +
		"\t[FALSE] x:FALSE\n" +
		"\t[TRUE] x:TRUE\n" +
		"}"

	exps := []Token{
		tok{MATCH, "MATCH", 0, 0}, // Line start
		tok{WHITESPACE, " ", 0, 5},
		tok{BLOCK_OPEN, "{", 0, 6},
		tok{NEWLINE, "\n", 0, 7}, // Line start
		tok{WHITESPACE, "\t", 1, 0},
		tok{GUARD_OPEN, "[", 1, 1},
		tok{BOOL, "FALSE", 1, 2},
		tok{GUARD_CLOSE, "]", 1, 7},
		tok{WHITESPACE, " ", 1, 8},
		tok{IDENTIFIER, "x", 1, 9},
		tok{ASSIGN, ":", 1, 10},
		tok{BOOL, "FALSE", 1, 11},
		tok{NEWLINE, "\n", 1, 16},
		tok{WHITESPACE, "\t", 2, 0}, // Line start
		tok{GUARD_OPEN, "[", 2, 1},
		tok{BOOL, "TRUE", 2, 2},
		tok{GUARD_CLOSE, "]", 2, 6},
		tok{WHITESPACE, " ", 2, 7},
		tok{IDENTIFIER, "x", 2, 8},
		tok{ASSIGN, ":", 2, 9},
		tok{BOOL, "TRUE", 2, 10},
		tok{NEWLINE, "\n", 2, 14},
		tok{BLOCK_CLOSE, "}", 3, 0}, // Line start
	}

	checkMany(t, exps, f(in))
}

func S5_FuncDef(t *testing.T, f ScanFunc) {

	in := "F(a,b,^c,^d)"

	exps := []Token{
		tok{FUNC, "F", 0, 0},
		tok{PAREN_OPEN, "(", 0, 1},
		tok{IDENTIFIER, "a", 0, 2},
		tok{DELIMITER, ",", 0, 3},
		tok{IDENTIFIER, "b", 0, 4},
		tok{DELIMITER, ",", 0, 5},
		tok{OUTPUT, "^", 0, 6},
		tok{IDENTIFIER, "c", 0, 7},
		tok{DELIMITER, ",", 0, 8},
		tok{OUTPUT, "^", 0, 9},
		tok{IDENTIFIER, "d", 0, 10},
		tok{PAREN_CLOSE, ")", 0, 11},
	}

	checkMany(t, exps, f(in))
}

func S6_FuncCall(t *testing.T, f ScanFunc) {

	in := "xyz(a,b)"

	exps := []Token{
		tok{IDENTIFIER, "xyz", 0, 0},
		tok{PAREN_OPEN, "(", 0, 3},
		tok{IDENTIFIER, "a", 0, 4},
		tok{DELIMITER, ",", 0, 5},
		tok{IDENTIFIER, "b", 0, 6},
		tok{PAREN_CLOSE, ")", 0, 7},
	}

	checkMany(t, exps, f(in))
}

func S7_Expression(t *testing.T, f ScanFunc) {

	in := "1+2-3*4/5%6"

	exps := []Token{
		tok{NUMBER, "1", 0, 0},
		tok{ADD, "+", 0, 1},
		tok{NUMBER, "2", 0, 2},
		tok{SUBTRACT, "-", 0, 3},
		tok{NUMBER, "3", 0, 4},
		tok{MULTIPLY, "*", 0, 5},
		tok{NUMBER, "4", 0, 6},
		tok{DIVIDE, "/", 0, 7},
		tok{NUMBER, "5", 0, 8},
		tok{REMAINDER, "%", 0, 9},
		tok{NUMBER, "6", 0, 10},
	}

	checkMany(t, exps, f(in))
}

func S8_Block(t *testing.T, f ScanFunc) {

	in := "{\n" +
		"\tx:1\n" +
		"\ty:2\n" +
		"}"

	exps := []Token{
		tok{BLOCK_OPEN, "{", 0, 0}, // Line Start
		tok{NEWLINE, "\n", 0, 1},
		tok{WHITESPACE, "\t", 1, 0}, // Line Start
		tok{IDENTIFIER, "x", 1, 1},
		tok{ASSIGN, ":", 1, 2},
		tok{NUMBER, "1", 1, 3},
		tok{NEWLINE, "\n", 1, 4},
		tok{WHITESPACE, "\t", 2, 0}, // Line Start
		tok{IDENTIFIER, "y", 2, 1},
		tok{ASSIGN, ":", 2, 2},
		tok{NUMBER, "2", 2, 3},
		tok{NEWLINE, "\n", 2, 4},
		tok{BLOCK_CLOSE, "}", 3, 0}, // Line Start
	}

	checkMany(t, exps, f(in))
}

func S9_List(t *testing.T, f ScanFunc) {

	in := "LIST {\n" +
		"\t`There's a snake in my boot`,\n" +
		"\t" + `"{x} + {y} = {x + y}"` + ",\n" +
		"}"

	exps := []Token{
		tok{LIST, "LIST", 0, 0},
		tok{WHITESPACE, " ", 0, 4},
		tok{BLOCK_OPEN, "{", 0, 5},
		tok{NEWLINE, "\n", 0, 6},
		tok{WHITESPACE, "\t", 1, 0}, // Line Start
		tok{STRING, "`There's a snake in my boot`", 1, 1},
		tok{DELIMITER, ",", 1, 29},
		tok{NEWLINE, "\n", 1, 30},
		tok{WHITESPACE, "\t", 2, 0}, // Line Start
		tok{TEMPLATE, `"{x} + {y} = {x + y}"`, 2, 1},
		tok{DELIMITER, ",", 2, 22},
		tok{NEWLINE, "\n", 2, 23},
		tok{BLOCK_CLOSE, "}", 3, 0}, // Line Start
	}

	checkMany(t, exps, f(in))
}

func S10_Loop(t *testing.T, f ScanFunc) {

	in := "LOOP i [i<5] {}"

	exps := []Token{
		tok{LOOP, "LOOP", 0, 0},
		tok{WHITESPACE, " ", 0, 4},
		tok{IDENTIFIER, "i", 0, 5},
		tok{WHITESPACE, " ", 0, 6},
		tok{GUARD_OPEN, "[", 0, 7},
		tok{IDENTIFIER, "i", 0, 8},
		tok{LESS_THAN, "<", 0, 9},
		tok{NUMBER, "5", 0, 10},
		tok{GUARD_CLOSE, "]", 0, 11},
		tok{WHITESPACE, " ", 0, 12},
		tok{BLOCK_OPEN, "{", 0, 13},
		tok{BLOCK_CLOSE, "}", 0, 14},
	}

	checkMany(t, exps, f(in))
}

func S11_ModifyList(t *testing.T, f ScanFunc) {

	in := "x[3],x[>>]:1,99"

	exps := []Token{
		tok{IDENTIFIER, "x", 0, 0},
		tok{GUARD_OPEN, "[", 0, 1},
		tok{NUMBER, "3", 0, 2},
		tok{GUARD_CLOSE, "]", 0, 3},
		tok{DELIMITER, ",", 0, 4},
		tok{IDENTIFIER, "x", 0, 5},
		tok{GUARD_OPEN, "[", 0, 6},
		tok{LIST_END, ">>", 0, 7},
		tok{GUARD_CLOSE, "]", 0, 9},
		tok{ASSIGN, ":", 0, 10},
		tok{NUMBER, "1", 0, 11},
		tok{DELIMITER, ",", 0, 12},
		tok{NUMBER, "99", 0, 13},
	}

	checkMany(t, exps, f(in))
}
