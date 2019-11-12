package cookies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNewline_1(t *testing.T) {
	assert.Equal(t, 1, NewlineRunes([]rune("\n"), 0))
	assert.Equal(t, 1, NewlineRunes([]rune("Scarlet\n"), 7))
	assert.Equal(t, 2, NewlineRunes([]rune("Scarlet\r\n"), 7))
	assert.Equal(t, 0, NewlineRunes([]rune("Scarlet"), 0))
	assert.Equal(t, 0, NewlineRunes([]rune("Scarlet\r"), 7))
	assert.Equal(t, 0, NewlineRunes([]rune("Scarlet\n"), 0))
}
