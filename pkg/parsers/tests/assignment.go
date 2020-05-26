package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func A1_Assignment(t *testing.T, f ParseFunc) {

	// x := 1

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(IDENTIFIER, "x"), nil},
	}

	exp := Assignment{
		false,
		targets,
		tok(ASSIGN, ":="),
		[]Expression{Value{tok(NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A2_MultiAssignment(t *testing.T, f ParseFunc) {

	// x, y := 1, 2

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "y"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(NUMBER, "2"),
		tok(TERMINATOR, "\n"),
	}

	targets := []AssignTarget{
		AssignTarget{tok(IDENTIFIER, "x"), nil},
		AssignTarget{tok(IDENTIFIER, "y"), nil},
	}

	exp := Assignment{
		false,
		targets,
		tok(ASSIGN, ":="),
		[]Expression{
			Value{tok(NUMBER, "1")},
			Value{tok(NUMBER, "2")},
		},
	}

	expectOneStat(t, exp, f(given))
}

func A3_Assignment(t *testing.T, f ParseFunc) {

	// DEF x := 1

	given := []Token{
		tok(DEF, "DEF"),
		tok(IDENTIFIER, "x"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(IDENTIFIER, "x"), nil},
	}

	exp := Assignment{
		true,
		targets,
		tok(ASSIGN, ":="),
		[]Expression{Value{tok(NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A4_MultiAssignment(t *testing.T, f ParseFunc) {

	// x, y := f()

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "y"),
		tok(ASSIGN, ":="),
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(IDENTIFIER, "x"), nil},
		AssignTarget{tok(IDENTIFIER, "y"), nil},
	}

	exp := Assignment{
		false,
		targets,
		tok(ASSIGN, ":="),
		[]Expression{
			FuncCall{
				Identifier{tok(IDENTIFIER, "f")},
				nil,
			},
		},
	}

	expectOneStat(t, exp, f(given))
}

func A5_Panics(t *testing.T, f ParseFunc) {

	// x, := 1

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(DELIMITER, ","),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func A6_Panics(t *testing.T, f ParseFunc) {

	// x, 1 := 1

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(DELIMITER, ","),
		tok(NUMBER, "1"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func A7_Panics(t *testing.T, f ParseFunc) {

	// x, F := 1

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(DELIMITER, ","),
		tok(FUNC, "F"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func A8_ListItem(t *testing.T, f ParseFunc) {

	// a[0] := 1

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(GUARD_OPEN, "["),
		tok(NUMBER, "0"),
		tok(GUARD_CLOSE, "]"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	target := AssignTarget{
		tok(IDENTIFIER, "a"),
		Value{tok(NUMBER, "0")},
	}

	exp := Assignment{
		false,
		[]AssignTarget{target},
		tok(ASSIGN, ":="),
		[]Expression{Value{tok(NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A9_ListItem(t *testing.T, f ParseFunc) {

	// a[b] := 1

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(GUARD_OPEN, "["),
		tok(IDENTIFIER, "b"),
		tok(GUARD_CLOSE, "]"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	target := AssignTarget{
		tok(IDENTIFIER, "a"),
		Identifier{tok(IDENTIFIER, "b")},
	}

	exp := Assignment{
		false,
		[]AssignTarget{target},
		tok(ASSIGN, ":="),
		[]Expression{Value{tok(NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A10_ListItem(t *testing.T, f ParseFunc) {

	// a[1+2] := 1

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(GUARD_OPEN, "["),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(GUARD_CLOSE, "]"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	index := Operation{
		Value{tok(NUMBER, "1")},
		tok(ADD, "+"),
		Value{tok(NUMBER, "2")},
	}

	target := AssignTarget{
		tok(IDENTIFIER, "a"),
		index,
	}

	exp := Assignment{
		false,
		[]AssignTarget{target},
		tok(ASSIGN, ":="),
		[]Expression{Value{tok(NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A11_ListItems(t *testing.T, f ParseFunc) {

	// a[<<], a[>>] := 1, 2

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(GUARD_OPEN, "["),
		tok(LIST_START, "<<"),
		tok(GUARD_CLOSE, "]"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "a"),
		tok(GUARD_OPEN, "["),
		tok(LIST_END, ">>"),
		tok(GUARD_CLOSE, "]"),
		tok(ASSIGN, ":="),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(NUMBER, "2"),
		tok(TERMINATOR, ""),
	}

	firstTarget := AssignTarget{
		ID:    tok(IDENTIFIER, "a"),
		Index: ListItemRef{tok(LIST_START, "<<")},
	}

	secondTarget := AssignTarget{
		ID:    tok(IDENTIFIER, "a"),
		Index: ListItemRef{tok(LIST_END, ">>")},
	}

	values := []Expression{
		Value{tok(NUMBER, "1")},
		Value{tok(NUMBER, "2")},
	}

	a := Assignment{
		false,
		[]AssignTarget{firstTarget, secondTarget},
		tok(ASSIGN, ":="),
		values,
	}

	expectOneStat(t, a, f(given))
}
