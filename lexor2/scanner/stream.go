package scanner

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token2"
)

// TokenFinder is a function prototype that identifies the kind of the next
// token and counts the number of runes in it.
type TokenFinder func([]rune) (int, token.Kind, error)

// stream represents some source code and provides functionality to remove
// slices of it and return them as tokens.
type stream struct {
	runes []rune
	line  int
	col   int
}

// IsEmpty returns true if there are no more terminal characters to parse.
func (s *stream) IsEmpty() bool {
	return len(s.runes) == 0
}

// Identify accepts a TokenFinder function and returns the kind and length of
// the next token from it.
func (s *stream) Identify(f TokenFinder) (int, token.Kind, error) {
	return f(s.runes)
}

// SliceBy accepts a TokenFinder function and slices off a token based on the
// result.
func (s *stream) SliceBy(f TokenFinder) (_ token.Token, e error) {
	n, k, e := s.Identify(f)

	if e != nil || k == token.UNDEFINED {
		return
	}

	s.checkSize(n)
	return s.tokenise(n, k), nil
}

// checkSize validates that `n` is greater than zero and less than the number of
// remaining runes. If either case is false then a panic ensues.
func (s *stream) checkSize(n int) {
	if n < 1 {
		panic("Reading a zero or a negative number of runes is not allowed")
	} else if n > len(s.runes) {
		panic("Cannot read more runes than are available")
	}
}

// tokenise slices `n` runes from the front of the source code and uses them to
// construct a new token. If token value ends in a linefeed the internal line
// count is incremented and internal column index reset to zero else the column
// index is increased by the number of runes sliced off. If `n` is less than
// one or greater than the number of remaining runes then a panic ensues.
func (s *stream) tokenise(n int, k token.Kind) token.Token {

	str, line, col := s.slice(n)

	if strings.HasSuffix(str, "\n") {
		s.line++
		s.col = 0
	}

	return token.NewToken(k, str, line, col)
}

// slice slices `n` runes from the front of the source code and updates the
// column field accordingly. Linefeeds must be sliced using the sliceNewline
// function.
func (s *stream) slice(n int) (str string, line int, col int) {

	str = string(s.runes[:n])
	line = s.line
	col = s.col

	s.runes = s.runes[n:]
	s.col = s.col + n

	return
}
