package container

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"

	"github.com/stretchr/testify/require"
)

func setupIterator(nodes ...*node) *Iterator {
	con := setupContainer(nodes...)
	return con.Iterator()
}

func Test_Iterator_Remove(t *testing.T) {

	a, b, c, _ := dummyNodes()
	it := setupIterator(a, b, c)
	it.Next()
	it.Next()

	l := it.Remove()
	require.Equal(t, b.data, l)
	require.Equal(t, a, c.prev)
	require.Equal(t, c, a.next)
	require.Equal(t, a, it.con.head)
	require.Equal(t, c, it.con.tail)
}

func Test_Iterator_InsertBefore(t *testing.T) {

	a, b, c, _ := dummyNodes()
	it := setupIterator(a, c)
	it.Next()
	it.Next()

	it.InsertBefore(b.data)
	b.prev = a
	b.next = c

	require.Equal(t, b, it.prev)
	require.Equal(t, c, it.curr)
	require.Nil(t, it.next)
}

func Test_Iterator_InsertAfter(t *testing.T) {

	a, b, c, _ := dummyNodes()
	it := setupIterator(a, c)
	it.Next()

	it.InsertAfter(b.data)
	b.prev = a
	b.next = c

	require.Nil(t, it.prev)
	require.Equal(t, a, it.curr)
	require.Equal(t, b, it.next)
}

func dummyItrData() (a, b, c *node) {
	a = &node{
		data: lexeme.New("true", lexeme.BOOL, 0, 0),
	}

	b = &node{
		data: lexeme.New("1", lexeme.NUMBER, 0, 4),
	}

	c = &node{
		data: lexeme.New("abc", lexeme.STRING, 0, 5),
	}

	return
}

func Test_Iterator_JumpToNext(t *testing.T) {

	a, b, c := dummyItrData()
	it := setupIterator(a, b, c)

	it.JumpToNext(func(v View) bool {
		return v.Item().Type() == lexeme.NUMBER
	})
	require.Equal(t, b.data, it.Item())
	require.True(t, it.More())

	it.JumpToNext(func(v View) bool {
		return v.Item().Type() == lexeme.NUMBER
	})
	require.False(t, it.More())
}

func Test_Iterator_JumpToPrev(t *testing.T) {

	a, b, c := dummyItrData()
	it := setupIterator(a, b, c)
	it.jumpToEnd()

	it.JumpToPrev(func(v View) bool {
		return v.Item().Type() == lexeme.NUMBER
	})
	require.Equal(t, b.data, it.Item())
	require.True(t, !it.IsFirst())

	it.JumpToPrev(func(v View) bool {
		return v.Item().Type() == lexeme.NUMBER
	})
	require.False(t, !it.IsFirst())
}
