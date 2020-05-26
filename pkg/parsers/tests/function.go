package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func F2_FuncDef(t *testing.T, f ParseFunc) {

	// f: F(a, b, ^c) {
	//	c: a
	// }

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":"),
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(DELIMITER, ","),
		tok(OUTPUT, "^"),
		tok(IDENTIFIER, "c"),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(IDENTIFIER, "c"),
		tok(ASSIGN, ":"),
		tok(IDENTIFIER, "a"),
		tok(TERMINATOR, "\n"),
		tok(BLOCK_CLOSE, "}"),
	}

	targets := []AssignTarget{
		AssignTarget{tok(IDENTIFIER, "f"), nil},
	}

	funcBody := Block{ // { c := a }
		tok(BLOCK_OPEN, "{"),
		[]Statement{
			Assignment{
				false,
				[]AssignTarget{AssignTarget{tok(IDENTIFIER, "c"), nil}},
				tok(ASSIGN, ":"),
				[]Expression{Identifier{tok(IDENTIFIER, "a")}},
			},
		},
		tok(BLOCK_CLOSE, "}"),
	}

	funcExpr := []Expression{ // F(a, b, ^c) { c := a }
		FuncDef{
			tok(FUNC, "F"),
			[]Token{ // a, b
				tok(IDENTIFIER, "a"),
				tok(IDENTIFIER, "b"),
			},
			[]Token{ // ^c
				tok(IDENTIFIER, "c"),
			},
			funcBody, // { c := a }
		},
	}

	exp := Assignment{
		false,
		targets,
		tok(ASSIGN, ":"),
		funcExpr,
	}

	expectOneStat(t, exp, f(given))
}

func F3_FuncCall(t *testing.T, f ParseFunc) {

	// f()

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, "\n"),
	}

	exp := FuncCall{
		Identifier{tok(IDENTIFIER, "f")},
		nil,
	}

	expectOneStat(t, exp, f(given))
}

func F4_FuncCall(t *testing.T, f ParseFunc) {

	// f(a, b)

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, "\n"),
	}

	exp := FuncCall{
		Identifier{tok(IDENTIFIER, "f")},
		[]Expression{
			Identifier{tok(IDENTIFIER, "a")},
			Identifier{tok(IDENTIFIER, "b")},
		},
	}

	expectOneStat(t, exp, f(given))
}

func F5_FuncCall(t *testing.T, f ParseFunc) {

	// f(1 + 2 - 3)

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(SUBTRACT, "-"),
		tok(NUMBER, "3"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, "\n"),
	}

	exp := FuncCall{
		ID: Identifier{tok(IDENTIFIER, "f")},
	}

	op := Operation{
		Value{tok(NUMBER, "1")},
		tok(ADD, "+"),
		Value{tok(NUMBER, "2")},
	}

	op = Operation{
		op,
		tok(SUBTRACT, "-"),
		Value{tok(NUMBER, "3")},
	}

	exp.Inputs = []Expression{op}

	expectOneStat(t, exp, f(given))
}

func F6_FuncCall(t *testing.T, f ParseFunc) {

	// f(abc())

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "abc"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, "\n"),
	}

	outer := FuncCall{
		ID: Identifier{tok(IDENTIFIER, "f")},
	}

	inner := FuncCall{
		Identifier{tok(IDENTIFIER, "abc")},
		nil,
	}

	outer.Inputs = []Expression{inner}

	expectOneStat(t, outer, f(given))
}

func F7_FuncCallPanics(t *testing.T, f ParseFunc) {

	// f(

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func F8_FuncCallPanics(t *testing.T, f ParseFunc) {

	// f)

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func F9_FuncCallPanics(t *testing.T, f ParseFunc) {

	// f(a,)

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func F10_FuncCallPanics(t *testing.T, f ParseFunc) {

	// f(a a)

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(IDENTIFIER, "a"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func F11_FuncDef(t *testing.T, f ParseFunc) {

	// f := F() {}

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":="),
		tok(FUNC, "F"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_OPEN, "{"),
		tok(BLOCK_CLOSE, "}"),
	}

	targets := []AssignTarget{
		AssignTarget{tok(IDENTIFIER, "f"), nil},
	}

	funcBody := Block{
		tok(BLOCK_OPEN, "{"),
		nil,
		tok(BLOCK_CLOSE, "}"),
	}

	funcDef := []Expression{
		FuncDef{
			tok(FUNC, "F"),
			nil,
			nil,
			funcBody, // c := a
		},
	}

	exp := Assignment{
		false,
		targets,
		tok(ASSIGN, ":="),
		funcDef,
	}

	expectOneStat(t, exp, f(given))
}
