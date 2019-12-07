package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindKeyword_1(t *testing.T) {
	// Check it is a type of source.TokenFinder.
	var _ source.TokenFinder = findKeyword
}

func TestFindKeyword_2(t *testing.T) {
	// Check it works when a keyword is the only input token.

	r := []rune("F")
	n, k, e := findKeyword(r)

	require.Nil(t, e)
	assert.Equal(t, 1, n)
	assert.Equal(t, token.FUNC, k)
}

func TestFindKeyword_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a keyword
	// is the first.

	r := []rune("F END")
	n, k, e := findKeyword(r)

	require.Nil(t, e)
	assert.Equal(t, 1, n)
	assert.Equal(t, token.FUNC, k)
}

func TestFindKeyword_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a keyword.

	r := []rune("   ")
	n, k, e := findKeyword(r)

	require.Nil(t, e)
	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
