package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindId_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findId
}

func TestFindId_2(t *testing.T) {
	// Check it works when an ID is the only input token.

	r := []rune("abc")
	n, k, e := findId(r)

	require.Nil(t, e)
	assert.Equal(t, 3, n)
	assert.Equal(t, token.ID, k)
}

func TestFindId_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and an ID is
	// the first.

	r := []rune("abc efg")
	n, k, e := findId(r)

	require.Nil(t, e)
	assert.Equal(t, 3, n)
	assert.Equal(t, token.ID, k)
}

func TestFindId_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not an ID.

	r := []rune("   ")
	n, k, e := findId(r)

	require.Nil(t, e)
	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
