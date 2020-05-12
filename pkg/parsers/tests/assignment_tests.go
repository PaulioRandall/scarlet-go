package tests

import (
	"testing"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func A1_Assignment(t *testing.T, f ParseFunc) {

	// x := 1

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	targets := []st.AssignTarget{
		st.AssignTarget{Token{ID, "x", 0, 0}, nil},
	}

	exp := st.Assignment{
		false,
		targets,
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{st.Value(Token{NUMBER, "1", 0, 0})},
	}

	expectOneStat(t, exp, f(given))
}

func A2_MultiAssignment(t *testing.T, f ParseFunc) {

	// x, y := 1, 2

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ID, "y", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{TERMINATOR, "\n", 0, 0},
		Token{EOF, "", 0, 0},
	}

	targets := []st.AssignTarget{
		st.AssignTarget{Token{ID, "x", 0, 0}, nil},
		st.AssignTarget{Token{ID, "y", 0, 0}, nil},
	}

	exp := st.Assignment{
		false,
		targets,
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{
			st.Value(Token{NUMBER, "1", 0, 0}),
			st.Value(Token{NUMBER, "2", 0, 0}),
		},
	}

	expectOneStat(t, exp, f(given))
}

func A3_Assignment(t *testing.T, f ParseFunc) {

	// FIX x := 1

	given := []Token{
		Token{FIX, "FIX", 0, 0},
		Token{ID, "x", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	targets := []st.AssignTarget{
		st.AssignTarget{Token{ID, "x", 0, 0}, nil},
	}

	exp := st.Assignment{
		true,
		targets,
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{st.Value(Token{NUMBER, "1", 0, 0})},
	}

	expectOneStat(t, exp, f(given))
}

func A4_MultiAssignment(t *testing.T, f ParseFunc) {

	// x, y := f()

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ID, "y", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{ID, "f", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	targets := []st.AssignTarget{
		st.AssignTarget{Token{ID, "x", 0, 0}, nil},
		st.AssignTarget{Token{ID, "y", 0, 0}, nil},
	}

	exp := st.Assignment{
		false,
		targets,
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{
			st.FuncCall{
				st.Identifier(Token{ID, "f", 0, 0}),
				nil,
			},
		},
	}

	expectOneStat(t, exp, f(given))
}

func A5_Panics(t *testing.T, f ParseFunc) {

	// x, := 1

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	expectPanic(t, func() { f(given) })
}

func A6_Panics(t *testing.T, f ParseFunc) {

	// x, 1 := 1

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	expectPanic(t, func() { f(given) })
}

func A7_Panics(t *testing.T, f ParseFunc) {

	// x, F := 1

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{FUNC, "F", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	expectPanic(t, func() { f(given) })
}

func A8_Assignment(t *testing.T, f ParseFunc) {

	// a[0] := 1

	given := []Token{
		Token{ID, "a", 0, 0},
		Token{GUARD_OPEN, "[", 0, 0},
		Token{NUMBER, "0", 0, 0},
		Token{GUARD_CLOSE, "]", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	target := st.AssignTarget{
		Token{ID, "a", 0, 0},
		st.Value(Token{NUMBER, "0", 0, 0}),
	}

	exp := st.Assignment{
		false,
		[]st.AssignTarget{target},
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{st.Value(Token{NUMBER, "1", 0, 0})},
	}

	expectOneStat(t, exp, f(given))
}

func A9_Assignment(t *testing.T, f ParseFunc) {

	// a[b] := 1

	given := []Token{
		Token{ID, "a", 0, 0},
		Token{GUARD_OPEN, "[", 0, 0},
		Token{ID, "b", 0, 0},
		Token{GUARD_CLOSE, "]", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	target := st.AssignTarget{
		Token{ID, "a", 0, 0},
		st.Identifier(Token{ID, "b", 0, 0}),
	}

	exp := st.Assignment{
		false,
		[]st.AssignTarget{target},
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{st.Value(Token{NUMBER, "1", 0, 0})},
	}

	expectOneStat(t, exp, f(given))
}
