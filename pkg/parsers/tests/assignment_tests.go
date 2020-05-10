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

	act := f(given)

	expectOneStat(t, exp, act)
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

	act := f(given)

	expectOneStat(t, exp, act)
}

func F1_FuncInline(t *testing.T, f ParseFunc) {

	// f := F(a, b, ^c) c := a

	given := []Token{
		Token{ID, "f", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{FUNC, "F", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{ID, "a", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ID, "b", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{OUTPUT, "^", 0, 0},
		Token{ID, "c", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{ID, "c", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{ID, "a", 0, 0},
		Token{TERMINATOR, "\n", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Assignment{
		false,
		[]Token{Token{ID, "f", 0, 0}},
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{
			st.FuncDef{ // F(a, b, ^c) c := a
				Open: Token{FUNC, "F", 0, 0},
				Input: []Token{ // a, b
					Token{ID, "a", 0, 0},
					Token{ID, "b", 0, 0},
				},
				Output: []Token{ // ^c
					Token{ID, "c", 0, 0},
				},
				Body: st.Block{ // c := a
					Open: Token{ID, "c", 0, 0},
					Stats: []st.Statement{
						st.Assignment{
							false,
							[]Token{Token{ID, "c", 0, 0}},
							Token{ASSIGN, ":=", 0, 0},
							[]st.Expression{st.Identifier(Token{ID, "a", 0, 0})},
						},
					},
					Close: Token{TERMINATOR, "\n", 0, 0},
				},
			},
		},
	}

	act := f(given)

	expectOneStat(t, exp, act)
}

func F2_Func(t *testing.T, f ParseFunc) {

	// f := F(a, b, ^c) {
	//	c := a
	// }

	given := []Token{
		Token{ID, "f", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{FUNC, "F", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{ID, "a", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ID, "b", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{OUTPUT, "^", 0, 0},
		Token{ID, "c", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{BLOCK_OPEN, "{", 0, 0},
		Token{ID, "c", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{ID, "a", 0, 0},
		Token{TERMINATOR, "\n", 0, 0},
		Token{BLOCK_CLOSE, "}", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.Assignment{
		false,
		[]Token{Token{ID, "f", 0, 0}},
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{
			st.FuncDef{ // F(a, b, ^c) { c := a }
				Open: Token{FUNC, "F", 0, 0},
				Input: []Token{ // a, b
					Token{ID, "a", 0, 0},
					Token{ID, "b", 0, 0},
				},
				Output: []Token{ // ^c
					Token{ID, "c", 0, 0},
				},
				Body: st.Block{ // c := a
					Open: Token{BLOCK_OPEN, "{", 0, 0},
					Stats: []st.Statement{
						st.Assignment{
							false,
							[]Token{Token{ID, "c", 0, 0}},
							Token{ASSIGN, ":=", 0, 0},
							[]st.Expression{st.Identifier(Token{ID, "a", 0, 0})},
						},
					},
					Close: Token{BLOCK_CLOSE, "}", 0, 0},
				},
			},
		},
	}

	act := f(given)

	expectOneStat(t, exp, act)
}

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
