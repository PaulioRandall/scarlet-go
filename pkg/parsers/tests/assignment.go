package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func A1_Assignment(t *testing.T, f ParseFunc) {

	// x := 1

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "x"), nil},
	}

	exp := Assignment{
		false,
		targets,
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{Value{tok(TK_NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A2_MultiAssignment(t *testing.T, f ParseFunc) {

	// x, y := 1, 2

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "y"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, "\n"),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "x"), nil},
		AssignTarget{tok(TK_IDENTIFIER, "y"), nil},
	}

	exp := Assignment{
		false,
		targets,
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{
			Value{tok(TK_NUMBER, "1")},
			Value{tok(TK_NUMBER, "2")},
		},
	}

	expectOneStat(t, exp, f(given))
}

func A3_Assignment(t *testing.T, f ParseFunc) {

	// DEF x := 1

	given := []Token{
		tok(TK_DEFINITION, "DEF"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "x"), nil},
	}

	exp := Assignment{
		true,
		targets,
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{Value{tok(TK_NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A4_MultiAssignment(t *testing.T, f ParseFunc) {

	// x, y := f()

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "y"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "x"), nil},
		AssignTarget{tok(TK_IDENTIFIER, "y"), nil},
	}

	exp := Assignment{
		false,
		targets,
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{
			FuncCall{
				Identifier{tok(TK_IDENTIFIER, "f")},
				nil,
			},
		},
	}

	expectOneStat(t, exp, f(given))
}

func A5_Panics(t *testing.T, f ParseFunc) {

	// x, := 1

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_DELIMITER, ","),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func A6_Panics(t *testing.T, f ParseFunc) {

	// x, 1 := 1

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "1"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func A7_Panics(t *testing.T, f ParseFunc) {

	// x, F := 1

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_DELIMITER, ","),
		tok(TK_FUNCTION, "F"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	expectPanic(t, func() { f(given) })
}

func A8_ListItem(t *testing.T, f ParseFunc) {

	// a[0] := 1

	given := []Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "0"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	target := AssignTarget{
		tok(TK_IDENTIFIER, "a"),
		Value{tok(TK_NUMBER, "0")},
	}

	exp := Assignment{
		false,
		[]AssignTarget{target},
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{Value{tok(TK_NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A9_ListItem(t *testing.T, f ParseFunc) {

	// a[b] := 1

	given := []Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	target := AssignTarget{
		tok(TK_IDENTIFIER, "a"),
		Identifier{tok(TK_IDENTIFIER, "b")},
	}

	exp := Assignment{
		false,
		[]AssignTarget{target},
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{Value{tok(TK_NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A10_ListItem(t *testing.T, f ParseFunc) {

	// a[1+2] := 1

	given := []Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_TERMINATOR, ""),
	}

	index := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_PLUS, "+"),
		Value{tok(TK_NUMBER, "2")},
	}

	target := AssignTarget{
		tok(TK_IDENTIFIER, "a"),
		index,
	}

	exp := Assignment{
		false,
		[]AssignTarget{target},
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{Value{tok(TK_NUMBER, "1")}},
	}

	expectOneStat(t, exp, f(given))
}

func A11_ListItems(t *testing.T, f ParseFunc) {

	// a[<<], a[>>] := 1, 2

	given := []Token{
		tok(TK_IDENTIFIER, "a"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_LIST_START, "<<"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_LIST_END, ">>"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_TERMINATOR, ""),
	}

	firstTarget := AssignTarget{
		ID:    tok(TK_IDENTIFIER, "a"),
		Index: ListItemRef{tok(TK_LIST_START, "<<")},
	}

	secondTarget := AssignTarget{
		ID:    tok(TK_IDENTIFIER, "a"),
		Index: ListItemRef{tok(TK_LIST_END, ">>")},
	}

	values := []Expression{
		Value{tok(TK_NUMBER, "1")},
		Value{tok(TK_NUMBER, "2")},
	}

	a := Assignment{
		false,
		[]AssignTarget{firstTarget, secondTarget},
		tok(TK_ASSIGNMENT, ":="),
		values,
	}

	expectOneStat(t, a, f(given))
}
