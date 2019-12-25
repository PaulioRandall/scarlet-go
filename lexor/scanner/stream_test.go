package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func tokenFinderTest(
	t *testing.T,
	f TokenFinder,
	in string,
	expN int,
	expK token.Kind,
) {

	n, k, e := f([]rune(in))

	require.Nil(t, e)
	assert.Equal(t, expN, n)
	assert.Equal(t, expK, k)
}

func tokenFinderErrTest(
	t *testing.T,
	f TokenFinder,
	in string,
) {

	n, k, e := f([]rune(in))

	require.NotNil(t, e, "Expected error")
	assert.Empty(t, n, "Expected `n` to be 0")
	assert.Empty(t, k, "Expected token.UNDEFINED")
}

func TestStream_SliceBy_1(t *testing.T) {
	// Typicial usage.

	s := stream{
		runes: []rune("Scarlet"),
	}

	exp := token.NewToken(token.ID, "Scar", 0, 0)

	act, e := s.SliceBy(func(r []rune) (int, token.Kind, error) {
		assert.Equal(t, []rune("Scarlet"), r)
		return 4, token.ID, nil
	})

	require.Nil(t, e)
	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("let"), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 4, s.col)
}

func TestStream_SliceBy_2(t *testing.T) {
	// Out of range slice indexes panic.

	s := stream{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.SliceBy(func(r []rune) (int, token.Kind, error) {
			return 99, token.ID, nil
		})
	})

	assert.Panics(t, func() {
		s.SliceBy(func(r []rune) (int, token.Kind, error) {
			return -1, token.ID, nil
		})
	})
}

func TestStream_SliceBy_3(t *testing.T) {
	// Updates line index correctly.

	s := stream{
		runes: []rune("\r\nScarlet"),
	}

	exp := token.NewToken(token.NEWLINE, "\r\n", 0, 0)

	act, e := s.SliceBy(func(r []rune) (int, token.Kind, error) {
		return 2, token.NEWLINE, nil
	})

	require.Nil(t, e)
	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("Scarlet"), s.runes)
	assert.Equal(t, 1, s.line)
	assert.Equal(t, 0, s.col)
}
