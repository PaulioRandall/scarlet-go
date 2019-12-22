package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindStrTemplate_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findStrTemplate
}

func TestFindStrTemplate_2(t *testing.T) {
	// Check it works when a string template is the only input token.

	r := []rune(`"abc @~\""`)
	n, k, e := findStrTemplate(r)

	require.Nil(t, e)
	assert.Equal(t, 10, n)
	assert.Equal(t, token.STR_TEMPLATE, k)
}

func TestFindStrTemplate_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a string
	// template is the first.

	r := []rune(`"abc" efg`)
	n, k, e := findStrTemplate(r)

	require.Nil(t, e)
	assert.Equal(t, 5, n)
	assert.Equal(t, token.STR_TEMPLATE, k)
}

func TestFindStrTemplate_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a string
	// template.

	r := []rune("   ")
	n, k, e := findStrTemplate(r)

	require.Nil(t, e)
	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
