package series

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"

	"github.com/stretchr/testify/require"
)

func requireSeries(t *testing.T, s *Series, lexs ...lexeme.Lexeme) {
	expHead, expTail, _ := chainLexemes(lexs...)
	requireChain(t, expHead, s.head)
	requireChain(t, expTail, s.tail)
}

func TestSeries_Prepend(t *testing.T) {

	a, b, c, d := dummyLexemes()
	s := New()

	s.Prepend(d)
	requireSeries(t, s, d)

	s.Prepend(c)
	requireSeries(t, s, c, d)

	s.Prepend(b)
	requireSeries(t, s, b, c, d)

	s.Prepend(a)
	requireSeries(t, s, a, b, c, d)
}

func TestSeries_Append(t *testing.T) {

	a, b, c, d := dummyLexemes()
	s := New()

	s.Append(a)
	requireSeries(t, s, a)

	s.Append(b)
	requireSeries(t, s, a, b)

	s.Append(c)
	requireSeries(t, s, a, b, c)

	s.Append(d)
	requireSeries(t, s, a, b, c, d)
}

func TestSeries_InsertAfter(t *testing.T) {

	la, lb, lc, _ := dummyLexemes()
	na, _, nc, _ := dummyNodes()

	s := new(na)
	s.Next()
	s.InsertAfter(lb)
	requireSeries(t, s, la, lb)

	na.unlink()

	s = new(na, nc)
	s.Next()
	s.InsertAfter(lb)
	requireSeries(t, s, la, lb, lc)
}

func TestSeries_InsertBefore(t *testing.T) {

	la, lb, lc, _ := dummyLexemes()
	na, nb, nc, _ := dummyNodes()

	s := new(nb)
	s.Next()
	s.InsertBefore(la)
	requireSeries(t, s, la, lb)

	nb.unlink()

	s = new(na, nc)
	s.Next()
	s.Next()
	s.InsertBefore(lb)
	requireSeries(t, s, la, lb, lc)
}

func TestSeries_Remove(t *testing.T) {

	la, lb, lc, _ := dummyLexemes()
	na, nb, nc, _ := dummyNodes()

	s := new(na, nb, nc)
	s.Next()
	act := s.Remove()
	require.Equal(t, la, act)
	requireSeries(t, s, lb, lc)

	unlinkAll(na, nb, nc)

	s = new(na, nb, nc)
	s.Next()
	s.Next()
	act = s.Remove()
	require.Equal(t, lb, act)
	requireSeries(t, s, la, lc)
}
