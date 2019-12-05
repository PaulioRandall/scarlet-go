package tokeniser

import (
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/tokeniser/source"
)

// newlineThunk returns a TokenThunk that returns a newline token along with
// the next emitter to use `after`.
func newlineThunk(s *source.Source, after TokenThunk) TokenThunk {
	return func() (token.Token, TokenThunk, perror.Perror) {
		n, k := s.Identify(lenOfNewline)

		if n == 0 {
			return token.Empty(), nil, perror.Newish(
				"Expected newline characters, i.e. LF or CRLF",
				s.Where(),
			)
		}

		t := s.SliceNewline(n, k)
		return t, after, nil
	}
}

// lenOfNewline is a TokenFinder function that identifies the kind and number of
// runes in the next newline token.
func lenOfNewline(r []rune) (n int, k token.Kind) {
	k = token.NEWLINE
	max := len(r)

	switch {
	case max == 0:
	case r[0] == '\n':
		n = 1
	case max > 0 && r[0] == '\r' && r[1] == '\n':
		n = 2
	}

	return
}
