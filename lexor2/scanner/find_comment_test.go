package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindComment_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findComment
}

func TestFindComment_2(t *testing.T) {
	// Check it works when a comment is the only input token.

	r := []rune("// Die Hard is a Christmas movie")
	n, k, e := findComment(r)

	require.Nil(t, e)
	assert.Equal(t, 32, n)
	assert.Equal(t, token.COMMENT, k)
}

func TestFindComment_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a comment is
	// the first.

	r := []rune("// abc\nefg")
	n, k, e := findComment(r)

	require.Nil(t, e)
	assert.Equal(t, 6, n)
	assert.Equal(t, token.COMMENT, k)
}

func TestFindComment_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a comment.

	r := []rune("   ")
	n, k, e := findComment(r)

	require.Nil(t, e)
	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
