package source

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/where"
	"github.com/stretchr/testify/assert"
)

func TestSource_Slice_1(t *testing.T) {
	s := Source{
		runes: []rune("Scarlet"),
	}

	exp := token.New(
		"Scar",
		token.UNDEFINED,
		where.New(0, 0, 4),
	)

	act := s.Slice(4, token.UNDEFINED)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("let"), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 4, s.col)
}

func TestSource_Slice_2(t *testing.T) {
	s := Source{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.Slice(0, token.UNDEFINED)
	})
}

func TestSource_Slice_3(t *testing.T) {
	s := Source{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.Slice(99, token.UNDEFINED)
	})

	assert.Panics(t, func() {
		s.Slice(0, token.UNDEFINED)
	})
}

func TestSource_SliceNewline_1(t *testing.T) {
	s := Source{
		runes: []rune("\nScarlet"),
	}

	exp := token.New(
		"\n",
		token.NEWLINE,
		where.New(0, 0, 1),
	)

	act := s.SliceNewline(1, token.NEWLINE)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("Scarlet"), s.runes)
	assert.Equal(t, 1, s.line)
	assert.Equal(t, 0, s.col)
}

func TestSource_SliceNewline_2(t *testing.T) {
	s := Source{
		runes: []rune("\r\nScarlet"),
	}

	exp := token.New(
		"\r\n",
		token.NEWLINE,
		where.New(0, 0, 2),
	)

	act := s.SliceNewline(2, token.NEWLINE)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("Scarlet"), s.runes)
	assert.Equal(t, 1, s.line)
	assert.Equal(t, 0, s.col)
}
