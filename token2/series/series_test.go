package series

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"

	"github.com/stretchr/testify/require"
)

func requireSeries(t *testing.T, s Series, lexs ...lexeme.Lexeme) {
	expHead, expTail, _ := chainLexemes(lexs...)
	requireChain(t, expHead, s.head)
	requireChain(t, expTail, s.tail)
}

func TestSeries_Prepend(t *testing.T) {

	l1, l2, l3, l4 := dummyLexemes()
	s := Make()

	s.Prepend(l4)
	requireSeries(t, s, l4)

	s.Prepend(l3)
	requireSeries(t, s, l3, l4)

	s.Prepend(l2)
	requireSeries(t, s, l2, l3, l4)

	s.Prepend(l1)
	requireSeries(t, s, l1, l2, l3, l4)
}

func TestSeries_Append(t *testing.T) {

	l1, l2, l3, l4 := dummyLexemes()
	s := Make()

	s.Append(l1)
	requireSeries(t, s, l1)

	s.Append(l2)
	requireSeries(t, s, l1, l2)

	s.Append(l3)
	requireSeries(t, s, l1, l2, l3)

	s.Append(l4)
	requireSeries(t, s, l1, l2, l3, l4)
}

func TestSeries_InsertAfter(t *testing.T) {

	l1, l2, l3, _ := dummyLexemes()
	n1, _, n3, _ := dummyNodes()

	s := makeWith(n1)
	s.Next()
	s.InsertAfter(l2)
	requireSeries(t, s, l1, l2)

	unlinkAll(n1, n3)

	s = makeWith(n1, n3)
	s.Next()
	s.InsertAfter(l2)
	requireSeries(t, s, l1, l2, l3)
}

func TestSeries_InsertBefore(t *testing.T) {

	l1, l2, l3, _ := dummyLexemes()
	n1, n2, n3, _ := dummyNodes()

	s := makeWith(n2)
	s.Next()
	s.InsertBefore(l1)
	requireSeries(t, s, l1, l2)

	unlinkAll(n1, n2, n3)

	s = makeWith(n1, n3)
	s.Next()
	s.Next()
	s.InsertBefore(l2)
	requireSeries(t, s, l1, l2, l3)
}

func TestSeries_Remove(t *testing.T) {

	l1, l2, l3, _ := dummyLexemes()
	n1, n2, n3, _ := dummyNodes()

	s := makeWith(n1, n2, n3)
	s.Next()
	act := s.Remove()
	require.Equal(t, l1, act)
	requireSeries(t, s, l2, l3)

	unlinkAll(n1, n2, n3)

	s = makeWith(n1, n2, n3)
	s.Next()
	s.Next()
	act = s.Remove()
	require.Equal(t, l2, act)
	requireSeries(t, s, l1, l3)
}

func TestSeries_JumpToNext(t *testing.T) {

	l1, _, l3, _ := dummyLexemes()
	n1, n2, n3, n4 := dummyNodes()
	s := makeWith(n1, n2, n3, n4)

	s.JumpToNext(func(ro ReadOnly) bool {
		return ro.Get() == l1
	})
	require.Equal(t, l1, s.Get())

	s.JumpToNext(func(ro ReadOnly) bool {
		return ro.Get() == l3
	})
	require.Equal(t, l3, s.Get())

	s.JumpToNext(func(ro ReadOnly) bool {
		return false
	})
	require.False(t, s.More())
	require.Empty(t, s.Get())
}

func TestSeries_JumpToPrev(t *testing.T) {

	_, l2, _, l4 := dummyLexemes()
	n1, n2, n3, n4 := dummyNodes()
	s := makeWith(n1, n2, n3, n4)
	s.JumpToEnd()

	s.JumpToPrev(func(ro ReadOnly) bool {
		return ro.Get() == l4
	})
	require.Equal(t, l4, s.Get())

	s.JumpToPrev(func(ro ReadOnly) bool {
		return ro.Get() == l2
	})
	require.Equal(t, l2, s.Get())

	s.JumpToPrev(func(ro ReadOnly) bool {
		return false
	})
	require.True(t, s.More())
	require.Empty(t, s.Get())
}
