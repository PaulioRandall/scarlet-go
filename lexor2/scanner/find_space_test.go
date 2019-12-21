package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindSpace_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findSpace
}

func TestFindSpace_2(t *testing.T) {
	// Check it works when whitespace is the only input token.

	r := []rune(" \t\v\f")
	n, k, e := findSpace(r)

	require.Nil(t, e)
	assert.Equal(t, 4, n)
	assert.Equal(t, token.WHITESPACE, k)
}

func TestFindSpace_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and whitespace
	// is the first.

	r := []rune("  ab")
	n, k, e := findSpace(r)

	require.Nil(t, e)
	assert.Equal(t, 2, n)
	assert.Equal(t, token.WHITESPACE, k)
}

func TestFindSpace_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not whitespace.

	r := []rune("abc")
	n, k, e := findSpace(r)

	require.Nil(t, e)
	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
