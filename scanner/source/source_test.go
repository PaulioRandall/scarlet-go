package source

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/cookies/where"
	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/stretchr/testify/assert"
)

func TestScanRunes_1(t *testing.T) {
	s := source{
		runes: []rune("Scarlet"),
	}

	exp := token.New(
		"Scar",
		token.UNDEFINED,
		where.New(0, 0, 4),
	)

	act := s.scanRunes(4, token.UNDEFINED)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("let"), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 4, s.col)
}

func TestScanRunes_2(t *testing.T) {
	s := source{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.scanRunes(0, token.UNDEFINED)
	})
}

func TestScanRunes_3(t *testing.T) {
	s := source{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.scanRunes(99, token.UNDEFINED)
	})

	assert.Panics(t, func() {
		s.scanRunes(0, token.UNDEFINED)
	})
}

func TestScanNewline_1(t *testing.T) {
	s := source{
		runes: []rune("\nScarlet"),
	}

	exp := token.New(
		"\n",
		token.NEWLINE,
		where.New(0, 0, 1),
	)

	act := s.scanNewline()

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("Scarlet"), s.runes)
	assert.Equal(t, 1, s.line)
	assert.Equal(t, 0, s.col)
}

func TestScanNewline_2(t *testing.T) {
	s := source{
		runes: []rune("\r\nScarlet"),
	}

	exp := token.New(
		"\r\n",
		token.NEWLINE,
		where.New(0, 0, 2),
	)

	act := s.scanNewline()

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("Scarlet"), s.runes)
	assert.Equal(t, 1, s.line)
	assert.Equal(t, 0, s.col)
}

func TestScanNewline_3(t *testing.T) {
	s := source{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.scanNewline()
	})
}

func TestScanWord_1(t *testing.T) {
	s := source{
		runes: []rune("END"),
	}

	exp := token.New(
		"END",
		token.END,
		where.New(0, 0, 3),
	)

	act := s.scanWord(3)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune(""), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 3, s.col)
}

func TestScanWord_2(t *testing.T) {
	s := source{
		runes: []rune("PROCEDURE END"),
	}

	exp := token.New(
		"PROCEDURE",
		token.PROCEDURE,
		where.New(0, 0, 9),
	)

	act := s.scanWord(9)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune(" END"), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 9, s.col)
}

func TestScanWord_3(t *testing.T) {
	s := source{
		runes: []rune("Anything"),
	}

	exp := token.New(
		"Anything",
		token.UNDEFINED,
		where.New(0, 0, 8),
	)

	act := s.scanWord(8)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune(""), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 8, s.col)
}

func TestScanWord_4(t *testing.T) {
	s := source{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.scanWord(99)
	})
}
