package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/source"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func TestFindNewline_1(t *testing.T) {
	var _ source.TokenFinder = findNewline
}

func TestFindNewline_2(t *testing.T) {
	r := []rune("\n")
	n, k := findNewline(r)

	assert.Equal(t, 1, n)
	assert.Equal(t, token.NEWLINE, k)
}

func TestFindNewline_3(t *testing.T) {
	r := []rune("\r\n")
	n, k := findNewline(r)

	assert.Equal(t, 2, n)
	assert.Equal(t, token.NEWLINE, k)
}

func TestFindNewline_4(t *testing.T) {
	r := []rune("\r\nabc")
	n, k := findNewline(r)

	assert.Equal(t, 2, n)
	assert.Equal(t, token.NEWLINE, k)
}

func TestFindNewline_5(t *testing.T) {
	r := []rune("   ")
	n, k := findNewline(r)

	assert.Equal(t, 0, n)
	assert.Equal(t, token.UNDEFINED, k)
}
