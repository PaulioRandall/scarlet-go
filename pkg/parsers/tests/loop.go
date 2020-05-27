package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func LP1_Conditional(t *testing.T, f ParseFunc) {

	// LOOP i: 0 [i < 5] {
	//	 x := i
	// }

	given := []Token{
		tok(LOOP, "LOOP"),
		tok(IDENTIFIER, "i"),
		tok(ASSIGN, ":"),
		tok(NUMBER, "0"),
		tok(GUARD_OPEN, "["),
		tok(IDENTIFIER, "i"),
		tok(LESS_THAN, "<"),
		tok(NUMBER, "5"),
		tok(GUARD_CLOSE, "]"),
		tok(BLOCK_OPEN, "{"),
		tok(IDENTIFIER, "x"),
		tok(ASSIGN, ":"),
		tok(IDENTIFIER, "i"),
		tok(TERMINATOR, ""),
		tok(BLOCK_CLOSE, "}"),
	}

	loop := Loop{
		Open:      tok(LOOP, "LOOP"),
		IndexId:   tok(IDENTIFIER, "i"),
		InitIndex: Value{tok(NUMBER, "0")},
	}

	condition := Operation{
		Identifier{tok(IDENTIFIER, "i")},
		tok(LESS_THAN, "<"),
		Value{tok(NUMBER, "5")},
	}

	guard := Guard{
		Open:      tok(GUARD_OPEN, "["),
		Condition: condition,
		Close:     tok(GUARD_CLOSE, "]"),
	}

	stat := Assignment{
		false,
		[]AssignTarget{AssignTarget{tok(IDENTIFIER, "x"), nil}},
		tok(ASSIGN, ":"),
		[]Expression{Identifier{tok(IDENTIFIER, "i")}},
	}

	guard.Block = Block{
		tok(BLOCK_OPEN, "{"),
		[]Statement{stat},
		tok(BLOCK_CLOSE, "}"),
	}

	loop.Guard = guard

	expectOneStat(t, loop, f(given))
}

func LP2_ForEach(t *testing.T, f ParseFunc) {

	// LOOP i, v, m <- list {
	//	 x := i
	// }

	given := []Token{
		tok(LOOP, "LOOP"),
		tok(IDENTIFIER, "i"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "v"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "m"),
		tok(UPDATES, "<-"),
		tok(IDENTIFIER, "list"),
		tok(BLOCK_OPEN, "{"),
		tok(IDENTIFIER, "x"),
		tok(ASSIGN, ":"),
		tok(IDENTIFIER, "i"),
		tok(TERMINATOR, ""),
		tok(BLOCK_CLOSE, "}"),
	}

	forEach := ForEach{
		Open:    tok(LOOP, "LOOP"),
		IndexId: tok(IDENTIFIER, "i"),
		ValueId: tok(IDENTIFIER, "v"),
		MoreId:  tok(IDENTIFIER, "m"),
		List:    Identifier{tok(IDENTIFIER, "list")},
	}

	stat := Assignment{
		false,
		[]AssignTarget{AssignTarget{tok(IDENTIFIER, "x"), nil}},
		tok(ASSIGN, ":"),
		[]Expression{Identifier{tok(IDENTIFIER, "i")}},
	}

	forEach.Block = Block{
		tok(BLOCK_OPEN, "{"),
		[]Statement{stat},
		tok(BLOCK_CLOSE, "}"),
	}

	expectOneStat(t, forEach, f(given))
}
