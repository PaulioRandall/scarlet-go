package source

import (
	"github.com/PaulioRandall/scarlet-go/cookies"
	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/where"
)

// Source represents some source code and provides some functionality to
// convert it into tokens.
type Source struct {
	runes []rune
	line  int
	col   int
}

// sliceNewline slices the next newline (LF or CRLF) from the front of the
// source code and uses them to construct a newline token; the source line and
// column indexes are updated accordingly. If the next sequence of runes do not
// form a newline token then a panic ensues.
func (s *Source) SliceNewline() token.Token {

	n := cookies.NewlineRunes(s.runes, 0)
	if n == 0 {
		panic("Expected characters representing a newline, LF or CRLF")
	}

	t := s.Slice(n, token.NEWLINE)
	s.line++

	return t
}

// Slice slices `n` runes from the front of the source code and uses them to
// construct a new token; the source line and column indexes are updated
// accordingly. If `n` is less than one or greater than the number of remaining
// runes then a panic ensues.
func (s *Source) Slice(n int, k token.Kind) token.Token {

	s.checkSize(n)
	str, start, end := s.slice(n)

	w := where.New(s.line, start, end)
	return token.New(str, k, w)
}

// checkSize validates that `n` is greater than zero and less than the number of
// remaining runes. If either case is false then a panic ensues.
func (s *Source) checkSize(n int) {
	if n < 1 {
		panic("Reading a zero or a negative number of runes is not allowed")
	} else if n > len(s.runes) {
		panic("Cannot read more runes than are available")
	}
}

// slice slices `n` runes from the front of the source code and updates the
// column field accordingly. Linefeeds must be sliced using the sliceNewline
// function.
func (s *Source) slice(n int) (str string, start int, end int) {

	str = string(s.runes[:n])
	start = s.col
	end = s.col + n

	s.runes = s.runes[n:]
	s.col = end

	return
}
