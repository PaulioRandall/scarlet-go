package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var testFunc func([]Token) ([]Expression, error) = ParseStatements

func Test_S1_1(t *testing.T) {

	quickSoloTokenTest := func(t *testing.T, exp Expression, tk Token) {

		var given []Token
		given = append(given, tk)
		given = append(given, tok(TERMINATOR, ""))

		act, e := testFunc(given)
		expectOneStat(t, exp, act, e)
	}

	// GIVEN an identifier only
	// THEN only an identifier expression is returned

	// a
	quickSoloTokenTest(t,
		newIdentifier(tok(IDENTIFIER, "a")),
		tok(IDENTIFIER, "a"),
	)

	// true
	quickSoloTokenTest(t,
		newLiteral(tok(BOOL, "true")),
		tok(BOOL, "true"),
	)

	// 1
	quickSoloTokenTest(t,
		newLiteral(tok(NUMBER, "1")),
		tok(NUMBER, "1"),
	)

	// abc
	quickSoloTokenTest(t,
		newLiteral(tok(STRING, "abc")),
		tok(STRING, "abc"),
	)
}

func Test_S2_1(t *testing.T) {

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

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "a")),
				newLiteral(tok(NUMBER, "1")),
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S2_2(t *testing.T) {

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

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "a")),
				newLiteral(tok(NUMBER, "1")),
			),
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "b")),
				newLiteral(tok(BOOL, "TRUE")),
			),
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "c")),
				newLiteral(tok(STRING, "abc")),
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S3_1(t *testing.T) {

	// GIVEN a negation
	// WITH a following expression
	// THEN only the parsed negation expression is returned

	// -2
	given := []Token{
		tok(SUBTRACT, "-"),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
	}

	exp := newNegation(
		newLiteral(tok(NUMBER, "2")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S4_1(t *testing.T) {

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

	exp := newListAccessor(
		newIdentifier(tok(IDENTIFIER, "abc")),
		newLiteral(tok(NUMBER, "1")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S4_2(t *testing.T) {

	// GIVEN an operation
	// WITH a list accessor index being an operation
	// THEN a single parsed list accessor is expected
	// WITH individual operations nested in the correct order

	// abc[1 + 2]
	given := []Token{
		tok(IDENTIFIER, "abc"),
		tok(GUARD_OPEN, "["),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(GUARD_CLOSE, "]"),
		tok(TERMINATOR, ""),
	}

	first := newOperation(
		tok(ADD, "+"),
		newLiteral(tok(NUMBER, "1")),
		newLiteral(tok(NUMBER, "2")),
	)

	exp := newListAccessor(
		newIdentifier(tok(IDENTIFIER, "abc")),
		first,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S5_1(t *testing.T) {

	// GIVEN only a comment
	// THEN no statements are returned

	// // abc
	given := []Token{
		tok(COMMENT, "// abc"),
		tok(TERMINATOR, ""),
	}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_2(t *testing.T) {

	// GIVEN only whitespace
	// THEN no statements are returned

	given := []Token{
		tok(WHITESPACE, "    "),
		tok(TERMINATOR, ""),
	}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_3(t *testing.T) {

	// GIVEN only one terminator
	// THEN no statements are returned

	given := []Token{
		tok(TERMINATOR, ""),
	}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_4(t *testing.T) {

	// GIVEN only terminators
	// THEN no statements are returned

	given := []Token{
		tok(TERMINATOR, ""),
		tok(TERMINATOR, ""),
		tok(TERMINATOR, ""),
	}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func quickOperationTest(t *testing.T, left, operator, right Token) {

	express := func(tk Token) Expression {
		switch tk.Morpheme() {
		case IDENTIFIER:
			return newIdentifier(tk)
		case BOOL, NUMBER:
			return newLiteral(tk)
		default:
			panic("SANITY CHECK! Unknown token type: " + tk.Morpheme().String())
		}
	}

	given := []Token{
		left,
		operator,
		right,
		tok(TERMINATOR, ""),
	}

	exp := newOperation(
		operator,
		express(left),
		express(right),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_1(t *testing.T) {

	// GIVEN a simple addition
	// THEN a single parsed operation is expected

	// a + 1
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(ADD, "+"),
		tok(NUMBER, "1"),
	)
}

func Test_S6_2(t *testing.T) {

	// GIVEN a simple subtraction
	// THEN a single parsed operation is expected

	// a - 1
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
	)
}

func Test_S6_3(t *testing.T) {

	// GIVEN a simple multiplication
	// THEN a single parsed operation is expected

	// a * 1
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "1"),
	)
}

func Test_S6_4(t *testing.T) {

	// GIVEN a simple division
	// THEN a single parsed operation is expected

	// a / 1
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(DIVIDE, "/"),
		tok(NUMBER, "1"),
	)
}

func Test_S6_5(t *testing.T) {

	// GIVEN a simple division
	// THEN a single parsed operation is expected

	// a % 1
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "1"),
	)
}

func Test_S6_6(t *testing.T) {

	// GIVEN a simple logical AND operation
	// THEN a single parsed operation is expected

	// a & b
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(AND, "&"),
		tok(IDENTIFIER, "b"),
	)
}

func Test_S6_7(t *testing.T) {

	// GIVEN a simple logical OR operation
	// THEN a single parsed operation is expected

	// a | b
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(OR, "|"),
		tok(IDENTIFIER, "b"),
	)
}

func Test_S6_8(t *testing.T) {

	// GIVEN a simple logical == operation
	// THEN a single parsed operation is expected

	// a == b
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(EQUAL, "=="),
		tok(IDENTIFIER, "b"),
	)
}

func Test_S6_9(t *testing.T) {

	// GIVEN a simple logical != operation
	// THEN a single parsed operation is expected

	// a != b
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(NOT_EQUAL, "!="),
		tok(IDENTIFIER, "b"),
	)
}

func Test_S6_10(t *testing.T) {

	// GIVEN a simple logical < operation
	// THEN a single parsed operation is expected

	// a < b
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(LESS_THAN, "<"),
		tok(IDENTIFIER, "b"),
	)
}

func Test_S6_11(t *testing.T) {

	// GIVEN a simple logical > operation
	// THEN a single parsed operation is expected

	// a > b
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(MORE_THAN, ">"),
		tok(IDENTIFIER, "b"),
	)
}

func Test_S6_12(t *testing.T) {

	// GIVEN a simple logical <= operation
	// THEN a single parsed operation is expected

	// a <= b
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(LESS_THAN_OR_EQUAL, "<="),
		tok(IDENTIFIER, "b"),
	)
}

func Test_S6_13(t *testing.T) {

	// GIVEN a simple logical >= operation
	// THEN a single parsed operation is expected

	// a >= b
	quickOperationTest(t,
		tok(IDENTIFIER, "a"),
		tok(MORE_THAN_OR_EQUAL, ">="),
		tok(IDENTIFIER, "b"),
	)
}

func Test_S6_14(t *testing.T) {

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

	exp := newOperation(
		tok(ADD, "+"),
		newIdentifier(tok(IDENTIFIER, "a")),
		newNegation(
			newLiteral(tok(NUMBER, "1")),
		),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_15(t *testing.T) {

	// GIVEN a complex operation
	// WITH a additions and subtractions
	// THEN a parsed operation expression is expected
	// WITH individual operations nested in the correct order

	// a + 1 - b
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(ADD, "+"),
		tok(NUMBER, "1"),
		tok(SUBTRACT, "-"),
		tok(IDENTIFIER, "b"),
		tok(TERMINATOR, ""),
	}

	left := newOperation(
		tok(ADD, "+"),
		newIdentifier(tok(IDENTIFIER, "a")),
		newLiteral(tok(NUMBER, "1")),
	)

	exp := newOperation(
		tok(SUBTRACT, "-"),
		left,
		newIdentifier(tok(IDENTIFIER, "b")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_16(t *testing.T) {

	// GIVEN a complex operation
	// WITH a multiplication and division
	// THEN a parsed operation expression is expected
	// WITH individual operations nested in the correct order

	// a * 1 / b
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "1"),
		tok(DIVIDE, "/"),
		tok(IDENTIFIER, "b"),
		tok(TERMINATOR, ""),
	}

	left := newOperation(
		tok(MULTIPLY, "*"),
		newIdentifier(tok(IDENTIFIER, "a")),
		newLiteral(tok(NUMBER, "1")),
	)

	exp := newOperation(
		tok(DIVIDE, "/"),
		left,
		newIdentifier(tok(IDENTIFIER, "b")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_17(t *testing.T) {

	// GIVEN a complex operation
	// WHERE a multiplicative preceeds an addative
	// THEN a parsed operation expression is expected
	// WITH individual operations nested in the correct order

	// a * 1 + b
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(IDENTIFIER, "b"),
		tok(TERMINATOR, ""),
	}

	left := newOperation(
		tok(MULTIPLY, "*"),
		newIdentifier(tok(IDENTIFIER, "a")),
		newLiteral(tok(NUMBER, "1")),
	)

	exp := newOperation(
		tok(ADD, "+"),
		left,
		newIdentifier(tok(IDENTIFIER, "b")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_18(t *testing.T) {

	// GIVEN a complex operation
	// WHERE an additive preceeds a multiplicative
	// THEN a parsed operation expression is expected
	// WITH individual operations nested in the correct order

	// a + 1 * b
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(ADD, "+"),
		tok(NUMBER, "1"),
		tok(MULTIPLY, "*"),
		tok(IDENTIFIER, "b"),
		tok(TERMINATOR, ""),
	}

	left := newOperation(
		tok(MULTIPLY, "*"),
		newLiteral(tok(NUMBER, "1")),
		newIdentifier(tok(IDENTIFIER, "b")),
	)

	exp := newOperation(
		tok(ADD, "+"),
		newIdentifier(tok(IDENTIFIER, "a")),
		left,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_19(t *testing.T) {

	// GIVEN a complex operation
	// WITH multiple sub-expressions with different precedence
	// THEN a parsed operation expression is expected
	// WITH individual operations nested in the correct order

	// a - 1 * b % 2 + 1
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
		tok(MULTIPLY, "*"),
		tok(IDENTIFIER, "b"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "2"),
		tok(ADD, "+"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	first := newOperation(
		tok(MULTIPLY, "*"),
		newLiteral(tok(NUMBER, "1")),
		newIdentifier(tok(IDENTIFIER, "b")),
	)

	second := newOperation(
		tok(REMAINDER, "%"),
		first,
		newLiteral(tok(NUMBER, "2")),
	)

	third := newOperation(
		tok(SUBTRACT, "-"),
		newIdentifier(tok(IDENTIFIER, "a")),
		second,
	)

	exp := newOperation(
		tok(ADD, "+"),
		third,
		newLiteral(tok(NUMBER, "1")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_20(t *testing.T) {

	// GIVEN a really complex operation
	// WITH multiple sub-expressions with different precedence
	// THEN a parsed operation expression is expected
	// WITH individual operations nested in the correct order

	// a - 1 * b % 2 + 1 == 2 | c > 5 & c % 2 != 0
	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
		tok(MULTIPLY, "*"),
		tok(IDENTIFIER, "b"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "2"),
		tok(ADD, "+"),
		tok(NUMBER, "1"),
		tok(EQUAL, "=="),
		tok(NUMBER, "2"),
		tok(OR, "|"),
		tok(IDENTIFIER, "c"),
		tok(MORE_THAN, ">"),
		tok(NUMBER, "5"),
		tok(AND, "&"),
		tok(IDENTIFIER, "c"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "2"),
		tok(NOT_EQUAL, "!="),
		tok(NUMBER, "0"),
		tok(TERMINATOR, ""),
	}

	first := newOperation(
		tok(MULTIPLY, "*"),
		newLiteral(tok(NUMBER, "1")),
		newIdentifier(tok(IDENTIFIER, "b")),
	)

	second := newOperation(
		tok(REMAINDER, "%"),
		first,
		newLiteral(tok(NUMBER, "2")),
	)

	third := newOperation(
		tok(SUBTRACT, "-"),
		newIdentifier(tok(IDENTIFIER, "a")),
		second,
	)

	fourth := newOperation(
		tok(ADD, "+"),
		third,
		newLiteral(tok(NUMBER, "1")),
	)

	fifth := newOperation(
		tok(EQUAL, "=="),
		fourth,
		newLiteral(tok(NUMBER, "2")),
	)

	sixth := newOperation(
		tok(MORE_THAN, ">"),
		newIdentifier(tok(IDENTIFIER, "c")),
		newLiteral(tok(NUMBER, "5")),
	)

	seventh := newOperation(
		tok(REMAINDER, "%"),
		newIdentifier(tok(IDENTIFIER, "c")),
		newLiteral(tok(NUMBER, "2")),
	)

	eigth := newOperation(
		tok(NOT_EQUAL, "!="),
		seventh,
		newLiteral(tok(NUMBER, "0")),
	)

	ninth := newOperation(
		tok(AND, "&"),
		sixth,
		eigth,
	)

	exp := newOperation(
		tok(OR, "|"),
		fifth,
		ninth,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_21(t *testing.T) {

	quickParenTest := func(t *testing.T, exp Expression, tks ...Token) {

		var given []Token
		given = append(given, tok(PAREN_OPEN, "("))
		given = append(given, tks...)
		given = append(given, tok(PAREN_CLOSE, ")"))
		given = append(given, tok(TERMINATOR, ""))

		act, e := testFunc(given)
		expectOneStat(t, exp, act, e)
	}

	// GIVEN a prioritised operation group
	// WITH a single identifier or literal
	// THEN a single parsed expression is expected

	// (a)
	quickParenTest(t,
		newIdentifier(tok(IDENTIFIER, "a")),
		tok(IDENTIFIER, "a"),
	)

	// (true)
	quickParenTest(t,
		newLiteral(tok(BOOL, "true")),
		tok(BOOL, "true"),
	)

	// (1)
	quickParenTest(t,
		newLiteral(tok(NUMBER, "1")),
		tok(NUMBER, "1"),
	)

	// ("abc")
	quickParenTest(t,
		newLiteral(tok(STRING, "abc")),
		tok(STRING, "abc"),
	)

	// (-1)
	quickParenTest(t,
		newNegation(
			newLiteral(tok(NUMBER, "1")),
		),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
	)
}

func Test_S6_22(t *testing.T) {

	// GIVEN an operation
	// WITH some operations grouped as priority
	// THEN a single parsed expression is expected
	// WITH individual operations nested in the correct order

	// a * (1 + b)
	given := []Token{

		tok(IDENTIFIER, "a"),
		tok(MULTIPLY, "*"),
		tok(PAREN_OPEN, "("),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(IDENTIFIER, "b"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, ""),
	}

	first := newOperation(
		tok(ADD, "+"),
		newLiteral(tok(NUMBER, "1")),
		newIdentifier(tok(IDENTIFIER, "b")),
	)

	exp := newOperation(
		tok(MULTIPLY, "*"),
		newIdentifier(tok(IDENTIFIER, "a")),
		first,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_23(t *testing.T) {

	// GIVEN an operation
	// WITH some operations grouped as priority
	// THEN a single parsed expression is expected
	// WITH individual operations nested in the correct order

	// (a * 1) + b
	given := []Token{
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "1"),
		tok(PAREN_CLOSE, ")"),
		tok(ADD, "+"),
		tok(IDENTIFIER, "b"),
		tok(TERMINATOR, ""),
	}

	first := newOperation(
		tok(MULTIPLY, "*"),
		newIdentifier(tok(IDENTIFIER, "a")),
		newLiteral(tok(NUMBER, "1")),
	)

	exp := newOperation(
		tok(ADD, "+"),
		first,
		newIdentifier(tok(IDENTIFIER, "b")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_24(t *testing.T) {

	// GIVEN a really complex operation
	// WITH multiple sub-expressions with different precedence
	// WITH some sub-expressions grouped for priority
	// THEN a parsed operation expression is expected
	// WITH individual operations nested in the correct order

	// (a - 1 * ((b % 2) + (1 == 2 | c > 5)) & c % 2) != 0
	given := []Token{
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
		tok(MULTIPLY, "*"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "b"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "2"),
		tok(PAREN_CLOSE, ")"),
		tok(ADD, "+"),
		tok(PAREN_OPEN, "("),
		tok(NUMBER, "1"),
		tok(EQUAL, "=="),
		tok(NUMBER, "2"),
		tok(OR, "|"),
		tok(IDENTIFIER, "c"),
		tok(MORE_THAN, ">"),
		tok(NUMBER, "5"),
		tok(PAREN_CLOSE, ")"),
		tok(PAREN_CLOSE, ")"),
		tok(AND, "&"),
		tok(IDENTIFIER, "c"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "2"),
		tok(PAREN_CLOSE, ")"),
		tok(NOT_EQUAL, "!="),
		tok(NUMBER, "0"),
		tok(TERMINATOR, ""),
	}

	first := newOperation(
		tok(REMAINDER, "%"),
		newIdentifier(tok(IDENTIFIER, "b")),
		newLiteral(tok(NUMBER, "2")),
	)

	second := newOperation(
		tok(EQUAL, "=="),
		newLiteral(tok(NUMBER, "1")),
		newLiteral(tok(NUMBER, "2")),
	)

	third := newOperation(
		tok(MORE_THAN, ">"),
		newIdentifier(tok(IDENTIFIER, "c")),
		newLiteral(tok(NUMBER, "5")),
	)

	fourth := newOperation(
		tok(OR, "|"),
		second,
		third,
	)

	fifth := newOperation(
		tok(ADD, "+"),
		first,
		fourth,
	)

	sixth := newOperation(
		tok(MULTIPLY, "*"),
		newLiteral(tok(NUMBER, "1")),
		fifth,
	)

	seventh := newOperation(
		tok(SUBTRACT, "-"),
		newIdentifier(tok(IDENTIFIER, "a")),
		sixth,
	)

	eigth := newOperation(
		tok(REMAINDER, "%"),
		newIdentifier(tok(IDENTIFIER, "c")),
		newLiteral(tok(NUMBER, "2")),
	)

	ninth := newOperation(
		tok(AND, "&"),
		seventh,
		eigth,
	)

	exp := newOperation(
		tok(NOT_EQUAL, "!="),
		ninth,
		newLiteral(tok(NUMBER, "0")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_1(t *testing.T) {

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

	f := newFunction(
		tok(FUNC, "F"),
		newParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		newBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_2(t *testing.T) {

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

	f := newFunction(
		tok(FUNC, "F"),
		newParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{
				tok(IDENTIFIER, "a"),
			},
			[]Token{},
		),
		newBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_3(t *testing.T) {

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

	f := newFunction(
		tok(FUNC, "F"),
		newParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{
				tok(IDENTIFIER, "a"),
			},
		),
		newBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_4(t *testing.T) {

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

	f := newFunction(
		tok(FUNC, "F"),
		newParameters(
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
		newBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_5(t *testing.T) {

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

	body := newBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Expression{
			newAssignmentBlock(
				[]Assignment{
					newAssignment(
						newIdentifier(tok(IDENTIFIER, "a")),
						newLiteral(tok(NUMBER, "1")),
					),
				},
			),
		},
	)

	f := newFunction(
		tok(FUNC, "F"),
		newParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		body,
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_6(t *testing.T) {

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

	body := newBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Expression{
			newAssignmentBlock(
				[]Assignment{
					newAssignment(
						newIdentifier(tok(IDENTIFIER, "a")),
						newLiteral(tok(NUMBER, "1")),
					),
				},
			),
		},
	)

	f := newFunction(
		tok(FUNC, "F"),
		newParameters(
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

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S8_1(t *testing.T) {

	// GIVEN an expression function
	// WITH no parameters
	// WITH a simple expression
	// THEN only the parsed expression function is returned

	// f := E() 1
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(EXPR_FUNC, "E"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	f := newExpressionFunction(
		tok(EXPR_FUNC, "E"),
		[]Token{},
		newLiteral(tok(NUMBER, "1")),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S8_2(t *testing.T) {

	// GIVEN an expression function
	// WITH one parameter
	// WITH a simple expression
	// THEN only the parsed expression function is returned

	// f := E(a) 1
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(EXPR_FUNC, "E"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(PAREN_CLOSE, ")"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	f := newExpressionFunction(
		tok(EXPR_FUNC, "E"),
		[]Token{
			tok(IDENTIFIER, "a"),
		},
		newLiteral(tok(NUMBER, "1")),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S8_3(t *testing.T) {

	// GIVEN an expression function
	// WITH multiple parameters
	// WITH a simple expression
	// THEN only the parsed expression function is returned

	// f := E(a) 1
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(EXPR_FUNC, "E"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "c"),
		tok(PAREN_CLOSE, ")"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	f := newExpressionFunction(
		tok(EXPR_FUNC, "E"),
		[]Token{
			tok(IDENTIFIER, "a"),
			tok(IDENTIFIER, "b"),
			tok(IDENTIFIER, "c"),
		},
		newLiteral(tok(NUMBER, "1")),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S8_4(t *testing.T) {

	// GIVEN an expression function
	// WITH no parameters
	// WITH a complex expression
	// THEN only the parsed expression function is returned

	// f := E() 1
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(EXPR_FUNC, "E"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "3"),
		tok(TERMINATOR, ""),
	}

	expr := newOperation(
		tok(ADD, "+"),
		newLiteral(tok(NUMBER, "1")),
		newOperation(
			tok(MULTIPLY, "*"),
			newLiteral(tok(NUMBER, "2")),
			newLiteral(tok(NUMBER, "3")),
		),
	)

	f := newExpressionFunction(
		tok(EXPR_FUNC, "E"),
		[]Token{},
		expr,
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S9_1(t *testing.T) {

	// GIVEN a watch block
	// WITH one identifier
	// WITH no statements
	// THEN only the parsed watch statement is returned

	// watch a {}
	given := []Token{
		tok(WATCH, "watch"),
		tok(IDENTIFIER, "a"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	body := newBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := newWatch(
		tok(WATCH, "watch"),
		[]Token{
			tok(IDENTIFIER, "a"),
		},
		body,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S9_2(t *testing.T) {

	// GIVEN a watch block
	// WITH multiple identifiers
	// WITH no statements
	// THEN only the parsed watch statement is returned

	// watch a, b, c {}
	given := []Token{
		tok(WATCH, "watch"),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "c"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	body := newBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := newWatch(
		tok(WATCH, "watch"),
		[]Token{
			tok(IDENTIFIER, "a"),
			tok(IDENTIFIER, "b"),
			tok(IDENTIFIER, "c"),
		},
		body,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S9_3(t *testing.T) {

	// GIVEN a watch block
	// WITH one identifier
	// WITH several statements
	// THEN only the parsed watch statement is returned

	// watch a, b, c {}
	given := []Token{
		tok(WATCH, "watch"),
		tok(IDENTIFIER, "a"),
		tok(BLOCK_OPEN, "{"),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
		tok(IDENTIFIER, "a"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "3"),
		tok(TERMINATOR, ""),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	first := newOperation(
		tok(ADD, "+"),
		newLiteral(tok(NUMBER, "1")),
		newLiteral(tok(NUMBER, "2")),
	)

	second := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(IDENTIFIER, "a")),
				newLiteral(tok(NUMBER, "3")),
			),
		},
	)

	body := newBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Expression{
			first,
			second,
		},
	)

	exp := newWatch(
		tok(WATCH, "watch"),
		[]Token{
			tok(IDENTIFIER, "a"),
		},
		body,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S10_1(t *testing.T) {

	// GIVEN a guarded block
	// WITH a simple condition
	// WITH no statements
	// THEN only the parsed guard statement is returned

	// [true] {}
	given := []Token{
		tok(GUARD_OPEN, "["),
		tok(BOOL, "true"),
		tok(GUARD_CLOSE, "]"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	body := newBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := newGuard(
		tok(GUARD_OPEN, "["),
		newLiteral(tok(BOOL, "true")),
		body,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S10_2(t *testing.T) {

	// GIVEN a guarded block
	// WITH a complex condition
	// WITH no statements
	// THEN only the parsed guard statement is returned

	// [1 == 2] {}
	given := []Token{
		tok(GUARD_OPEN, "["),
		tok(NUMBER, "1"),
		tok(EQUAL, "=="),
		tok(NUMBER, "2"),
		tok(GUARD_CLOSE, "]"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	condition := newOperation(
		tok(EQUAL, "=="),
		newLiteral(tok(NUMBER, "1")),
		newLiteral(tok(NUMBER, "2")),
	)

	body := newBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := newGuard(
		tok(GUARD_OPEN, "["),
		condition,
		body,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S10_3(t *testing.T) {

	// GIVEN a guarded block
	// WITH a simple condition
	// WITH several statements
	// THEN only the parsed guard statement is returned

	// [true] {
	//   1 + 2
	//   3 * 4
	// }
	given := []Token{
		tok(GUARD_OPEN, "["),
		tok(BOOL, "true"),
		tok(GUARD_CLOSE, "]"),
		tok(BLOCK_OPEN, "{"),
		tok(TERMINATOR, ""),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
		tok(NUMBER, "3"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "4"),
		tok(TERMINATOR, ""),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	statements := []Expression{
		newOperation(
			tok(ADD, "+"),
			newLiteral(tok(NUMBER, "1")),
			newLiteral(tok(NUMBER, "2")),
		),
		newOperation(
			tok(MULTIPLY, "*"),
			newLiteral(tok(NUMBER, "3")),
			newLiteral(tok(NUMBER, "4")),
		),
	}

	body := newBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		statements,
	)

	exp := newGuard(
		tok(GUARD_OPEN, "["),
		newLiteral(tok(BOOL, "true")),
		body,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_1(t *testing.T) {

	// GIVEN a match block
	// WITH a simple condition
	// WITH no statements
	// THEN then the correct match statement is returned

	// match true {
	// }
	given := []Token{
		tok(MATCH, "match"),
		tok(BOOL, "true"),
		tok(BLOCK_OPEN, "{"),
		tok(TERMINATOR, ""),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	exp := newMatch(
		tok(MATCH, "match"),
		tok(BLOCK_CLOSE, "}"),
		newLiteral(tok(BOOL, "true")),
		[]MatchCase{},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_2(t *testing.T) {

	// GIVEN a match block
	// WITH a complex condition
	// WITH no statements
	// THEN then the correct match statement is returned

	// match true {
	// }
	given := []Token{
		tok(MATCH, "match"),
		tok(NUMBER, "1"),
		tok(EQUAL, "=="),
		tok(NUMBER, "2"),
		tok(BLOCK_OPEN, "{"),
		tok(TERMINATOR, ""),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	condition := newOperation(
		tok(EQUAL, "=="),
		newLiteral(tok(NUMBER, "1")),
		newLiteral(tok(NUMBER, "2")),
	)

	exp := newMatch(
		tok(MATCH, "match"),
		tok(BLOCK_CLOSE, "}"),
		condition,
		[]MatchCase{},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_3(t *testing.T) {

	// GIVEN a match block
	// WITH a simple condition
	// WITH one case
	// THEN then the correct match statement is returned

	// match 1 {
	//   1 -> 3 * 4
	// }
	given := []Token{
		tok(MATCH, "match"),
		tok(NUMBER, "1"),
		tok(BLOCK_OPEN, "{"),
		tok(TERMINATOR, "\n"),
		tok(NUMBER, "1"),
		tok(DO, "->"),
		tok(NUMBER, "3"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "4"),
		tok(TERMINATOR, "\n"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	firstCase := newLiteral(tok(NUMBER, "1"))

	firstBlock := newUnDelimiteredBlockExpr(
		[]Expression{
			newOperation(
				tok(MULTIPLY, "*"),
				newLiteral(tok(NUMBER, "3")),
				newLiteral(tok(NUMBER, "4")),
			),
		},
	)

	exp := newMatch(
		tok(MATCH, "match"),
		tok(BLOCK_CLOSE, "}"),
		newLiteral(tok(NUMBER, "1")),
		[]MatchCase{
			newMatchCase(firstCase, firstBlock),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_4(t *testing.T) {

	// GIVEN a match block
	// WITH a simple condition
	// WITH a guard case
	// THEN then the correct match statement is returned

	// match 1 {
	//   [1 == 2] -> 3 * 4
	// }
	given := []Token{
		tok(MATCH, "match"),
		tok(NUMBER, "1"),
		tok(BLOCK_OPEN, "{"),
		tok(TERMINATOR, "\n"),
		tok(GUARD_OPEN, "["),
		tok(NUMBER, "1"),
		tok(EQUAL, "=="),
		tok(NUMBER, "2"),
		tok(GUARD_CLOSE, "]"),
		tok(DO, "->"),
		tok(NUMBER, "3"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "4"),
		tok(TERMINATOR, "\n"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	firstBlock := newUnDelimiteredBlockExpr(
		[]Expression{
			newOperation(
				tok(MULTIPLY, "*"),
				newLiteral(tok(NUMBER, "3")),
				newLiteral(tok(NUMBER, "4")),
			),
		},
	)

	firstCase := newGuard(
		tok(GUARD_OPEN, "["),
		newOperation(
			tok(EQUAL, "=="),
			newLiteral(tok(NUMBER, "1")),
			newLiteral(tok(NUMBER, "2")),
		),
		firstBlock,
	)

	exp := newMatch(
		tok(MATCH, "match"),
		tok(BLOCK_CLOSE, "}"),
		newLiteral(tok(NUMBER, "1")),
		[]MatchCase{
			firstCase,
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_F1(t *testing.T) {

	// GIVEN an invalid statement or expression starting token
	// THEN parser returns error

	given := []Token{
		tok(ASSIGN, ":="),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
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

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F15(t *testing.T) {

	// GIVEN an expression function
	// WITH no expression
	// THEN parser returns error

	// f: E()
	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(EXPR_FUNC, "E"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F16(t *testing.T) {

	// GIVEN a watch block
	// WITH no body
	// THEN parser returns error

	// watch a
	given := []Token{
		tok(WATCH, "watch"),
		tok(IDENTIFIER, "a"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F17(t *testing.T) {

	// GIVEN a guard block
	// WITH no body
	// THEN parser returns error

	// watch a
	given := []Token{
		tok(GUARD_OPEN, "["),
		tok(BOOL, "true"),
		tok(GUARD_CLOSE, "]"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F18(t *testing.T) {

	// GIVEN a guard block
	// WITH no condition
	// THEN parser returns error

	// watch a
	given := []Token{
		tok(GUARD_OPEN, "["),
		tok(GUARD_CLOSE, "]"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(given)
	expectError(t, act, e)
}
