package tests

import (
	"testing"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func L1_ListDef(t *testing.T, f ParseFunc) {

	// LIST {1,2,3}

	given := []Token{
		Token{LIST, "LIST", 0, 0},
		Token{BLOCK_OPEN, "{", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{BLOCK_CLOSE, "}", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exprs := []st.Expression{
		st.Value(Token{NUMBER, "1", 0, 0}),
		st.Value(Token{NUMBER, "2", 0, 0}),
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	list := st.List{
		Key:   Token{LIST, "LIST", 0, 0},
		Open:  Token{BLOCK_OPEN, "{", 0, 0},
		Exprs: exprs,
		Close: Token{BLOCK_CLOSE, "}", 0, 0},
	}

	expectOneStat(t, list, f(given))
}

func L2_ListDef(t *testing.T, f ParseFunc) {

	// LIST {1+2,f()}

	given := []Token{
		Token{LIST, "LIST", 0, 0},
		Token{BLOCK_OPEN, "{", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ID, "f", 0, 0},
		Token{PAREN_OPEN, "(", 0, 0},
		Token{PAREN_CLOSE, ")", 0, 0},
		Token{BLOCK_CLOSE, "}", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	add := st.Operation{
		Left:     st.Value(Token{NUMBER, "1", 0, 0}),
		Operator: Token{ADD, "+", 0, 0},
		Right:    st.Value(Token{NUMBER, "2", 0, 0}),
	}

	funcCall := st.FuncCall{
		ID: st.Identifier(Token{ID, "f", 0, 0}),
	}

	list := st.List{
		Key:   Token{LIST, "LIST", 0, 0},
		Open:  Token{BLOCK_OPEN, "{", 0, 0},
		Exprs: []st.Expression{add, funcCall},
		Close: Token{BLOCK_CLOSE, "}", 0, 0},
	}

	expectOneStat(t, list, f(given))
}

func L3_ListAccess(t *testing.T, f ParseFunc) {

	// x := y[1]

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{ID, "y", 0, 0},
		Token{GUARD_OPEN, "[", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{GUARD_CLOSE, "]", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	target := st.AssignTarget{Token{ID, "x", 0, 0}, nil}

	listItem := st.ListAccess{
		ID:    st.Identifier(Token{ID, "y", 0, 0}),
		Index: st.Value(Token{NUMBER, "1", 0, 0}),
	}

	a := st.Assignment{
		false,
		[]st.AssignTarget{target},
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{listItem},
	}

	expectOneStat(t, a, f(given))
}

func L4_ListAccess(t *testing.T, f ParseFunc) {

	// x := y[1+2]

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{ID, "y", 0, 0},
		Token{GUARD_OPEN, "[", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{ADD, "+", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{GUARD_CLOSE, "]", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	target := st.AssignTarget{Token{ID, "x", 0, 0}, nil}

	add := st.Operation{
		Left:     st.Value(Token{NUMBER, "1", 0, 0}),
		Operator: Token{ADD, "+", 0, 0},
		Right:    st.Value(Token{NUMBER, "2", 0, 0}),
	}

	listItem := st.ListAccess{
		ID:    st.Identifier(Token{ID, "y", 0, 0}),
		Index: add,
	}

	a := st.Assignment{
		false,
		[]st.AssignTarget{target},
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{listItem},
	}

	expectOneStat(t, a, f(given))
}

func L5_ListAccess(t *testing.T, f ParseFunc) {

	// x, y := z[<<], z[>>]

	given := []Token{
		Token{ID, "x", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ID, "y", 0, 0},
		Token{ASSIGN, ":=", 0, 0},
		Token{ID, "z", 0, 0},
		Token{GUARD_OPEN, "[", 0, 0},
		Token{PREPEND, "<<", 0, 0},
		Token{GUARD_CLOSE, "]", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{ID, "z", 0, 0},
		Token{GUARD_OPEN, "[", 0, 0},
		Token{APPEND, ">>", 0, 0},
		Token{GUARD_CLOSE, "]", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	targets := []st.AssignTarget{
		st.AssignTarget{Token{ID, "x", 0, 0}, nil},
		st.AssignTarget{Token{ID, "y", 0, 0}, nil},
	}

	firstItem := st.ListAccess{
		ID:    st.Identifier(Token{ID, "z", 0, 0}),
		Index: st.ListItemRef(Token{PREPEND, "<<", 0, 0}),
	}

	lastItem := st.ListAccess{
		ID:    st.Identifier(Token{ID, "z", 0, 0}),
		Index: st.ListItemRef(Token{APPEND, ">>", 0, 0}),
	}

	a := st.Assignment{
		false,
		targets,
		Token{ASSIGN, ":=", 0, 0},
		[]st.Expression{firstItem, lastItem},
	}

	expectOneStat(t, a, f(given))
}
