package tests

import (
	"testing"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

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
