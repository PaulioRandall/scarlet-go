package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindNewline_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findNewline
}

func TestFindNewline_2(t *testing.T) {
	// Check it works when `\n` is the only input token.

	r := []rune("\n")
	n, k, e := findNewline(r)

	require.Nil(t, e)
	assert.Equal(t, 1, n)
	assert.Equal(t, token.NEWLINE, k)
}

func TestFindNewline_3(t *testing.T) {
	// Check it works when `\r\n` is the only input token.

	r := []rune("\r\n")
	n, k, e := findNewline(r)

	require.Nil(t, e)
	assert.Equal(t, 2, n)
	assert.Equal(t, token.NEWLINE, k)
}

func TestFindNewline_4(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a newline
	// is the first.

	r := []rune("\r\nabc")
	n, k, e := findNewline(r)

	require.Nil(t, e)
	assert.Equal(t, 2, n)
	assert.Equal(t, token.NEWLINE, k)
}

func TestFindNewline_5(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a newline.

	r := []rune("   ")
	n, k, e := findNewline(r)

	require.Nil(t, e)
	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
