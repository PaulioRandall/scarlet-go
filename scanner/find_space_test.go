package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func TestFindSpace_1(t *testing.T) {
	var _ source.TokenFinder = findSpace
}

func TestFindSpace_2(t *testing.T) {
	r := []rune(" \t\v\f")
	n, k := findSpace(r)

	assert.Equal(t, 4, n)
	assert.Equal(t, token.WHITESPACE, k)
}

func TestFindSpace_3(t *testing.T) {
	r := []rune("  ab")
	n, k := findSpace(r)

	assert.Equal(t, 2, n)
	assert.Equal(t, token.WHITESPACE, k)
}

func TestFindSpace_4(t *testing.T) {
	r := []rune("abc")
	n, k := findSpace(r)

	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
