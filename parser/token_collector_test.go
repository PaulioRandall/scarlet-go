package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func setupTokenCollector(stream []token.Token) *TokenCollector {
	st := lexor.DummyScanToken(stream)
	tr := NewTokenReader(st)
	return NewTokenCollector(tr)
}

func TestTokenCollector_Peek_1(t *testing.T) {

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.ID, "efg", 0, 3),
	}

	tc := setupTokenCollector(stream)

	doTest := func(exp token.Token) {
		act := tc.Peek()

		require.Equal(t, exp, act)
		require.Nil(t, tc.Err())
		require.Equal(t, 0, tc.index)
		require.Equal(t, 1, len(tc.buffer))
	}

	doTest(stream[0])
	doTest(stream[0])
}

func TestTokenCollector_Read_1(t *testing.T) {

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.ID, "efg", 0, 3),
		token.NewToken(token.ID, "hij", 0, 6),
		token.Token{},
	}

	tc := setupTokenCollector(stream)

	doTest := func(expMore bool, expT token.Token, expBufIndex int) {
		require.Equal(t, expMore, tc.HasMore())
		act := tc.Read()

		require.Equal(t, expT, act)
		require.Nil(t, tc.Err())
		require.Equal(t, expBufIndex, tc.index)
		require.Equal(t, expBufIndex, len(tc.buffer))
	}

	doTest(true, stream[0], 1)
	doTest(true, stream[1], 2)
	doTest(true, stream[2], 3)
	doTest(true, stream[3], 3)

	doTest(false, token.Token{}, 3)
}

func TestTokenCollector_PutBack_1(t *testing.T) {

	stream := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.ID, "efg", 0, 3),
		token.NewToken(token.ID, "hij", 0, 6),
		token.Token{},
	}

	tc := setupTokenCollector(stream)

	doTest := func(expMore bool, expBufIndex, expBufLen int) {
		require.Equal(t, expMore, tc.HasMore())
		require.Nil(t, tc.Err())
		require.Equal(t, expBufIndex, tc.index)
		require.Equal(t, expBufLen, len(tc.buffer))
	}

	_ = tc.Read()
	_ = tc.Read()
	_ = tc.Read()

	doTest(true, 3, 3)

	tc.PutBack(2)

	doTest(true, 1, 3)
	require.Equal(t, stream[1], tc.Peek())
}
