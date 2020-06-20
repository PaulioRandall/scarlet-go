package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func L1_ListDef(t *testing.T, f ParseFunc) {

	// LIST {1,2,3}

	given := []Token{
		tok(TK_LIST, "LIST"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_NUMBER, "1"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "2"),
		tok(TK_DELIMITER, ","),
		tok(TK_NUMBER, "3"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	exprs := []Expression{
		Value{tok(TK_NUMBER, "1")},
		Value{tok(TK_NUMBER, "2")},
		Value{tok(TK_NUMBER, "3")},
	}

	list := List{
		Key:   tok(TK_LIST, "LIST"),
		Open:  tok(TK_BLOCK_OPEN, "{"),
		Exprs: exprs,
		Close: tok(TK_BLOCK_CLOSE, "}"),
	}

	expectOneStat(t, list, f(given))
}

func L2_ListDef(t *testing.T, f ParseFunc) {

	// LIST {1+2,f()}

	given := []Token{
		tok(TK_LIST, "LIST"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_BLOCK_CLOSE, "}"),
		tok(TK_TERMINATOR, ""),
	}

	add := Operation{
		Left:     Value{tok(TK_NUMBER, "1")},
		Operator: tok(TK_PLUS, "+"),
		Right:    Value{tok(TK_NUMBER, "2")},
	}

	funcCall := FuncCall{
		ID: Identifier{tok(TK_IDENTIFIER, "f")},
	}

	list := List{
		Key:   tok(TK_LIST, "LIST"),
		Open:  tok(TK_BLOCK_OPEN, "{"),
		Exprs: []Expression{add, funcCall},
		Close: tok(TK_BLOCK_CLOSE, "}"),
	}

	expectOneStat(t, list, f(given))
}

func L3_ListAccess(t *testing.T, f ParseFunc) {

	// x := y[1]

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "y"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_TERMINATOR, ""),
	}

	target := AssignTarget{tok(TK_IDENTIFIER, "x"), nil}

	listItem := ListAccess{
		ID:    Identifier{tok(TK_IDENTIFIER, "y")},
		Index: Value{tok(TK_NUMBER, "1")},
	}

	a := Assignment{
		false,
		[]AssignTarget{target},
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{listItem},
	}

	expectOneStat(t, a, f(given))
}

func L4_ListAccess(t *testing.T, f ParseFunc) {

	// x := y[1+2]

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "y"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_NUMBER, "1"),
		tok(TK_PLUS, "+"),
		tok(TK_NUMBER, "2"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_TERMINATOR, ""),
	}

	target := AssignTarget{tok(TK_IDENTIFIER, "x"), nil}

	add := Operation{
		Left:     Value{tok(TK_NUMBER, "1")},
		Operator: tok(TK_PLUS, "+"),
		Right:    Value{tok(TK_NUMBER, "2")},
	}

	listItem := ListAccess{
		ID:    Identifier{tok(TK_IDENTIFIER, "y")},
		Index: add,
	}

	a := Assignment{
		false,
		[]AssignTarget{target},
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{listItem},
	}

	expectOneStat(t, a, f(given))
}

func L5_ListAccess(t *testing.T, f ParseFunc) {

	// x, y := z[<<], z[>>]

	given := []Token{
		tok(TK_IDENTIFIER, "x"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "y"),
		tok(TK_ASSIGNMENT, ":="),
		tok(TK_IDENTIFIER, "z"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_LIST_START, "<<"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "z"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_LIST_END, ">>"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "x"), nil},
		AssignTarget{tok(TK_IDENTIFIER, "y"), nil},
	}

	firstItem := ListAccess{
		ID:    Identifier{tok(TK_IDENTIFIER, "z")},
		Index: ListItemRef{tok(TK_LIST_START, "<<")},
	}

	lastItem := ListAccess{
		ID:    Identifier{tok(TK_IDENTIFIER, "z")},
		Index: ListItemRef{tok(TK_LIST_END, ">>")},
	}

	a := Assignment{
		false,
		targets,
		tok(TK_ASSIGNMENT, ":="),
		[]Expression{firstItem, lastItem},
	}

	expectOneStat(t, a, f(given))
}
