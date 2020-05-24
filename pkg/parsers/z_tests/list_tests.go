package z_tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func L1_ListDef(t *testing.T, f ParseFunc) {

	// LIST {1,2,3}

	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(NUMBER, "1"),
		tok(DELIMITER, ","),
		tok(NUMBER, "2"),
		tok(DELIMITER, ","),
		tok(NUMBER, "3"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	exprs := []Expression{
		Value{tok(NUMBER, "1")},
		Value{tok(NUMBER, "2")},
		Value{tok(NUMBER, "3")},
	}

	list := List{
		Key:   tok(LIST, "LIST"),
		Open:  tok(BLOCK_OPEN, "{"),
		Exprs: exprs,
		Close: tok(BLOCK_CLOSE, "}"),
	}

	expectOneStat(t, list, f(given))
}

func L2_ListDef(t *testing.T, f ParseFunc) {

	// LIST {1+2,f()}

	given := []Token{
		tok(LIST, "LIST"),
		tok(BLOCK_OPEN, "{"),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(BLOCK_CLOSE, "}"),
		tok(TERMINATOR, ""),
	}

	add := Operation{
		Left:     Value{tok(NUMBER, "1")},
		Operator: tok(ADD, "+"),
		Right:    Value{tok(NUMBER, "2")},
	}

	funcCall := FuncCall{
		ID: Identifier{tok(IDENTIFIER, "f")},
	}

	list := List{
		Key:   tok(LIST, "LIST"),
		Open:  tok(BLOCK_OPEN, "{"),
		Exprs: []Expression{add, funcCall},
		Close: tok(BLOCK_CLOSE, "}"),
	}

	expectOneStat(t, list, f(given))
}

func L3_ListAccess(t *testing.T, f ParseFunc) {

	// x := y[1]

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(ASSIGN, ":="),
		tok(IDENTIFIER, "y"),
		tok(GUARD_OPEN, "["),
		tok(NUMBER, "1"),
		tok(GUARD_CLOSE, "]"),
		tok(TERMINATOR, ""),
	}

	target := AssignTarget{tok(IDENTIFIER, "x"), nil}

	listItem := ListAccess{
		ID:    Identifier{tok(IDENTIFIER, "y")},
		Index: Value{tok(NUMBER, "1")},
	}

	a := Assignment{
		false,
		[]AssignTarget{target},
		tok(ASSIGN, ":="),
		[]Expression{listItem},
	}

	expectOneStat(t, a, f(given))
}

func L4_ListAccess(t *testing.T, f ParseFunc) {

	// x := y[1+2]

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(ASSIGN, ":="),
		tok(IDENTIFIER, "y"),
		tok(GUARD_OPEN, "["),
		tok(NUMBER, "1"),
		tok(ADD, "+"),
		tok(NUMBER, "2"),
		tok(GUARD_CLOSE, "]"),
		tok(TERMINATOR, ""),
	}

	target := AssignTarget{tok(IDENTIFIER, "x"), nil}

	add := Operation{
		Left:     Value{tok(NUMBER, "1")},
		Operator: tok(ADD, "+"),
		Right:    Value{tok(NUMBER, "2")},
	}

	listItem := ListAccess{
		ID:    Identifier{tok(IDENTIFIER, "y")},
		Index: add,
	}

	a := Assignment{
		false,
		[]AssignTarget{target},
		tok(ASSIGN, ":="),
		[]Expression{listItem},
	}

	expectOneStat(t, a, f(given))
}

func L5_ListAccess(t *testing.T, f ParseFunc) {

	// x, y := z[<<], z[>>]

	given := []Token{
		tok(IDENTIFIER, "x"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "y"),
		tok(ASSIGN, ":="),
		tok(IDENTIFIER, "z"),
		tok(GUARD_OPEN, "["),
		tok(LIST_START, "<<"),
		tok(GUARD_CLOSE, "]"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "z"),
		tok(GUARD_OPEN, "["),
		tok(LIST_END, ">>"),
		tok(GUARD_CLOSE, "]"),
		tok(TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(IDENTIFIER, "x"), nil},
		AssignTarget{tok(IDENTIFIER, "y"), nil},
	}

	firstItem := ListAccess{
		ID:    Identifier{tok(IDENTIFIER, "z")},
		Index: ListItemRef{tok(LIST_START, "<<")},
	}

	lastItem := ListAccess{
		ID:    Identifier{tok(IDENTIFIER, "z")},
		Index: ListItemRef{tok(LIST_END, ">>")},
	}

	a := Assignment{
		false,
		targets,
		tok(ASSIGN, ":="),
		[]Expression{firstItem, lastItem},
	}

	expectOneStat(t, a, f(given))
}
