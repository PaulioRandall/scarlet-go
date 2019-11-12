package count

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWord_1(t *testing.T) {
	n := CountWord([]rune(""))
	assert.Equal(t, 0, n)
}

func TestCountWord_2(t *testing.T) {
	n := CountWord([]rune("Happy"))
	assert.Equal(t, 5, n)
}

func TestCountWord_3(t *testing.T) {
	n := CountWord([]rune("Left Right"))
	assert.Equal(t, 4, n)
}
