package tests

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func S1() (string, []Token) {

	in := "x := 1"

	exps := []Token{
		NewToken(IDENTIFIER, "x", 0, 0),
		NewToken(WHITESPACE, " ", 0, 1),
		NewToken(ASSIGN, ":=", 0, 2),
		NewToken(WHITESPACE, " ", 0, 4),
		NewToken(NUMBER, "1", 0, 5),
	}

	return in, exps
}

func S2() (string, []Token) {

	in := "x,y:=1,true"

	exps := []Token{
		NewToken(IDENTIFIER, "x", 0, 0),
		NewToken(DELIMITER, ",", 0, 1),
		NewToken(IDENTIFIER, "y", 0, 2),
		NewToken(ASSIGN, ":=", 0, 3),
		NewToken(NUMBER, "1", 0, 5),
		NewToken(DELIMITER, ",", 0, 6),
		NewToken(BOOL, "true", 0, 7),
	}

	return in, exps
}

func S3() (string, []Token) {

	in := "[1<2] x:=true"

	exps := []Token{
		NewToken(GUARD_OPEN, "[", 0, 0),
		NewToken(NUMBER, "1", 0, 1),
		NewToken(LESS_THAN, "<", 0, 2),
		NewToken(NUMBER, "2", 0, 3),
		NewToken(GUARD_CLOSE, "]", 0, 4),
		NewToken(WHITESPACE, " ", 0, 5),
		NewToken(IDENTIFIER, "x", 0, 6),
		NewToken(ASSIGN, ":=", 0, 7),
		NewToken(BOOL, "true", 0, 9),
	}

	return in, exps
}

func S4() (string, []Token) {

	in := "match {\n" +
		"\t[false] x:=false\n" +
		"\t[true] x:=true\n" +
		"}"

	exps := []Token{
		NewToken(MATCH, "match", 0, 0), // Line start
		NewToken(WHITESPACE, " ", 0, 5),
		NewToken(BLOCK_OPEN, "{", 0, 6),
		NewToken(NEWLINE, "\n", 0, 7), // Line start
		NewToken(WHITESPACE, "\t", 1, 0),
		NewToken(GUARD_OPEN, "[", 1, 1),
		NewToken(BOOL, "false", 1, 2),
		NewToken(GUARD_CLOSE, "]", 1, 7),
		NewToken(WHITESPACE, " ", 1, 8),
		NewToken(IDENTIFIER, "x", 1, 9),
		NewToken(ASSIGN, ":=", 1, 10),
		NewToken(BOOL, "false", 1, 12),
		NewToken(NEWLINE, "\n", 1, 17),
		NewToken(WHITESPACE, "\t", 2, 0), // Line start
		NewToken(GUARD_OPEN, "[", 2, 1),
		NewToken(BOOL, "true", 2, 2),
		NewToken(GUARD_CLOSE, "]", 2, 6),
		NewToken(WHITESPACE, " ", 2, 7),
		NewToken(IDENTIFIER, "x", 2, 8),
		NewToken(ASSIGN, ":=", 2, 9),
		NewToken(BOOL, "true", 2, 11),
		NewToken(NEWLINE, "\n", 2, 15),
		NewToken(BLOCK_CLOSE, "}", 3, 0), // Line start
	}

	return in, exps
}

func S5() (string, []Token) {

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

	return in, exps
}

func S6() (string, []Token) {

	in := "xyz(a,b)"

	exps := []Token{
		NewToken(IDENTIFIER, "xyz", 0, 0),
		NewToken(PAREN_OPEN, "(", 0, 3),
		NewToken(IDENTIFIER, "a", 0, 4),
		NewToken(DELIMITER, ",", 0, 5),
		NewToken(IDENTIFIER, "b", 0, 6),
		NewToken(PAREN_CLOSE, ")", 0, 7),
	}

	return in, exps
}

func S7() (string, []Token) {

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

	return in, exps
}

func S8() (string, []Token) {

	in := "{\n" +
		"\tx:=1\n" +
		"\ty:=2\n" +
		"}"

	exps := []Token{
		NewToken(BLOCK_OPEN, "{", 0, 0), // Line Start
		NewToken(NEWLINE, "\n", 0, 1),
		NewToken(WHITESPACE, "\t", 1, 0), // Line Start
		NewToken(IDENTIFIER, "x", 1, 1),
		NewToken(ASSIGN, ":=", 1, 2),
		NewToken(NUMBER, "1", 1, 4),
		NewToken(NEWLINE, "\n", 1, 5),
		NewToken(WHITESPACE, "\t", 2, 0), // Line Start
		NewToken(IDENTIFIER, "y", 2, 1),
		NewToken(ASSIGN, ":=", 2, 2),
		NewToken(NUMBER, "2", 2, 4),
		NewToken(NEWLINE, "\n", 2, 5),
		NewToken(BLOCK_CLOSE, "}", 3, 0), // Line Start
	}

	return in, exps
}

func S9() (string, []Token) {

	in := "list {\n" +
		"\t" + `"There's a snake in my boot",` + "\n" +
		"\t" + `"{x} + {y} = {x + y}"` + ",\n" +
		"}"

	exps := []Token{
		NewToken(LIST, "list", 0, 0),
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

	return in, exps
}

func S10() (string, []Token) {

	in := "loop i := 0 [i<5] {}"

	exps := []Token{
		NewToken(LOOP, "loop", 0, 0),
		NewToken(WHITESPACE, " ", 0, 4),
		NewToken(IDENTIFIER, "i", 0, 5),
		NewToken(WHITESPACE, " ", 0, 6),
		NewToken(ASSIGN, ":=", 0, 7),
		NewToken(WHITESPACE, " ", 0, 9),
		NewToken(NUMBER, "0", 0, 10),
		NewToken(WHITESPACE, " ", 0, 11),
		NewToken(GUARD_OPEN, "[", 0, 12),
		NewToken(IDENTIFIER, "i", 0, 13),
		NewToken(LESS_THAN, "<", 0, 14),
		NewToken(NUMBER, "5", 0, 15),
		NewToken(GUARD_CLOSE, "]", 0, 16),
		NewToken(WHITESPACE, " ", 0, 17),
		NewToken(BLOCK_OPEN, "{", 0, 18),
		NewToken(BLOCK_CLOSE, "}", 0, 19),
	}

	return in, exps
}
