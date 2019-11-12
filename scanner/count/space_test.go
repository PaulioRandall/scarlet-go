package count

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoutSpaces_1(t *testing.T) {
	assert.Equal(t, 0, CountSpaces([]rune("")))
	assert.Equal(t, 0, CountSpaces([]rune("Scarlet")))
	assert.Equal(t, 3, CountSpaces([]rune("   ")))
	assert.Equal(t, 3, CountSpaces([]rune("   Scarlet")))
	assert.Equal(t, 3, CountSpaces([]rune("\t\f\v")))
	assert.Equal(t, 1, CountSpaces([]rune(" \n")))
	assert.Equal(t, 1, CountSpaces([]rune(" \r\n")))
	assert.Equal(t, 2, CountSpaces([]rune(" \r")))
}
