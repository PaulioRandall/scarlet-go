package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken_IsSignificant_1(t *testing.T) {
	assert.False(t, Token{Kind: UNDEFINED}.IsSignificant())
	assert.False(t, Token{Kind: WHITESPACE}.IsSignificant())
	assert.True(t, Token{Kind: PROCEDURE}.IsSignificant())
	assert.True(t, Token{Kind: END}.IsSignificant())
}
