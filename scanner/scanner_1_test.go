package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/where"
	"github.com/stretchr/testify/assert"
)

func TestScanner_scanToken_1(t *testing.T) {
	s := scanner{
		runes: []rune("Scarlet"),
	}

	exp := token.New(
		"Scar",
		token.UNDEFINED,
		where.New(0, 0, 4),
	)

	act := s.scanToken(4, token.UNDEFINED)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("let"), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 4, s.col)
}

func TestScanner_scanToken_2(t *testing.T) {
	s := scanner{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.scanToken(0, token.UNDEFINED)
	})
}

func TestScanner_scanToken_3(t *testing.T) {
	s := scanner{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.scanToken(99, token.UNDEFINED)
	})

	assert.Panics(t, func() {
		s.scanToken(0, token.UNDEFINED)
	})
}

func TestScanner_scanNewlineToken_1(t *testing.T) {
	s := scanner{
		runes: []rune("\nScarlet"),
	}

	exp := token.New(
		"\n",
		token.NEWLINE,
		where.New(0, 0, 1),
	)

	act := s.scanNewlineToken()

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("Scarlet"), s.runes)
	assert.Equal(t, 1, s.line)
	assert.Equal(t, 0, s.col)
}

func TestScanner_scanNewlineToken_2(t *testing.T) {
	s := scanner{
		runes: []rune("\r\nScarlet"),
	}

	exp := token.New(
		"\r\n",
		token.NEWLINE,
		where.New(0, 0, 2),
	)

	act := s.scanNewlineToken()

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune("Scarlet"), s.runes)
	assert.Equal(t, 1, s.line)
	assert.Equal(t, 0, s.col)
}

func TestScanner_scanNewlineToken_3(t *testing.T) {
	s := scanner{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.scanNewlineToken()
	})
}

func TestScanner_scanWordToken_1(t *testing.T) {
	s := scanner{
		runes: []rune("END"),
	}

	exp := token.New(
		"END",
		token.END,
		where.New(0, 0, 3),
	)

	act := s.scanWordToken(3)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune(""), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 3, s.col)
}

func TestScanner_scanWordToken_2(t *testing.T) {
	s := scanner{
		runes: []rune("FUNC END"),
	}

	exp := token.New(
		"FUNC",
		token.FUNC,
		where.New(0, 0, 4),
	)

	act := s.scanWordToken(4)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune(" END"), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 4, s.col)
}

func TestScanner_scanWordToken_3(t *testing.T) {
	s := scanner{
		runes: []rune("Anything"),
	}

	exp := token.New(
		"Anything",
		token.UNDEFINED,
		where.New(0, 0, 8),
	)

	act := s.scanWordToken(8)

	assert.Equal(t, exp, act)
	assert.Equal(t, []rune(""), s.runes)
	assert.Equal(t, 0, s.line)
	assert.Equal(t, 8, s.col)
}

func TestScanner_scanWordToken_4(t *testing.T) {
	s := scanner{
		runes: []rune("Scarlet"),
	}

	assert.Panics(t, func() {
		s.scanWordToken(99)
	})
}
