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
		given = append(given, tok(TK_TERMINATOR, ""))

		act, e := testFunc(given)
		expectOneStat(t, exp, act, e)
	}

	// GIVEN an identifier only
	// THEN only an identifier expression is returned

	// a
	quickSoloTokenTest(t,
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		tok(TK_IDENTIFIER, "a"),
	)

	// true
	quickSoloTokenTest(t,
		newLiteral(tok(TK_BOOL, "true")),
		tok(TK_BOOL, "true"),
	)

	// 1
	quickSoloTokenTest(t,
		newLiteral(tok(TK_NUMBER, "1")),
		tok(TK_NUMBER, "1"),
	)

	// abc
	quickSoloTokenTest(t,
		newLiteral(tok(TK_STRING, "abc")),
		tok(TK_STRING, "abc"),
	)
}

func Test_S2_1(t *testing.T) {

	// GIVEN an assignment
	// WITH only one identifier and one expression
	// THEN only the parsed assignment is returned

	// a: 1
	given := []Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "a")),
				newLiteral(tok(TK_NUMBER, "1")),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_BOOL, "TRUE"),
		tok(TK_DELIMITER, ","),
		tok(TK_STRING, "abc"),
		tok(TK_TERMINATOR, ""),
	}

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "a")),
				newLiteral(tok(TK_NUMBER, "1")),
			),
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "b")),
				newLiteral(tok(TK_BOOL, "TRUE")),
			),
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "c")),
				newLiteral(tok(TK_STRING, "abc")),
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
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
	}

	exp := newNegation(
		newLiteral(tok(TK_NUMBER, "2")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S4_1(t *testing.T) {

	// GIVEN a list identifier with a number literal index
	// THEN only the parsed list accessor is returned

	// abc[1]
	given := []Token{
		tok(TK_IDENTIFIER, "abc"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_TERMINATOR, ""),
	}

	exp := newCollectionAccessor(
		newIdentifier(tok(TK_IDENTIFIER, "abc")),
		newLiteral(tok(TK_NUMBER, "1")),
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
		tok(TK_IDENTIFIER, "abc"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_TERMINATOR, ""),
	}

	first := newOperation(
		tok(TK_PLUS, "+"),
		newLiteral(tok(TK_NUMBER, "1")),
		newLiteral(tok(TK_NUMBER, "2")),
	)

	exp := newCollectionAccessor(
		newIdentifier(tok(TK_IDENTIFIER, "abc")),
		first,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S4_3(t *testing.T) {

	// GIVEN a list accessor
	// WHICH returns another list
	// WHICH is then accessed
	// THEN the correctly nested parsed list accessors are returned

	// abc[1][2][3]
	given := []Token{
		tok(TK_IDENTIFIER, "abc"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "2"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "3"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_TERMINATOR, ""),
	}

	var first Expression = newCollectionAccessor(
		newIdentifier(tok(TK_IDENTIFIER, "abc")),
		newLiteral(tok(TK_NUMBER, "1")),
	)

	var second Expression = newCollectionAccessor(
		first,
		newLiteral(tok(TK_NUMBER, "2")),
	)

	exp := newCollectionAccessor(
		second,
		newLiteral(tok(TK_NUMBER, "3")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S5_1(t *testing.T) {

	// GIVEN only a comment
	// THEN no statements are returned

	// // abc
	given := []Token{
		tok(TK_COMMENT, "// abc"),
		tok(TK_TERMINATOR, ""),
	}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_2(t *testing.T) {

	// GIVEN only whitespace
	// THEN no statements are returned

	given := []Token{
		tok(TK_WHITESPACE, "    "),
		tok(TK_TERMINATOR, ""),
	}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_3(t *testing.T) {

	// GIVEN only one terminator
	// THEN no statements are returned

	given := []Token{
		tok(TK_TERMINATOR, ""),
	}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_4(t *testing.T) {

	// GIVEN only terminators
	// THEN no statements are returned

	given := []Token{
		tok(TK_TERMINATOR, ""),
		tok(TK_TERMINATOR, ""),
		tok(TK_TERMINATOR, ""),
	}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func quickOperationTest(t *testing.T, left, operator, right Token) {

	express := func(tk Token) Expression {

		switch tk.Type() {
		case TK_IDENTIFIER:
			return newIdentifier(tk)

		case TK_BOOL, TK_NUMBER:
			return newLiteral(tk)

		case TK_VOID:
			return newVoid(tk)

		default:
			panic("SANITY CHECK! Unknown token type: " + tk.Type().String())
		}
	}

	given := []Token{
		left,
		operator,
		right,
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
	)
}

func Test_S6_2(t *testing.T) {

	// GIVEN a simple subtraction
	// THEN a single parsed operation is expected

	// a - 1
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
	)
}

func Test_S6_3(t *testing.T) {

	// GIVEN a simple multiplication
	// THEN a single parsed operation is expected

	// a * 1
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "1"),
	)
}

func Test_S6_4(t *testing.T) {

	// GIVEN a simple division
	// THEN a single parsed operation is expected

	// a / 1
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DIVIDE, "/"),
		tok(TK_NUMBER, "1"),
	)
}

func Test_S6_5(t *testing.T) {

	// GIVEN a simple division
	// THEN a single parsed operation is expected

	// a % 1
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "1"),
	)
}

func Test_S6_6(t *testing.T) {

	// GIVEN a simple logical AND operation
	// THEN a single parsed operation is expected

	// a & b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_AND, "&"),
		tok(TK_IDENTIFIER, "b"),
	)
}

func Test_S6_7(t *testing.T) {

	// GIVEN a simple logical OR operation
	// THEN a single parsed operation is expected

	// a | b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_OR, "|"),
		tok(TK_IDENTIFIER, "b"),
	)
}

func Test_S6_8(t *testing.T) {

	// GIVEN a simple logical == operation
	// THEN a single parsed operation is expected

	// a == b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_EQUAL, "=="),
		tok(TK_IDENTIFIER, "b"),
	)
}

func Test_S6_9(t *testing.T) {

	// GIVEN a simple logical != operation
	// THEN a single parsed operation is expected

	// a != b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_NOT_EQUAL, "!="),
		tok(TK_IDENTIFIER, "b"),
	)
}

func Test_S6_10(t *testing.T) {

	// GIVEN a simple logical < operation
	// THEN a single parsed operation is expected

	// a < b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_LESS_THAN, "<"),
		tok(TK_IDENTIFIER, "b"),
	)
}

func Test_S6_11(t *testing.T) {

	// GIVEN a simple logical > operation
	// THEN a single parsed operation is expected

	// a > b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MORE_THAN, ">"),
		tok(TK_IDENTIFIER, "b"),
	)
}

func Test_S6_12(t *testing.T) {

	// GIVEN a simple logical <= operation
	// THEN a single parsed operation is expected

	// a <= b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_LESS_THAN_OR_EQUAL, "<="),
		tok(TK_IDENTIFIER, "b"),
	)
}

func Test_S6_13(t *testing.T) {

	// GIVEN a simple logical >= operation
	// THEN a single parsed operation is expected

	// a >= b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MORE_THAN_OR_EQUAL, ">="),
		tok(TK_IDENTIFIER, "b"),
	)
}

func Test_S6_14(t *testing.T) {

	// GIVEN a simple addition
	// WITH a negated right operand
	// THEN a single parsed operation is expected
	// WITH an parsed negation a the right operand

	// a + -1
	given := []Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	exp := newOperation(
		tok(TK_PLUS, "+"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		newNegation(
			newLiteral(tok(TK_NUMBER, "1")),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
		tok(TK_MINUS, "-"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}

	left := newOperation(
		tok(TK_PLUS, "+"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newOperation(
		tok(TK_MINUS, "-"),
		left,
		newIdentifier(tok(TK_IDENTIFIER, "b")),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "1"),
		tok(TK_DIVIDE, "/"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}

	left := newOperation(
		tok(TK_MULTIPLY, "*"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newOperation(
		tok(TK_DIVIDE, "/"),
		left,
		newIdentifier(tok(TK_IDENTIFIER, "b")),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}

	left := newOperation(
		tok(TK_MULTIPLY, "*"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newOperation(
		tok(TK_PLUS, "+"),
		left,
		newIdentifier(tok(TK_IDENTIFIER, "b")),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}

	left := newOperation(
		tok(TK_MULTIPLY, "*"),
		newLiteral(tok(TK_NUMBER, "1")),
		newIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	exp := newOperation(
		tok(TK_PLUS, "+"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "2"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	first := newOperation(
		tok(TK_MULTIPLY, "*"),
		newLiteral(tok(TK_NUMBER, "1")),
		newIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	second := newOperation(
		tok(TK_REMAINDER, "%"),
		first,
		newLiteral(tok(TK_NUMBER, "2")),
	)

	third := newOperation(
		tok(TK_MINUS, "-"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		second,
	)

	exp := newOperation(
		tok(TK_PLUS, "+"),
		third,
		newLiteral(tok(TK_NUMBER, "1")),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "2"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
		tok(TK_EQUAL, "=="),
		tok(TK_NUMBER, "2"),
		tok(TK_OR, "|"),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_MORE_THAN, ">"),
		tok(TK_NUMBER, "5"),
		tok(TK_AND, "&"),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "2"),
		tok(TK_NOT_EQUAL, "!="),
		tok(TK_NUMBER, "0"),
		tok(TK_TERMINATOR, ""),
	}

	first := newOperation(
		tok(TK_MULTIPLY, "*"),
		newLiteral(tok(TK_NUMBER, "1")),
		newIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	second := newOperation(
		tok(TK_REMAINDER, "%"),
		first,
		newLiteral(tok(TK_NUMBER, "2")),
	)

	third := newOperation(
		tok(TK_MINUS, "-"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		second,
	)

	fourth := newOperation(
		tok(TK_PLUS, "+"),
		third,
		newLiteral(tok(TK_NUMBER, "1")),
	)

	fifth := newOperation(
		tok(TK_EQUAL, "=="),
		fourth,
		newLiteral(tok(TK_NUMBER, "2")),
	)

	sixth := newOperation(
		tok(TK_MORE_THAN, ">"),
		newIdentifier(tok(TK_IDENTIFIER, "c")),
		newLiteral(tok(TK_NUMBER, "5")),
	)

	seventh := newOperation(
		tok(TK_REMAINDER, "%"),
		newIdentifier(tok(TK_IDENTIFIER, "c")),
		newLiteral(tok(TK_NUMBER, "2")),
	)

	eigth := newOperation(
		tok(TK_NOT_EQUAL, "!="),
		seventh,
		newLiteral(tok(TK_NUMBER, "0")),
	)

	ninth := newOperation(
		tok(TK_AND, "&"),
		sixth,
		eigth,
	)

	exp := newOperation(
		tok(TK_OR, "|"),
		fifth,
		ninth,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_21(t *testing.T) {

	quickParenTest := func(t *testing.T, exp Expression, tks ...Token) {

		var given []Token
		given = append(given, tok(TK_PAREN_OPEN, "("))
		given = append(given, tks...)
		given = append(given, tok(TK_PAREN_CLOSE, ")"))
		given = append(given, tok(TK_TERMINATOR, ""))

		act, e := testFunc(given)
		expectOneStat(t, exp, act, e)
	}

	// GIVEN a prioritised operation group
	// WITH a single identifier or literal
	// THEN a single parsed expression is expected

	// (a)
	quickParenTest(t,
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		tok(TK_IDENTIFIER, "a"),
	)

	// (true)
	quickParenTest(t,
		newLiteral(tok(TK_BOOL, "true")),
		tok(TK_BOOL, "true"),
	)

	// (1)
	quickParenTest(t,
		newLiteral(tok(TK_NUMBER, "1")),
		tok(TK_NUMBER, "1"),
	)

	// ("abc")
	quickParenTest(t,
		newLiteral(tok(TK_STRING, "abc")),
		tok(TK_STRING, "abc"),
	)

	// (-1)
	quickParenTest(t,
		newNegation(
			newLiteral(tok(TK_NUMBER, "1")),
		),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
	)
}

func Test_S6_22(t *testing.T) {

	// GIVEN an operation
	// WITH some operations grouped as priority
	// THEN a single parsed expression is expected
	// WITH individual operations nested in the correct order

	// a * (1 + b)
	given := []Token{

		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	first := newOperation(
		tok(TK_PLUS, "+"),
		newLiteral(tok(TK_NUMBER, "1")),
		newIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	exp := newOperation(
		tok(TK_MULTIPLY, "*"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
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
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "1"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}

	first := newOperation(
		tok(TK_MULTIPLY, "*"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newOperation(
		tok(TK_PLUS, "+"),
		first,
		newIdentifier(tok(TK_IDENTIFIER, "b")),
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
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "2"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_PLUS, "+"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_EQUAL, "=="),
		tok(TK_NUMBER, "2"),
		tok(TK_OR, "|"),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_MORE_THAN, ">"),
		tok(TK_NUMBER, "5"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_AND, "&"),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "2"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_NOT_EQUAL, "!="),
		tok(TK_NUMBER, "0"),
		tok(TK_TERMINATOR, ""),
	}

	first := newOperation(
		tok(TK_REMAINDER, "%"),
		newIdentifier(tok(TK_IDENTIFIER, "b")),
		newLiteral(tok(TK_NUMBER, "2")),
	)

	second := newOperation(
		tok(TK_EQUAL, "=="),
		newLiteral(tok(TK_NUMBER, "1")),
		newLiteral(tok(TK_NUMBER, "2")),
	)

	third := newOperation(
		tok(TK_MORE_THAN, ">"),
		newIdentifier(tok(TK_IDENTIFIER, "c")),
		newLiteral(tok(TK_NUMBER, "5")),
	)

	fourth := newOperation(
		tok(TK_OR, "|"),
		second,
		third,
	)

	fifth := newOperation(
		tok(TK_PLUS, "+"),
		first,
		fourth,
	)

	sixth := newOperation(
		tok(TK_MULTIPLY, "*"),
		newLiteral(tok(TK_NUMBER, "1")),
		fifth,
	)

	seventh := newOperation(
		tok(TK_MINUS, "-"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		sixth,
	)

	eigth := newOperation(
		tok(TK_REMAINDER, "%"),
		newIdentifier(tok(TK_IDENTIFIER, "c")),
		newLiteral(tok(TK_NUMBER, "2")),
	)

	ninth := newOperation(
		tok(TK_AND, "&"),
		seventh,
		eigth,
	)

	exp := newOperation(
		tok(TK_NOT_EQUAL, "!="),
		ninth,
		newLiteral(tok(TK_NUMBER, "0")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_25(t *testing.T) {

	// GIVEN a simple logical != operation
	// WITH a void as an operand
	// THEN a single parsed operation is expected

	// a <= b
	quickOperationTest(t,
		tok(TK_IDENTIFIER, "a"),
		tok(TK_NOT_EQUAL, "!="),
		tok(TK_VOID, "_"),
	)
}

func Test_S7_1(t *testing.T) {

	// GIVEN a function
	// WITH no parameters
	// AND no statements in the body
	// THEN only the parsed function is returned

	// f := F() {}
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	f := newFunction(
		tok(TK_FUNCTION, "F"),
		newParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		newBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	f := newFunction(
		tok(TK_FUNCTION, "F"),
		newParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{
				tok(TK_IDENTIFIER, "a"),
			},
			[]Token{},
		),
		newBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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

	// f := F(-> a) {}
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_OUTPUTS, "->"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	f := newFunction(
		tok(TK_FUNCTION, "F"),
		newParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{
				tok(TK_IDENTIFIER, "a"),
			},
		),
		newBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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

	// f := F(a, b -> c, d) {}
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_DELIMITER, ","),
		tok(TK_OUTPUTS, "->"),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "d"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	f := newFunction(
		tok(TK_FUNCTION, "F"),
		newParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{
				tok(TK_IDENTIFIER, "a"),
				tok(TK_IDENTIFIER, "b"),
			},
			[]Token{
				tok(TK_IDENTIFIER, "c"),
				tok(TK_IDENTIFIER, "d"),
			},
		),
		newBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	body := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{
			newAssignmentBlock(
				[]Assignment{
					newAssignment(
						newIdentifier(tok(TK_IDENTIFIER, "a")),
						newLiteral(tok(TK_NUMBER, "1")),
					),
				},
			),
		},
	)

	f := newFunction(
		tok(TK_FUNCTION, "F"),
		newParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		body,
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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
	// -> b,
	// ) {
	// a: b
	// }
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_OUTPUTS, "->"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_DELIMITER, ","),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	body := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{
			newAssignmentBlock(
				[]Assignment{
					newAssignment(
						newIdentifier(tok(TK_IDENTIFIER, "a")),
						newLiteral(tok(TK_NUMBER, "1")),
					),
				},
			),
		},
	)

	f := newFunction(
		tok(TK_FUNCTION, "F"),
		newParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{
				tok(TK_IDENTIFIER, "a"),
			},
			[]Token{
				tok(TK_IDENTIFIER, "b"),
			},
		),
		body,
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	f := newExpressionFunction(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{},
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	f := newExpressionFunction(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
		},
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	f := newExpressionFunction(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "3"),
		tok(TK_TERMINATOR, ""),
	}

	expr := newOperation(
		tok(TK_PLUS, "+"),
		newLiteral(tok(TK_NUMBER, "1")),
		newOperation(
			tok(TK_MULTIPLY, "*"),
			newLiteral(tok(TK_NUMBER, "2")),
			newLiteral(tok(TK_NUMBER, "3")),
		),
	)

	f := newExpressionFunction(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{},
		expr,
	)

	exp := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "f")),
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
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	body := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := newWatch(
		tok(TK_WATCH, "watch"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
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
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	body := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := newWatch(
		tok(TK_WATCH, "watch"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
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
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "3"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	first := newOperation(
		tok(TK_PLUS, "+"),
		newLiteral(tok(TK_NUMBER, "1")),
		newLiteral(tok(TK_NUMBER, "2")),
	)

	second := newAssignmentBlock(
		[]Assignment{
			newAssignment(
				newIdentifier(tok(TK_IDENTIFIER, "a")),
				newLiteral(tok(TK_NUMBER, "3")),
			),
		},
	)

	body := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{
			first,
			second,
		},
	)

	exp := newWatch(
		tok(TK_WATCH, "watch"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
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
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	body := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := newGuard(
		tok(TK_GUARD_OPEN, "["),
		newLiteral(tok(TK_BOOL, "true")),
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
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_EQUAL, "=="),
		tok(TK_NUMBER, "2"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	condition := newOperation(
		tok(TK_EQUAL, "=="),
		newLiteral(tok(TK_NUMBER, "1")),
		newLiteral(tok(TK_NUMBER, "2")),
	)

	body := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := newGuard(
		tok(TK_GUARD_OPEN, "["),
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
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, ""),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
		tok(TK_NUMBER, "3"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "4"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	statements := []Expression{
		newOperation(
			tok(TK_PLUS, "+"),
			newLiteral(tok(TK_NUMBER, "1")),
			newLiteral(tok(TK_NUMBER, "2")),
		),
		newOperation(
			tok(TK_MULTIPLY, "*"),
			newLiteral(tok(TK_NUMBER, "3")),
			newLiteral(tok(TK_NUMBER, "4")),
		),
	}

	body := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		statements,
	)

	exp := newGuard(
		tok(TK_GUARD_OPEN, "["),
		newLiteral(tok(TK_BOOL, "true")),
		body,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S10_4(t *testing.T) {

	// GIVEN a guarded block
	// WITH a simple condition
	// WITH inline body
	// THEN only the parsed guard statement is returned

	// [true] {}
	given := []Token{
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	body := newUnDelimiteredBlock(
		[]Expression{
			newLiteral(tok(TK_NUMBER, "1")),
		},
	)

	exp := newGuard(
		tok(TK_GUARD_OPEN, "["),
		newLiteral(tok(TK_BOOL, "true")),
		body,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_1(t *testing.T) {

	// GIVEN a when block
	// WITH a simple condition
	// WITH no statements
	// THEN then the correct when statement is returned

	// when true {
	// }
	given := []Token{
		tok(TK_WHEN, "when"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_BOOL, "true"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	init := newAssignment(
		newIdentifier(tok(TK_IDENTIFIER, "x")),
		newLiteral(tok(TK_BOOL, "true")),
	)

	exp := newWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "}"),
		init,
		[]WhenCase{},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_2(t *testing.T) {

	// GIVEN a when block
	// WITH a complex condition
	// WITH no statements
	// THEN then the correct when statement is returned

	// when true {
	// }
	given := []Token{
		tok(TK_WHEN, "when"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_EQUAL, "=="),
		tok(TK_NUMBER, "2"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	condition := newOperation(
		tok(TK_EQUAL, "=="),
		newLiteral(tok(TK_NUMBER, "1")),
		newLiteral(tok(TK_NUMBER, "2")),
	)

	init := newAssignment(
		newIdentifier(tok(TK_IDENTIFIER, "x")),
		condition,
	)

	exp := newWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "}"),
		init,
		[]WhenCase{},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_3(t *testing.T) {

	// GIVEN a when block
	// WITH a simple condition
	// WITH one case
	// THEN then the correct when statement is returned

	// when 1 {
	//   1 -> 3 * 4
	// }
	given := []Token{
		tok(TK_WHEN, "when"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_NUMBER, "1"),
		tok(TK_THEN, "->"),
		tok(TK_NUMBER, "3"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "4"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	firstCase := newLiteral(tok(TK_NUMBER, "1"))

	firstBlock := newUnDelimiteredBlock(
		[]Expression{
			newOperation(
				tok(TK_MULTIPLY, "*"),
				newLiteral(tok(TK_NUMBER, "3")),
				newLiteral(tok(TK_NUMBER, "4")),
			),
		},
	)

	init := newAssignment(
		newIdentifier(tok(TK_IDENTIFIER, "x")),
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "}"),
		init,
		[]WhenCase{
			newWhenCase(firstCase, firstBlock),
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_4(t *testing.T) {

	// GIVEN a when block
	// WITH a simple condition
	// WITH a guard case
	// THEN then the correct when statement is returned

	// when 1 {
	//   [1 == 2] -> 3 * 4
	// }
	given := []Token{
		tok(TK_WHEN, "when"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_EQUAL, "=="),
		tok(TK_NUMBER, "2"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_THEN, "->"),
		tok(TK_NUMBER, "3"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "4"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	firstBlock := newUnDelimiteredBlock(
		[]Expression{
			newOperation(
				tok(TK_MULTIPLY, "*"),
				newLiteral(tok(TK_NUMBER, "3")),
				newLiteral(tok(TK_NUMBER, "4")),
			),
		},
	)

	firstCase := newGuard(
		tok(TK_GUARD_OPEN, "["),
		newOperation(
			tok(TK_EQUAL, "=="),
			newLiteral(tok(TK_NUMBER, "1")),
			newLiteral(tok(TK_NUMBER, "2")),
		),
		firstBlock,
	)

	init := newAssignment(
		newIdentifier(tok(TK_IDENTIFIER, "x")),
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "}"),
		init,
		[]WhenCase{
			firstCase,
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S11_5(t *testing.T) {

	// GIVEN a when block
	// WITH a simple condition
	// WITH a severakl when and guard cases
	// THEN then the correct when statement is returned

	// when 3 {
	//   1        -> a
	//   [a == b] -> b
	//	 2        -> c
	//   [true]   -> {
	//                 d
	//               }
	// }
	given := []Token{
		tok(TK_WHEN, "when"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "3"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_NUMBER, "1"),
		tok(TK_THEN, "->"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_EQUAL, "=="),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_THEN, "->"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_NUMBER, "2"),
		tok(TK_THEN, "->"),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_THEN, "->"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_IDENTIFIER, "d"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	firstCase := newWhenCase(
		newLiteral(tok(TK_NUMBER, "1")),
		newUnDelimiteredBlock(
			[]Expression{
				newIdentifier(tok(TK_IDENTIFIER, "a")),
			},
		),
	)

	secondCase := newGuard(
		tok(TK_GUARD_OPEN, "["),
		newOperation(
			tok(TK_EQUAL, "=="),
			newIdentifier(tok(TK_IDENTIFIER, "a")),
			newIdentifier(tok(TK_IDENTIFIER, "b")),
		),
		newUnDelimiteredBlock(
			[]Expression{
				newIdentifier(tok(TK_IDENTIFIER, "b")),
			},
		),
	)

	thirdCase := newWhenCase(
		newLiteral(tok(TK_NUMBER, "2")),
		newUnDelimiteredBlock(
			[]Expression{
				newIdentifier(tok(TK_IDENTIFIER, "c")),
			},
		),
	)

	fourthCase := newGuard(
		tok(TK_GUARD_OPEN, "["),
		newLiteral(tok(TK_BOOL, "true")),
		newBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{
				newIdentifier(tok(TK_IDENTIFIER, "d")),
			},
		),
	)

	init := newAssignment(
		newIdentifier(tok(TK_IDENTIFIER, "x")),
		newLiteral(tok(TK_NUMBER, "3")),
	)

	exp := newWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "}"),
		init,
		[]WhenCase{
			firstCase,
			secondCase,
			thirdCase,
			fourthCase,
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S12_1(t *testing.T) {

	// GIVEN a loop
	// WITH a simple initialiser
	// WITH a simple guard
	// WITH no body statements
	// THEN then the correct loop statement is returned

	// loop i := 0 [true] {
	// }
	given := []Token{
		tok(TK_LOOP, "loop"),
		tok(TK_IDENTIFIER, "i"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "0"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	init := newAssignment(
		newIdentifier(tok(TK_IDENTIFIER, "i")),
		newLiteral(tok(TK_NUMBER, "0")),
	)

	guard := newGuard(
		tok(TK_GUARD_OPEN, "["),
		newLiteral(tok(TK_BOOL, "true")),
		newBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := newLoop(
		tok(TK_LOOP, "loop"),
		init,
		guard,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S12_2(t *testing.T) {

	// GIVEN a loop
	// WITH a complex initialiser
	// WITH a complex guard
	// WITH several body statements
	// THEN then the correct loop statement is returned

	// loop i := a - 1 [i < 10] {
	//   a + b
	//   c * d
	// }
	given := []Token{
		tok(TK_LOOP, "loop"),
		tok(TK_IDENTIFIER, "i"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_IDENTIFIER, "i"),
		tok(TK_LESS_THAN, "<"),
		tok(TK_NUMBER, "10"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, ""),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_IDENTIFIER, "d"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	init := newAssignment(
		newIdentifier(tok(TK_IDENTIFIER, "i")),
		newOperation(
			tok(TK_MINUS, "-"),
			newIdentifier(tok(TK_IDENTIFIER, "a")),
			newLiteral(tok(TK_NUMBER, "1")),
		),
	)

	condition := newOperation(
		tok(TK_LESS_THAN, "<"),
		newIdentifier(tok(TK_IDENTIFIER, "i")),
		newLiteral(tok(TK_NUMBER, "10")),
	)

	firstStat := newOperation(
		tok(TK_PLUS, "+"),
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		newIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	secondStat := newOperation(
		tok(TK_MULTIPLY, "*"),
		newIdentifier(tok(TK_IDENTIFIER, "c")),
		newIdentifier(tok(TK_IDENTIFIER, "d")),
	)

	guard := newGuard(
		tok(TK_GUARD_OPEN, "["),
		condition,
		newBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{
				firstStat,
				secondStat,
			},
		),
	)

	exp := newLoop(
		tok(TK_LOOP, "loop"),
		init,
		guard,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S13_1(t *testing.T) {

	// GIVEN a function call
	// WITH no arguments
	// THEN then the correct function expression is returned

	// f()
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	var f Expression = newIdentifier(tok(TK_IDENTIFIER, "f"))

	exp := newFunctionCall(
		tok(TK_PAREN_CLOSE, ")"),
		f,
		[]Expression{},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S13_2(t *testing.T) {

	// GIVEN a function call
	// WITH one simple argument
	// THEN then the correct function expression is returned

	// f(1)
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	var f Expression = newIdentifier(tok(TK_IDENTIFIER, "f"))

	args := []Expression{
		newLiteral(tok(TK_NUMBER, "1")),
	}

	exp := newFunctionCall(
		tok(TK_PAREN_CLOSE, ")"),
		f,
		args,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S13_3(t *testing.T) {

	// GIVEN a function call
	// WITH one complex argument
	// THEN then the correct function expression is returned

	// f(1+2)
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	var f Expression = newIdentifier(tok(TK_IDENTIFIER, "f"))

	args := []Expression{
		newOperation(
			tok(TK_PLUS, "+"),
			newLiteral(tok(TK_NUMBER, "1")),
			newLiteral(tok(TK_NUMBER, "2")),
		),
	}

	exp := newFunctionCall(
		tok(TK_PAREN_CLOSE, ")"),
		f,
		args,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S13_4(t *testing.T) {

	// GIVEN a function call
	// WITH several simple arguments
	// THEN then the correct function expression is returned

	// f(a, true, 1, "abc")
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_BOOL, "true"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_STRING, "abc"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	var f Expression = newIdentifier(tok(TK_IDENTIFIER, "f"))

	args := []Expression{
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		newLiteral(tok(TK_BOOL, "true")),
		newLiteral(tok(TK_NUMBER, "1")),
		newLiteral(tok(TK_STRING, "abc")),
	}

	exp := newFunctionCall(
		tok(TK_PAREN_CLOSE, ")"),
		f,
		args,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S13_5(t *testing.T) {

	// GIVEN a function call
	// WITH several simple arguments over several lines
	// THEN then the correct function expression is returned

	// f(
	//   a,
	//   true,
	//   1,
	//   "abc",
	// )
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BOOL, "true"),
		tok(TK_DELIMITER, ","),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_STRING, "abc"),
		tok(TK_DELIMITER, ","),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	var f Expression = newIdentifier(tok(TK_IDENTIFIER, "f"))

	args := []Expression{
		newIdentifier(tok(TK_IDENTIFIER, "a")),
		newLiteral(tok(TK_BOOL, "true")),
		newLiteral(tok(TK_NUMBER, "1")),
		newLiteral(tok(TK_STRING, "abc")),
	}

	exp := newFunctionCall(
		tok(TK_PAREN_CLOSE, ")"),
		f,
		args,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S13_6(t *testing.T) {

	// GIVEN a function call
	// WHICH returns function
	// WHICH is then called within the same expression
	// THEN then the correct function expression is returned

	// f()()()
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "(a"),
		tok(TK_PAREN_CLOSE, "a)"),
		tok(TK_PAREN_OPEN, "(b"),
		tok(TK_PAREN_CLOSE, "b)"),
		tok(TK_PAREN_OPEN, "(c"),
		tok(TK_PAREN_CLOSE, "c)"),
		tok(TK_TERMINATOR, ""),
	}

	var f Expression = newIdentifier(tok(TK_IDENTIFIER, "f"))

	var first Expression = newFunctionCall(
		tok(TK_PAREN_CLOSE, "a)"),
		f,
		[]Expression{},
	)

	var second Expression = newFunctionCall(
		tok(TK_PAREN_CLOSE, "b)"),
		first,
		[]Expression{},
	)

	exp := newFunctionCall(
		tok(TK_PAREN_CLOSE, "c)"),
		second,
		[]Expression{},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S14_1(t *testing.T) {

	// GIVEN a spell call
	// WITH no arguments
	// THEN then the correct spell expression is returned

	// @s()
	given := []Token{
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	exp := newSpellCall(
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_CLOSE, ")"),
		[]Expression{},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S14_2(t *testing.T) {

	// GIVEN a spell call
	// WHICH returns a list
	// WHICH is being queried for an item
	// WHICH returns another function
	// THEN then the correct spell expression is returned

	// @s()()()
	given := []Token{
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_OPEN, "(a"),
		tok(TK_PAREN_CLOSE, "a)"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_PAREN_OPEN, "(c"),
		tok(TK_PAREN_CLOSE, "c)"),
		tok(TK_TERMINATOR, ""),
	}

	var first Expression = newSpellCall(
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_CLOSE, "a)"),
		[]Expression{},
	)

	var second Expression = newCollectionAccessor(
		first,
		newLiteral(tok(TK_NUMBER, "1")),
	)

	exp := newFunctionCall(
		tok(TK_PAREN_CLOSE, "c)"),
		second,
		[]Expression{},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S14_3(t *testing.T) {

	// GIVEN a spell call
	// WHICH containing a block argument
	// THEN then the correct spell expression is returned

	// @s(1, {
	//   "abc"
	// })
	given := []Token{
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_STRING, "abc"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	first := newLiteral(tok(TK_NUMBER, "1"))

	second := newBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{
			newLiteral(tok(TK_STRING, "abc")),
		},
	)

	exp := newSpellCall(
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_CLOSE, ")"),
		[]Expression{
			first,
			second,
		},
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S15_1(t *testing.T) {

	// GIVEN an exit
	// THEN then the correct exit statement is returned

	// exit
	given := []Token{
		tok(TK_EXIT, "exit"),
		tok(TK_NUMBER, "0"),
		tok(TK_TERMINATOR, ""),
	}

	exp := newExit(
		tok(TK_EXIT, "exit"),
		newLiteral(tok(TK_NUMBER, "0")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_F1(t *testing.T) {

	// GIVEN an invalid statement or expression starting token
	// THEN parser returns error

	given := []Token{
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_MINUS, "-"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_LIST, "LIST"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_LIST, "LIST"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_LIST, "LIST"),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_LIST, "LIST"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "x"),
		tok(TK_PLUS, "+"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "x"),
		tok(TK_PLUS, "+"),
		tok(TK_PLUS, "+"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_GUARD_OPEN, "["),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F19(t *testing.T) {

	// GIVEN a function call
	// WITH no closing parenthesis
	// THEN parser returns error

	// f(loop
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_LOOP, "loop"),
		tok(TK_TERMINATOR, ""),
	}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F20(t *testing.T) {

	// GIVEN a function call
	// WITH a missig delimiter between arguments
	// THEN parser returns error

	// f(1, 2 3)
	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_NUMBER, "3"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	act, e := testFunc(given)
	expectError(t, act, e)
}
