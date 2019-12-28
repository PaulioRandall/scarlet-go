package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func TestTokenReader_Peek_1(t *testing.T) {

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.ID, "efg", 0, 3),
	}

	tr := NewTokenReader(lexor.DummyScanToken(stream))

	doTest := func(exp token.Token) {
		act := tr.Peek()
		require.Equal(t, exp, act)
	}

	doTest(stream[0])
	doTest(stream[0])
}

func TestTokenReader_Read_1(t *testing.T) {

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.ID, "efg", 0, 3),
		token.NewToken(token.ID, "hij", 0, 6),
		token.Token{},
	}

	tr := NewTokenReader(lexor.DummyScanToken(stream))

	doTest := func(expMore bool, exp token.Token) {
		require.Equal(t, expMore, tr.HasMore())

		act := tr.Read()
		require.Equal(t, exp, act)
	}

	doTest(true, stream[0])
	doTest(true, stream[1])
	doTest(true, stream[2])
	doTest(true, stream[3])

	doTest(false, token.Token{})
}
