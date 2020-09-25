package container

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func dummyNodes() (a, b, c, d *node) {

	a = &node{}
	b = &node{}
	c = &node{}
	d = &node{}

	a.data = token.New("true", token.BOOL, 0, 0)
	a.next = b

	b.prev = a
	b.data = token.New("1", token.NUMBER, 0, 4)
	b.next = c

	c.prev = b
	c.data = token.New("abc", token.STRING, 0, 5)
	c.next = d

	d.prev = c
	d.data = token.New("i", token.IDENT, 0, 8)

	return
}

func dummyItr() (itr *Iterator, a, b, c, d *node) {
	a, b, c, d = dummyNodes()
	itr = &Iterator{}
	return
}

func Test_Iterator_Remove(t *testing.T) {

	var (
		l       token.Lexeme
		a, b, c *node
		itr     *Iterator
	)

	itr, a, b, _, _ = dummyItr()
	itr.curr = a

	l = itr.Remove()
	require.Equal(t, a.data, l)
	require.Nil(t, b.prev)

	itr, a, b, c, _ = dummyItr()
	itr.curr = b

	l = itr.Remove()
	require.Equal(t, b.data, l)
	require.Equal(t, c, a.next)
	require.Equal(t, a, c.prev)
}

func Test_Iterator_InsertBefore(t *testing.T) {

}

/*
func Test_Iterator_2_1(t *testing.T) {

	itr, a, b, c, _ := setupItrAt("front")

	require.True(t, itr.HasNext())
	require.True(t, itr.Next())
	halfEqual(t, nil, itr.Before())
	halfEqual(t, a, itr.Curr())
	halfEqual(t, b, itr.After())

	require.True(t, itr.HasNext())
	require.True(t, itr.Next())
	halfEqual(t, a, itr.Before())
	halfEqual(t, b, itr.Curr())
	halfEqual(t, c, itr.After())

	require.True(t, itr.HasNext())
	require.True(t, itr.Next())
	halfEqual(t, b, itr.Before())
	halfEqual(t, c, itr.Curr())
	halfEqual(t, nil, itr.After())

	require.False(t, itr.HasNext())
	require.False(t, itr.Next())
	halfEqual(t, c, itr.Before())
	halfEqual(t, nil, itr.Curr())
	halfEqual(t, nil, itr.After())

	require.False(t, itr.HasNext())
	require.False(t, itr.Next())
	halfEqual(t, c, itr.Before())
	halfEqual(t, nil, itr.Curr())
	halfEqual(t, nil, itr.After())
}

func Test_Iterator_2_2(t *testing.T) {

	itr, a, b, c, _ := setupItrAt("back")

	halfEqual(t, c, itr.Before())
	halfEqual(t, nil, itr.Curr())
	halfEqual(t, nil, itr.After())

	require.True(t, itr.HasPrev())
	require.True(t, itr.Prev())
	halfEqual(t, b, itr.Before())
	halfEqual(t, c, itr.Curr())
	halfEqual(t, nil, itr.After())

	require.True(t, itr.HasPrev())
	require.True(t, itr.Prev())
	halfEqual(t, a, itr.Before())
	halfEqual(t, b, itr.Curr())
	halfEqual(t, c, itr.After())

	require.True(t, itr.HasPrev())
	require.True(t, itr.Prev())
	halfEqual(t, nil, itr.Before())
	halfEqual(t, a, itr.Curr())
	halfEqual(t, b, itr.After())

	require.False(t, itr.HasPrev())
	require.False(t, itr.Prev())
	halfEqual(t, nil, itr.Before())
	halfEqual(t, nil, itr.Curr())
	halfEqual(t, a, itr.After())

	require.False(t, itr.HasPrev())
	require.False(t, itr.Prev())
	halfEqual(t, nil, itr.Before())
	halfEqual(t, nil, itr.Curr())
	halfEqual(t, a, itr.After())
}

func Test_Iterator_3_1(t *testing.T) {

	itr, a, b, c, _ := setupItrAt("mid")

	z := itr.Remove()
	fullEqual(t, b, nil, nil, z)
	halfEqual(t, a, itr.Before())
	halfEqual(t, nil, itr.Curr())
	halfEqual(t, c, itr.After())

	itr.Next()
	halfEqual(t, a, itr.Before())
	halfEqual(t, c, itr.Curr())
	halfEqual(t, nil, itr.After())

	z = itr.Remove()
	fullEqual(t, c, nil, nil, z)
	halfEqual(t, a, itr.Before())
	halfEqual(t, nil, itr.Curr())
	halfEqual(t, nil, itr.After())
}
*/
