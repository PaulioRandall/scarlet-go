package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func S1_Assignment(t *testing.T, f ScanFunc) {

	in := "x : 1"

	exps := []Token{
		NewToken(IDENTIFIER, "x", 0, 0),
		NewToken(WHITESPACE, " ", 0, 1),
		NewToken(ASSIGN, ":", 0, 2),
		NewToken(WHITESPACE, " ", 0, 3),
		NewToken(NUMBER, "1", 0, 4),
	}

	checkMany(t, exps, f(in))
}

func S2_MultiAssignment(t *testing.T, f ScanFunc) {

	in := "x,y:1,TRUE"

	exps := []Token{
		NewToken(IDENTIFIER, "x", 0, 0),
		NewToken(DELIMITER, ",", 0, 1),
		NewToken(IDENTIFIER, "y", 0, 2),
		NewToken(ASSIGN, ":", 0, 3),
		NewToken(NUMBER, "1", 0, 4),
		NewToken(DELIMITER, ",", 0, 5),
		NewToken(BOOL, "TRUE", 0, 6),
	}

	checkMany(t, exps, f(in))
}

func S3_GuardBlock(t *testing.T, f ScanFunc) {

	in := "[1<2] x:TRUE"

	exps := []Token{
		NewToken(GUARD_OPEN, "[", 0, 0),
		NewToken(NUMBER, "1", 0, 1),
		NewToken(LESS_THAN, "<", 0, 2),
		NewToken(NUMBER, "2", 0, 3),
		NewToken(GUARD_CLOSE, "]", 0, 4),
		NewToken(WHITESPACE, " ", 0, 5),
		NewToken(IDENTIFIER, "x", 0, 6),
		NewToken(ASSIGN, ":", 0, 7),
		NewToken(BOOL, "TRUE", 0, 8),
	}

	checkMany(t, exps, f(in))
}

func S4_MatchBlock(t *testing.T, f ScanFunc) {

	in := "MATCH {\n" +
		"\t[FALSE] x:FALSE\n" +
		"\t[TRUE] x:TRUE\n" +
		"}"

	exps := []Token{
		NewToken(MATCH, "MATCH", 0, 0), // Line start
		NewToken(WHITESPACE, " ", 0, 5),
		NewToken(BLOCK_OPEN, "{", 0, 6),
		NewToken(NEWLINE, "\n", 0, 7), // Line start
		NewToken(WHITESPACE, "\t", 1, 0),
		NewToken(GUARD_OPEN, "[", 1, 1),
		NewToken(BOOL, "FALSE", 1, 2),
		NewToken(GUARD_CLOSE, "]", 1, 7),
		NewToken(WHITESPACE, " ", 1, 8),
		NewToken(IDENTIFIER, "x", 1, 9),
		NewToken(ASSIGN, ":", 1, 10),
		NewToken(BOOL, "FALSE", 1, 11),
		NewToken(NEWLINE, "\n", 1, 16),
		NewToken(WHITESPACE, "\t", 2, 0), // Line start
		NewToken(GUARD_OPEN, "[", 2, 1),
		NewToken(BOOL, "TRUE", 2, 2),
		NewToken(GUARD_CLOSE, "]", 2, 6),
		NewToken(WHITESPACE, " ", 2, 7),
		NewToken(IDENTIFIER, "x", 2, 8),
		NewToken(ASSIGN, ":", 2, 9),
		NewToken(BOOL, "TRUE", 2, 10),
		NewToken(NEWLINE, "\n", 2, 14),
		NewToken(BLOCK_CLOSE, "}", 3, 0), // Line start
	}

	checkMany(t, exps, f(in))
}

func S5_FuncDef(t *testing.T, f ScanFunc) {

	in := "F(a,b,^c,^d)"

	exps := []Token{
		NewToken(FUNC, "F", 0, 0),
		NewToken(PAREN_OPEN, "(", 0, 1),
		NewToken(IDENTIFIER, "a", 0, 2),
		NewToken(DELIMITER, ",", 0, 3),
		NewToken(IDENTIFIER, "b", 0, 4),
		NewToken(DELIMITER, ",", 0, 5),
		NewToken(OUTPUT, "^", 0, 6),
		NewToken(IDENTIFIER, "c", 0, 7),
		NewToken(DELIMITER, ",", 0, 8),
		NewToken(OUTPUT, "^", 0, 9),
		NewToken(IDENTIFIER, "d", 0, 10),
		NewToken(PAREN_CLOSE, ")", 0, 11),
	}

	checkMany(t, exps, f(in))
}

func S6_FuncCall(t *testing.T, f ScanFunc) {

	in := "xyz(a,b)"

	exps := []Token{
		NewToken(IDENTIFIER, "xyz", 0, 0),
		NewToken(PAREN_OPEN, "(", 0, 3),
		NewToken(IDENTIFIER, "a", 0, 4),
		NewToken(DELIMITER, ",", 0, 5),
		NewToken(IDENTIFIER, "b", 0, 6),
		NewToken(PAREN_CLOSE, ")", 0, 7),
	}

	checkMany(t, exps, f(in))
}

func S7_Expression(t *testing.T, f ScanFunc) {

	in := "1+2-3*4/5%6"

	exps := []Token{
		NewToken(NUMBER, "1", 0, 0),
		NewToken(ADD, "+", 0, 1),
		NewToken(NUMBER, "2", 0, 2),
		NewToken(SUBTRACT, "-", 0, 3),
		NewToken(NUMBER, "3", 0, 4),
		NewToken(MULTIPLY, "*", 0, 5),
		NewToken(NUMBER, "4", 0, 6),
		NewToken(DIVIDE, "/", 0, 7),
		NewToken(NUMBER, "5", 0, 8),
		NewToken(REMAINDER, "%", 0, 9),
		NewToken(NUMBER, "6", 0, 10),
	}

	checkMany(t, exps, f(in))
}

func S8_Block(t *testing.T, f ScanFunc) {

	in := "{\n" +
		"\tx:1\n" +
		"\ty:2\n" +
		"}"

	exps := []Token{
		NewToken(BLOCK_OPEN, "{", 0, 0), // Line Start
		NewToken(NEWLINE, "\n", 0, 1),
		NewToken(WHITESPACE, "\t", 1, 0), // Line Start
		NewToken(IDENTIFIER, "x", 1, 1),
		NewToken(ASSIGN, ":", 1, 2),
		NewToken(NUMBER, "1", 1, 3),
		NewToken(NEWLINE, "\n", 1, 4),
		NewToken(WHITESPACE, "\t", 2, 0), // Line Start
		NewToken(IDENTIFIER, "y", 2, 1),
		NewToken(ASSIGN, ":", 2, 2),
		NewToken(NUMBER, "2", 2, 3),
		NewToken(NEWLINE, "\n", 2, 4),
		NewToken(BLOCK_CLOSE, "}", 3, 0), // Line Start
	}

	checkMany(t, exps, f(in))
}

func S9_List(t *testing.T, f ScanFunc) {

	in := "LIST {\n" +
		"\t" + `"There's a snake in my boot",` + "\n" +
		"\t" + `"{x} + {y} = {x + y}"` + ",\n" +
		"}"

	exps := []Token{
		NewToken(LIST, "LIST", 0, 0),
		NewToken(WHITESPACE, " ", 0, 4),
		NewToken(BLOCK_OPEN, "{", 0, 5),
		NewToken(NEWLINE, "\n", 0, 6),
		NewToken(WHITESPACE, "\t", 1, 0), // Line Start
		NewToken(STRING, `"There's a snake in my boot"`, 1, 1),
		NewToken(DELIMITER, ",", 1, 29),
		NewToken(NEWLINE, "\n", 1, 30),
		NewToken(WHITESPACE, "\t", 2, 0), // Line Start
		NewToken(STRING, `"{x} + {y} = {x + y}"`, 2, 1),
		NewToken(DELIMITER, ",", 2, 22),
		NewToken(NEWLINE, "\n", 2, 23),
		NewToken(BLOCK_CLOSE, "}", 3, 0), // Line Start
	}

	checkMany(t, exps, f(in))
}

func S10_Loop(t *testing.T, f ScanFunc) {

	in := "LOOP i [i<5] {}"

	exps := []Token{
		NewToken(LOOP, "LOOP", 0, 0),
		NewToken(WHITESPACE, " ", 0, 4),
		NewToken(IDENTIFIER, "i", 0, 5),
		NewToken(WHITESPACE, " ", 0, 6),
		NewToken(GUARD_OPEN, "[", 0, 7),
		NewToken(IDENTIFIER, "i", 0, 8),
		NewToken(LESS_THAN, "<", 0, 9),
		NewToken(NUMBER, "5", 0, 10),
		NewToken(GUARD_CLOSE, "]", 0, 11),
		NewToken(WHITESPACE, " ", 0, 12),
		NewToken(BLOCK_OPEN, "{", 0, 13),
		NewToken(BLOCK_CLOSE, "}", 0, 14),
	}

	checkMany(t, exps, f(in))
}

func S11_ModifyList(t *testing.T, f ScanFunc) {

	in := "x[3],x[>>]:1,99"

	exps := []Token{
		NewToken(IDENTIFIER, "x", 0, 0),
		NewToken(GUARD_OPEN, "[", 0, 1),
		NewToken(NUMBER, "3", 0, 2),
		NewToken(GUARD_CLOSE, "]", 0, 3),
		NewToken(DELIMITER, ",", 0, 4),
		NewToken(IDENTIFIER, "x", 0, 5),
		NewToken(GUARD_OPEN, "[", 0, 6),
		NewToken(LIST_END, ">>", 0, 7),
		NewToken(GUARD_CLOSE, "]", 0, 9),
		NewToken(ASSIGN, ":", 0, 10),
		NewToken(NUMBER, "1", 0, 11),
		NewToken(DELIMITER, ",", 0, 12),
		NewToken(NUMBER, "99", 0, 13),
	}

	checkMany(t, exps, f(in))
}

func S12_ForEach(t *testing.T, f ScanFunc) {

	in := "LOOP i,v,m<-list"

	exps := []Token{
		NewToken(LOOP, "LOOP", 0, 0),
		NewToken(WHITESPACE, " ", 0, 4),
		NewToken(IDENTIFIER, "i", 0, 5),
		NewToken(DELIMITER, ",", 0, 6),
		NewToken(IDENTIFIER, "v", 0, 7),
		NewToken(DELIMITER, ",", 0, 8),
		NewToken(IDENTIFIER, "m", 0, 9),
		NewToken(UPDATES, "<-", 0, 10),
		NewToken(IDENTIFIER, "list", 0, 12),
	}

	checkMany(t, exps, f(in))
}
