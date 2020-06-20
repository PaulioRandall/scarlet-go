package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func S1_Assignment(t *testing.T, f ScanFunc) {

	in := "x : 1"

	exps := []Token{
		NewToken(TK_IDENTIFIER, "x", 0, 0),
		NewToken(TK_WHITESPACE, " ", 0, 1),
		NewToken(TK_ASSIGNMENT, ":", 0, 2),
		NewToken(TK_WHITESPACE, " ", 0, 3),
		NewToken(TK_NUMBER, "1", 0, 4),
	}

	checkMany(t, exps, f(in))
}

func S2_MultiAssignment(t *testing.T, f ScanFunc) {

	in := "x,y:1,TRUE"

	exps := []Token{
		NewToken(TK_IDENTIFIER, "x", 0, 0),
		NewToken(TK_DELIMITER, ",", 0, 1),
		NewToken(TK_IDENTIFIER, "y", 0, 2),
		NewToken(TK_ASSIGNMENT, ":", 0, 3),
		NewToken(TK_NUMBER, "1", 0, 4),
		NewToken(TK_DELIMITER, ",", 0, 5),
		NewToken(TK_BOOL, "TRUE", 0, 6),
	}

	checkMany(t, exps, f(in))
}

func S3_GuardBlock(t *testing.T, f ScanFunc) {

	in := "[1<2] x:TRUE"

	exps := []Token{
		NewToken(TK_GUARD_OPEN, "[", 0, 0),
		NewToken(TK_NUMBER, "1", 0, 1),
		NewToken(TK_LESS_THAN, "<", 0, 2),
		NewToken(TK_NUMBER, "2", 0, 3),
		NewToken(TK_GUARD_CLOSE, "]", 0, 4),
		NewToken(TK_WHITESPACE, " ", 0, 5),
		NewToken(TK_IDENTIFIER, "x", 0, 6),
		NewToken(TK_ASSIGNMENT, ":", 0, 7),
		NewToken(TK_BOOL, "TRUE", 0, 8),
	}

	checkMany(t, exps, f(in))
}

func S4_WhenBlock(t *testing.T, f ScanFunc) {

	in := "WHEN {\n" +
		"\t[FALSE] x:FALSE\n" +
		"\t[TRUE] x:TRUE\n" +
		"}"

	exps := []Token{
		NewToken(TK_WHEN, "WHEN", 0, 0), // Line start
		NewToken(TK_WHITESPACE, " ", 0, 4),
		NewToken(TK_BLOCK_OPEN, "{", 0, 5),
		NewToken(TK_NEWLINE, "\n", 0, 6), // Line start
		NewToken(TK_WHITESPACE, "\t", 1, 0),
		NewToken(TK_GUARD_OPEN, "[", 1, 1),
		NewToken(TK_BOOL, "FALSE", 1, 2),
		NewToken(TK_GUARD_CLOSE, "]", 1, 7),
		NewToken(TK_WHITESPACE, " ", 1, 8),
		NewToken(TK_IDENTIFIER, "x", 1, 9),
		NewToken(TK_ASSIGNMENT, ":", 1, 10),
		NewToken(TK_BOOL, "FALSE", 1, 11),
		NewToken(TK_NEWLINE, "\n", 1, 16),
		NewToken(TK_WHITESPACE, "\t", 2, 0), // Line start
		NewToken(TK_GUARD_OPEN, "[", 2, 1),
		NewToken(TK_BOOL, "TRUE", 2, 2),
		NewToken(TK_GUARD_CLOSE, "]", 2, 6),
		NewToken(TK_WHITESPACE, " ", 2, 7),
		NewToken(TK_IDENTIFIER, "x", 2, 8),
		NewToken(TK_ASSIGNMENT, ":", 2, 9),
		NewToken(TK_BOOL, "TRUE", 2, 10),
		NewToken(TK_NEWLINE, "\n", 2, 14),
		NewToken(TK_BLOCK_CLOSE, "}", 3, 0), // Line start
	}

	checkMany(t, exps, f(in))
}

func S5_FuncDef(t *testing.T, f ScanFunc) {

	in := "F(a,b,^c,^d)"

	exps := []Token{
		NewToken(TK_FUNCTION, "F", 0, 0),
		NewToken(TK_PAREN_OPEN, "(", 0, 1),
		NewToken(TK_IDENTIFIER, "a", 0, 2),
		NewToken(TK_DELIMITER, ",", 0, 3),
		NewToken(TK_IDENTIFIER, "b", 0, 4),
		NewToken(TK_DELIMITER, ",", 0, 5),
		NewToken(TK_OUTPUT, "^", 0, 6),
		NewToken(TK_IDENTIFIER, "c", 0, 7),
		NewToken(TK_DELIMITER, ",", 0, 8),
		NewToken(TK_OUTPUT, "^", 0, 9),
		NewToken(TK_IDENTIFIER, "d", 0, 10),
		NewToken(TK_PAREN_CLOSE, ")", 0, 11),
	}

	checkMany(t, exps, f(in))
}

func S6_FuncCall(t *testing.T, f ScanFunc) {

	in := "xyz(a,b)"

	exps := []Token{
		NewToken(TK_IDENTIFIER, "xyz", 0, 0),
		NewToken(TK_PAREN_OPEN, "(", 0, 3),
		NewToken(TK_IDENTIFIER, "a", 0, 4),
		NewToken(TK_DELIMITER, ",", 0, 5),
		NewToken(TK_IDENTIFIER, "b", 0, 6),
		NewToken(TK_PAREN_CLOSE, ")", 0, 7),
	}

	checkMany(t, exps, f(in))
}

func S7_Expression(t *testing.T, f ScanFunc) {

	in := "1+2-3*4/5%6"

	exps := []Token{
		NewToken(TK_NUMBER, "1", 0, 0),
		NewToken(TK_PLUS, "+", 0, 1),
		NewToken(TK_NUMBER, "2", 0, 2),
		NewToken(TK_MINUS, "-", 0, 3),
		NewToken(TK_NUMBER, "3", 0, 4),
		NewToken(TK_MULTIPLY, "*", 0, 5),
		NewToken(TK_NUMBER, "4", 0, 6),
		NewToken(TK_DIVIDE, "/", 0, 7),
		NewToken(TK_NUMBER, "5", 0, 8),
		NewToken(TK_REMAINDER, "%", 0, 9),
		NewToken(TK_NUMBER, "6", 0, 10),
	}

	checkMany(t, exps, f(in))
}

func S8_Block(t *testing.T, f ScanFunc) {

	in := "{\n" +
		"\tx:1\n" +
		"\ty:2\n" +
		"}"

	exps := []Token{
		NewToken(TK_BLOCK_OPEN, "{", 0, 0), // Line Start
		NewToken(TK_NEWLINE, "\n", 0, 1),
		NewToken(TK_WHITESPACE, "\t", 1, 0), // Line Start
		NewToken(TK_IDENTIFIER, "x", 1, 1),
		NewToken(TK_ASSIGNMENT, ":", 1, 2),
		NewToken(TK_NUMBER, "1", 1, 3),
		NewToken(TK_NEWLINE, "\n", 1, 4),
		NewToken(TK_WHITESPACE, "\t", 2, 0), // Line Start
		NewToken(TK_IDENTIFIER, "y", 2, 1),
		NewToken(TK_ASSIGNMENT, ":", 2, 2),
		NewToken(TK_NUMBER, "2", 2, 3),
		NewToken(TK_NEWLINE, "\n", 2, 4),
		NewToken(TK_BLOCK_CLOSE, "}", 3, 0), // Line Start
	}

	checkMany(t, exps, f(in))
}

func S9_List(t *testing.T, f ScanFunc) {

	in := "LIST {\n" +
		"\t" + `"There's a snake in my boot",` + "\n" +
		"\t" + `"{x} + {y} = {x + y}"` + ",\n" +
		"}"

	exps := []Token{
		NewToken(TK_LIST, "LIST", 0, 0),
		NewToken(TK_WHITESPACE, " ", 0, 4),
		NewToken(TK_BLOCK_OPEN, "{", 0, 5),
		NewToken(TK_NEWLINE, "\n", 0, 6),
		NewToken(TK_WHITESPACE, "\t", 1, 0), // Line Start
		NewToken(TK_STRING, `"There's a snake in my boot"`, 1, 1),
		NewToken(TK_DELIMITER, ",", 1, 29),
		NewToken(TK_NEWLINE, "\n", 1, 30),
		NewToken(TK_WHITESPACE, "\t", 2, 0), // Line Start
		NewToken(TK_STRING, `"{x} + {y} = {x + y}"`, 2, 1),
		NewToken(TK_DELIMITER, ",", 2, 22),
		NewToken(TK_NEWLINE, "\n", 2, 23),
		NewToken(TK_BLOCK_CLOSE, "}", 3, 0), // Line Start
	}

	checkMany(t, exps, f(in))
}

func S10_Loop(t *testing.T, f ScanFunc) {

	in := "LOOP i [i<5] {}"

	exps := []Token{
		NewToken(TK_LOOP, "LOOP", 0, 0),
		NewToken(TK_WHITESPACE, " ", 0, 4),
		NewToken(TK_IDENTIFIER, "i", 0, 5),
		NewToken(TK_WHITESPACE, " ", 0, 6),
		NewToken(TK_GUARD_OPEN, "[", 0, 7),
		NewToken(TK_IDENTIFIER, "i", 0, 8),
		NewToken(TK_LESS_THAN, "<", 0, 9),
		NewToken(TK_NUMBER, "5", 0, 10),
		NewToken(TK_GUARD_CLOSE, "]", 0, 11),
		NewToken(TK_WHITESPACE, " ", 0, 12),
		NewToken(TK_BLOCK_OPEN, "{", 0, 13),
		NewToken(TK_BLOCK_CLOSE, "}", 0, 14),
	}

	checkMany(t, exps, f(in))
}

func S11_ModifyList(t *testing.T, f ScanFunc) {

	in := "x[3],x[>>]:1,99"

	exps := []Token{
		NewToken(TK_IDENTIFIER, "x", 0, 0),
		NewToken(TK_GUARD_OPEN, "[", 0, 1),
		NewToken(TK_NUMBER, "3", 0, 2),
		NewToken(TK_GUARD_CLOSE, "]", 0, 3),
		NewToken(TK_DELIMITER, ",", 0, 4),
		NewToken(TK_IDENTIFIER, "x", 0, 5),
		NewToken(TK_GUARD_OPEN, "[", 0, 6),
		NewToken(TK_LIST_END, ">>", 0, 7),
		NewToken(TK_GUARD_CLOSE, "]", 0, 9),
		NewToken(TK_ASSIGNMENT, ":", 0, 10),
		NewToken(TK_NUMBER, "1", 0, 11),
		NewToken(TK_DELIMITER, ",", 0, 12),
		NewToken(TK_NUMBER, "99", 0, 13),
	}

	checkMany(t, exps, f(in))
}

func S12_ForEach(t *testing.T, f ScanFunc) {

	in := "LOOP i,v,m<-list"

	exps := []Token{
		NewToken(TK_LOOP, "LOOP", 0, 0),
		NewToken(TK_WHITESPACE, " ", 0, 4),
		NewToken(TK_IDENTIFIER, "i", 0, 5),
		NewToken(TK_DELIMITER, ",", 0, 6),
		NewToken(TK_IDENTIFIER, "v", 0, 7),
		NewToken(TK_DELIMITER, ",", 0, 8),
		NewToken(TK_IDENTIFIER, "m", 0, 9),
		NewToken(TK_UPDATES, "<-", 0, 10),
		NewToken(TK_IDENTIFIER, "list", 0, 12),
	}

	checkMany(t, exps, f(in))
}
