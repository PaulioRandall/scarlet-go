package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var testFunc func(Factory, []Token) ([]Statement, error) = ParseStatements
var testFactory Factory = NewFactory()

func Test_S1(t *testing.T) {

	// GIVEN an identifier only
	// THEN only an identifier expression is returned

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

	// GIVEN a bool literal only
	// THEN only a bool literal expression is returned

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

	// GIVEN a number literal only
	// THEN only a number literal expression is returned

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

	// GIVEN a string literal
	// THEN only a string literal expression is returned

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

	// GIVEN an assignment
	// WITH only one identifier and one expression
	// THEN only the parsed assignment is returned

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

	// GIVEN an assignment
	// WITH multiple identifiers and expressions
	// AND there are an equal number of identifiers and expressions
	// THEN all and only the parsed assignments are returned

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

	// GIVEN a negation
	// WITH a following expression
	// THEN only the parsed negation expression is returned

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

	// GIVEN an empty list
	// THEN only the parsed empty list is returned

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

	// GIVEN an empty list
	// WITH a terminator after the block open
	// THEN only the parsed empty list is returned

	// LIST {}
	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(TERMINATOR, "\n"),
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

func Test_S11(t *testing.T) {

	// GIVEN a list
	// WITH multiple inline items
	// THEN only the parsed list is returned

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

func Test_S12(t *testing.T) {

	// GIVEN a list
	// WITH multiple items each ending in a delimiter and terminator
	// THEN only the parsed list is returned

	// LIST {
	// 	1,
	// 	TRUE,
	// 	"abc",
	// }
	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
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

	// GIVEN an assignment
	// WHEN a token not found at the start of a statement
	// THEN parser returns error

	given := []Token{
		tok(ASSIGN, ":"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F2(t *testing.T) {

	// GIVEN an assignment
	// WHEN not enough expressions are present to pair with identifiers
	// THEN parser returns error

	// a:
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(ASSIGN, ":"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F3(t *testing.T) {

	// GIVEN an assignment
	// WHEN not enough identifiers are present to pair with expressions
	// THEN parser returns error

	// a: 1, 2
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

	// GIVEN an assignment
	// WHEN the assignment token is missing
	// THEN parser returns error

	// a 1
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F5(t *testing.T) {

	// GIVEN an assignment
	// WHEN an identifier delimiter token is missing
	// THEN parser returns error

	// a b: 1, 2
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

	// GIVEN an assignment
	// WHEN an expression delimiter token is missing
	// THEN parser returns error

	// a, b: 1 2
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

	// GIVEN a negation prefix
	// WHEN an expression does follow
	// THEN parser returns error

	// -
	given := []Token{
		tok(SUBTRACT, "-"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F8(t *testing.T) {

	// GIVEN a list
	// WHEN the last item ends in a delimiter without a following terminator
	// THEN parser returns error

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

	// GIVEN a list
	// WHEN the first token is not an expression or block close
	// THEN parser returns error

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

	// GIVEN a list
	// WHEN the last item doesn't have a delimiter but a terminator follows
	// THEN parser returns error

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
