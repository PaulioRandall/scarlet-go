package parser

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var testFunc func(TokenStream) ([]Expression, error) = ParseStatements

func Test_S2_1(t *testing.T) {

	// GIVEN an assignment
	// WITH only one identifier and one expression
	// THEN only the parsed assignment is returned

	// a: 1
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}}

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "a"))},
		[]Expression{NewLiteral(tok(TK_NUMBER, "1"))},
		1,
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
	given := &tkStream{[]Token{
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
	}}

	exp := NewAssignmentBlock(
		false,
		[]Expression{
			NewIdentifier(tok(TK_IDENTIFIER, "a")),
			NewIdentifier(tok(TK_IDENTIFIER, "b")),
			NewIdentifier(tok(TK_IDENTIFIER, "c")),
		},
		[]Expression{
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_BOOL, "TRUE")),
			NewLiteral(tok(TK_STRING, "abc")),
		},
		3,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S2_3(t *testing.T) {

	// GIVEN an assignment
	// WITH a preceeding definition keyword
	// THEN parsed assignment are returned

	// a, b, c: 1, TRUE, "abc"
	given := &tkStream{[]Token{
		tok(TK_DEFINITION, "def"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}}

	exp := NewAssignmentBlock(
		true,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "a"))},
		[]Expression{NewLiteral(tok(TK_NUMBER, "1"))},
		1,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S5_1(t *testing.T) {

	// GIVEN only a comment
	// THEN no statements are returned

	// // abc
	given := &tkStream{[]Token{
		tok(TK_COMMENT, "// abc"),
		tok(TK_TERMINATOR, ""),
	}}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_2(t *testing.T) {

	// GIVEN only whitespace
	// THEN no statements are returned

	given := &tkStream{[]Token{
		tok(TK_WHITESPACE, "    "),
		tok(TK_TERMINATOR, ""),
	}}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_3(t *testing.T) {

	// GIVEN only one terminator
	// THEN no statements are returned

	given := &tkStream{[]Token{
		tok(TK_TERMINATOR, ""),
	}}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func Test_S5_4(t *testing.T) {

	// GIVEN only terminators
	// THEN no statements are returned

	given := &tkStream{[]Token{
		tok(TK_TERMINATOR, ""),
		tok(TK_TERMINATOR, ""),
		tok(TK_TERMINATOR, ""),
	}}

	exp := []Expression{}

	act, e := testFunc(given)
	expectStats(t, exp, act, e)
}

func quickOperationTest(t *testing.T, left, operator, right Token) {

	express := func(tk Token) Expression {

		switch tk.Type() {
		case TK_IDENTIFIER:
			return NewIdentifier(tk)

		case TK_BOOL, TK_NUMBER:
			return NewLiteral(tk)

		case TK_VOID:
			return NewVoid(tk)

		default:
			panic("SANITY CHECK! Unknown token type: " + tk.Type().String())
		}
	}

	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		left,
		operator,
		right,
		tok(TK_TERMINATOR, ""),
	}}

	op := NewOperation(
		operator,
		express(left),
		express(right),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}}

	op := NewOperation(
		tok(TK_PLUS, "+"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		NewNegation(
			NewLiteral(tok(TK_NUMBER, "1")),
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
		tok(TK_MINUS, "-"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}}

	left := NewOperation(
		tok(TK_PLUS, "+"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	op := NewOperation(
		tok(TK_MINUS, "-"),
		left,
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "1"),
		tok(TK_DIVIDE, "/"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}}

	left := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	op := NewOperation(
		tok(TK_DIVIDE, "/"),
		left,
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}}

	left := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	op := NewOperation(
		tok(TK_PLUS, "+"),
		left,
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}}

	left := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	op := NewOperation(
		tok(TK_PLUS, "+"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		left,
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
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
	}}

	first := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	second := NewOperation(
		tok(TK_REMAINDER, "%"),
		first,
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	third := NewOperation(
		tok(TK_MINUS, "-"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		second,
	)

	op := NewOperation(
		tok(TK_PLUS, "+"),
		third,
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
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
	}}

	first := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	second := NewOperation(
		tok(TK_REMAINDER, "%"),
		first,
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	third := NewOperation(
		tok(TK_MINUS, "-"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		second,
	)

	fourth := NewOperation(
		tok(TK_PLUS, "+"),
		third,
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	fifth := NewOperation(
		tok(TK_EQUAL, "=="),
		fourth,
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	sixth := NewOperation(
		tok(TK_MORE_THAN, ">"),
		NewIdentifier(tok(TK_IDENTIFIER, "c")),
		NewLiteral(tok(TK_NUMBER, "5")),
	)

	seventh := NewOperation(
		tok(TK_REMAINDER, "%"),
		NewIdentifier(tok(TK_IDENTIFIER, "c")),
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	eigth := NewOperation(
		tok(TK_NOT_EQUAL, "!="),
		seventh,
		NewLiteral(tok(TK_NUMBER, "0")),
	)

	ninth := NewOperation(
		tok(TK_AND, "&"),
		sixth,
		eigth,
	)

	op := NewOperation(
		tok(TK_OR, "|"),
		fifth,
		ninth,
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_21(t *testing.T) {

	quickParenTest := func(t *testing.T, op Expression, tks ...Token) {

		inTks := []Token{
			tok(TK_IDENTIFIER, "x"),
			tok(TK_ASSIGNMENT, ":="),
		}
		inTks = append(inTks, tok(TK_PAREN_OPEN, "("))
		inTks = append(inTks, tks...)
		inTks = append(inTks, tok(TK_PAREN_CLOSE, ")"))
		inTks = append(inTks, tok(TK_TERMINATOR, ""))

		given := &tkStream{inTks}

		exp := NewAssignmentBlock(
			false,
			[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
			[]Expression{op},
			1,
		)

		act, e := testFunc(given)
		expectOneStat(t, exp, act, e)
	}

	// GIVEN a prioritised operation group
	// WITH a single identifier or literal
	// THEN a single parsed expression is expected

	// (a)
	quickParenTest(t,
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		tok(TK_IDENTIFIER, "a"),
	)

	// (true)
	quickParenTest(t,
		NewLiteral(tok(TK_BOOL, "true")),
		tok(TK_BOOL, "true"),
	)

	// (1)
	quickParenTest(t,
		NewLiteral(tok(TK_NUMBER, "1")),
		tok(TK_NUMBER, "1"),
	)

	// ("abc")
	quickParenTest(t,
		NewLiteral(tok(TK_STRING, "abc")),
		tok(TK_STRING, "abc"),
	)

	// (-1)
	quickParenTest(t,
		NewNegation(
			NewLiteral(tok(TK_NUMBER, "1")),
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}}

	first := NewOperation(
		tok(TK_PLUS, "+"),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	op := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		first,
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "1"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
	}}

	first := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	op := NewOperation(
		tok(TK_PLUS, "+"),
		first,
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
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
	}}

	first := NewOperation(
		tok(TK_REMAINDER, "%"),
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	second := NewOperation(
		tok(TK_EQUAL, "=="),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	third := NewOperation(
		tok(TK_MORE_THAN, ">"),
		NewIdentifier(tok(TK_IDENTIFIER, "c")),
		NewLiteral(tok(TK_NUMBER, "5")),
	)

	fourth := NewOperation(
		tok(TK_OR, "|"),
		second,
		third,
	)

	fifth := NewOperation(
		tok(TK_PLUS, "+"),
		first,
		fourth,
	)

	sixth := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewLiteral(tok(TK_NUMBER, "1")),
		fifth,
	)

	seventh := NewOperation(
		tok(TK_MINUS, "-"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		sixth,
	)

	eigth := NewOperation(
		tok(TK_REMAINDER, "%"),
		NewIdentifier(tok(TK_IDENTIFIER, "c")),
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	ninth := NewOperation(
		tok(TK_AND, "&"),
		seventh,
		eigth,
	)

	op := NewOperation(
		tok(TK_NOT_EQUAL, "!="),
		ninth,
		NewLiteral(tok(TK_NUMBER, "0")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		NewBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{
				tok(TK_IDENTIFIER, "a"),
			},
			[]Token{},
		),
		NewBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
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
	}}

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{
				tok(TK_IDENTIFIER, "a"),
			},
		),
		NewBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
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
	}}

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
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
		NewBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
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
	}}

	body := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{
			NewAssignmentBlock(
				false,
				[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "a"))},
				[]Expression{NewLiteral(tok(TK_NUMBER, "1"))},
				1,
			),
		},
	)

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		body,
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
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
	}}

	body := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{
			NewAssignmentBlock(
				false,
				[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "a"))},
				[]Expression{NewLiteral(tok(TK_NUMBER, "1"))},
				1,
			),
		},
	)

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
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

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_7(t *testing.T) {

	// GIVEN a function
	// WITH no parameters
	// AND a watch as the body
	// THEN the correctly parsed function is returned

	// f := F() {}
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		NewWatch(
			tok(TK_WATCH, "watch"),
			[]Token{
				tok(TK_IDENTIFIER, "a"),
			},
			NewBlock(
				tok(TK_BLOCK_OPEN, "{"),
				tok(TK_BLOCK_CLOSE, "}"),
				[]Expression{},
			),
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_8(t *testing.T) {

	// GIVEN a function
	// WITH no parameters
	// AND a when as the body
	// THEN the correctly parsed function is returned

	// f := F() {}
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_WHEN, "when"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		NewWhen(
			tok(TK_WHEN, "when"),
			tok(TK_BLOCK_CLOSE, "}"),
			NewAssignment(
				NewIdentifier(tok(TK_IDENTIFIER, "a")),
				NewLiteral(tok(TK_NUMBER, "1")),
			),
			[]WhenCase{},
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S7_9(t *testing.T) {

	// GIVEN a function
	// WITH no parameters
	// AND a guard as the body
	// THEN the correctly parsed function is returned

	// f := F() {}
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_EQUAL, "=="),
		tok(TK_NUMBER, "1"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	f := NewFunction(
		tok(TK_FUNCTION, "F"),
		NewParameters(
			tok(TK_PAREN_OPEN, "("),
			tok(TK_PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		NewGuard(
			tok(TK_GUARD_OPEN, "["),
			NewOperation(
				tok(TK_EQUAL, "=="),
				NewLiteral(tok(TK_NUMBER, "1")),
				NewLiteral(tok(TK_NUMBER, "1")),
			),
			NewBlock(
				tok(TK_BLOCK_OPEN, "{"),
				tok(TK_BLOCK_CLOSE, "}"),
				[]Expression{},
			),
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}}

	f := NewExpressionFunction(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{},
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}}

	f := NewExpressionFunction(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
		},
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
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
	}}

	f := NewExpressionFunction(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
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
	}}

	expr := NewOperation(
		tok(TK_PLUS, "+"),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewOperation(
			tok(TK_MULTIPLY, "*"),
			NewLiteral(tok(TK_NUMBER, "2")),
			NewLiteral(tok(TK_NUMBER, "3")),
		),
	)

	f := NewExpressionFunction(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{},
		expr,
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "f"))},
		[]Expression{f},
		1,
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
	given := &tkStream{[]Token{
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	body := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := NewWatch(
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
	given := &tkStream{[]Token{
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	body := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := NewWatch(
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
	given := &tkStream{[]Token{
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
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
	}}

	op := NewOperation(
		tok(TK_PLUS, "+"),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	first := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{op},
		1,
	)

	second := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "a"))},
		[]Expression{NewLiteral(tok(TK_NUMBER, "3"))},
		1,
	)

	body := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{
			first,
			second,
		},
	)

	exp := NewWatch(
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
	given := &tkStream{[]Token{
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	body := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewLiteral(tok(TK_BOOL, "true")),
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
	given := &tkStream{[]Token{
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_EQUAL, "=="),
		tok(TK_NUMBER, "2"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	condition := NewOperation(
		tok(TK_EQUAL, "=="),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	body := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)

	exp := NewGuard(
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
	given := &tkStream{[]Token{
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, ""),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "3"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "4"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	first := NewAssignmentBlock(
		false,
		[]Expression{
			NewIdentifier(tok(TK_IDENTIFIER, "x")),
		},
		[]Expression{
			NewOperation(
				tok(TK_PLUS, "+"),
				NewLiteral(tok(TK_NUMBER, "1")),
				NewLiteral(tok(TK_NUMBER, "2")),
			),
		},
		1,
	)

	second := NewAssignmentBlock(
		false,
		[]Expression{
			NewIdentifier(tok(TK_IDENTIFIER, "x")),
		},
		[]Expression{
			NewOperation(
				tok(TK_MULTIPLY, "*"),
				NewLiteral(tok(TK_NUMBER, "3")),
				NewLiteral(tok(TK_NUMBER, "4")),
			),
		},
		1,
	)

	statements := []Expression{
		first,
		second,
	}

	body := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		statements,
	)

	exp := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewLiteral(tok(TK_BOOL, "true")),
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
	given := &tkStream{[]Token{
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_EXIT, "exit"),
		tok(TK_NUMBER, "0"),
		tok(TK_TERMINATOR, ""),
	}}

	body := NewUnDelimiteredBlock(
		[]Expression{
			NewExit(
				tok(TK_EXIT, "exit"),
				NewLiteral(tok(TK_NUMBER, "0")),
			),
		},
	)

	exp := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewLiteral(tok(TK_BOOL, "true")),
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
	given := &tkStream{[]Token{
		tok(TK_WHEN, "when"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_BOOL, "true"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	init := NewAssignment(
		NewIdentifier(tok(TK_IDENTIFIER, "x")),
		NewLiteral(tok(TK_BOOL, "true")),
	)

	exp := NewWhen(
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
	given := &tkStream{[]Token{
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
	}}

	condition := NewOperation(
		tok(TK_EQUAL, "=="),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewLiteral(tok(TK_NUMBER, "2")),
	)

	init := NewAssignment(
		NewIdentifier(tok(TK_IDENTIFIER, "x")),
		condition,
	)

	exp := NewWhen(
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
	given := &tkStream{[]Token{
		tok(TK_WHEN, "when"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_NUMBER, "1"),
		tok(TK_THEN, "->"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "3"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "4"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	firstCase := NewLiteral(tok(TK_NUMBER, "1"))

	op := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewLiteral(tok(TK_NUMBER, "3")),
		NewLiteral(tok(TK_NUMBER, "4")),
	)

	firstBlock := NewUnDelimiteredBlock(
		[]Expression{
			NewAssignmentBlock(
				false,
				[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
				[]Expression{op},
				1,
			),
		},
	)

	init := NewAssignment(
		NewIdentifier(tok(TK_IDENTIFIER, "x")),
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	exp := NewWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "}"),
		init,
		[]WhenCase{
			NewWhenCase(firstCase, firstBlock),
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
	given := &tkStream{[]Token{
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
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "3"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "4"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	op := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewLiteral(tok(TK_NUMBER, "3")),
		NewLiteral(tok(TK_NUMBER, "4")),
	)

	firstBlock := NewUnDelimiteredBlock(
		[]Expression{
			NewAssignmentBlock(
				false,
				[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
				[]Expression{op},
				1,
			),
		},
	)

	firstCase := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewOperation(
			tok(TK_EQUAL, "=="),
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_NUMBER, "2")),
		),
		firstBlock,
	)

	init := NewAssignment(
		NewIdentifier(tok(TK_IDENTIFIER, "x")),
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	exp := NewWhen(
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
	given := &tkStream{[]Token{
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
	}}

	firstCase := NewWhenCase(
		NewLiteral(tok(TK_NUMBER, "1")),
		NewUnDelimiteredBlock(
			[]Expression{
				NewIdentifier(tok(TK_IDENTIFIER, "a")),
			},
		),
	)

	secondCase := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewOperation(
			tok(TK_EQUAL, "=="),
			NewIdentifier(tok(TK_IDENTIFIER, "a")),
			NewIdentifier(tok(TK_IDENTIFIER, "b")),
		),
		NewUnDelimiteredBlock(
			[]Expression{
				NewIdentifier(tok(TK_IDENTIFIER, "b")),
			},
		),
	)

	thirdCase := NewWhenCase(
		NewLiteral(tok(TK_NUMBER, "2")),
		NewUnDelimiteredBlock(
			[]Expression{
				NewIdentifier(tok(TK_IDENTIFIER, "c")),
			},
		),
	)

	fourthCase := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewLiteral(tok(TK_BOOL, "true")),
		NewBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{
				NewIdentifier(tok(TK_IDENTIFIER, "d")),
			},
		),
	)

	init := NewAssignment(
		NewIdentifier(tok(TK_IDENTIFIER, "x")),
		NewLiteral(tok(TK_NUMBER, "3")),
	)

	exp := NewWhen(
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
	given := &tkStream{[]Token{
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
	}}

	init := NewAssignment(
		NewIdentifier(tok(TK_IDENTIFIER, "i")),
		NewLiteral(tok(TK_NUMBER, "0")),
	)

	guard := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewLiteral(tok(TK_BOOL, "true")),
		NewBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{},
		),
	)

	exp := NewLoop(
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
	given := &tkStream{[]Token{
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
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, ""),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_IDENTIFIER, "d"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	init := NewAssignment(
		NewIdentifier(tok(TK_IDENTIFIER, "i")),
		NewOperation(
			tok(TK_MINUS, "-"),
			NewIdentifier(tok(TK_IDENTIFIER, "a")),
			NewLiteral(tok(TK_NUMBER, "1")),
		),
	)

	condition := NewOperation(
		tok(TK_LESS_THAN, "<"),
		NewIdentifier(tok(TK_IDENTIFIER, "i")),
		NewLiteral(tok(TK_NUMBER, "10")),
	)

	firstOp := NewOperation(
		tok(TK_PLUS, "+"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		NewIdentifier(tok(TK_IDENTIFIER, "b")),
	)

	firstStat := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{firstOp},
		1,
	)

	secondOp := NewOperation(
		tok(TK_MULTIPLY, "*"),
		NewIdentifier(tok(TK_IDENTIFIER, "c")),
		NewIdentifier(tok(TK_IDENTIFIER, "d")),
	)

	secondStat := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{secondOp},
		1,
	)

	guard := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		condition,
		NewBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expression{
				firstStat,
				secondStat,
			},
		),
	)

	exp := NewLoop(
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}}

	var f Expression = NewIdentifier(tok(TK_IDENTIFIER, "f"))

	exp := NewFunctionCall(
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}}

	var f Expression = NewIdentifier(tok(TK_IDENTIFIER, "f"))

	args := []Expression{
		NewLiteral(tok(TK_NUMBER, "1")),
	}

	exp := NewFunctionCall(
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}}

	var f Expression = NewIdentifier(tok(TK_IDENTIFIER, "f"))

	args := []Expression{
		NewOperation(
			tok(TK_PLUS, "+"),
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_NUMBER, "2")),
		),
	}

	exp := NewFunctionCall(
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
	given := &tkStream{[]Token{
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
	}}

	var f Expression = NewIdentifier(tok(TK_IDENTIFIER, "f"))

	args := []Expression{
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		NewLiteral(tok(TK_BOOL, "true")),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewLiteral(tok(TK_STRING, "abc")),
	}

	exp := NewFunctionCall(
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
	given := &tkStream{[]Token{
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
	}}

	var f Expression = NewIdentifier(tok(TK_IDENTIFIER, "f"))

	args := []Expression{
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
		NewLiteral(tok(TK_BOOL, "true")),
		NewLiteral(tok(TK_NUMBER, "1")),
		NewLiteral(tok(TK_STRING, "abc")),
	}

	exp := NewFunctionCall(
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
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "(a"),
		tok(TK_PAREN_CLOSE, "a)"),
		tok(TK_PAREN_OPEN, "(b"),
		tok(TK_PAREN_CLOSE, "b)"),
		tok(TK_PAREN_OPEN, "(c"),
		tok(TK_PAREN_CLOSE, "c)"),
		tok(TK_TERMINATOR, ""),
	}}

	var f Expression = NewIdentifier(tok(TK_IDENTIFIER, "f"))

	var first Expression = NewFunctionCall(
		tok(TK_PAREN_CLOSE, "a)"),
		f,
		[]Expression{},
	)

	var second Expression = NewFunctionCall(
		tok(TK_PAREN_CLOSE, "b)"),
		first,
		[]Expression{},
	)

	exp := NewFunctionCall(
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
	given := &tkStream{[]Token{
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}}

	exp := NewSpellCall(
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
	given := &tkStream{[]Token{
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_OPEN, "(a"),
		tok(TK_PAREN_CLOSE, "a)"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_PAREN_OPEN, "(c"),
		tok(TK_PAREN_CLOSE, "c)"),
		tok(TK_TERMINATOR, ""),
	}}

	var first Expression = NewSpellCall(
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_CLOSE, "a)"),
		[]Expression{},
	)

	var second Expression = NewCollectionAccessor(
		first,
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	exp := NewFunctionCall(
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
	given := &tkStream{[]Token{
		tok(TK_SPELL, "@s"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_STRING, "abc"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}}

	first := NewLiteral(tok(TK_NUMBER, "1"))

	assign := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{NewLiteral(tok(TK_STRING, "abc"))},
		1,
	)

	second := NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{assign},
	)

	exp := NewSpellCall(
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
	given := &tkStream{[]Token{
		tok(TK_EXIT, "exit"),
		tok(TK_NUMBER, "0"),
		tok(TK_TERMINATOR, ""),
	}}

	exp := NewExit(
		tok(TK_EXIT, "exit"),
		NewLiteral(tok(TK_NUMBER, "0")),
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S16_1(t *testing.T) {

	// GIVEN an exists test
	// WITH a literal subject
	// THEN then the correct exists statement is returned

	// 0?
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "0"),
		tok(TK_EXISTS, "?"),
		tok(TK_TERMINATOR, ""),
	}}

	exists := NewExists(
		tok(TK_EXISTS, "?"),
		NewLiteral(tok(TK_NUMBER, "0")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{exists},
		1,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S16_2(t *testing.T) {

	// GIVEN an exists test
	// WITH an identifier subject
	// THEN then the correct exists statement is returned

	// a?
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_EXISTS, "?"),
		tok(TK_TERMINATOR, ""),
	}}

	exists := NewExists(
		tok(TK_EXISTS, "?"),
		NewIdentifier(tok(TK_IDENTIFIER, "a")),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{exists},
		1,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_S16_3(t *testing.T) {

	// GIVEN an exists test
	// WITH a group as the subject
	// THEN then the correct exists statement is returned

	// (a + b)?
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_PAREN_CLOSE, "("),
		tok(TK_EXISTS, "?"),
		tok(TK_TERMINATOR, ""),
	}}

	exists := NewExists(
		tok(TK_EXISTS, "?"),
		NewOperation(
			tok(TK_PLUS, "+"),
			NewIdentifier(tok(TK_IDENTIFIER, "a")),
			NewIdentifier(tok(TK_IDENTIFIER, "b")),
		),
	)

	exp := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(TK_IDENTIFIER, "x"))},
		[]Expression{exists},
		1,
	)

	act, e := testFunc(given)
	expectOneStat(t, exp, act, e)
}

func Test_F1(t *testing.T) {

	// GIVEN an invalid statement or expression starting token
	// THEN parser returns error

	given := &tkStream{[]Token{
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F2(t *testing.T) {

	// GIVEN an assignment
	// WITHOUT enough expressions
	// THEN parser returns error

	// a:
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F3(t *testing.T) {

	// GIVEN an assignment
	// WITHOUT enough identifiers
	// THEN parser returns error

	// a: 1, 2
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F4(t *testing.T) {

	// GIVEN an assignment
	// WITH the assignment token missing
	// THEN parser returns error

	// a 1
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F5(t *testing.T) {

	// GIVEN an assignment
	// WITH an delimiter token missing from the assignment targets
	// THEN parser returns error

	// a b: 1, 2
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F6(t *testing.T) {

	// GIVEN an assignment
	// WITH an delimiter token missing from the expressions
	// THEN parser returns error

	// a, b: 1 2
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F7(t *testing.T) {

	// GIVEN a negation prefix
	// WITHOUT a following expression
	// THEN parser returns error

	// -
	given := &tkStream{[]Token{
		tok(TK_MINUS, "-"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F9(t *testing.T) {

	// GIVEN a list
	// WITHOUT an expression or block close following the block open
	// THEN parser returns error

	// LIST {,1}
	given := &tkStream{[]Token{
		tok(TK_LIST, "LIST"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F10(t *testing.T) {

	// GIVEN a list
	// WITHOUT a delimiter after an expression but with a terminator
	// THEN parser returns error

	// LIST {1
	// }
	given := &tkStream{[]Token{
		tok(TK_LIST, "LIST"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F11(t *testing.T) {

	// GIVEN a list
	// WITHOUT a block open
	// THEN parser returns error

	// LIST 1}
	given := &tkStream{[]Token{
		tok(TK_LIST, "LIST"),
		tok(TK_NUMBER, "1"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F12(t *testing.T) {

	// GIVEN a list
	// WITHOUT a block close
	// THEN parser returns error

	// LIST {1
	given := &tkStream{[]Token{
		tok(TK_LIST, "LIST"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F13(t *testing.T) {

	// GIVEN an operation
	// WITHOUT a right operand
	// THEN parser returns error

	// x +
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_PLUS, "+"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F14(t *testing.T) {

	// GIVEN an operation
	// WITH two operators
	// THEN parser returns error

	// x + +
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_PLUS, "+"),
		tok(TK_PLUS, "+"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F15(t *testing.T) {

	// GIVEN an expression function
	// WITH no expression
	// THEN parser returns error

	// f: E()
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F16(t *testing.T) {

	// GIVEN a watch block
	// WITH no body
	// THEN parser returns error

	// watch a
	given := &tkStream{[]Token{
		tok(TK_WATCH, "watch"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F17(t *testing.T) {

	// GIVEN a guard block
	// WITH no body
	// THEN parser returns error

	// watch a
	given := &tkStream{[]Token{
		tok(TK_GUARD_OPEN, "["),
		tok(TK_BOOL, "true"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F18(t *testing.T) {

	// GIVEN a guard block
	// WITH no condition
	// THEN parser returns error

	// watch a
	given := &tkStream{[]Token{
		tok(TK_GUARD_OPEN, "["),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F19(t *testing.T) {

	// GIVEN a function call
	// WITH no closing parenthesis
	// THEN parser returns error

	// f(loop
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_LOOP, "loop"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}

func Test_F20(t *testing.T) {

	// GIVEN a function call
	// WITH a missig delimiter between arguments
	// THEN parser returns error

	// f(1, 2 3)
	given := &tkStream{[]Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_NUMBER, "3"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}}

	act, e := testFunc(given)
	expectError(t, act, e)
}
