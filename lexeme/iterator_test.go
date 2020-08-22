package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func setupItrAt(start string) (_ *Iterator, a, b, c, d *Lexeme) {

	a = lex(0, 0, "1st", BOOL)
	b = lex(0, 4, "2nd", NUMBER)
	c = lex(0, 5, "3rd", STRING)
	d = lex(0, 9, "4th", IDENT)

	a.prev, a.next = nil, b
	b.prev, b.next = a, c
	c.prev, c.next = b, nil

	con := &Container{
		size: 3,
		head: a,
		tail: c,
	}

	itr := &Iterator{
		con: con,
	}

	switch start {
	case "front":
		itr.after = a
	case "mid":
		itr.before = a
		itr.curr = b
		itr.after = c
	case "back":
		itr.before = c
	default:
		panic("Unknown iterator starting location '" + start + "'")
	}

	return itr, a, b, c, d
}

func Test_Iterator_1_1(t *testing.T) {

	itr, a, _, _, _ := setupItrAt("front")

	halfEqual(t, nil, itr.Before())
	halfEqual(t, nil, itr.Curr())
	halfEqual(t, a, itr.After())
}

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
