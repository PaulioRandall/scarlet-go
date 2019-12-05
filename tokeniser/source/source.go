// source provides a structure to store the state of source code parsing.
package source

import (
	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/where"
)

// Source represents some source code and provides some functionality to
// convert slices of it into tokens.
type Source struct {
	runes []rune
	line  int
	col   int
}

// New creates a new source instance.
func New(src string) *Source {
	return &Source{
		runes: []rune(src),
	}
}

// Where returns the current location in the source code.
func (s *Source) Where() where.Where {
	return where.New(s.line, s.col, s.col)
}

// TokenFinder represents a function that identifies the kind of the next token
// and counts the number of runes in it.
type TokenFinder func([]rune) (int, token.Kind)

// SliceBy accepts a TokenFinder function and slices off a token based on the
// result.
func (s *Source) SliceBy(f TokenFinder) token.Token {
	n, k := s.Identify(f)
	return s.Slice(n, k)
}

// Identify accepts a TokenFinder function and returns the kind and length of
// the next token from it.
func (s *Source) Identify(f TokenFinder) (int, token.Kind) {
	return f(s.runes)
}

// SliceNewline performs the same action the Slice function but increments the
// line number and resets the coloumn index afterwards.
func (s *Source) SliceNewline(n int, k token.Kind) token.Token {

	t := s.Slice(n, k)
	s.line++
	s.col = 0

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
	return token.Newish(str, k, w)
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
