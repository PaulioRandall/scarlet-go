package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func TestFindSymbol_1(t *testing.T) {
	// Check it is a type of source.TokenFinder.
	var _ source.TokenFinder = findSymbol
}

func TestFindSymbol_2(t *testing.T) {
	// Check it works on a range of valid inputs.

	f := func(s string, expN int, expK token.Kind) {
		r := []rune(s)
		n, k := findSymbol(r)

		assert.Equal(t, expN, n,
			"Odd number of runes in symbol")
		assert.Equal(t, expK, k,
			"Expected: %s, actual: %s", expK.Name(), k.Name())
	}

	// When input contains only one token, a symbol token
	f(":=", 2, token.ASSIGN)
	f("(", 1, token.OPEN_PAREN)
	f(")", 1, token.CLOSE_PAREN)
	f(",", 1, token.ID_DELIM)
	f("@", 1, token.SPELL)

	// When input contains multiple tokens, but the first is a symbol token
	f(":= 123.456", 2, token.ASSIGN)
	f("@Abc", 1, token.SPELL)
}

func TestFindSymbol_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a symbol.

	r := []rune("  :=")
	n, k := findId(r)

	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
