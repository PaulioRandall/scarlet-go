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

	act := f(given)

	expectOneStat(t, exp, act)
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

	act := f(given)

	expectOneStat(t, exp, act)
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

	act := f(given)

	expectOneStat(t, exp, act)
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

	act := f(given)

	expectOneStat(t, exp, act)
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

	act := f(given)

	expectOneStat(t, exp, act)
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

	act := f(given)

	expectOneStat(t, exp, act)
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

	act := f(given)

	expectOneStat(t, exp, act)
}

// TODO: 1*2%3/4 -> 1*2%3/4
// TODO: 1%2/3*4 -> 1%2/3*4
// TODO: 1%2*3/4 -> 1%2*3/4

// TODO: 1*2+3 -> 1*2+3
// TODO: 1+2*3 -> 1+(2*3)
// TODO: 1/2+3 -> 1/2+3
// TODO: 1+2/3 -> 1+(2/3)

// TODO: 1+2*3%4/5-6 -> 1+((2*3)%(4/5))-6
