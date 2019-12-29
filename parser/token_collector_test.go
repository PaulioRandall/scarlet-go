package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func TestTokenCollector_Peek_1(t *testing.T) {

	stream := []token.Token{
		token.OfValue(token.ID, "abc"),
		token.OfValue(token.ID, "efg"),
	}

	tc := dummyTC(stream)

	doTest := func(exp token.Token) {
		act := tc.Peek()

		require.Equal(t, exp, act)
		require.Equal(t, 0, tc.index)
		require.Equal(t, 1, len(tc.buffer))
	}

	doTest(stream[0])
	doTest(stream[0])
}

func TestTokenCollector_Read_1(t *testing.T) {

	stream := []token.Token{
		token.OfValue(token.ID, "abc"),
		token.OfValue(token.ID, "efg"),
		token.OfValue(token.ID, "hij"),
		token.Token{},
	}

	tc := dummyTC(stream)

	doTest := func(expMore bool, expT token.Token, expBufIndex int) {
		require.Equal(t, expMore, tc.HasMore())
		act := tc.Read()

		require.Equal(t, expT, act)
		require.Equal(t, expBufIndex, tc.index)
		require.Equal(t, expBufIndex, len(tc.buffer))
	}

	doTest(true, stream[0], 1)
	doTest(true, stream[1], 2)
	doTest(true, stream[2], 3)
	doTest(true, stream[3], 3)

	doTest(false, token.Token{}, 3)
}

func TestTokenCollector_Unread_1(t *testing.T) {

	stream := []token.Token{
		token.OfValue(token.ID, "abc"),
		token.OfValue(token.ID, "efg"),
		token.OfValue(token.ID, "hij"),
		token.Token{},
	}

	tc := dummyTC(stream)

	doTest := func(expMore bool, expBufIndex, expBufLen int) {
		require.Equal(t, expMore, tc.HasMore())
		require.Equal(t, expBufIndex, tc.index)
		require.Equal(t, expBufLen, len(tc.buffer))
	}

	_ = tc.Read()
	_ = tc.Read()
	_ = tc.Read()

	doTest(true, 3, 3)

	tc.Unread(2)

	doTest(true, 1, 3)
	require.Equal(t, stream[1], tc.Peek())
}
