package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func LP1_Conditional(t *testing.T, f ParseFunc) {

	// LOOP i: 0 [i < 5] {
	//	 x := i
	// }

	given := []Token{
		tok(TK_LOOP, "LOOP"),
		tok(TK_IDENTIFIER, "i"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_NUMBER, "0"),
		tok(TK_GUARD_OPEN, "["),
		tok(TK_IDENTIFIER, "i"),
		tok(TK_LESS_THAN, "<"),
		tok(TK_NUMBER, "5"),
		tok(TK_GUARD_CLOSE, "]"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_IDENTIFIER, "i"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
	}

	loop := Loop{
		Open:      tok(TK_LOOP, "LOOP"),
		IndexId:   tok(TK_IDENTIFIER, "i"),
		InitIndex: Value{tok(TK_NUMBER, "0")},
	}

	condition := Operation{
		Identifier{tok(TK_IDENTIFIER, "i")},
		tok(TK_LESS_THAN, "<"),
		Value{tok(TK_NUMBER, "5")},
	}

	guard := Guard{
		Open:      tok(TK_GUARD_OPEN, "["),
		Condition: condition,
		Close:     tok(TK_GUARD_CLOSE, "]"),
	}

	stat := Assignment{
		false,
		[]AssignTarget{AssignTarget{tok(TK_IDENTIFIER, "x"), nil}},
		tok(TK_ASSIGNMENT, ":"),
		[]Expression{Identifier{tok(TK_IDENTIFIER, "i")}},
	}

	guard.Block = Block{
		tok(TK_BLOCK_OPEN, "{"),
		[]Statement{stat},
		tok(TK_BLOCK_CLOSE, "}"),
	}

	loop.Guard = guard

	expectOneStat(t, loop, f(given))
}

func LP2_ForEach(t *testing.T, f ParseFunc) {

	// LOOP i, v, m <- list {
	//	 x := i
	// }

	given := []Token{
		tok(TK_LOOP, "LOOP"),
		tok(TK_IDENTIFIER, "i"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "v"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "m"),
		tok(TK_UPDATES, "<-"),
		tok(TK_IDENTIFIER, "list"),
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_IDENTIFIER, "x"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_IDENTIFIER, "i"),
		tok(TK_TERMINATOR, ""),
		tok(TK_BLOCK_CLOSE, "}"),
	}

	forEach := ForEach{
		Open:    tok(TK_LOOP, "LOOP"),
		IndexId: tok(TK_IDENTIFIER, "i"),
		ValueId: tok(TK_IDENTIFIER, "v"),
		MoreId:  tok(TK_IDENTIFIER, "m"),
		List:    Identifier{tok(TK_IDENTIFIER, "list")},
	}

	stat := Assignment{
		false,
		[]AssignTarget{AssignTarget{tok(TK_IDENTIFIER, "x"), nil}},
		tok(TK_ASSIGNMENT, ":"),
		[]Expression{Identifier{tok(TK_IDENTIFIER, "i")}},
	}

	forEach.Block = Block{
		tok(TK_BLOCK_OPEN, "{"),
		[]Statement{stat},
		tok(TK_BLOCK_CLOSE, "}"),
	}

	expectOneStat(t, forEach, f(given))
}
