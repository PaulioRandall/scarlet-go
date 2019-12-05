package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenSimple_1(t *testing.T) {
	_ = Token(tokenSimple{})
}

func TestTokenSimple_IsSignificant_1(t *testing.T) {
	kinds := map[Kind]bool{
		// Kinds that are significant
		FUNC: true,
		END:  true,
		// Kinds that are NOT significant
		UNDEFINED:  false,
		WHITESPACE: false,
	}

	for k, exp := range kinds {
		act := tokenSimple{k: k}.IsSignificant()
		assert.Equal(t, exp, act)
	}
}
