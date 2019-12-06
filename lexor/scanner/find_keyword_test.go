package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func TestFindKeyword_1(t *testing.T) {
	// Check it is a type of source.TokenFinder.
	var _ source.TokenFinder = findKeyword
}

func TestFindKeyword_2(t *testing.T) {
	// Check it works when a keyword is the only input token.

	r := []rune("FUNC")
	n, k := findKeyword(r)

	assert.Equal(t, 4, n)
	assert.Equal(t, token.FUNC, k)
}

func TestFindKeyword_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a keyword
	// is the first.

	r := []rune("FUNC END")
	n, k := findKeyword(r)

	assert.Equal(t, 4, n)
	assert.Equal(t, token.FUNC, k)
}

func TestFindKeyword_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a keyword.

	r := []rune("   ")
	n, k := findKeyword(r)

	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
