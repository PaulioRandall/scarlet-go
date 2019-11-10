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
	assert.True(t, isNewline(0, []rune("\n")))
	assert.True(t, isNewline(7, []rune("Scarlet\n")))
	assert.True(t, isNewline(7, []rune("Scarlet\r\n")))
	assert.False(t, isNewline(0, []rune("Scarlet")))
	assert.False(t, isNewline(7, []rune("Scarlet\r")))
	assert.False(t, isNewline(0, []rune("Scarlet\n")))
}
