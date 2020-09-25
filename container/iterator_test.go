package container

import (
	"testing"

	//	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func setupIterator(nodes ...*node) *Iterator {
	con := setupContainer(nodes...)
	return con.Iterator()
}

func Test_Iterator_Remove_1(t *testing.T) {

	a, b, _, _ := dummyNodes()
	itr := setupIterator(a, b)
	itr.Next()

	l := itr.Remove()
	require.Equal(t, a.data, l)
	require.Nil(t, b.prev)
	require.Equal(t, b, itr.con.head)
	require.Equal(t, b, itr.con.tail)
}

func Test_Iterator_Remove_2(t *testing.T) {

	a, b, c, _ := dummyNodes()
	itr := setupIterator(a, b, c)
	itr.Next()
	itr.Next()

	l := itr.Remove()
	require.Equal(t, b.data, l)
	require.Equal(t, a, c.prev)
	require.Equal(t, c, a.next)
	require.Equal(t, a, itr.con.head)
	require.Equal(t, c, itr.con.tail)
}

func Test_Iterator_InsertBefore_1(t *testing.T) {

	a, b, c, _ := dummyNodes()
	itr := setupIterator(b, c)
	itr.Next()

	itr.InsertBefore(a.data)
	a.next = b

	require.Equal(t, a, itr.prev)
	require.Equal(t, b, itr.curr)
	require.Equal(t, c, itr.next)
	require.Equal(t, a, itr.con.head)
}

func Test_Iterator_InsertBefore_2(t *testing.T) {

	a, b, c, _ := dummyNodes()
	itr := setupIterator(a, c)
	itr.Next()
	itr.Next()

	itr.InsertBefore(b.data)
	b.prev = a
	b.next = c

	require.Equal(t, b, itr.prev)
	require.Equal(t, c, itr.curr)
	require.Nil(t, itr.next)
	require.Equal(t, a, itr.con.head)
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
