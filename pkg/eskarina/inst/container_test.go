package inst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {

	con := &Container{}

	_ = Collection(con)
	_ = List(con)
	_ = Queue(con)
}

func Test_NewContainer(t *testing.T) {

	a, b, c, _ := setupList()
	con := NewContainer(a)

	fullEqual(t, a, b, con.head)
	fullEqual(t, b, c, con.head.Next)
	fullEqual(t, c, nil, con.head.Next.Next)
	fullEqual(t, c, nil, con.tail)
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

func Test_Container_Put(t *testing.T) {

	a, b, c, _ := setup()
	con := Container{}

	con.Put(a)
	fullEqual(t, a, nil, con.head)
	fullEqual(t, a, nil, con.tail)
	require.Equal(t, 1, con.size)

	con.Put(b)
	fullEqual(t, a, b, con.head)
	fullEqual(t, b, nil, con.tail)
	require.Equal(t, 2, con.size)

	con.Put(c)
	fullEqual(t, a, b, con.head)
	fullEqual(t, b, c, con.head.Next)
	fullEqual(t, c, nil, con.tail)
	require.Equal(t, 3, con.size)
}

func Test_Container_Take(t *testing.T) {

	nil := (*Instruction)(nil)
	a, b, c, _ := setupList()
	con := NewContainer(a)

	z := con.Take()
	fullEqual(t, a, nil, z)
	fullEqual(t, b, c, con.head)
	fullEqual(t, c, nil, con.tail)
	require.Equal(t, 2, con.size)

	z = con.Take()
	fullEqual(t, b, nil, z)
	fullEqual(t, c, nil, con.head)
	fullEqual(t, c, nil, con.tail)
	require.Equal(t, 1, con.size)

	z = con.Take()
	fullEqual(t, c, nil, z)
	require.Equal(t, nil, con.head)
	require.Equal(t, nil, con.tail)
	require.Equal(t, 0, con.size)
}
