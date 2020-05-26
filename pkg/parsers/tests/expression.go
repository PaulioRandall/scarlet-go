package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func E1_Add(t *testing.T, f ParseFunc) {

	// 1 + 2

	given := []Token{
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "1")},
		tok(ADD, "+"),
		Value{tok(NUMBER, "2")},
	}

	expectOneStat(t, exp, f(given))
}

func E2_Subtract(t *testing.T, f ParseFunc) {

	// 2 - 1

	given := []Token{
		tok(NUMBER, "2"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "2")},
		tok(SUBTRACT, "-"),
		Value{tok(NUMBER, "1")},
	}

	expectOneStat(t, exp, f(given))
}

func E3_Multiply(t *testing.T, f ParseFunc) {

	// 6 * 7

	given := []Token{
		tok(NUMBER, "6"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "7"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "6")},
		tok(MULTIPLY, "*"),
		Value{tok(NUMBER, "7")},
	}

	expectOneStat(t, exp, f(given))
}

func E4_Divide(t *testing.T, f ParseFunc) {

	// 12 / 3

	given := []Token{
		tok(NUMBER, "12"),
		tok(DIVIDE, "/"),
		tok(NUMBER, "3"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "12")},
		tok(DIVIDE, "/"),
		Value{tok(NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E5_AdditiveOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 - 3

	given := []Token{
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "3"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "1")},
		tok(ADD, "+"),
		Value{tok(NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(SUBTRACT, "-"),
		Value{tok(NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E6_AdditiveOrdering(t *testing.T, f ParseFunc) {

	// 1 - 2 + 3

	given := []Token{
		tok(NUMBER, "1"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "2"),
		tok(ADD, "+"),
		tok(NUMBER, "3"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "1")},
		tok(SUBTRACT, "-"),
		Value{tok(NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(ADD, "+"),
		Value{tok(NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E7_MultiplicativeOrdering(t *testing.T, f ParseFunc) {

	// 1 * 2 % 3 4

	given := []Token{
		tok(NUMBER, "1"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "2"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "3"),
		tok(DIVIDE, "/"),
		tok(NUMBER, "4"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "1")},
		tok(MULTIPLY, "*"),
		Value{tok(NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(REMAINDER, "%"),
		Value{tok(NUMBER, "3")},
	}

	exp = Operation{
		exp,
		tok(DIVIDE, "/"),
		Value{tok(NUMBER, "4")},
	}

	expectOneStat(t, exp, f(given))
}

func E8_MultiplicativeOrdering(t *testing.T, f ParseFunc) {

	// 1 % 2 / 3 * 4

	given := []Token{
		tok(NUMBER, "1"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "2"),
		tok(DIVIDE, "/"),
		tok(NUMBER, "3"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "4"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "1")},
		tok(REMAINDER, "%"),
		Value{tok(NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(DIVIDE, "/"),
		Value{tok(NUMBER, "3")},
	}

	exp = Operation{
		exp,
		tok(MULTIPLY, "*"),
		Value{tok(NUMBER, "4")},
	}

	expectOneStat(t, exp, f(given))
}

func E9_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 * 2 + 3

	given := []Token{
		tok(NUMBER, "1"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "2"),
		tok(ADD, "+"),
		tok(NUMBER, "3"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(NUMBER, "1")},
		tok(MULTIPLY, "*"),
		Value{tok(NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(ADD, "+"),
		Value{tok(NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E10_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 * 3

	given := []Token{
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "3"),
		tok(TERMINATOR, ""),
	}

	// 1 + (2 * 3)

	exp := Operation{
		Left:     Value{tok(NUMBER, "1")},
		Operator: tok(ADD, "+"),
	}

	exp.Right = Operation{
		Value{tok(NUMBER, "2")},
		tok(MULTIPLY, "*"),
		Value{tok(NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E11_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 * 3 - 4 % 5 / 6

	given := []Token{
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(MULTIPLY, "*"),
		tok(NUMBER, "3"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "4"),
		tok(REMAINDER, "%"),
		tok(NUMBER, "5"),
		tok(DIVIDE, "/"),
		tok(NUMBER, "6"),
		tok(TERMINATOR, ""),
	}

	// 1 + (2 * 3) - ((4 % 5) / 6)

	add := Operation{
		Left:     Value{tok(NUMBER, "1")},
		Operator: tok(ADD, "+"),
	}

	mul := Operation{
		Value{tok(NUMBER, "2")},
		tok(MULTIPLY, "*"),
		Value{tok(NUMBER, "3")},
	}

	add.Right = mul

	sub := Operation{
		Left:     add,
		Operator: tok(SUBTRACT, "-"),
	}

	rem := Operation{
		Value{tok(NUMBER, "4")},
		tok(REMAINDER, "%"),
		Value{tok(NUMBER, "5")},
	}

	div := Operation{
		rem,
		tok(DIVIDE, "/"),
		Value{tok(NUMBER, "6")},
	}

	sub.Right = div

	expectOneStat(t, sub, f(given))
}

func E12_FuncCall(t *testing.T, f ParseFunc) {

	// 1 + f(a,b)

	given := []Token{
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, ""),
	}

	exp := Operation{
		Left:     Value{tok(NUMBER, "1")},
		Operator: tok(ADD, "+"),
	}

	exp.Right = FuncCall{
		Identifier{tok(IDENTIFIER, "f")},
		[]Expression{
			Identifier{tok(IDENTIFIER, "a")},
			Identifier{tok(IDENTIFIER, "b")},
		},
	}

	expectOneStat(t, exp, f(given))
}

func E13_Panics(t *testing.T, f ParseFunc) {

	// 1 + +

	given := []Token{
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(ADD, "+"),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func E14_Panics(t *testing.T, f ParseFunc) {

	// + 1 + 1

	given := []Token{
		tok(ADD, "+"),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func E15_Negation(t *testing.T, f ParseFunc) {

	// -1

	given := []Token{
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := Negation{
		Tk:   tok(SUBTRACT, "-"),
		Expr: Value{tok(NUMBER, "1")},
	}

	expectOneStat(t, exp, f(given))
}

func E16_Negation(t *testing.T, f ParseFunc) {

	// -1 - -2

	given := []Token{
		tok(SUBTRACT, "-"),
		tok(NUMBER, "1"),
		tok(SUBTRACT, "-"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
	}

	left := Negation{
		Tk:   tok(SUBTRACT, "-"),
		Expr: Value{tok(NUMBER, "1")},
	}

	right := Negation{
		Tk:   tok(SUBTRACT, "-"),
		Expr: Value{tok(NUMBER, "2")},
	}

	op := Operation{
		Left:     left,
		Operator: tok(SUBTRACT, "-"),
		Right:    right,
	}

	expectOneStat(t, op, f(given))
}

func E17_Negation(t *testing.T, f ParseFunc) {

	// -(a == b)

	given := []Token{
		tok(SUBTRACT, "-"),
		tok(PAREN_OPEN, "-"),
		tok(IDENTIFIER, "a"),
		tok(EQUAL, "=="),
		tok(IDENTIFIER, "b"),
		tok(PAREN_CLOSE, "-"),
		tok(TERMINATOR, ""),
	}

	op := Operation{
		Left:     Identifier{tok(IDENTIFIER, "a")},
		Operator: tok(EQUAL, "=="),
		Right:    Identifier{tok(IDENTIFIER, "b")},
	}

	neg := Negation{
		Tk:   tok(SUBTRACT, "-"),
		Expr: op,
	}

	expectOneStat(t, neg, f(given))
}

func E18_Increment(t *testing.T, f ParseFunc) {

	// i++

	given := []Token{
		tok(IDENTIFIER, "i"),
		tok(INCREMENT, "++"),
		tok(TERMINATOR, ""),
	}

	inc := Increment{
		ID:        Identifier{tok(IDENTIFIER, "i")},
		Direction: tok(INCREMENT, "++"),
	}

	expectOneStat(t, inc, f(given))
}

func E19_Decrement(t *testing.T, f ParseFunc) {

	// i--

	given := []Token{
		tok(IDENTIFIER, "i"),
		tok(DECREMENT, "--"),
		tok(TERMINATOR, ""),
	}

	dec := Increment{
		ID:        Identifier{tok(IDENTIFIER, "i")},
		Direction: tok(DECREMENT, "--"),
	}

	expectOneStat(t, dec, f(given))
}
