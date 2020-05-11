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

	exp := st.Assignment{
		false,
		[]Token{Token{ID, "x", 0, 0}},
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

	exp := st.Assignment{
		false,
		[]Token{
			Token{ID, "x", 0, 0},
			Token{ID, "y", 0, 0},
		},
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{
			st.Value(Token{NUMBER, "1", 0, 0}),
			st.Value(Token{NUMBER, "2", 0, 0}),
		},
	}

	expectOneStat(t, exp, f(given))
}
