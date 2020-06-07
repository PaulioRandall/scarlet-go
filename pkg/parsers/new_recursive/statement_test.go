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
		tok(ASSIGN, ":="),
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
		tok(ASSIGN, ":="),
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
	// THEN only the parsed function is returned

	// f := F() {}
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	f := testFactory.NewFunction(
		tok(FUNC, "F"),
		testFactory.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		testFactory.NewBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Statement{},
		),
	)

	exp := testFactory.NewNonWrappedBlock(
		[]Statement{
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S20(t *testing.T) {

	// GIVEN a function
	// WITH one input parameter
	// AND no statements in the body
	// THEN only the parsed function is returned

	// f := F(a) {}
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	f := testFactory.NewFunction(
		tok(FUNC, "F"),
		testFactory.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{
				tok(IDENTIFIER, "a"),
			},
			[]Token{},
		),
		testFactory.NewBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Statement{},
		),
	)

	exp := testFactory.NewNonWrappedBlock(
		[]Statement{
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S21(t *testing.T) {

	// GIVEN a function
	// WITH one output parameter
	// AND no statements in the body
	// THEN only the parsed function is returned

	// f := F(^a) {}
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(OUTPUT, "^"),
		tok(IDENTIFIER, "a"),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	f := testFactory.NewFunction(
		tok(FUNC, "F"),
		testFactory.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{
				tok(IDENTIFIER, "a"),
			},
		),
		testFactory.NewBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Statement{},
		),
	)

	exp := testFactory.NewNonWrappedBlock(
		[]Statement{
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S22(t *testing.T) {

	// GIVEN a function
	// WITH multiple parameters
	// AND no statements in the body
	// THEN only the parsed function is returned

	// f := F(a, b, ^c, ^d) {}
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(DELIMITER, ","),
		tok(OUTPUT, "^"),
		tok(IDENTIFIER, "c"),
		tok(DELIMITER, ","),
		tok(OUTPUT, "^"),
		tok(IDENTIFIER, "d"),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	f := testFactory.NewFunction(
		tok(FUNC, "F"),
		testFactory.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{
				tok(IDENTIFIER, "a"),
				tok(IDENTIFIER, "b"),
			},
			[]Token{
				tok(IDENTIFIER, "c"),
				tok(IDENTIFIER, "d"),
			},
		),
		testFactory.NewBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Statement{},
		),
	)

	exp := testFactory.NewNonWrappedBlock(
		[]Statement{
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S23(t *testing.T) {

	// GIVEN a function
	// WITH no parameters
	// AND a statement in the body
	// THEN only the parsed function is returned

	// f := F() {a := 1}
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(IDENTIFIER, "a"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	body := testFactory.NewBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Statement{
			testFactory.NewNonWrappedBlock(
				[]Statement{
					testFactory.NewAssignment(
						testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
						testFactory.NewLiteral(tok(NUMBER, "1")),
					),
				},
			),
		},
	)

	f := testFactory.NewFunction(
		tok(FUNC, "F"),
		testFactory.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		body,
	)

	exp := testFactory.NewNonWrappedBlock(
		[]Statement{
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S24(t *testing.T) {

	// GIVEN a function
	// WITH input and output parameters over several lines
	// AND a statement in the body with leading and trailing linefeeds
	// THEN only the parsed function is returned

	// f := F(
	// a,
	// ^b,
	// ) {
	// a: b
	// }
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(TERMINATOR, "\n"),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(TERMINATOR, "\n"),
		tok(OUTPUT, "^"),
		tok(IDENTIFIER, "b"),
		tok(DELIMITER, ","),
		tok(TERMINATOR, "\n"),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(TERMINATOR, "\n"),
		tok(IDENTIFIER, "a"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, "\n"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	body := testFactory.NewBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Statement{
			testFactory.NewNonWrappedBlock(
				[]Statement{
					testFactory.NewAssignment(
						testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
						testFactory.NewLiteral(tok(NUMBER, "1")),
					),
				},
			),
		},
	)

	f := testFactory.NewFunction(
		tok(FUNC, "F"),
		testFactory.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{
				tok(IDENTIFIER, "a"),
			},
			[]Token{
				tok(IDENTIFIER, "b"),
			},
		),
		body,
	)

	exp := testFactory.NewNonWrappedBlock(
		[]Statement{
			testFactory.NewAssignment(
				testFactory.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S25(t *testing.T) {

	// GIVEN a simple addition
	// THEN a single parsed operation is expected

	// a + 1
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(ADD, "+"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewOperation(
		tok(ADD, "+"),
		testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
		testFactory.NewLiteral(tok(NUMBER, "1")),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S26(t *testing.T) {

	// GIVEN a simple subtraction
	// THEN a single parsed operation is expected

	// a - 1
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewOperation(
		tok(SUBTRACT, "-"),
		testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
		testFactory.NewLiteral(tok(NUMBER, "1")),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S27(t *testing.T) {

	// GIVEN a simple multiplication
	// THEN a single parsed operation is expected

	// a * 1
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewOperation(
		tok(MULTIPLY, "*"),
		testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
		testFactory.NewLiteral(tok(NUMBER, "1")),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S28(t *testing.T) {

	// GIVEN a simple division
	// THEN a single parsed operation is expected

	// a / 1
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(DIVIDE, "/"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewOperation(
		tok(DIVIDE, "/"),
		testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
		testFactory.NewLiteral(tok(NUMBER, "1")),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S29(t *testing.T) {

	// GIVEN a simple division
	// THEN a single parsed operation is expected

	// a % 1
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewOperation(
		tok(REMAINDER, "%"),
		testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
		testFactory.NewLiteral(tok(NUMBER, "1")),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_S30(t *testing.T) {

	// GIVEN a simple addition
	// WITH a negated right operand
	// THEN a single parsed operation is expected
	// WITH an parsed negation a the right operand

	// a + -1
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(ADD, "+"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := testFactory.NewOperation(
		tok(ADD, "+"),
		testFactory.NewIdentifier(tok(IDENTIFIER, "a")),
		testFactory.NewNegation(
			testFactory.NewLiteral(tok(NUMBER, "1")),
		),
	)

	act, e := testFunc(testFactory, given)
	expectOneStat(t, exp, act, e)
}

func Test_F1(t *testing.T) {

	// GIVEN an invalid statement or expression starting token
	// THEN parser returns error

	given := []Token{
		tok(ASSIGN, ":="),
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
		tok(ASSIGN, ":="),
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
		tok(ASSIGN, ":="),
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
		tok(ASSIGN, ":="),
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
		tok(ASSIGN, ":="),
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

func Test_F13(t *testing.T) {

	// GIVEN an operation
	// WITHOUT a right operand
	// THEN parser returns error

	// x +
	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(ADD, "+"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

func Test_F14(t *testing.T) {

	// GIVEN an operation
	// WITH two operators
	// THEN parser returns error

	// x + +
	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(ADD, "+"),
		tok(ADD, "+"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}
