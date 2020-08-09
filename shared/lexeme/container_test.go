package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func setupContainer() (_ *Container, a, b, c, d *Lexeme) {

	a = lex(0, 0, "1st", BOOL)
	b = lex(0, 4, "2nd", NUMBER)
	c = lex(0, 5, "3rd", STRING)
	d = lex(0, 9, "4th", IDENTIFIER)

	a.prev, a.next = nil, b
	b.prev, b.next = a, c
	c.prev, c.next = b, nil

	con := &Container{
		size: 3,
		head: a,
		tail: c,
	}

	return con, a, b, c, d
}

func Test_Container_1_1(t *testing.T) {

	con, a, b, c, _ := setupContainer()

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

	a := lex(0, 0, "1st", BOOL)
	b := lex(0, 4, "2nd", NUMBER)
	c := lex(0, 5, "3rd", STRING)
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
