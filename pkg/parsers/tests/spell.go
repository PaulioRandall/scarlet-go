package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func SP1(t *testing.T, f ParseFunc) {

	// @f()

	given := []Token{
		tok(SPELL, "@"),
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, ""),
	}

	sp := SpellCall{
		Identifier{tok(IDENTIFIER, "f")},
		nil,
	}

	expectOneStat(t, sp, f(given))
}

func SP2(t *testing.T, f ParseFunc) {

	// @f(a, b)

	given := []Token{
		tok(SPELL, "@"),
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, "\n"),
	}

	exp := SpellCall{
		Identifier{tok(IDENTIFIER, "f")},
		[]Expression{
			Identifier{tok(IDENTIFIER, "a")},
			Identifier{tok(IDENTIFIER, "b")},
		},
	}

	expectOneStat(t, exp, f(given))
}

func SP3(t *testing.T, f ParseFunc) {

	// @f(1 + 2 - 3)

	given := []Token{
		tok(SPELL, "@"),
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

	add := Operation{
		Value{tok(NUMBER, "1")},
		tok(ADD, "+"),
		Value{tok(NUMBER, "2")},
	}

	sub := Operation{
		Left:     add,
		Operator: tok(SUBTRACT, "-"),
		Right:    Value{tok(NUMBER, "3")},
	}

	sp := SpellCall{
		ID:     Identifier{tok(IDENTIFIER, "f")},
		Inputs: []Expression{sub},
	}

	expectOneStat(t, sp, f(given))
}

func SP4(t *testing.T, f ParseFunc) {

	// @f(@abc())

	given := []Token{
		tok(SPELL, "@"),
		tok(IDENTIFIER, "f"),
		tok(PAREN_OPEN, "("),
		tok(SPELL, "@"),
		tok(IDENTIFIER, "abc"),
		tok(PAREN_OPEN, "("),
		tok(PAREN_CLOSE, ")"),
		tok(PAREN_CLOSE, ")"),
		tok(TERMINATOR, "\n"),
	}

	inner := SpellCall{
		ID:     Identifier{tok(IDENTIFIER, "abc")},
		Inputs: nil,
	}

	outer := SpellCall{
		ID:     Identifier{tok(IDENTIFIER, "f")},
		Inputs: []Expression{inner},
	}

	expectOneStat(t, outer, f(given))
}
