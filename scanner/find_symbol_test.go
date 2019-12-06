package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func TestFindSymbol_1(t *testing.T) {
	// Check it is a type of source.TokenFinder.
	var _ source.TokenFinder = findSymbol
}

func TestFindSymbol_2(t *testing.T) {
	// Check it works when a symbol is the only input token.

	r := []rune(":=")
	n, k := findSymbol(r)

	assert.Equal(t, 2, n)
	assert.Equal(t, token.ASSIGN, k)
}

func TestFindSymbol_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a symbol is
	// the first.

	r := []rune(":= 123.456")
	n, k := findSymbol(r)

	assert.Equal(t, 2, n)
	assert.Equal(t, token.ASSIGN, k)
}

func TestFindSymbol_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a symbol.

	r := []rune("  :=")
	n, k := findId(r)

	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
