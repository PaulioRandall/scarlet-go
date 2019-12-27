package parser

/*
import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenBuffer_1(t *testing.T) {
	// Test single token

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
	}

	tb := NewTokenBuffer(lexor.DummyScanToken(stream))
	doTestTokenBuffer(t, tb, stream[0])

	doTestTokenBufferIsEmpty(t, tb)
}

func TestTokenBuffer_2(t *testing.T) {
	// Test multiple tokens

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.NEWLINE, "\n", 0, 3),
		token.NewToken(token.STR_LITERAL, "`efg`", 1, 3),
		token.Token{},
	}

	tb := NewTokenBuffer(lexor.DummyScanToken(stream))

	doTestTokenBuffer(t, tb, stream[0])
	doTestTokenBuffer(t, tb, stream[1])
	doTestTokenBuffer(t, tb, stream[2])
	doTestTokenBuffer(t, tb, stream[3])

	doTestTokenBufferIsEmpty(t, tb)
}

func doTestTokenBuffer(t *testing.T, tb *TokenReader, exp token.Token) {

	require.True(t, tb.HasMore())
	doTestTokenBuffer_Peek(t, tb, exp)
	doTestTokenBuffer_Reed(t, tb, exp)

	// Push it back on
	ok := tb.Push(exp)
	require.True(t, ok)

	// Read again, re-removing it from the buffer
	doTestTokenBuffer_Reed(t, tb, exp)
}

func doTestTokenBuffer_Reed(t *testing.T, tb *TokenReader, exp token.Token) {
	act := tb.Read()
	require.Equal(t, exp, act)
	require.Nil(t, tb.Err())
}

func doTestTokenBuffer_Peek(t *testing.T, tb *TokenReader, exp token.Token) {
	act := tb.Peek()
	require.Equal(t, exp, act)
	require.Nil(t, tb.Err())
}

func doTestTokenBufferIsEmpty(t *testing.T, tb *TokenReader) {
	assert.False(t, tb.HasMore())
	doTestTokenBuffer_Peek(t, tb, token.Token{})
	doTestTokenBuffer_Reed(t, tb, token.Token{})
}
*/
