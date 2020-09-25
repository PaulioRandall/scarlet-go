package container

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func setupContainer(nodes ...*node) *Container {
	head, tail, size := chain(nodes...)
	con := &Container{
		head: head,
		tail: tail,
		size: size,
	}
	return con
}

func Test_Container_prepend(t *testing.T) {

	a, b, c, d := dummyNodes()
	con := &Container{}

	con.prepend(d)
	require.Equal(t, d, con.head)
	require.Equal(t, d, con.tail)

	con.prepend(c)
	require.Equal(t, c, con.head)
	require.Equal(t, d, con.tail)

	con.prepend(b)
	require.Equal(t, b, con.head)
	require.Equal(t, d, con.tail)

	con.prepend(a)
	require.Equal(t, a, con.head)
	require.Equal(t, d, con.tail)
}

func Test_Container_append(t *testing.T) {

	a, b, c, d := dummyNodes()
	con := &Container{}

	con.append(a)
	require.Equal(t, a, con.head)
	require.Equal(t, a, con.tail)

	con.append(b)
	require.Equal(t, a, con.head)
	require.Equal(t, b, con.tail)

	con.append(c)
	require.Equal(t, a, con.head)
	require.Equal(t, c, con.tail)

	con.append(d)
	require.Equal(t, a, con.head)
	require.Equal(t, d, con.tail)
}

func Test_Container_pop(t *testing.T) {

	var n *node
	a, b, c, d := dummyNodes()
	con := setupContainer(a, b, c, d)

	n = con.pop()
	require.Equal(t, a, n)
	require.Equal(t, b, con.head)
	require.Equal(t, d, con.tail)
	require.Equal(t, 3, con.size)

	n = con.pop()
	require.Equal(t, b, n)
	require.Equal(t, c, con.head)
	require.Equal(t, d, con.tail)
	require.Equal(t, 2, con.size)

	n = con.pop()
	require.Equal(t, c, n)
	require.Equal(t, d, con.head)
	require.Equal(t, d, con.tail)
	require.Equal(t, 1, con.size)

	n = con.pop()
	require.Equal(t, d, n)
	require.Nil(t, con.head)
	require.Nil(t, con.tail)
	require.Equal(t, 0, con.size)
}

func Test_Container_remove_1(t *testing.T) {

	a, b, _, _ := dummyNodes()
	con := setupContainer(a, b)

	con.remove(a)
	require.Nil(t, b.prev)
	require.Equal(t, b, con.head)
	require.Equal(t, b, con.tail)
	require.Equal(t, 1, con.size)
}

func Test_Container_remove_2(t *testing.T) {

	a, b, _, _ := dummyNodes()
	con := setupContainer(a, b)

	con.remove(b)
	require.Nil(t, a.next)
	require.Equal(t, a, con.head)
	require.Equal(t, a, con.tail)
	require.Equal(t, 1, con.size)
}

func Test_Container_remove_3(t *testing.T) {

	a, b, c, _ := dummyNodes()
	con := setupContainer(a, b, c)

	con.remove(b)
	require.Equal(t, a, c.prev)
	require.Equal(t, c, a.next)
	require.Equal(t, a, con.head)
	require.Equal(t, c, con.tail)
	require.Equal(t, 2, con.size)
}

func Test_Container_InsertBefore_1(t *testing.T) {

	a, b, _, _ := dummyNodes()
	con := setupContainer(b)

	con.insertBefore(con.head, a)
	require.Equal(t, a, con.head)
	require.Equal(t, b, con.tail)
}

func Test_Container_InsertBefore_2(t *testing.T) {

	a, b, c, _ := dummyNodes()
	con := setupContainer(a, c)

	con.insertBefore(c, b)
	require.Equal(t, a, con.head)
	require.Equal(t, a, b.prev)
	require.Equal(t, b, a.next)
	require.Equal(t, b, c.prev)
	require.Equal(t, c, b.next)
	require.Equal(t, c, con.tail)
}
