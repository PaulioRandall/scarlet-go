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

	return func() (t token.Token, sc lexor.ScanToken, e lexor.ScanErr) {

		fs := []TokenFinder{
			findNewline,     // 1
			findSpace,       // 2
			findComment,     // 3
			findWord,        // 4
			findStrLiteral,  // 5
			findStrTemplate, // 6
			findSymbol,      // 7
		}

		var err error

		for _, f := range fs {
			t, err = s.SliceBy(f)

			if err != nil {
				e = lexor.NewScanErr("Scanning error", err, s.line, s.col)
				return
			}

			if t != (token.Token{}) {
				sc = scan(s)
				return
			}
		}

		e = lexor.NewScanErr("Unknown token", nil, s.line, s.col)
		return
	}
}
