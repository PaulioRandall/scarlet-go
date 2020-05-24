package z_tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func S1_Assignment(t *testing.T, f ScanFunc) {

	in := "x := 1"

	exps := []Token{
		tok{M_IDENTIFIER, "x", 0, 0},
		tok{M_WHITESPACE, " ", 0, 1},
		tok{M_ASSIGN, ":=", 0, 2},
		tok{M_WHITESPACE, " ", 0, 4},
		tok{M_NUMBER, "1", 0, 5},
	}

	checkMany(t, exps, f(in))
}

func S2_MultiAssignment(t *testing.T, f ScanFunc) {

	in := "x,y:=1,TRUE"

	exps := []Token{
		tok{M_IDENTIFIER, "x", 0, 0},
		tok{M_DELIMITER, ",", 0, 1},
		tok{M_IDENTIFIER, "y", 0, 2},
		tok{M_ASSIGN, ":=", 0, 3},
		tok{M_NUMBER, "1", 0, 5},
		tok{M_DELIMITER, ",", 0, 6},
		tok{M_BOOL, "TRUE", 0, 7},
	}

	checkMany(t, exps, f(in))
}

func S3_GuardBlock(t *testing.T, f ScanFunc) {

	in := "[1<2] x:=TRUE"

	exps := []Token{
		tok{M_GUARD_OPEN, "[", 0, 0},
		tok{M_NUMBER, "1", 0, 1},
		tok{M_LESS_THAN, "<", 0, 2},
		tok{M_NUMBER, "2", 0, 3},
		tok{M_GUARD_CLOSE, "]", 0, 4},
		tok{M_WHITESPACE, " ", 0, 5},
		tok{M_IDENTIFIER, "x", 0, 6},
		tok{M_ASSIGN, ":=", 0, 7},
		tok{M_BOOL, "TRUE", 0, 9},
	}

	checkMany(t, exps, f(in))
}

func S4_MatchBlock(t *testing.T, f ScanFunc) {

	in := "MATCH {\n" +
		"\t[FALSE] x:=FALSE\n" +
		"\t[TRUE] x:=TRUE\n" +
		"}"

	exps := []Token{
		tok{M_MATCH, "MATCH", 0, 0}, // Line start
		tok{M_WHITESPACE, " ", 0, 5},
		tok{M_BLOCK_OPEN, "{", 0, 6},
		tok{M_NEWLINE, "\n", 0, 7}, // Line start
		tok{M_WHITESPACE, "\t", 1, 0},
		tok{M_GUARD_OPEN, "[", 1, 1},
		tok{M_BOOL, "FALSE", 1, 2},
		tok{M_GUARD_CLOSE, "]", 1, 7},
		tok{M_WHITESPACE, " ", 1, 8},
		tok{M_IDENTIFIER, "x", 1, 9},
		tok{M_ASSIGN, ":=", 1, 10},
		tok{M_BOOL, "FALSE", 1, 12},
		tok{M_NEWLINE, "\n", 1, 17},
		tok{M_WHITESPACE, "\t", 2, 0}, // Line start
		tok{M_GUARD_OPEN, "[", 2, 1},
		tok{M_BOOL, "TRUE", 2, 2},
		tok{M_GUARD_CLOSE, "]", 2, 6},
		tok{M_WHITESPACE, " ", 2, 7},
		tok{M_IDENTIFIER, "x", 2, 8},
		tok{M_ASSIGN, ":=", 2, 9},
		tok{M_BOOL, "TRUE", 2, 11},
		tok{M_NEWLINE, "\n", 2, 15},
		tok{M_BLOCK_CLOSE, "}", 3, 0}, // Line start
	}

	checkMany(t, exps, f(in))
}

func S5_FuncDef(t *testing.T, f ScanFunc) {

	in := "F(a,b,^c,^d)"

	exps := []Token{
		tok{M_FUNC, "F", 0, 0},
		tok{M_PAREN_OPEN, "(", 0, 1},
		tok{M_IDENTIFIER, "a", 0, 2},
		tok{M_DELIMITER, ",", 0, 3},
		tok{M_IDENTIFIER, "b", 0, 4},
		tok{M_DELIMITER, ",", 0, 5},
		tok{M_OUTPUT, "^", 0, 6},
		tok{M_IDENTIFIER, "c", 0, 7},
		tok{M_DELIMITER, ",", 0, 8},
		tok{M_OUTPUT, "^", 0, 9},
		tok{M_IDENTIFIER, "d", 0, 10},
		tok{M_PAREN_CLOSE, ")", 0, 11},
	}

	checkMany(t, exps, f(in))
}

func S6_FuncCall(t *testing.T, f ScanFunc) {

	in := "xyz(a,b)"

	exps := []Token{
		tok{M_IDENTIFIER, "xyz", 0, 0},
		tok{M_PAREN_OPEN, "(", 0, 3},
		tok{M_IDENTIFIER, "a", 0, 4},
		tok{M_DELIMITER, ",", 0, 5},
		tok{M_IDENTIFIER, "b", 0, 6},
		tok{M_PAREN_CLOSE, ")", 0, 7},
	}

	checkMany(t, exps, f(in))
}

func S7_Expression(t *testing.T, f ScanFunc) {

	in := "1+2-3*4/5%6"

	exps := []Token{
		tok{M_NUMBER, "1", 0, 0},
		tok{M_ADD, "+", 0, 1},
		tok{M_NUMBER, "2", 0, 2},
		tok{M_SUBTRACT, "-", 0, 3},
		tok{M_NUMBER, "3", 0, 4},
		tok{M_MULTIPLY, "*", 0, 5},
		tok{M_NUMBER, "4", 0, 6},
		tok{M_DIVIDE, "/", 0, 7},
		tok{M_NUMBER, "5", 0, 8},
		tok{M_REMAINDER, "%", 0, 9},
		tok{M_NUMBER, "6", 0, 10},
	}

	checkMany(t, exps, f(in))
}

func S8_Block(t *testing.T, f ScanFunc) {

	in := "{\n" +
		"\tx:=1\n" +
		"\ty:=2\n" +
		"}"

	exps := []Token{
		tok{M_BLOCK_OPEN, "{", 0, 0}, // Line Start
		tok{M_NEWLINE, "\n", 0, 1},
		tok{M_WHITESPACE, "\t", 1, 0}, // Line Start
		tok{M_IDENTIFIER, "x", 1, 1},
		tok{M_ASSIGN, ":=", 1, 2},
		tok{M_NUMBER, "1", 1, 4},
		tok{M_NEWLINE, "\n", 1, 5},
		tok{M_WHITESPACE, "\t", 2, 0}, // Line Start
		tok{M_IDENTIFIER, "y", 2, 1},
		tok{M_ASSIGN, ":=", 2, 2},
		tok{M_NUMBER, "2", 2, 4},
		tok{M_NEWLINE, "\n", 2, 5},
		tok{M_BLOCK_CLOSE, "}", 3, 0}, // Line Start
	}

	checkMany(t, exps, f(in))
}

func S9_List(t *testing.T, f ScanFunc) {

	in := "LIST {\n" +
		"\t`There's a snake in my boot`,\n" +
		"\t" + `"{x} + {y} = {x + y}"` + ",\n" +
		"}"

	exps := []Token{
		tok{M_LIST, "LIST", 0, 0},
		tok{M_WHITESPACE, " ", 0, 4},
		tok{M_BLOCK_OPEN, "{", 0, 5},
		tok{M_NEWLINE, "\n", 0, 6},
		tok{M_WHITESPACE, "\t", 1, 0}, // Line Start
		tok{M_STRING, "`There's a snake in my boot`", 1, 1},
		tok{M_DELIMITER, ",", 1, 29},
		tok{M_NEWLINE, "\n", 1, 30},
		tok{M_WHITESPACE, "\t", 2, 0}, // Line Start
		tok{M_TEMPLATE, `"{x} + {y} = {x + y}"`, 2, 1},
		tok{M_DELIMITER, ",", 2, 22},
		tok{M_NEWLINE, "\n", 2, 23},
		tok{M_BLOCK_CLOSE, "}", 3, 0}, // Line Start
	}

	checkMany(t, exps, f(in))
}

func S10_Loop(t *testing.T, f ScanFunc) {

	in := "LOOP i [i<5] {}"

	exps := []Token{
		tok{M_LOOP, "LOOP", 0, 0},
		tok{M_WHITESPACE, " ", 0, 4},
		tok{M_IDENTIFIER, "i", 0, 5},
		tok{M_WHITESPACE, " ", 0, 6},
		tok{M_GUARD_OPEN, "[", 0, 7},
		tok{M_IDENTIFIER, "i", 0, 8},
		tok{M_LESS_THAN, "<", 0, 9},
		tok{M_NUMBER, "5", 0, 10},
		tok{M_GUARD_CLOSE, "]", 0, 11},
		tok{M_WHITESPACE, " ", 0, 12},
		tok{M_BLOCK_OPEN, "{", 0, 13},
		tok{M_BLOCK_CLOSE, "}", 0, 14},
	}

	checkMany(t, exps, f(in))
}

func S11_ModifyList(t *testing.T, f ScanFunc) {

	in := "x[3],x[>>]:=1,99"

	exps := []Token{
		tok{M_IDENTIFIER, "x", 0, 0},
		tok{M_GUARD_OPEN, "[", 0, 1},
		tok{M_NUMBER, "3", 0, 2},
		tok{M_GUARD_CLOSE, "]", 0, 3},
		tok{M_DELIMITER, ",", 0, 4},
		tok{M_IDENTIFIER, "x", 0, 5},
		tok{M_GUARD_OPEN, "[", 0, 6},
		tok{M_LIST_END, ">>", 0, 7},
		tok{M_GUARD_CLOSE, "]", 0, 9},
		tok{M_ASSIGN, ":=", 0, 10},
		tok{M_NUMBER, "1", 0, 12},
		tok{M_DELIMITER, ",", 0, 13},
		tok{M_NUMBER, "99", 0, 14},
	}

	checkMany(t, exps, f(in))
}
