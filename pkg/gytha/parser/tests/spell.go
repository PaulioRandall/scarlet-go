package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func SP1(t *testing.T, f ParseFunc) {

	// @f()

	given := []Token{
		tok(TK_SPELL, "@"),
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, ""),
	}

	sp := SpellCall{
		tok(TK_IDENTIFIER, "f"),
		nil,
	}

	expectOneStat(t, sp, f(given))
}

func SP2(t *testing.T, f ParseFunc) {

	// @f(a, b)

	given := []Token{
		tok(TK_SPELL, "@"),
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, "\n"),
	}

	exp := SpellCall{
		tok(TK_IDENTIFIER, "f"),
		[]Expression{
			Identifier{tok(TK_IDENTIFIER, "a")},
			Identifier{tok(TK_IDENTIFIER, "b")},
		},
	}

	expectOneStat(t, exp, f(given))
}

func SP3(t *testing.T, f ParseFunc) {

	// @f(1 + 2 - 3)

	given := []Token{
		tok(TK_SPELL, "@"),
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

	add := Operation{
		Value{tok(TK_NUMBER, "1")},
		tok(TK_PLUS, "+"),
		Value{tok(TK_NUMBER, "2")},
	}

	sub := Operation{
		Left:     add,
		Operator: tok(TK_MINUS, "-"),
		Right:    Value{tok(TK_NUMBER, "3")},
	}

	sp := SpellCall{
		ID:     tok(TK_IDENTIFIER, "f"),
		Inputs: []Expression{sub},
	}

	expectOneStat(t, sp, f(given))
}

func SP4(t *testing.T, f ParseFunc) {

	// @f(@abc())

	given := []Token{
		tok(TK_SPELL, "@"),
		tok(TK_IDENTIFIER, "f"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_SPELL, "@"),
		tok(TK_IDENTIFIER, "abc"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_TERMINATOR, "\n"),
	}

	inner := SpellCall{
		ID:     tok(TK_IDENTIFIER, "abc"),
		Inputs: nil,
	}

	outer := SpellCall{
		ID:     tok(TK_IDENTIFIER, "f"),
		Inputs: []Expression{inner},
	}

	expectOneStat(t, outer, f(given))
}
