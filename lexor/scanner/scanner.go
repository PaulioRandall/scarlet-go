// scanner scans source code for tokens. The API is a `lexor.ScanToken`
// function returned by the public New function.
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

	if s.isEmpty() {
		return nil
	}

	return func() (t token.Token, sc lexor.ScanToken, e lexor.ScanErr) {

		fs := []tokenFinder{
			findSpace,       // 1
			findComment,     // 2
			findWord,        // 3
			findStrLiteral,  // 4
			findStrTemplate, // 5
			findSymbol,      // 6
		}

		var err error

		for _, f := range fs {
			t, err = s.sliceBy(f)

			if err != nil {
				e = lexor.NewScanErr("Scanning error", err, s.line, s.col)
				return
			}

			if t.IsNotZero() {
				sc = scan(s)
				return
			}
		}

		e = lexor.NewScanErr("Unknown token", nil, s.line, s.col)
		return
	}
}
