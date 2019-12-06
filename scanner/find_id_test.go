package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func TestFindId_1(t *testing.T) {
	var _ source.TokenFinder = findId
}

func TestFindId_2(t *testing.T) {
	r := []rune("abc")
	n, k := findId(r)

	assert.Equal(t, 3, n)
	assert.Equal(t, token.ID, k)
}

func TestFindId_3(t *testing.T) {
	r := []rune("abc efg")
	n, k := findId(r)

	assert.Equal(t, 3, n)
	assert.Equal(t, token.ID, k)
}

func TestFindId_4(t *testing.T) {
	r := []rune("   ")
	n, k := findId(r)

	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
