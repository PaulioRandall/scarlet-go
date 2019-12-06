package scanner

import (
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/source"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New returns a ScanToken thunk that will return the first token in the input
// source.
func New(src string) token.ScanToken {
	s := source.New(src)
	return scan(s)
}

// scan returns a ScanToken thunk that returns the next token in the source.
func scan(s *source.Source) token.ScanToken {

	if s.IsEmpty() {
		return nil
	}

	return func() (token.Token, token.ScanToken, perror.Perror) {

		fs := []source.TokenFinder{
			findNewline, // 1
			findSpace,   // 2
			findKeyword, // 3
			findId,      // 4
			findSymbol,  // 5
		}

		for _, f := range fs {
			if t := s.SliceBy(f); t != token.Empty() {
				return t, scan(s), nil
			}
		}

		return token.Empty(), nil, perror.Newish(
			"Unknown token",
			s.Where(),
		)
	}
}
