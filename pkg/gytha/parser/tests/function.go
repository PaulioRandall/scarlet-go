package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func F2_FuncDef(t *testing.T, f ParseFunc) {

	// f: F(a, b, ^c) {
	//	c: a
	// }

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_DELIMITER, ","),
		tok(TK_OUTPUT, "^"),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_IDENTIFIER, "c"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_TERMINATOR, "\n"),
		tok(TK_BLOCK_CLOSE, "}"),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "f"), nil},
	}

	funcBody := Block{ // { c := a }
		tok(TK_BLOCK_OPEN, "{"),
		[]Statement{
			Assignment{
				false,
				[]AssignTarget{AssignTarget{tok(TK_IDENTIFIER, "c"), nil}},
				tok(TK_ASSIGNMENT, ":"),
				[]Expression{Identifier{tok(TK_IDENTIFIER, "a")}},
			},
		},
		tok(TK_BLOCK_CLOSE, "}"),
	}

	funcExpr := []Expression{ // F(a, b, ^c) { c := a }
		FuncDef{
			tok(TK_FUNCTION, "F"),
			[]Token{ // a, b
				tok(TK_IDENTIFIER, "a"),
				tok(TK_IDENTIFIER, "b"),
			},
			[]OutputParam{ // ^c
				OutputParam{
					Identifier{tok(TK_IDENTIFIER, "c")},
					nil,
				},
			},
			funcBody, // { c := a }
		},
	}

	exp := Assignment{
		false,
		targets,
		tok(TK_ASSIGNMENT, ":"),
		funcExpr,
	}

	expectOneStat(t, exp, f(given))
}

func F3_FuncCall(t *testing.T, f ParseFunc) {

	// f()

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, "\n"),
	}

	exp := FuncCall{
		Identifier{tok(TK_IDENTIFIER, "f")},
		nil,
	}

	expectOneStat(t, exp, f(given))
}

func F4_FuncCall(t *testing.T, f ParseFunc) {

	// f(a, b)

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, "\n"),
	}

	exp := FuncCall{
		Identifier{tok(TK_IDENTIFIER, "f")},
		[]Expression{
			Identifier{tok(TK_IDENTIFIER, "a")},
			Identifier{tok(TK_IDENTIFIER, "b")},
		},
	}

	expectOneStat(t, exp, f(given))
}

func F5_FuncCall(t *testing.T, f ParseFunc) {

	// f(1 + 2 - 3)

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_MINUS, "-"),
		tok(TK_NUMBER, "3"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, "\n"),
	}

	exp := FuncCall{
		ID: Identifier{tok(TK_IDENTIFIER, "f")},
	}

	op := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_PLUS, "+"),
		Value{tok(TK_NUMBER, "2")},
	}

	op = Operation{
		op,
		tok(TK_MINUS, "-"),
		Value{tok(TK_NUMBER, "3")},
	}

	exp.Inputs = []Expression{op}

	expectOneStat(t, exp, f(given))
}

func F6_FuncCall(t *testing.T, f ParseFunc) {

	// f(abc())

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "abc"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, "\n"),
	}

	outer := FuncCall{
		ID: Identifier{tok(TK_IDENTIFIER, "f")},
	}

	inner := FuncCall{
		Identifier{tok(TK_IDENTIFIER, "abc")},
		nil,
	}

	outer.Inputs = []Expression{inner}

	expectOneStat(t, outer, f(given))
}

func F7_FuncCallPanics(t *testing.T, f ParseFunc) {

	// f(

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func F8_FuncCallPanics(t *testing.T, f ParseFunc) {

	// f)

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func F9_FuncCallPanics(t *testing.T, f ParseFunc) {

	// f(a,)

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func F10_FuncCallPanics(t *testing.T, f ParseFunc) {

	// f(a a)

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func F11_FuncDef(t *testing.T, f ParseFunc) {

	// f := F() {}

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "f"), nil},
	}

	funcBody := Block{
		tok(TK_BLOCK_OPEN, "{"),
		nil,
		tok(TK_BLOCK_CLOSE, "}"),
	}

	funcDef := []Expression{
		FuncDef{
			tok(TK_FUNCTION, "F"),
			nil,
			nil,
			funcBody, // c := a
		},
	}

	exp := Assignment{
		false,
		targets,
		tok(TK_ASSIGNMENT, ":="),
		funcDef,
	}

	expectOneStat(t, exp, f(given))
}

func F12_AssignOutput(t *testing.T, f ParseFunc) {

	// f: F(^a: 5, ^b: 1 + 2) {}

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_FUNCTION, "F"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_OUTPUT, "^"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_NUMBER, "5"),
		tok(TK_DELIMITER, ","),
		tok(TK_OUTPUT, "^"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "f"), nil},
	}

	funcDef := FuncDef{
		Key: tok(TK_FUNCTION, "F"),
	}

	funcDef.Outputs = []OutputParam{
		OutputParam{
			Identifier{tok(TK_IDENTIFIER, "a")},
			Value{tok(TK_NUMBER, "5")},
		},
		OutputParam{
			Identifier{tok(TK_IDENTIFIER, "b")},
			Operation{
				Left:     Value{tok(TK_NUMBER, "1")},
				Operator: tok(TK_PLUS, "+"),
				Right:    Value{tok(TK_NUMBER, "2")},
			},
		},
	}

	funcDef.Body = Block{
		tok(TK_BLOCK_OPEN, "{"),
		nil,
		tok(TK_BLOCK_CLOSE, "}"),
	}

	a := Assignment{
		Fixed:   false,
		Targets: targets,
		Assign:  tok(TK_ASSIGNMENT, ":"),
		Exprs:   []Expression{funcDef},
	}

	expectOneStat(t, a, f(given))
}
