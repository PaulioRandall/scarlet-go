package tests

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func S1() (in string, expects []Token) {

	in = "x := 1"

	expects = []Token{
		NewToken(IDENTIFIER, "x", 0, 0),
		NewToken(WHITESPACE, " ", 0, 1),
		NewToken(ASSIGN, ":=", 0, 2),
		NewToken(WHITESPACE, " ", 0, 4),
		NewToken(NUMBER, "1", 0, 5),
	}

	return in, expects
}

func S2() (in string, expects []Token) {

	in = "x,y:=1,true"

	expects = []Token{
		NewToken(IDENTIFIER, "x", 0, 0),
		NewToken(DELIMITER, ",", 0, 1),
		NewToken(IDENTIFIER, "y", 0, 2),
		NewToken(ASSIGN, ":=", 0, 3),
		NewToken(NUMBER, "1", 0, 5),
		NewToken(DELIMITER, ",", 0, 6),
		NewToken(BOOL, "true", 0, 7),
	}

	return in, expects
}

func S3() (in string, expects []Token) {

	in = "[1<2] x:=true"

	expects = []Token{
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

	return in, expects
}

func S4() (in string, expects []Token) {

	in = "match abc {\n" +
		"\t1 -> x:=true\n" +
		"\t[false] -> x:=false\n" +
		"\t[true] -> x:=true\n" +
		"}"

	expects = []Token{
		NewToken(MATCH, "match", 0, 0), // Line start
		NewToken(WHITESPACE, " ", 0, 5),
		NewToken(IDENTIFIER, "abc", 0, 6),
		NewToken(WHITESPACE, " ", 0, 9),
		NewToken(BLOCK_OPEN, "{", 0, 10),
		NewToken(NEWLINE, "\n", 0, 11), // Line start
		NewToken(WHITESPACE, "\t", 1, 0),
		NewToken(NUMBER, "1", 1, 1),
		NewToken(WHITESPACE, " ", 1, 2),
		NewToken(THEN, "->", 1, 3),
		NewToken(WHITESPACE, " ", 1, 5),
		NewToken(IDENTIFIER, "x", 1, 6),
		NewToken(ASSIGN, ":=", 1, 7),
		NewToken(BOOL, "true", 1, 9),
		NewToken(NEWLINE, "\n", 1, 13), // Line start
		NewToken(WHITESPACE, "\t", 2, 0),
		NewToken(GUARD_OPEN, "[", 2, 1),
		NewToken(BOOL, "false", 2, 2),
		NewToken(GUARD_CLOSE, "]", 2, 7),
		NewToken(WHITESPACE, " ", 2, 8),
		NewToken(THEN, "->", 2, 9),
		NewToken(WHITESPACE, " ", 2, 11),
		NewToken(IDENTIFIER, "x", 2, 12),
		NewToken(ASSIGN, ":=", 2, 13),
		NewToken(BOOL, "false", 2, 15),
		NewToken(NEWLINE, "\n", 2, 20),
		NewToken(WHITESPACE, "\t", 3, 0), // Line start
		NewToken(GUARD_OPEN, "[", 3, 1),
		NewToken(BOOL, "true", 3, 2),
		NewToken(GUARD_CLOSE, "]", 3, 6),
		NewToken(WHITESPACE, " ", 3, 7),
		NewToken(THEN, "->", 3, 8),
		NewToken(WHITESPACE, " ", 3, 10),
		NewToken(IDENTIFIER, "x", 3, 11),
		NewToken(ASSIGN, ":=", 3, 12),
		NewToken(BOOL, "true", 3, 14),
		NewToken(NEWLINE, "\n", 3, 18),
		NewToken(BLOCK_CLOSE, "}", 4, 0), // Line start
	}

	return in, expects
}

func S5() (in string, expects []Token) {

	in = "F(a,b,^c,^d)"

	expects = []Token{
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

	return in, expects
}

func S6() (in string, expects []Token) {

	in = "xyz(a,b)"

	expects = []Token{
		NewToken(IDENTIFIER, "xyz", 0, 0),
		NewToken(PAREN_OPEN, "(", 0, 3),
		NewToken(IDENTIFIER, "a", 0, 4),
		NewToken(DELIMITER, ",", 0, 5),
		NewToken(IDENTIFIER, "b", 0, 6),
		NewToken(PAREN_CLOSE, ")", 0, 7),
	}

	return in, expects
}

func S7() (in string, expects []Token) {

	in = "1+2-3*4/5%6"

	expects = []Token{
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

	return in, expects
}

func S8() (in string, expects []Token) {

	in = "{\n" +
		"\tx:=1\n" +
		"\ty:=2\n" +
		"}"

	expects = []Token{
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

	return in, expects
}

func S10() (in string, expects []Token) {

	in = "loop i := 0 [i<5] {}"

	expects = []Token{
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

	return in, expects
}
