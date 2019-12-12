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
			findNewline,    // 1
			findSpace,      // 2
			findComment,    // 3
			findKeyword,    // 4
			findId,         // 5
			findStrLiteral, // 6
			findSymbol,     // 7
		}

		var err error

		for _, f := range fs {
			t, err = s.SliceBy(f)

			if err != nil {
				e = lexor.WrapScanErr("Scanning error", s.Where(), err)
				return
			}

			if t != nil {
				sc = scan(s)
				return
			}
		}

		e = lexor.NewScanErr_2("Unknown token", s.Where())
		return
	}
}
