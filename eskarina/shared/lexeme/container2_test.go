package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Container2_Pop(t *testing.T) {

	con, a, b, c, _ := setupContainer2()

	z := con.pop(false)
	fullEqual2(t, a, nil, nil, z)
	fullEqual2(t, b, nil, c, con.head)
	fullEqual2(t, c, b, nil, con.tail)
	require.Equal(t, 2, con.size)

	z = con.pop(true)
	fullEqual2(t, b, nil, nil, con.head)
	fullEqual2(t, b, nil, nil, con.tail)
	fullEqual2(t, c, nil, nil, z)
	require.Equal(t, 1, con.size)

	z = con.pop(false)
	halfEqual(t, nil, con.head)
	halfEqual(t, nil, con.tail)
	fullEqual2(t, b, nil, nil, z)
	require.Equal(t, 0, con.size)

	z = con.pop(true)
	require.Nil(t, z)
}

func Test_Container2_push(t *testing.T) {

	a, b, c, _ := setup2()
	con := &Container2{}

	con.push(b, true)
	fullEqual2(t, b, nil, nil, con.head)
	fullEqual2(t, b, nil, nil, con.tail)
	require.Equal(t, 1, con.size)

	con.push(a, false)
	fullEqual2(t, a, nil, b, con.head)
	fullEqual2(t, b, a, nil, con.tail)
	require.Equal(t, 2, con.size)

	con.push(c, true)
	fullEqual2(t, a, nil, b, con.head)
	fullEqual2(t, b, a, c, con.head.next)
	fullEqual2(t, b, a, c, con.tail.prev)
	fullEqual2(t, c, b, nil, con.tail)
	require.Equal(t, 3, con.size)
}
