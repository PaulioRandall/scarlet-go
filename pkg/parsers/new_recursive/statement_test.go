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

	exp := testFactory.NewIdentifier(tok(IDENTIFIER, "a"))

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

	exp := testFactory.NewLiteral(tok(BOOL, "TRUE"))

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

	exp := testFactory.NewLiteral(tok(NUMBER, "1"))

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

	exp := testFactory.NewLiteral(tok(STRING, "abc"))

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

	exp := testFactory.NewNonWrappedBlock(
		[]Statement{
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
				testFactory.NewLiteral(tok(NUMBER, "1")),
			),
		},
	)

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

	exp := testFactory.NewNonWrappedBlock(
		[]Statement{
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
				testFactory.NewLiteral(tok(NUMBER, "1")),
			),
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "b")),
				testFactory.NewLiteral(tok(BOOL, "TRUE")),
			),
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "c")),
				testFactory.NewLiteral(tok(STRING, "abc")),
			),
		},
	)

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

	exp := testFactory.NewNegation(
		testFactory.NewLiteral(tok(NUMBER, "2")),
	)

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

	exp := testFactory.NewList(
		tok(BLOCK_OPEN, "{"),
		[]Expression{},
		tok(BLOCK_CLOSE, "}"),
	)

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

	exp := testFactory.NewList(
		tok(BLOCK_OPEN, "{"),
		[]Expression{},
		tok(BLOCK_CLOSE, "}"),
	)

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

	exp := testFactory.NewList(
		tok(BLOCK_OPEN, "{"),
		[]Expression{
			Literal{tok(NUMBER, "1")},
			Literal{tok(BOOL, "TRUE")},
			Literal{tok(STRING, "abc")},
		},
		tok(BLOCK_CLOSE, "}"),
	)

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

	exp := testFactory.NewList(
		tok(BLOCK_OPEN, "{"),
		[]Expression{
			Literal{tok(NUMBER, "1")},
			Literal{tok(BOOL, "TRUE")},
			Literal{tok(STRING, "abc")},
		},
		tok(BLOCK_CLOSE, "}"),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S13(t *testing.T) {

	// GIVEN a non-empty list
	// WITH a delimiter after the last item but without a following terminator
	// THEN only the parsed list is returned

	// LIST {1,}
	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewList(
		tok(BLOCK_OPEN, "{"),
		[]Expression{
			Literal{tok(NUMBER, "1")},
		},
		tok(BLOCK_CLOSE, "}"),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S14(t *testing.T) {

	// GIVEN a list identifier with a number literal index
	// THEN only the parsed list accessor is returned

	// abc[1]
	given := []Token{
		tok(IDENTIFIER, "abc"),
		tok(GUARD_OPEN, "["),
		tok(NUMBER, "1"),
		tok(GUARD_CLOSE, "]"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewListAccessor(
		Identifier{tok(IDENTIFIER, "abc")},
		Literal{tok(NUMBER, "1")},
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S15(t *testing.T) {

	// GIVEN only a comment
	// THEN no statements are returned

	// // abc
	given := []Token{
		tok(COMMENT, "// abc"),
		tok(TERMINATOR, ""),
	}

	exp := []Statement{}

	act, e := testFunc(testFactory, given)
	expectStats(t, exp, act, e)
}

func Test_S16(t *testing.T) {

	// GIVEN only whitespace
	// THEN no statements are returned

	given := []Token{
		tok(WHITESPACE, "    "),
		tok(TERMINATOR, ""),
	}

	exp := []Statement{}

	act, e := testFunc(testFactory, given)
	expectStats(t, exp, act, e)
}

func Test_S17(t *testing.T) {

	// GIVEN only one terminator
	// THEN no statements are returned

	given := []Token{
		tok(TERMINATOR, ""),
	}

	exp := []Statement{}

	act, e := testFunc(testFactory, given)
	expectStats(t, exp, act, e)
}

func Test_S18(t *testing.T) {

	// GIVEN only terminators
	// THEN no statements are returned

	given := []Token{
		tok(TERMINATOR, ""),
		tok(TERMINATOR, ""),
		tok(TERMINATOR, ""),
	}

	exp := []Statement{}

	act, e := testFunc(testFactory, given)
	expectStats(t, exp, act, e)
}

func Test_S19(t *testing.T) {

	// GIVEN a function
	// WITH no parameters
	// AND no statements in the body
	// THEN no statements are returned

	given := []Token{
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewFunction(
		tok(FUNC, "F"),
		testFactory.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		testFactory.NewBlock(
			tok(BLOCK_OPEN, "{"),
			[]Statement{},
			tok(BLOCK_CLOSE, "}"),
		),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_F1(t *testing.T) {

	// GIVEN an invalid statement or expression starting token
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
	// WITHOUT enough expressions
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
	// WITHOUT enough identifiers
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
	// WITH the assignment token missing
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
	// WITH an delimiter token missing from the assignment targets
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
	// WITH an delimiter token missing from the expressions
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
	// WITHOUT a following expression
	// THEN parser returns error

	// -
	given := []Token{
		tok(SUBTRACT, "-"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F9(t *testing.T) {

	// GIVEN a list
	// WITHOUT an expression or block close following the block open
	// THEN parser returns error

	// LIST {,1}
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
	// WITHOUT a delimiter after an expression but with a terminator
	// THEN parser returns error

	// LIST {1
	// }
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

func Test_F11(t *testing.T) {

	// GIVEN a list
	// WITHOUT a block open
	// THEN parser returns error

	// LIST 1}
	given := []Token{
		tok(LIST, "LIST"),
		tok(NUMBER, "1"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F12(t *testing.T) {

	// GIVEN a list
	// WITHOUT a block close
	// THEN parser returns error

	// LIST {1
	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}
