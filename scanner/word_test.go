package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/stretchr/testify/assert"
)

func doTestIdentifyWord(t *testing.T, in string, expK token.Kind, expN int) {
	k, n := identifyWord([]rune(in))
	assert.Equal(t, expK, k, "Input: '%s'", in)
	assert.Equal(t, expN, n, "Input: '%s'", in)
}

func TestIdentifyWord_1(t *testing.T) {
	doTestIdentifyWord(t, "", token.UNDEFINED, 0)
	doTestIdentifyWord(t, "AAGGHH", token.UNDEFINED, 0)
	doTestIdentifyWord(t, "PROCEDURE", token.PROCEDURE, 9)
	doTestIdentifyWord(t, "END", token.END, 3)
}
