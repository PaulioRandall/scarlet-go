package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/stretchr/testify/assert"
)

func TestScan_1(t *testing.T) {
	s := source{
		runes: []rune("Scarlet"),
	}

	exp := token.Token{
		Kind:  token.UNDEFINED,
		Value: "Scar",
		Line:  0,
		Start: 0,
		End:   4,
	}

	act := s.scan(4, token.UNDEFINED)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("let"), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 4, s.col)
}
