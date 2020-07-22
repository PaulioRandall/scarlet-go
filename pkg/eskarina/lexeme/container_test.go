package lexeme

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"

	"github.com/stretchr/testify/require"
)

func setup() (a, b, c, d *Lexeme) {
	a = tok("true", prop.PR_BOOL)
	b = tok("1", prop.PR_NUMBER)
	c = tok(`"abc"`, prop.PR_STRING)
	d = tok("x", prop.PR_IDENTIFIER)
	return
}

func setupList() (a, b, c, d *Lexeme) {
	a, b, c, d = setup()
	_ = feign(a, b, c)
	return
}

func Test_NewContainer(t *testing.T) {

	a, b, c, _ := setupList()
	con := NewContainer(a)

	require.Equal(t, 3, con.size)
	fullEqual(t, a, nil, b, con.first)
	fullEqual(t, b, a, c, con.first.Next)
	fullEqual(t, b, a, c, con.last.Prev)
	fullEqual(t, c, b, nil, con.last)
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
	fullEqual(t, c, nil, nil, con.first)
	fullEqual(t, c, nil, nil, con.last)
	require.Equal(t, 1, con.size)

	con.Prepend(b)
	fullEqual(t, b, nil, c, con.first)
	fullEqual(t, c, b, nil, con.last)
	require.Equal(t, 2, con.size)

	con.Prepend(a)
	fullEqual(t, a, nil, b, con.first)
	fullEqual(t, b, a, c, con.first.Next)
	fullEqual(t, b, a, c, con.last.Prev)
	fullEqual(t, c, b, nil, con.last)
	require.Equal(t, 3, con.size)
}

func Test_Container_Append(t *testing.T) {

	a, b, c, _ := setup()
	con := Container{}

	con.Append(a)
	fullEqual(t, a, nil, nil, con.first)
	fullEqual(t, a, nil, nil, con.last)
	require.Equal(t, 1, con.size)

	con.Append(b)
	fullEqual(t, a, nil, b, con.first)
	fullEqual(t, b, a, nil, con.last)
	require.Equal(t, 2, con.size)

	con.Append(c)
	fullEqual(t, a, nil, b, con.first)
	fullEqual(t, b, a, c, con.first.Next)
	fullEqual(t, b, a, c, con.last.Prev)
	fullEqual(t, c, b, nil, con.last)
	require.Equal(t, 3, con.size)
}

func Test_Container_InsertBefore(t *testing.T) {

	a, b, c, d := setup()
	con := NewContainer(c)

	con.InsertBefore(0, a)
	fullEqual(t, a, nil, c, con.first)
	fullEqual(t, c, a, nil, con.last)
	require.Equal(t, 2, con.size)

	con.InsertBefore(1, b)
	fullEqual(t, a, nil, b, con.first)
	fullEqual(t, b, a, c, con.first.Next)
	fullEqual(t, b, a, c, con.last.Prev)
	fullEqual(t, c, b, nil, con.last)
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
	fullEqual(t, a, nil, b, con.first)
	fullEqual(t, b, a, nil, con.last)
	require.Equal(t, 2, con.size)

	con.InsertAfter(1, c)
	fullEqual(t, a, nil, b, con.first)
	fullEqual(t, b, a, c, con.first.Next)
	fullEqual(t, b, a, c, con.last.Prev)
	fullEqual(t, c, b, nil, con.last)
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
	fullEqual(t, a, nil, c, con.first)
	fullEqual(t, b, nil, nil, z)
	fullEqual(t, c, a, nil, con.last)
	require.Equal(t, 2, con.size)

	z = con.Remove(0)
	fullEqual(t, c, nil, nil, con.first)
	fullEqual(t, c, nil, nil, con.last)
	fullEqual(t, a, nil, nil, z)
	require.Equal(t, 1, con.size)

	z = con.Remove(0)
	halfEqual(t, nil, con.first)
	halfEqual(t, nil, con.last)
	fullEqual(t, c, nil, nil, z)
	require.Equal(t, 0, con.size)
}
