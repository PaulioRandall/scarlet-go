package source

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSource_SliceBy_1(t *testing.T) {
	// Typicial usage.

	s := Source{
		runes: []rune("Scarlet"),
	}

	exp := token.New(
		"Scar",
		token.ID,
		0, 0, 4,
	)

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

func TestSource_SliceBy_2(t *testing.T) {
	// Out of range slice indexes panic.

	s := Source{
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

func TestSource_SliceBy_3(t *testing.T) {
	// Updates line index correctly.

	s := Source{
		runes: []rune("\r\nScarlet"),
	}

	exp := token.New(
		"\r\n",
		token.NEWLINE,
		0, 0, 2,
	)

	act, e := s.SliceBy(func(r []rune) (int, token.Kind, error) {
		return 2, token.NEWLINE, nil
	})

	require.Nil(t, e)
	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("Scarlet"), s.runes)
	assert.Equal(t, 1, s.line)
	assert.Equal(t, 0, s.col)
}
