package aggregator

import (
	"github.com/PaulioRandall/scarlet-go/stat"
	"github.com/PaulioRandall/scarlet-go/token"
)

// SequenceFinder is a function prototype that identifies a sequence of tokens
// returning the number of tokens and if they represent a statement, the
// statement kind.
type SequenceFinder func([]token.Token) (int, stat.Kind, token.Perror)

// stream represents the tokens scanned from source and provides functionality
// to remove and return slices of them.
type stream struct {
	t []token.Token
}

// IsEmpty returns true if there is no more tokens to parse.
func (s *stream) IsEmpty() bool {
	return len(s.t) == 0
}

// Identify accepts an SequenceFinder function and returns the kind and length
// of the next statement in it.
func (s *stream) Identify(f SequenceFinder) (int, stat.Kind, token.Perror) {
	return f(s.t)
}

// SliceBy accepts a SequenceFinder function and slices off tokens based on the
// result.
func (s *stream) SliceBy(f SequenceFinder) (
	t []token.Token,
	k stat.Kind,
	e token.Perror) {

	n, k, e := s.Identify(f)

	if e != nil || k == stat.UNDEFINED {
		return
	}

	s.checkSize(n)
	t = s.slice(n)
	return
}

// checkSize validates that `n` is greater than zero and less than the number of
// remaining tokens. If either case is false then a panic ensues.
func (s *stream) checkSize(n int) {
	if n < 1 {
		panic("Reading a zero or a negative number of tokens is not allowed")
	} else if n > len(s.t) {
		panic("Cannot read more tokens than are available")
	}
}

// slice slices `n` tokens from the front of the token array. The index is up
// dated accordingly.
func (s *stream) slice(n int) []token.Token {
	t := s.t[:n]
	s.t = s.t[n:]
	return t
}
