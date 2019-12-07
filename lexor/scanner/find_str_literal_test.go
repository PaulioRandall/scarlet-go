package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindLiteral_1(t *testing.T) {
	// Check it is a type of source.TokenFinder.
	var _ source.TokenFinder = findStrLiteral
}

func TestFindLiteral_2(t *testing.T) {
	// Check it works when a string literal is the only input token.

	r := []rune("`abc @~\"`")
	n, k, e := findStrLiteral(r)

	require.Nil(t, e)
	assert.Equal(t, 9, n)
	assert.Equal(t, token.STR_LITERAL, k)
}

func TestFindLiteral_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a string
	// literal is the first.

	r := []rune("`abc` efg")
	n, k, e := findStrLiteral(r)

	require.Nil(t, e)
	assert.Equal(t, 5, n)
	assert.Equal(t, token.STR_LITERAL, k)
}

func TestFindLiteral_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a string
	// literal.

	r := []rune("   ")
	n, k, e := findStrLiteral(r)

	require.Nil(t, e)
	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
