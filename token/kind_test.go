package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func doTestFindWordKind(t *testing.T, in string, exp Kind) {
	act := FindWordKind(in)
	assert.Equal(t, exp, act, "Input: '%s'", in)
}

func TestFindWordKind_1(t *testing.T) {
	doTestFindWordKind(t, "", UNDEFINED)
	doTestFindWordKind(t, "AAGGHH", UNDEFINED)
	doTestFindWordKind(t, "ENDEND", UNDEFINED)
	doTestFindWordKind(t, "FUNC", FUNC)
	doTestFindWordKind(t, "END", END)
}
