package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {

	con := &Container{}

	_ = Collection(con)
	_ = List(con)
	_ = MutList(con)
	_ = Stack(con)
	_ = Queue(con)
}

func Test_NewContainer(t *testing.T) {

	a, b, c, _ := setupList()
	con := NewContainer(a)

	require.Equal(t, 3, con.size)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, c, con.head.Next)
	fullEqual(t, b, a, c, con.tail.Prev)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 3, con.size)
}

func Test_Container_Get(t *testing.T) {

	a, b, c, _ := setupList()
	con := NewContainer(a)

	halfEqual(t, a, con.Get(0))
	halfEqual(t, b, con.Get(1))
	halfEqual(t, c, con.Get(2))

	require.Panics(t, func() {
		con.Get(-1)
	})

	require.Panics(t, func() {
		con.Get(3)
	})
}

func Test_Container_Prepend(t *testing.T) {

	a, b, c, _ := setup()
	con := Container{}

	con.Prepend(c)
	fullEqual(t, c, nil, nil, con.head)
	fullEqual(t, c, nil, nil, con.tail)
	require.Equal(t, 1, con.size)

	con.Prepend(b)
	fullEqual(t, b, nil, c, con.head)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 2, con.size)

	con.Prepend(a)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, c, con.head.Next)
	fullEqual(t, b, a, c, con.tail.Prev)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 3, con.size)
}

func Test_Container_Append(t *testing.T) {

	a, b, c, _ := setup()
	con := Container{}

	con.Append(a)
	fullEqual(t, a, nil, nil, con.head)
	fullEqual(t, a, nil, nil, con.tail)
	require.Equal(t, 1, con.size)

	con.Append(b)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, nil, con.tail)
	require.Equal(t, 2, con.size)

	con.Append(c)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, c, con.head.Next)
	fullEqual(t, b, a, c, con.tail.Prev)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 3, con.size)
}

func Test_Container_InsertBefore(t *testing.T) {

	a, b, c, d := setup()
	con := NewContainer(c)

	con.InsertBefore(0, a)
	fullEqual(t, a, nil, c, con.head)
	fullEqual(t, c, a, nil, con.tail)
	require.Equal(t, 2, con.size)

	con.InsertBefore(1, b)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, c, con.head.Next)
	fullEqual(t, b, a, c, con.tail.Prev)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 3, con.size)

	require.Panics(t, func() {
		con.InsertBefore(-1, d)
	})

	require.Panics(t, func() {
		con.InsertBefore(3, d)
	})
}

func Test_Container_InsertAfter(t *testing.T) {

	a, b, c, d := setup()
	con := NewContainer(a)

	con.InsertAfter(0, b)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, nil, con.tail)
	require.Equal(t, 2, con.size)

	con.InsertAfter(1, c)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, c, con.head.Next)
	fullEqual(t, b, a, c, con.tail.Prev)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 3, con.size)

	require.Panics(t, func() {
		con.InsertAfter(-1, d)
	})

	require.Panics(t, func() {
		con.InsertAfter(3, d)
	})
}

func Test_Container_Remove(t *testing.T) {

	a, b, c, _ := setupList()
	con := NewContainer(a)

	z := con.Remove(1)
	fullEqual(t, a, nil, c, con.head)
	fullEqual(t, b, nil, nil, z)
	fullEqual(t, c, a, nil, con.tail)
	require.Equal(t, 2, con.size)

	z = con.Remove(0)
	fullEqual(t, c, nil, nil, con.head)
	fullEqual(t, c, nil, nil, con.tail)
	fullEqual(t, a, nil, nil, z)
	require.Equal(t, 1, con.size)

	z = con.Remove(0)
	halfEqual(t, nil, con.head)
	halfEqual(t, nil, con.tail)
	fullEqual(t, c, nil, nil, z)
	require.Equal(t, 0, con.size)
}
