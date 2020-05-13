package tests

import (
	"testing"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func LP1_Assignment(t *testing.T, f ParseFunc) {

	// LOOP i [i < 5] {
	//	 x := i
	// }

	given := []Token{
		Token{LOOP, "LOOP", 0, 0},
		Token{ID, "i", 0, 0},
		Token{GUARD_OPEN, "[", 0, 0},
		Token{ID, "i", 0, 0},
		Token{LESS_THAN, "<", 0, 0},
		Token{NUMBER, "5", 0, 0},
		Token{GUARD_CLOSE, "]", 0, 0},
		Token{BLOCK_OPEN, "{", 0, 0},
		Token{ID, "x", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{ID, "i", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{BLOCK_CLOSE, "}", 0, 0},
		Token{EOF, "", 0, 0},
	}

	loop := st.Loop{
		Open:     Token{LOOP, "LOOP", 0, 0},
		IndexVar: Token{ID, "i", 0, 0},
	}

	condition := st.Operation{
		st.Identifier(Token{ID, "i", 0, 0}),
		Token{LESS_THAN, "<", 0, 0},
		st.Value(Token{NUMBER, "5", 0, 0}),
	}

	guard := st.Guard{
		Open:      Token{GUARD_OPEN, "[", 0, 0},
		Condition: condition,
		Close:     Token{GUARD_CLOSE, "]", 0, 0},
	}

	stat := st.Assignment{
		false,
		[]st.AssignTarget{st.AssignTarget{Token{ID, "x", 0, 0}, nil}},
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{st.Identifier(Token{ID, "i", 0, 0})},
	}

	guard.Block = st.Block{
		Token{BLOCK_OPEN, "{", 0, 0},
		[]st.Statement{stat},
		Token{BLOCK_CLOSE, "}", 0, 0},
	}

	loop.Guard = guard

	expectOneStat(t, loop, f(given))
}
