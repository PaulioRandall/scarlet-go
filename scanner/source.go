package scanner

import (
	"github.com/PaulioRandall/scarlet-go/cookies/where"
	"github.com/PaulioRandall/scarlet-go/token"
)

// source represents the source code yet to be scanned.
type source struct {
	runes []rune // The source code
	line  int    // The line index
	col   int    // The column index of the line
}

// scan removes `n` runes from the unscanned source code and uses them to
// create a new token. The source line and column indexes are updated
// accordingly. If `n` is less than one or greater than the number of remaining
// runes then a panic ensues.
func (s *source) scan(n int, k token.Kind) token.Token {

	if n < 1 {
		panic("Scanning a negative or zero number of runes is not allowed")
	} else if n > len(s.runes) {
		panic("Cannot scan more runes than are available")
	}

	t := token.Token{
		Kind:  k,
		Where: where.New(s.line, s.col, s.col+n),
		Value: string(s.runes[:n]),
	}

	s.runes = s.runes[n:]
	s.col = t.Where.End()

	return t
}

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) ScanToken {
	s := source{
		runes: []rune(src),
	}
	return s.fileScope
}
