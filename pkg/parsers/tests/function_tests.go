package tests

import (
	"testing"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func F1_FuncDef(t *testing.T, f ParseFunc) {

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

	funcBody := st.Block{ // c := a
		Token{ID, "c", 0, 0},
		[]st.Statement{
			st.Assignment{
				false,
				[]Token{Token{ID, "c", 0, 0}},
				Token{ASSIGN, ":=", 0, 0},
				[]st.Expression{st.Identifier(Token{ID, "a", 0, 0})},
			},
		},
		Token{TERMINATOR, "\n", 0, 0},
	}

	funcExpr := []st.Expression{ // F(a, b, ^c) c := a
		st.FuncDef{
			Token{FUNC, "F", 0, 0},
			[]Token{ // a, b
				Token{ID, "a", 0, 0},
				Token{ID, "b", 0, 0},
			},
			[]Token{ // ^c
				Token{ID, "c", 0, 0},
			},
			funcBody, // c := a
		},
	}

	exp := st.Assignment{
		false,
		[]Token{Token{ID, "f", 0, 0}},
		Token{ASSIGN, ":=", 0, 0},
		funcExpr,
	}

	expectOneStat(t, exp, f(given))
}

func F2_FuncDef(t *testing.T, f ParseFunc) {

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

	funcBody := st.Block{ // { c := a }
		Token{BLOCK_OPEN, "{", 0, 0},
		[]st.Statement{
			st.Assignment{
				false,
				[]Token{Token{ID, "c", 0, 0}},
				Token{ASSIGN, ":=", 0, 0},
				[]st.Expression{st.Identifier(Token{ID, "a", 0, 0})},
			},
		},
		Token{BLOCK_CLOSE, "}", 0, 0},
	}

	funcExpr := []st.Expression{ // F(a, b, ^c) { c := a }
		st.FuncDef{
			Token{FUNC, "F", 0, 0},
			[]Token{ // a, b
				Token{ID, "a", 0, 0},
				Token{ID, "b", 0, 0},
			},
			[]Token{ // ^c
				Token{ID, "c", 0, 0},
			},
			funcBody, // { c := a }
		},
	}

	exp := st.Assignment{
		false,
		[]Token{Token{ID, "f", 0, 0}},
		Token{ASSIGN, ":=", 0, 0},
		funcExpr,
	}

	expectOneStat(t, exp, f(given))
}

func F3_FuncCall(t *testing.T, f ParseFunc) {

	// f()

	given := []Token{
		Token{ID, "f", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{TERMINATOR, "\n", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.FuncCall{
		st.Identifier(Token{ID, "f", 0, 0}),
		nil,
	}

	expectOneStat(t, exp, f(given))
}

func F4_FuncCall(t *testing.T, f ParseFunc) {

	// f(a, b)

	given := []Token{
		Token{ID, "f", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{ID, "a", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ID, "b", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{TERMINATOR, "\n", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.FuncCall{
		st.Identifier(Token{ID, "f", 0, 0}),
		[]st.Expression{
			st.Identifier(Token{ID, "a", 0, 0}),
			st.Identifier(Token{ID, "b", 0, 0}),
		},
	}

	expectOneStat(t, exp, f(given))
}

func F5_FuncCall(t *testing.T, f ParseFunc) {

	// f(1 + 2 - 3)

	given := []Token{
		Token{ID, "f", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{SUBTRACT, "-", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{TERMINATOR, "\n", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exp := st.FuncCall{
		ID: st.Identifier(Token{ID, "f", 0, 0}),
	}

	op := st.Operation{
		st.Value(Token{NUMBER, "1", 0, 0}),
		Token{ADD, "+", 0, 0},
		st.Value(Token{NUMBER, "2", 0, 0}),
	}

	op = st.Operation{
		op,
		Token{SUBTRACT, "-", 0, 0},
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	exp.Input = []st.Expression{op}

	expectOneStat(t, exp, f(given))
}

func F6_FuncCall(t *testing.T, f ParseFunc) {

	// f(abc())

	given := []Token{
		Token{ID, "f", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{ID, "abc", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{TERMINATOR, "\n", 0, 0},
		Token{EOF, "", 0, 0},
	}

	outer := st.FuncCall{
		ID: st.Identifier(Token{ID, "f", 0, 0}),
	}

	inner := st.FuncCall{
		st.Identifier(Token{ID, "abc", 0, 0}),
		nil,
	}

	outer.Input = []st.Expression{inner}

	expectOneStat(t, outer, f(given))
}

func F7_Panics(t *testing.T, f ParseFunc) {

	// f(

	given := []Token{
		Token{ID, "f", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	expectPanic(t, func() { f(given) })
}

func F8_Panics(t *testing.T, f ParseFunc) {

	// f)

	given := []Token{
		Token{ID, "f", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	expectPanic(t, func() { f(given) })
}
