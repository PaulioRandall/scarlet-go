package scanner

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New returns a ScanToken thunk that will return the first token in the input
// lexor.
func New(src string) lexor.ScanToken {
	s := &stream{
		runes: []rune(src),
	}
	return scan(s)
}

// scan returns a ScanToken thunk that returns the next token in the source.
func scan(s *stream) lexor.ScanToken {

	if s.IsEmpty() {
		return nil
	}

	return func() (t token.Token, sc lexor.ScanToken, e token.Perror) {

		fs := []TokenFinder{
			findNewline,    // 1
			findSpace,      // 2
			findKeyword,    // 3
			findId,         // 4
			findStrLiteral, // 5
			findSymbol,     // 6
		}

		var err error

		for _, f := range fs {
			t, err = s.SliceBy(f)

			if err != nil {
				e = token.WrapPerror("Scanning error", s.Where(), err)
				return
			}

			if t != nil {
				sc = scan(s)
				return
			}
		}

		e = token.PerrorBySnippet("Unknown token", s.Where())
		return
	}
}
