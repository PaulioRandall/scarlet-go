package lexeme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Itinerant2_1_1(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	it := newItinerant(a)

	halfEqual(t, nil, it.Prior())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.Ahead())
}

func Test_Itinerant2_2_1(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	it := newItinerant(a)

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, nil, it.Prior())
	halfEqual(t, a, it.Curr())
	halfEqual(t, b, it.Ahead())

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, a, it.Prior())
	halfEqual(t, b, it.Curr())
	halfEqual(t, c, it.Ahead())

	require.True(t, it.HasNext())
	require.True(t, it.Next())
	halfEqual(t, b, it.Prior())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.Ahead())

	require.False(t, it.HasNext())
	require.False(t, it.Next())
	halfEqual(t, c, it.Prior())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.Ahead())

	require.False(t, it.HasNext())
	require.False(t, it.Next())
	halfEqual(t, c, it.Prior())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.Ahead())
}

func Test_Itinerant2_2_2(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	it := Itinerant2{
		prior: c,
	}

	halfEqual(t, c, it.Prior())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.Ahead())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, b, it.Prior())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.Ahead())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, a, it.Prior())
	halfEqual(t, b, it.Curr())
	halfEqual(t, c, it.Ahead())

	require.True(t, it.HasPrev())
	require.True(t, it.Prev())
	halfEqual(t, nil, it.Prior())
	halfEqual(t, a, it.Curr())
	halfEqual(t, b, it.Ahead())

	require.False(t, it.HasPrev())
	require.False(t, it.Prev())
	halfEqual(t, nil, it.Prior())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.Ahead())

	require.False(t, it.HasPrev())
	require.False(t, it.Prev())
	halfEqual(t, nil, it.Prior())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, a, it.Ahead())
}

func Test_Itinerant2_3_1(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	it := Itinerant2{
		prior: a,
		curr:  b,
		ahead: c,
	}

	z := it.Remove()
	fullEqual2(t, b, nil, nil, z)
	halfEqual(t, a, it.Prior())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, c, it.Ahead())

	it.Next()
	halfEqual(t, a, it.Prior())
	halfEqual(t, c, it.Curr())
	halfEqual(t, nil, it.Ahead())

	z = it.Remove()
	fullEqual2(t, c, nil, nil, z)
	halfEqual(t, a, it.Prior())
	halfEqual(t, nil, it.Curr())
	halfEqual(t, nil, it.Ahead())
}
