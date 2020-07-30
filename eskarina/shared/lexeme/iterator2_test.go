package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Iterator2_1_1(t *testing.T) {

	con, a, _, _, _ := setupContainer()
	it := con.Iterator()

	halfEqual(t, nil, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.After())
}

func Test_Iterator2_2_1(t *testing.T) {

	con, a, b, c, _ := setupContainer()
	it := con.Iterator()

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, nil, it.Before())
	halfEqual(t, a, it.Curr())
	halfEqual(t, b, it.After())

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, a, it.Before())
	halfEqual(t, b, it.Curr())
	halfEqual(t, c, it.After())

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, b, it.Before())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.After())

	require.False(t, it.HasNext())
	require.False(t, it.Next())
	halfEqual(t, c, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.After())

	require.False(t, it.HasNext())
	require.False(t, it.Next())
	halfEqual(t, c, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.After())
}

func Test_Iterator2_2_2(t *testing.T) {

	con, a, b, c, _ := setupContainer()
	it := Iterator2{
		con:    con,
		before: c,
	}

	halfEqual(t, c, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.After())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, b, it.Before())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.After())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, a, it.Before())
	halfEqual(t, b, it.Curr())
	halfEqual(t, c, it.After())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, nil, it.Before())
	halfEqual(t, a, it.Curr())
	halfEqual(t, b, it.After())

	require.False(t, it.HasPrev())
	require.False(t, it.Prev())
	halfEqual(t, nil, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.After())

	require.False(t, it.HasPrev())
	require.False(t, it.Prev())
	halfEqual(t, nil, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.After())
}

func Test_Iterator2_3_1(t *testing.T) {

	con, a, b, c, _ := setupContainer()
	it := Iterator2{
		con:    con,
		before: a,
		curr:   b,
		after:  c,
	}

	z := it.Remove()
	fullEqual(t, b, nil, nil, z)
	halfEqual(t, a, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, c, it.After())

	it.Next()
	halfEqual(t, a, it.Before())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.After())

	z = it.Remove()
	fullEqual(t, c, nil, nil, z)
	halfEqual(t, a, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.After())
}
