package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenSimple_1(t *testing.T) {
	_ = Token(tokenSimple{})
}

func TestTokenSimple_IsSignificant_1(t *testing.T) {
	assert.False(t, tokenSimple{k: UNDEFINED}.IsSignificant())
	assert.False(t, tokenSimple{k: WHITESPACE}.IsSignificant())
	assert.True(t, tokenSimple{k: PROCEDURE}.IsSignificant())
	assert.True(t, tokenSimple{k: END}.IsSignificant())
}
