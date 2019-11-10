package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoutSpaces_1(t *testing.T) {
	assert.Equal(t, 0, countSpaces([]rune("")))
	assert.Equal(t, 0, countSpaces([]rune("Scarlet")))
	assert.Equal(t, 3, countSpaces([]rune("   ")))
	assert.Equal(t, 3, countSpaces([]rune("   Scarlet")))
	assert.Equal(t, 3, countSpaces([]rune("\t\f\v")))
	assert.Equal(t, 1, countSpaces([]rune(" \n")))
	assert.Equal(t, 1, countSpaces([]rune(" \r\n")))
	assert.Equal(t, 2, countSpaces([]rune(" \r")))
}

func TestIsNewline_1(t *testing.T) {
	assert.Equal(t, 1, newlineRunes([]rune("\n"), 0))
	assert.Equal(t, 1, newlineRunes([]rune("Scarlet\n"), 7))
	assert.Equal(t, 2, newlineRunes([]rune("Scarlet\r\n"), 7))
	assert.Equal(t, 0, newlineRunes([]rune("Scarlet"), 0))
	assert.Equal(t, 0, newlineRunes([]rune("Scarlet\r"), 7))
	assert.Equal(t, 0, newlineRunes([]rune("Scarlet\n"), 0))
}
