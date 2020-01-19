package parser

/*
import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func dummyTM(stream ...token.Token) *TokenMatcher {
	st := lexor.DummyScanToken(stream)
	return NewTokenMatcher(st)
}

func TestTokenMatcher_Read_1(t *testing.T) {

	stream := []token.Token{
		token.OfValue(token.ID, "abc"),
		token.OfValue(token.ID, "efg"),
	}

	sc := lexor.DummyScanToken(stream)
	tm := NewTokenMatcher(sc)

	doTest := func(exp token.Token, expOk bool) {
		act, ok := tm.Read()
		require.Equal(t, expOk, ok)
		require.Equal(t, exp, act)
		require.Equal(t, 0, len(tm.buffer))
	}

	doTest(stream[0], true)
	doTest(stream[1], true)
	doTest(token.ZERO(), false)
}

func TestTokenMatcher_ReadMany_1(t *testing.T) {

	stream := []token.Token{
		token.OfValue(token.ID, "abc"),
		token.OfValue(token.ID, "efg"),
		token.OfValue(token.ID, "hij"),
	}

	sc := lexor.DummyScanToken(stream)
	tm := NewTokenMatcher(sc)

	exp := stream[:2]

	act, n := tm.ReadMany(2)
	require.Equal(t, 2, n)
	require.Equal(t, exp, act)
	require.Equal(t, 0, len(tm.buffer))
}

func TestTokenMatcher_Match_1(t *testing.T) {

	stream := []token.Token{
		token.OfValue(token.ID, "abc"),
	}

	sc := lexor.DummyScanToken(stream)
	tm := NewTokenMatcher(sc)

	n := tm.Match(token.ID)
	require.Equal(t, 1, n)

	n = tm.Match(token.FUNC)
	require.Equal(t, 0, n)
}

func TestTokenMatcher_MatchAny_1(t *testing.T) {

	stream := []token.Token{
		token.OfValue(token.ID, "abc"),
	}

	sc := lexor.DummyScanToken(stream)
	tm := NewTokenMatcher(sc)

	n := tm.MatchAny(token.FUNC, token.ID)
	require.Equal(t, 1, n)

	n = tm.MatchAny(token.FUNC, token.END)
	require.Equal(t, 0, n)
}

func TestTokenMatcher_MatchSeq_1(t *testing.T) {

	stream := []token.Token{
		token.OfKind(token.ID),
		token.OfKind(token.ASSIGN),
		token.OfKind(token.BOOL_LITERAL),
	}

	sc := lexor.DummyScanToken(stream)
	tm := NewTokenMatcher(sc)

	n := tm.MatchSeq(token.ID, token.ASSIGN)
	require.Equal(t, 2, n)

	n = tm.MatchSeq(token.ID, token.DELIM)
	require.Equal(t, 0, n)
}
*/
