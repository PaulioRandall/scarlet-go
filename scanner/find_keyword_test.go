package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func TestFindKeyword_1(t *testing.T) {
	var _ source.TokenFinder = findKeyword
}

func TestFindKeyword_2(t *testing.T) {
	r := []rune("FUNC")
	n, k := findKeyword(r)

	assert.Equal(t, 4, n)
	assert.Equal(t, token.FUNC, k)
}

func TestFindKeyword_3(t *testing.T) {
	r := []rune("END")
	n, k := findKeyword(r)

	assert.Equal(t, 3, n)
	assert.Equal(t, token.END, k)
}

func TestFindKeyword_4(t *testing.T) {
	r := []rune("   ")
	n, k := findKeyword(r)

	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
