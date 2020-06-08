package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var testFunc func(Factory, []Token) ([]Statement, error) = ParseStatements
var testFac Factory = NewFactory()

func Test_S1_1(t *testing.T) {

	quickSoloTokenTest := func(t *testing.T, exp Statement, tk Token) {

		var given []Token
		given = append(given, tk)
		given = append(given, tok(TERMINATOR, ""))

		act, e := testFunc(testFac, given)
		expectOneStat(t, exp, act, e)
	}

	// GIVEN an identifier only
	// THEN only an identifier expression is returned

	// a
	quickSoloTokenTest(t,
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		tok(IDENTIFIER, "a"),
	)

	// true
	quickSoloTokenTest(t,
		testFac.NewLiteral(tok(BOOL, "true")),
		tok(BOOL, "true"),
	)

	// 1
	quickSoloTokenTest(t,
		testFac.NewLiteral(tok(NUMBER, "1")),
		tok(NUMBER, "1"),
	)

	// abc
	quickSoloTokenTest(t,
		testFac.NewLiteral(tok(STRING, "abc")),
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

	exp := testFac.NewNonWrappedBlock(
		[]Statement{
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "a")),
				testFac.NewLiteral(tok(NUMBER, "1")),
			),
		},
	)

	act, e := testFunc(testFac, given)
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

	exp := testFac.NewNonWrappedBlock(
		[]Statement{
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "a")),
				testFac.NewLiteral(tok(NUMBER, "1")),
			),
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "b")),
				testFac.NewLiteral(tok(BOOL, "TRUE")),
			),
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "c")),
				testFac.NewLiteral(tok(STRING, "abc")),
			),
		},
	)

	act, e := testFunc(testFac, given)
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

	exp := testFac.NewNegation(
		testFac.NewLiteral(tok(NUMBER, "2")),
	)

	act, e := testFunc(testFac, given)
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

	exp := testFac.NewListAccessor(
		Identifier{tok(IDENTIFIER, "abc")},
		Literal{tok(NUMBER, "1")},
	)

	act, e := testFunc(testFac, given)
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

	exp := []Statement{}

	act, e := testFunc(testFac, given)
	expectStats(t, exp, act, e)
}

func Test_S5_2(t *testing.T) {

	// GIVEN only whitespace
	// THEN no statements are returned

	given := []Token{
		tok(WHITESPACE, "    "),
		tok(TERMINATOR, ""),
	}

	exp := []Statement{}

	act, e := testFunc(testFac, given)
	expectStats(t, exp, act, e)
}

func Test_S5_3(t *testing.T) {

	// GIVEN only one terminator
	// THEN no statements are returned

	given := []Token{
		tok(TERMINATOR, ""),
	}

	exp := []Statement{}

	act, e := testFunc(testFac, given)
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

	exp := []Statement{}

	act, e := testFunc(testFac, given)
	expectStats(t, exp, act, e)
}

func quickOperationTest(t *testing.T, left, operator, right Token) {

	express := func(tk Token) Expression {
		switch tk.Morpheme() {
		case IDENTIFIER:
			return testFac.NewIdentifier(tk)
		case BOOL, NUMBER:
			return testFac.NewLiteral(tk)
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

	exp := testFac.NewOperation(
		operator,
		express(left),
		express(right),
	)

	act, e := testFunc(testFac, given)
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

	exp := testFac.NewOperation(
		tok(ADD, "+"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		testFac.NewNegation(
			testFac.NewLiteral(tok(NUMBER, "1")),
		),
	)

	act, e := testFunc(testFac, given)
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

	left := testFac.NewOperation(
		tok(ADD, "+"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		testFac.NewLiteral(tok(NUMBER, "1")),
	)

	exp := testFac.NewOperation(
		tok(SUBTRACT, "-"),
		left,
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
	)

	act, e := testFunc(testFac, given)
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

	left := testFac.NewOperation(
		tok(MULTIPLY, "*"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		testFac.NewLiteral(tok(NUMBER, "1")),
	)

	exp := testFac.NewOperation(
		tok(DIVIDE, "/"),
		left,
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
	)

	act, e := testFunc(testFac, given)
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

	left := testFac.NewOperation(
		tok(MULTIPLY, "*"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		testFac.NewLiteral(tok(NUMBER, "1")),
	)

	exp := testFac.NewOperation(
		tok(ADD, "+"),
		left,
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
	)

	act, e := testFunc(testFac, given)
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

	left := testFac.NewOperation(
		tok(MULTIPLY, "*"),
		testFac.NewLiteral(tok(NUMBER, "1")),
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
	)

	exp := testFac.NewOperation(
		tok(ADD, "+"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		left,
	)

	act, e := testFunc(testFac, given)
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

	first := testFac.NewOperation(
		tok(MULTIPLY, "*"),
		testFac.NewLiteral(tok(NUMBER, "1")),
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
	)

	second := testFac.NewOperation(
		tok(REMAINDER, "%"),
		first,
		testFac.NewLiteral(tok(NUMBER, "2")),
	)

	third := testFac.NewOperation(
		tok(SUBTRACT, "-"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		second,
	)

	exp := testFac.NewOperation(
		tok(ADD, "+"),
		third,
		testFac.NewLiteral(tok(NUMBER, "1")),
	)

	act, e := testFunc(testFac, given)
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

	first := testFac.NewOperation(
		tok(MULTIPLY, "*"),
		testFac.NewLiteral(tok(NUMBER, "1")),
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
	)

	second := testFac.NewOperation(
		tok(REMAINDER, "%"),
		first,
		testFac.NewLiteral(tok(NUMBER, "2")),
	)

	third := testFac.NewOperation(
		tok(SUBTRACT, "-"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		second,
	)

	fourth := testFac.NewOperation(
		tok(ADD, "+"),
		third,
		testFac.NewLiteral(tok(NUMBER, "1")),
	)

	fifth := testFac.NewOperation(
		tok(EQUAL, "=="),
		fourth,
		testFac.NewLiteral(tok(NUMBER, "2")),
	)

	sixth := testFac.NewOperation(
		tok(MORE_THAN, ">"),
		testFac.NewIdentifier(tok(IDENTIFIER, "c")),
		testFac.NewLiteral(tok(NUMBER, "5")),
	)

	seventh := testFac.NewOperation(
		tok(REMAINDER, "%"),
		testFac.NewIdentifier(tok(IDENTIFIER, "c")),
		testFac.NewLiteral(tok(NUMBER, "2")),
	)

	eigth := testFac.NewOperation(
		tok(NOT_EQUAL, "!="),
		seventh,
		testFac.NewLiteral(tok(NUMBER, "0")),
	)

	ninth := testFac.NewOperation(
		tok(AND, "&"),
		sixth,
		eigth,
	)

	exp := testFac.NewOperation(
		tok(OR, "|"),
		fifth,
		ninth,
	)

	act, e := testFunc(testFac, given)
	expectOneStat(t, exp, act, e)
}

func Test_S6_21(t *testing.T) {

	quickParenTest := func(t *testing.T, exp Statement, tks ...Token) {

		var given []Token
		given = append(given, tok(PAREN_OPEN, "("))
		given = append(given, tks...)
		given = append(given, tok(PAREN_CLOSE, ")"))
		given = append(given, tok(TERMINATOR, ""))

		act, e := testFunc(testFac, given)
		expectOneStat(t, exp, act, e)
	}

	// GIVEN a prioritised operation group
	// WITH a single identifier or literal
	// THEN a single parsed expression is expected

	// (a)
	quickParenTest(t,
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		tok(IDENTIFIER, "a"),
	)

	// (true)
	quickParenTest(t,
		testFac.NewLiteral(tok(BOOL, "true")),
		tok(BOOL, "true"),
	)

	// (1)
	quickParenTest(t,
		testFac.NewLiteral(tok(NUMBER, "1")),
		tok(NUMBER, "1"),
	)

	// ("abc")
	quickParenTest(t,
		testFac.NewLiteral(tok(STRING, "abc")),
		tok(STRING, "abc"),
	)

	// (-1)
	quickParenTest(t,
		testFac.NewNegation(
			testFac.NewLiteral(tok(NUMBER, "1")),
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

	first := testFac.NewOperation(
		tok(ADD, "+"),
		testFac.NewLiteral(tok(NUMBER, "1")),
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
	)

	exp := testFac.NewOperation(
		tok(MULTIPLY, "*"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		first,
	)

	act, e := testFunc(testFac, given)
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

	first := testFac.NewOperation(
		tok(MULTIPLY, "*"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		testFac.NewLiteral(tok(NUMBER, "1")),
	)

	exp := testFac.NewOperation(
		tok(ADD, "+"),
		first,
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
	)

	act, e := testFunc(testFac, given)
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

	first := testFac.NewOperation(
		tok(REMAINDER, "%"),
		testFac.NewIdentifier(tok(IDENTIFIER, "b")),
		testFac.NewLiteral(tok(NUMBER, "2")),
	)

	second := testFac.NewOperation(
		tok(EQUAL, "=="),
		testFac.NewLiteral(tok(NUMBER, "1")),
		testFac.NewLiteral(tok(NUMBER, "2")),
	)

	third := testFac.NewOperation(
		tok(MORE_THAN, ">"),
		testFac.NewIdentifier(tok(IDENTIFIER, "c")),
		testFac.NewLiteral(tok(NUMBER, "5")),
	)

	fourth := testFac.NewOperation(
		tok(OR, "|"),
		second,
		third,
	)

	fifth := testFac.NewOperation(
		tok(ADD, "+"),
		first,
		fourth,
	)

	sixth := testFac.NewOperation(
		tok(MULTIPLY, "*"),
		testFac.NewLiteral(tok(NUMBER, "1")),
		fifth,
	)

	seventh := testFac.NewOperation(
		tok(SUBTRACT, "-"),
		testFac.NewIdentifier(tok(IDENTIFIER, "a")),
		sixth,
	)

	eigth := testFac.NewOperation(
		tok(REMAINDER, "%"),
		testFac.NewIdentifier(tok(IDENTIFIER, "c")),
		testFac.NewLiteral(tok(NUMBER, "2")),
	)

	ninth := testFac.NewOperation(
		tok(AND, "&"),
		seventh,
		eigth,
	)

	exp := testFac.NewOperation(
		tok(NOT_EQUAL, "!="),
		ninth,
		testFac.NewLiteral(tok(NUMBER, "0")),
	)

	act, e := testFunc(testFac, given)
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

	f := testFac.NewFunction(
		tok(FUNC, "F"),
		testFac.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		testFac.NewBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Statement{},
		),
	)

	exp := testFac.NewNonWrappedBlock(
		[]Statement{
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFac, given)
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

	f := testFac.NewFunction(
		tok(FUNC, "F"),
		testFac.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{
				tok(IDENTIFIER, "a"),
			},
			[]Token{},
		),
		testFac.NewBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Statement{},
		),
	)

	exp := testFac.NewNonWrappedBlock(
		[]Statement{
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFac, given)
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

	f := testFac.NewFunction(
		tok(FUNC, "F"),
		testFac.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{
				tok(IDENTIFIER, "a"),
			},
		),
		testFac.NewBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Statement{},
		),
	)

	exp := testFac.NewNonWrappedBlock(
		[]Statement{
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFac, given)
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

	f := testFac.NewFunction(
		tok(FUNC, "F"),
		testFac.NewParameters(
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
		testFac.NewBlock(
			tok(BLOCK_OPEN, "{"),
			tok(BLOCK_CLOSE, "}"),
			[]Statement{},
		),
	)

	exp := testFac.NewNonWrappedBlock(
		[]Statement{
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFac, given)
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

	body := testFac.NewBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Statement{
			testFac.NewNonWrappedBlock(
				[]Statement{
					testFac.NewAssignment(
						testFac.NewIdentifier(tok(IDENTIFIER, "a")),
						testFac.NewLiteral(tok(NUMBER, "1")),
					),
				},
			),
		},
	)

	f := testFac.NewFunction(
		tok(FUNC, "F"),
		testFac.NewParameters(
			tok(PAREN_OPEN, "("),
			tok(PAREN_CLOSE, ")"),
			[]Token{},
			[]Token{},
		),
		body,
	)

	exp := testFac.NewNonWrappedBlock(
		[]Statement{
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFac, given)
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

	body := testFac.NewBlock(
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
		[]Statement{
			testFac.NewNonWrappedBlock(
				[]Statement{
					testFac.NewAssignment(
						testFac.NewIdentifier(tok(IDENTIFIER, "a")),
						testFac.NewLiteral(tok(NUMBER, "1")),
					),
				},
			),
		},
	)

	f := testFac.NewFunction(
		tok(FUNC, "F"),
		testFac.NewParameters(
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

	exp := testFac.NewNonWrappedBlock(
		[]Statement{
			testFac.NewAssignment(
				testFac.NewIdentifier(tok(IDENTIFIER, "f")),
				f,
			),
		},
	)

	act, e := testFunc(testFac, given)
	expectOneStat(t, exp, act, e)
}

func Test_F1(t *testing.T) {

	// GIVEN an invalid statement or expression starting token
	// THEN parser returns error

	given := []Token{
		tok(ASSIGN, ":="),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
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

	act, e := testFunc(testFac, given)
	expectError(t, act, e)
}
