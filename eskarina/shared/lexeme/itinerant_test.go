package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Itinerant2_1_1(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	it := NewItinerant(a)

	halfEqual(t, nil, it.Behind())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.Ahead())
}

func Test_Itinerant2_2_1(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	it := NewItinerant(a)

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, nil, it.Behind())
	halfEqual(t, a, it.Curr())
	halfEqual(t, b, it.Ahead())

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, a, it.Behind())
	halfEqual(t, b, it.Curr())
	halfEqual(t, c, it.Ahead())

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, b, it.Behind())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.Ahead())

	require.False(t, it.HasNext())
	require.False(t, it.Next())
	halfEqual(t, c, it.Behind())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.Ahead())

	require.False(t, it.HasNext())
	require.False(t, it.Next())
	halfEqual(t, c, it.Behind())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.Ahead())
}

func Test_Itinerant2_2_2(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	it := Itinerant2{
		behind: c,
	}

	halfEqual(t, c, it.Behind())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.Ahead())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, b, it.Behind())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.Ahead())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, a, it.Behind())
	halfEqual(t, b, it.Curr())
	halfEqual(t, c, it.Ahead())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, nil, it.Behind())
	halfEqual(t, a, it.Curr())
	halfEqual(t, b, it.Ahead())

	require.False(t, it.HasPrev())
	require.False(t, it.Prev())
	halfEqual(t, nil, it.Behind())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.Ahead())

	require.False(t, it.HasPrev())
	require.False(t, it.Prev())
	halfEqual(t, nil, it.Behind())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.Ahead())
}

func Test_Itinerant2_3_1(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	it := Itinerant2{
		behind: a,
		curr:   b,
		ahead:  c,
	}

	z := it.Remove()
	fullEqual2(t, b, nil, nil, z)
	halfEqual(t, a, it.Behind())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, c, it.Ahead())

	it.Next()
	halfEqual(t, a, it.Behind())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.Ahead())

	z = it.Remove()
	fullEqual2(t, c, nil, nil, z)
	halfEqual(t, a, it.Behind())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.Ahead())
}
