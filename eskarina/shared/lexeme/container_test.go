package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Container_1_1(t *testing.T) {

	a, b, c, _ := setup()
	feign(a, b, c)
	con := NewContainer(a)

	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, c, con.head.next)
	fullEqual(t, b, a, c, con.tail.prev)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 3, con.size)
}

func Test_Container_2_1(t *testing.T) {

	con, a, b, c, _ := setupContainer()

	z := con.pop(false)
	fullEqual(t, a, nil, nil, z)
	fullEqual(t, b, nil, c, con.head)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 2, con.size)

	z = con.pop(true)
	fullEqual(t, b, nil, nil, con.head)
	fullEqual(t, b, nil, nil, con.tail)
	fullEqual(t, c, nil, nil, z)
	require.Equal(t, 1, con.size)

	z = con.pop(false)
	halfEqual(t, nil, con.head)
	halfEqual(t, nil, con.tail)
	fullEqual(t, b, nil, nil, z)
	require.Equal(t, 0, con.size)

	z = con.pop(true)
	require.Nil(t, z)
}

func Test_Container_3_1(t *testing.T) {

	a, b, c, _ := setup()
	con := &Container{}

	con.push(b, true)
	fullEqual(t, b, nil, nil, con.head)
	fullEqual(t, b, nil, nil, con.tail)
	require.Equal(t, 1, con.size)

	con.push(a, false)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, nil, con.tail)
	require.Equal(t, 2, con.size)

	con.push(c, true)
	fullEqual(t, a, nil, b, con.head)
	fullEqual(t, b, a, c, con.head.next)
	fullEqual(t, b, a, c, con.tail.prev)
	fullEqual(t, c, b, nil, con.tail)
	require.Equal(t, 3, con.size)
}
