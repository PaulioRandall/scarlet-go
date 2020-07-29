package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Iterator_1_1(t *testing.T) {

	a, b, c, _ := setup()
	feign(a, b, c)
	it := NewIterator(a)

	halfEqual(t, nil, it.Before())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.After())
}

func Test_Iterator_2_1(t *testing.T) {

	a, b, c, _ := setup()
	feign(a, b, c)
	it := NewIterator(a)

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

func Test_Iterator_2_2(t *testing.T) {

	a, b, c, _ := setup()
	feign(a, b, c)
	it := Iterator{
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

func Test_Iterator_3_1(t *testing.T) {

	a, b, c, _ := setup()
	feign(a, b, c)
	it := Iterator{
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
