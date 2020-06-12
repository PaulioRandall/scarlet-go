package tests

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func S1() (in string, expects []Token) {

	in = "x := 1"

	expects = []Token{
		NewToken(TK_IDENTIFIER, "x", 0, 0),
		NewToken(TK_WHITESPACE, " ", 0, 1),
		NewToken(TK_ASSIGNMENT, ":=", 0, 2),
		NewToken(TK_WHITESPACE, " ", 0, 4),
		NewToken(TK_NUMBER, "1", 0, 5),
	}

	return in, expects
}

func S2() (in string, expects []Token) {

	in = "x,y:=1,true"

	expects = []Token{
		NewToken(TK_IDENTIFIER, "x", 0, 0),
		NewToken(TK_DELIMITER, ",", 0, 1),
		NewToken(TK_IDENTIFIER, "y", 0, 2),
		NewToken(TK_ASSIGNMENT, ":=", 0, 3),
		NewToken(TK_NUMBER, "1", 0, 5),
		NewToken(TK_DELIMITER, ",", 0, 6),
		NewToken(TK_BOOL, "true", 0, 7),
	}

	return in, expects
}

func S3() (in string, expects []Token) {

	in = "[1<2] x:=true"

	expects = []Token{
		NewToken(TK_GUARD_OPEN, "[", 0, 0),
		NewToken(TK_NUMBER, "1", 0, 1),
		NewToken(TK_LESS_THAN, "<", 0, 2),
		NewToken(TK_NUMBER, "2", 0, 3),
		NewToken(TK_GUARD_CLOSE, "]", 0, 4),
		NewToken(TK_WHITESPACE, " ", 0, 5),
		NewToken(TK_IDENTIFIER, "x", 0, 6),
		NewToken(TK_ASSIGNMENT, ":=", 0, 7),
		NewToken(TK_BOOL, "true", 0, 9),
	}

	return in, expects
}

func S4() (in string, expects []Token) {

	in = "when a := 1 {\n" +
		"\tb: x:=true\n" +
		"\t[a==b]:\n" +
		"\t[true]: x:=true\n" +
		"}"

	expects = []Token{
		NewToken(TK_WHEN, "when", 0, 0), // Line start
		NewToken(TK_WHITESPACE, " ", 0, 4),
		NewToken(TK_IDENTIFIER, "a", 0, 5),
		NewToken(TK_WHITESPACE, " ", 0, 6),
		NewToken(TK_ASSIGNMENT, ":=", 0, 7),
		NewToken(TK_WHITESPACE, " ", 0, 9),
		NewToken(TK_NUMBER, "1", 0, 10),
		NewToken(TK_WHITESPACE, " ", 0, 11),
		NewToken(TK_BLOCK_OPEN, "{", 0, 12),
		NewToken(TK_NEWLINE, "\n", 0, 13), // Line start
		NewToken(TK_WHITESPACE, "\t", 1, 0),
		NewToken(TK_IDENTIFIER, "b", 1, 1),
		NewToken(TK_THEN, ":", 1, 2),
		NewToken(TK_WHITESPACE, " ", 1, 3),
		NewToken(TK_IDENTIFIER, "x", 1, 4),
		NewToken(TK_ASSIGNMENT, ":=", 1, 5),
		NewToken(TK_BOOL, "true", 1, 7),
		NewToken(TK_NEWLINE, "\n", 1, 11), // Line start
		NewToken(TK_WHITESPACE, "\t", 2, 0),
		NewToken(TK_GUARD_OPEN, "[", 2, 1),
		NewToken(TK_IDENTIFIER, "a", 2, 2),
		NewToken(TK_EQUAL, "==", 2, 3),
		NewToken(TK_IDENTIFIER, "b", 2, 5),
		NewToken(TK_GUARD_CLOSE, "]", 2, 6),
		NewToken(TK_THEN, ":", 2, 7),
		NewToken(TK_NEWLINE, "\n", 2, 8),
		NewToken(TK_WHITESPACE, "\t", 3, 0), // Line start
		NewToken(TK_GUARD_OPEN, "[", 3, 1),
		NewToken(TK_BOOL, "true", 3, 2),
		NewToken(TK_GUARD_CLOSE, "]", 3, 6),
		NewToken(TK_THEN, ":", 3, 7),
		NewToken(TK_WHITESPACE, " ", 3, 8),
		NewToken(TK_IDENTIFIER, "x", 3, 9),
		NewToken(TK_ASSIGNMENT, ":=", 3, 10),
		NewToken(TK_BOOL, "true", 3, 12),
		NewToken(TK_NEWLINE, "\n", 3, 16),
		NewToken(TK_BLOCK_CLOSE, "}", 4, 0), // Line start
	}

	return in, expects
}

func S5() (in string, expects []Token) {

	in = "F(a,b,^c,^d)"

	expects = []Token{
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

	return in, expects
}

func S6() (in string, expects []Token) {

	in = "xyz(a,b)"

	expects = []Token{
		NewToken(TK_IDENTIFIER, "xyz", 0, 0),
		NewToken(TK_PAREN_OPEN, "(", 0, 3),
		NewToken(TK_IDENTIFIER, "a", 0, 4),
		NewToken(TK_DELIMITER, ",", 0, 5),
		NewToken(TK_IDENTIFIER, "b", 0, 6),
		NewToken(TK_PAREN_CLOSE, ")", 0, 7),
	}

	return in, expects
}

func S7() (in string, expects []Token) {

	in = "1+2-3*4/5%6"

	expects = []Token{
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

	return in, expects
}

func S8() (in string, expects []Token) {

	in = "{\n" +
		"\tx:=1\n" +
		"\ty:=2\n" +
		"}"

	expects = []Token{
		NewToken(TK_BLOCK_OPEN, "{", 0, 0), // Line Start
		NewToken(TK_NEWLINE, "\n", 0, 1),
		NewToken(TK_WHITESPACE, "\t", 1, 0), // Line Start
		NewToken(TK_IDENTIFIER, "x", 1, 1),
		NewToken(TK_ASSIGNMENT, ":=", 1, 2),
		NewToken(TK_NUMBER, "1", 1, 4),
		NewToken(TK_NEWLINE, "\n", 1, 5),
		NewToken(TK_WHITESPACE, "\t", 2, 0), // Line Start
		NewToken(TK_IDENTIFIER, "y", 2, 1),
		NewToken(TK_ASSIGNMENT, ":=", 2, 2),
		NewToken(TK_NUMBER, "2", 2, 4),
		NewToken(TK_NEWLINE, "\n", 2, 5),
		NewToken(TK_BLOCK_CLOSE, "}", 3, 0), // Line Start
	}

	return in, expects
}

func S10() (in string, expects []Token) {

	in = "loop i := 0 [i<5] {}"

	expects = []Token{
		NewToken(TK_LOOP, "loop", 0, 0),
		NewToken(TK_WHITESPACE, " ", 0, 4),
		NewToken(TK_IDENTIFIER, "i", 0, 5),
		NewToken(TK_WHITESPACE, " ", 0, 6),
		NewToken(TK_ASSIGNMENT, ":=", 0, 7),
		NewToken(TK_WHITESPACE, " ", 0, 9),
		NewToken(TK_NUMBER, "0", 0, 10),
		NewToken(TK_WHITESPACE, " ", 0, 11),
		NewToken(TK_GUARD_OPEN, "[", 0, 12),
		NewToken(TK_IDENTIFIER, "i", 0, 13),
		NewToken(TK_LESS_THAN, "<", 0, 14),
		NewToken(TK_NUMBER, "5", 0, 15),
		NewToken(TK_GUARD_CLOSE, "]", 0, 16),
		NewToken(TK_WHITESPACE, " ", 0, 17),
		NewToken(TK_BLOCK_OPEN, "{", 0, 18),
		NewToken(TK_BLOCK_CLOSE, "}", 0, 19),
	}

	return in, expects
}
