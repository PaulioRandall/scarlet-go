package z_tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func LP1_Assignment(t *testing.T, f ParseFunc) {

	// LOOP i [i < 5] {
	//	 x := i
	// }

	given := []Token{
		tok(LOOP, "LOOP"),
		tok(IDENTIFIER, "i"),
		tok(GUARD_OPEN, "["),
		tok(IDENTIFIER, "i"),
		tok(LESS_THAN, "<"),
		tok(NUMBER, "5"),
		tok(GUARD_CLOSE, "]"),
		tok(BLOCK_OPEN, "{"),
		tok(IDENTIFIER, "x"),
		tok(ASSIGN, ":="),
		tok(IDENTIFIER, "i"),
		tok(TERMINATOR, ""),
		tok(BLOCK_CLOSE, "}"),
	}

	loop := Loop{
		Open:     tok(LOOP, "LOOP"),
		IndexVar: tok(IDENTIFIER, "i"),
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
		tok(ASSIGN, ":="),
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
