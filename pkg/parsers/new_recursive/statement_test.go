package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var testFunc func(Factory, []Token) ([]Statement, error) = ParseStatements
var testFactory Factory = NewFactory()

func Test_S1(t *testing.T) {

	// a

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(TERMINATOR, ""),
	}

	exp := Identifier{tok(IDENTIFIER, "a")}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S3(t *testing.T) {

	// TRUE

	given := []Token{
		tok(BOOL, "TRUE"),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(BOOL, "TRUE")}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S4(t *testing.T) {

	// 1

	given := []Token{
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(NUMBER, "1")}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S5(t *testing.T) {

	// "abc"

	given := []Token{
		tok(STRING, "abc"),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(STRING, "abc")}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S6(t *testing.T) {

	// a: 1

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(ASSIGN, ":"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := AssignmentBlock{
		[]Assignment{
			Assignment{
				Target: tok(IDENTIFIER, "a"),
				Source: testFactory.NewLiteral(tok(NUMBER, "1")),
			},
		},
	}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S7(t *testing.T) {

	// a, b, c: 1, TRUE, "abc"

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "c"),
		tok(ASSIGN, ":"),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(BOOL, "TRUE"),
		tok(DELIMITER, ","),
		tok(STRING, "abc"),
		tok(TERMINATOR, ""),
	}

	exp := AssignmentBlock{
		[]Assignment{
			Assignment{
				Target: tok(IDENTIFIER, "a"),
				Source: testFactory.NewLiteral(tok(NUMBER, "1")),
			},
			Assignment{
				Target: tok(IDENTIFIER, "b"),
				Source: testFactory.NewLiteral(tok(BOOL, "TRUE")),
			},
			Assignment{
				Target: tok(IDENTIFIER, "c"),
				Source: testFactory.NewLiteral(tok(STRING, "abc")),
			},
		},
	}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S8(t *testing.T) {

	// -2

	given := []Token{
		tok(SUBTRACT, "-"),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
	}

	exp := Negation{
		Literal{tok(NUMBER, "2")},
	}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S9(t *testing.T) {

	// LIST {}

	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	exp := List{
		Open:  tok(BLOCK_OPEN, "{"),
		Items: []Expression{},
		Close: tok(BLOCK_CLOSE, "}"),
	}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S10(t *testing.T) {

	// LIST {1,TRUE,"abc"}

	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(BOOL, "TRUE"),
		tok(DELIMITER, ","),
		tok(STRING, "abc"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	exp := List{
		Open: tok(BLOCK_OPEN, "{"),
		Items: []Expression{
			Literal{tok(NUMBER, "1")},
			Literal{tok(BOOL, "TRUE")},
			Literal{tok(STRING, "abc")},
		},
		Close: tok(BLOCK_CLOSE, "}"),
	}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S11(t *testing.T) {

	// LIST {
	// 	1,
	// 	TRUE,
	// 	"abc",
	// }

	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(TERMINATOR, "\n"),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(TERMINATOR, "\n"),
		tok(BOOL, "TRUE"),
		tok(DELIMITER, ","),
		tok(TERMINATOR, "\n"),
		tok(STRING, "abc"),
		tok(DELIMITER, ","),
		tok(TERMINATOR, "\n"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	exp := List{
		Open: tok(BLOCK_OPEN, "{"),
		Items: []Expression{
			Literal{tok(NUMBER, "1")},
			Literal{tok(BOOL, "TRUE")},
			Literal{tok(STRING, "abc")},
		},
		Close: tok(BLOCK_CLOSE, "}"),
	}

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_F1(t *testing.T) {

	// Assignment token is never at the start of a statement
	// :

	given := []Token{
		tok(ASSIGN, ":"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F2(t *testing.T) {

	// Not enough expressions
	// :

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(ASSIGN, ":"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F3(t *testing.T) {

	// Not enough identifiers
	// :

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(ASSIGN, ":"),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F4(t *testing.T) {

	// Missing assignment token
	// :

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F5(t *testing.T) {

	// Missing identifier delimiter token
	// :

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(IDENTIFIER, "b"),
		tok(ASSIGN, ":"),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F6(t *testing.T) {

	// Missing expression delimiter token
	// :

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(ASSIGN, ":"),
		tok(NUMBER, "1"),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F7(t *testing.T) {

	// Negation without an expression
	// -

	given := []Token{
		tok(SUBTRACT, "-"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F8(t *testing.T) {

	// List items ends in delimiter without a terminator
	// LIST {1,}

	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F9(t *testing.T) {

	// List items start with a delimiter
	// LIST {1,}

	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(DELIMITER, ","),
		tok(NUMBER, "1"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F10(t *testing.T) {

	// Last list item doesn't have a delimiter when a terminator follows
	// LIST {1,}

	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, "\n"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}
