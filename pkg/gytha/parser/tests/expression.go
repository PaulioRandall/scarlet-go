package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func E1_Add(t *testing.T, f ParseFunc) {

	// 1 + 2

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_PLUS, "+"),
		Value{tok(TK_NUMBER, "2")},
	}

	expectOneStat(t, exp, f(given))
}

func E2_Subtract(t *testing.T, f ParseFunc) {

	// 2 - 1

	given := []Token{
		tok(TK_NUMBER, "2"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "2")},
		tok(TK_MINUS, "-"),
		Value{tok(TK_NUMBER, "1")},
	}

	expectOneStat(t, exp, f(given))
}

func E3_Multiply(t *testing.T, f ParseFunc) {

	// 6 * 7

	given := []Token{
		tok(TK_NUMBER, "6"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "7"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "6")},
		tok(TK_MULTIPLY, "*"),
		Value{tok(TK_NUMBER, "7")},
	}

	expectOneStat(t, exp, f(given))
}

func E4_Divide(t *testing.T, f ParseFunc) {

	// 12 / 3

	given := []Token{
		tok(TK_NUMBER, "12"),
		tok(TK_DIVIDE, "/"),
		tok(TK_NUMBER, "3"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "12")},
		tok(TK_DIVIDE, "/"),
		Value{tok(TK_NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E5_AdditiveOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 - 3

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "3"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_PLUS, "+"),
		Value{tok(TK_NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(TK_MINUS, "-"),
		Value{tok(TK_NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E6_AdditiveOrdering(t *testing.T, f ParseFunc) {

	// 1 - 2 + 3

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "2"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "3"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_MINUS, "-"),
		Value{tok(TK_NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(TK_PLUS, "+"),
		Value{tok(TK_NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E7_MultiplicativeOrdering(t *testing.T, f ParseFunc) {

	// 1 * 2 % 3 4

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "2"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "3"),
		tok(TK_DIVIDE, "/"),
		tok(TK_NUMBER, "4"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_MULTIPLY, "*"),
		Value{tok(TK_NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(TK_REMAINDER, "%"),
		Value{tok(TK_NUMBER, "3")},
	}

	exp = Operation{
		exp,
		tok(TK_DIVIDE, "/"),
		Value{tok(TK_NUMBER, "4")},
	}

	expectOneStat(t, exp, f(given))
}

func E8_MultiplicativeOrdering(t *testing.T, f ParseFunc) {

	// 1 % 2 / 3 * 4

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "2"),
		tok(TK_DIVIDE, "/"),
		tok(TK_NUMBER, "3"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "4"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_REMAINDER, "%"),
		Value{tok(TK_NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(TK_DIVIDE, "/"),
		Value{tok(TK_NUMBER, "3")},
	}

	exp = Operation{
		exp,
		tok(TK_MULTIPLY, "*"),
		Value{tok(TK_NUMBER, "4")},
	}

	expectOneStat(t, exp, f(given))
}

func E9_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 * 2 + 3

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "2"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "3"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_MULTIPLY, "*"),
		Value{tok(TK_NUMBER, "2")},
	}

	exp = Operation{
		exp,
		tok(TK_PLUS, "+"),
		Value{tok(TK_NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E10_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 * 3

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "3"),
		tok(TK_TERMINATOR, ""),
	}

	// 1 + (2 * 3)

	exp := Operation{
		Left:     Value{tok(TK_NUMBER, "1")},
		Operator: tok(TK_PLUS, "+"),
	}

	exp.Right = Operation{
		Value{tok(TK_NUMBER, "2")},
		tok(TK_MULTIPLY, "*"),
		Value{tok(TK_NUMBER, "3")},
	}

	expectOneStat(t, exp, f(given))
}

func E11_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 * 3 - 4 % 5 / 6

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_MULTIPLY, "*"),
		tok(TK_NUMBER, "3"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "4"),
		tok(TK_REMAINDER, "%"),
		tok(TK_NUMBER, "5"),
		tok(TK_DIVIDE, "/"),
		tok(TK_NUMBER, "6"),
		tok(TK_TERMINATOR, ""),
	}

	// 1 + (2 * 3) - ((4 % 5) / 6)

	add := Operation{
		Left:     Value{tok(TK_NUMBER, "1")},
		Operator: tok(TK_PLUS, "+"),
	}

	mul := Operation{
		Value{tok(TK_NUMBER, "2")},
		tok(TK_MULTIPLY, "*"),
		Value{tok(TK_NUMBER, "3")},
	}

	add.Right = mul

	sub := Operation{
		Left:     add,
		Operator: tok(TK_MINUS, "-"),
	}

	rem := Operation{
		Value{tok(TK_NUMBER, "4")},
		tok(TK_REMAINDER, "%"),
		Value{tok(TK_NUMBER, "5")},
	}

	div := Operation{
		rem,
		tok(TK_DIVIDE, "/"),
		Value{tok(TK_NUMBER, "6")},
	}

	sub.Right = div

	expectOneStat(t, sub, f(given))
}

func E12_FuncCall(t *testing.T, f ParseFunc) {

	// 1 + f(a,b)

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Operation{
		Left:     Value{tok(TK_NUMBER, "1")},
		Operator: tok(TK_PLUS, "+"),
	}

	exp.Right = FuncCall{
		Identifier{tok(TK_IDENTIFIER, "f")},
		[]Expression{
			Identifier{tok(TK_IDENTIFIER, "a")},
			Identifier{tok(TK_IDENTIFIER, "b")},
		},
	}

	expectOneStat(t, exp, f(given))
}

func E13_Panics(t *testing.T, f ParseFunc) {

	// 1 + +

	given := []Token{
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_PLUS, "+"),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func E14_Panics(t *testing.T, f ParseFunc) {

	// + 1 + 1

	given := []Token{
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func E15_Negation(t *testing.T, f ParseFunc) {

	// -1

	given := []Token{
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	exp := Negation{
		Tk:   tok(TK_MINUS, "-"),
		Expr: Value{tok(TK_NUMBER, "1")},
	}

	expectOneStat(t, exp, f(given))
}

func E16_Negation(t *testing.T, f ParseFunc) {

	// -1 - -2

	given := []Token{
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "1"),
		tok(TK_MINUS, "-"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
	}

	left := Negation{
		Tk:   tok(TK_MINUS, "-"),
		Expr: Value{tok(TK_NUMBER, "1")},
	}

	right := Negation{
		Tk:   tok(TK_MINUS, "-"),
		Expr: Value{tok(TK_NUMBER, "2")},
	}

	op := Operation{
		Left:     left,
		Operator: tok(TK_MINUS, "-"),
		Right:    right,
	}

	expectOneStat(t, op, f(given))
}

func E17_Negation(t *testing.T, f ParseFunc) {

	// -(a == b)

	given := []Token{
		tok(TK_MINUS, "-"),
		tok(TK_PAREN_OPEN, "-"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_EQUAL, "=="),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_PAREN_CLOSE, "-"),
		tok(TK_TERMINATOR, ""),
	}

	op := Operation{
		Left:     Identifier{tok(TK_IDENTIFIER, "a")},
		Operator: tok(TK_EQUAL, "=="),
		Right:    Identifier{tok(TK_IDENTIFIER, "b")},
	}

	neg := Negation{
		Tk:   tok(TK_MINUS, "-"),
		Expr: op,
	}

	expectOneStat(t, neg, f(given))
}
