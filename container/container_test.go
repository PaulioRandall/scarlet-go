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
