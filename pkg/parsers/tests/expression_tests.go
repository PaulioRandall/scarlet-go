package tests

import (
	"testing"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func E1_Add(t *testing.T, f ParseFunc) {

	// 1 + 2

	given := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "1", 0, 0}),
		Token{ADD, "+", 0, 0},
		st.Value(Token{NUMBER, "2", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E2_Subtract(t *testing.T, f ParseFunc) {

	// 2 - 1

	given := []Token{
		Token{NUMBER, "2", 0, 0},
		Token{SUBTRACT, "-", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "2", 0, 0}),
		Token{SUBTRACT, "-", 0, 0},
		st.Value(Token{NUMBER, "1", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E3_Multiply(t *testing.T, f ParseFunc) {

	// 6 * 7

	given := []Token{
		Token{NUMBER, "6", 0, 0},
		Token{MULTIPLY, "*", 0, 0},
		Token{NUMBER, "7", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "6", 0, 0}),
		Token{MULTIPLY, "*", 0, 0},
		st.Value(Token{NUMBER, "7", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E4_Divide(t *testing.T, f ParseFunc) {

	// 12 / 3

	given := []Token{
		Token{NUMBER, "12", 0, 0},
		Token{DIVIDE, "/", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "12", 0, 0}),
		Token{DIVIDE, "/", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E5_AdditiveOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 - 3

	given := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{SUBTRACT, "-", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "1", 0, 0}),
		Token{ADD, "+", 0, 0},
		st.Value(Token{NUMBER, "2", 0, 0}),
	}

	exp = st.Operation{
		exp,
		Token{SUBTRACT, "-", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E6_AdditiveOrdering(t *testing.T, f ParseFunc) {

	// 1 - 2 + 3

	given := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{SUBTRACT, "-", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "1", 0, 0}),
		Token{SUBTRACT, "-", 0, 0},
		st.Value(Token{NUMBER, "2", 0, 0}),
	}

	exp = st.Operation{
		exp,
		Token{ADD, "+", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E7_MultiplicativeOrdering(t *testing.T, f ParseFunc) {

	// 1 * 2 % 3 4

	given := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{MULTIPLY, "*", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{REMAINDER, "%", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{DIVIDE, "/", 0, 0},
		Token{NUMBER, "4", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "1", 0, 0}),
		Token{MULTIPLY, "*", 0, 0},
		st.Value(Token{NUMBER, "2", 0, 0}),
	}

	exp = st.Operation{
		exp,
		Token{REMAINDER, "%", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	exp = st.Operation{
		exp,
		Token{DIVIDE, "/", 0, 0},
		st.Value(Token{NUMBER, "4", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E8_MultiplicativeOrdering(t *testing.T, f ParseFunc) {

	// 1 % 2 / 3 * 4

	given := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{REMAINDER, "%", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{DIVIDE, "/", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{MULTIPLY, "*", 0, 0},
		Token{NUMBER, "4", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "1", 0, 0}),
		Token{REMAINDER, "%", 0, 0},
		st.Value(Token{NUMBER, "2", 0, 0}),
	}

	exp = st.Operation{
		exp,
		Token{DIVIDE, "/", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	exp = st.Operation{
		exp,
		Token{MULTIPLY, "*", 0, 0},
		st.Value(Token{NUMBER, "4", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E9_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 * 2 + 3

	given := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{MULTIPLY, "*", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Operation{
		st.Value(Token{NUMBER, "1", 0, 0}),
		Token{MULTIPLY, "*", 0, 0},
		st.Value(Token{NUMBER, "2", 0, 0}),
	}

	exp = st.Operation{
		exp,
		Token{ADD, "+", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E10_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 * 3

	given := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{MULTIPLY, "*", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	// 1 + (2 * 3)

	exp := st.Operation{
		Left:     st.Value(Token{NUMBER, "1", 0, 0}),
		Operator: Token{ADD, "+", 0, 0},
	}

	exp.Right = st.Operation{
		st.Value(Token{NUMBER, "2", 0, 0}),
		Token{MULTIPLY, "*", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	expectOneStat(t, exp, f(given))
}

func E11_OperationOrdering(t *testing.T, f ParseFunc) {

	// 1 + 2 * 3 - 4 % 5 / 6

	given := []Token{
		Token{NUMBER, "1", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{MULTIPLY, "*", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{SUBTRACT, "-", 0, 0},
		Token{NUMBER, "4", 0, 0},
		Token{REMAINDER, "%", 0, 0},
		Token{NUMBER, "5", 0, 0},
		Token{DIVIDE, "/", 0, 0},
		Token{NUMBER, "6", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	// 1 + (2 * 3) - ((4 % 5) / 6)

	add := st.Operation{
		Left:     st.Value(Token{NUMBER, "1", 0, 0}),
		Operator: Token{ADD, "+", 0, 0},
	}

	mul := st.Operation{
		st.Value(Token{NUMBER, "2", 0, 0}),
		Token{MULTIPLY, "*", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	add.Right = mul

	sub := st.Operation{
		Left:     add,
		Operator: Token{SUBTRACT, "-", 0, 0},
	}

	rem := st.Operation{
		st.Value(Token{NUMBER, "4", 0, 0}),
		Token{REMAINDER, "%", 0, 0},
		st.Value(Token{NUMBER, "5", 0, 0}),
	}

	div := st.Operation{
		rem,
		Token{DIVIDE, "/", 0, 0},
		st.Value(Token{NUMBER, "6", 0, 0}),
	}

	sub.Right = div

	expectOneStat(t, sub, f(given))
}
