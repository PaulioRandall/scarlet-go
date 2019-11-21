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

func TestCountWord_1(t *testing.T) {
	n := countWord([]rune(""))
	assert.Equal(t, 0, n)
}

func TestCountWord_2(t *testing.T) {
	n := countWord([]rune("Happy"))
	assert.Equal(t, 5, n)
}

func TestCountWord_3(t *testing.T) {
	n := countWord([]rune("Left Right"))
	assert.Equal(t, 4, n)
}
