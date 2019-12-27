package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenReader_Peek_1(t *testing.T) {
	// Test single token

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.ID, "efg", 0, 3),
	}

	tb := NewTokenReader(lexor.DummyScanToken(stream))

	doTest := func(exp token.Token) {
		act := tb.Peek()
		require.Equal(t, exp, act)
		require.Nil(t, tb.Err())
	}

	doTest(stream[0])
	doTest(stream[0])
}

func TestTokenReader_Read_1(t *testing.T) {
	// Test multiple tokens

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.ID, "efg", 0, 3),
		token.NewToken(token.ID, "hij", 0, 6),
		token.Token{},
	}

	tb := NewTokenReader(lexor.DummyScanToken(stream))

	doTest := func(expMore bool, exp token.Token) {
		assert.Equal(t, expMore, tb.HasMore())
		act := tb.Read()
		require.Equal(t, exp, act)
		require.Nil(t, tb.Err())
	}

	doTest(true, stream[0])
	doTest(true, stream[1])
	doTest(true, stream[2])
	doTest(true, stream[3])

	doTest(false, token.Token{})
}
