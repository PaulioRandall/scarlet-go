package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func doTestToken_IsSignificant(t *testing.T, k Kind, exp bool) {
	act := Token{Kind: k}.IsSignificant()
	assert.Equal(t, exp, act, "Input: '%s'", k.Name())
}

func TestToken_IsSignificant_1(t *testing.T) {
	doTestToken_IsSignificant(t, UNDEFINED, false)
	doTestToken_IsSignificant(t, WHITESPACE, false)
	doTestToken_IsSignificant(t, PROCEDURE, true)
	doTestToken_IsSignificant(t, END, true)
}
