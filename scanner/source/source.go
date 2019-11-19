package source

import (
	"github.com/PaulioRandall/scarlet-go/cookies"
	"github.com/PaulioRandall/scarlet-go/cookies/where"
	"github.com/PaulioRandall/scarlet-go/token"
)

// source represents the source code yet to be scanned.
type source struct {
	runes []rune // The source code
	line  int    // The line index
	col   int    // The column index of the line
}

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) token.ScanToken {
	s := source{
		runes: []rune(src),
	}
	return s.scan
}

// checkSize validates that `n` is greater than zero and less than the number of
// remaining runes. If either case is false then a panic ensues.
func (s *source) checkSize(n int) {
	if n < 1 {
		panic("Scanning a negative or zero number of runes is not allowed")
	} else if n > len(s.runes) {
		panic("Cannot scan more runes than are available")
	}
}

// scan removes `n` runes from the unscanned source code and uses them to
// create a new token. The source line and column indexes are updated
// accordingly. If `n` is less than one or greater than the number of remaining
// runes then a panic ensues.
func (s *source) scanRunes(n int, k token.Kind) token.Token {

	s.checkSize(n)

	t := token.New(
		string(s.runes[:n]),
		k,
		where.New(s.line, s.col, s.col+n),
	)

	s.runes = s.runes[n:]
	s.col = t.Where().End()

	return t
}

// scanNewline removes the next linefeed (or CRLF) runes from the unscanned
// source code and uses them to create a newline token. The source line and
// column indexes are updated accordingly. If the next sequence of runes do not
// form a newline token then a panic ensues.
func (s *source) scanNewline() token.Token {

	n := cookies.NewlineRunes(s.runes, 0)
	if n == 0 {
		panic("Expected characters representing a newline, LF or CRLF")
	}

	t := token.New(
		string(s.runes[:n]),
		token.NEWLINE,
		where.New(s.line, s.col, s.col+n),
	)

	s.runes = s.runes[n:]
	s.line++
	s.col = 0

	return t
}

// scanWord removes `n` runes from the unscanned source code and uses them to
// create a new word token. The kind is identified from the resultant word
// string. The source line and column indexes are updated accordingly. If `n` is
// less than one or greater than the number of remaining runes then a panic
// ensues.
func (s *source) scanWord(n int) token.Token {

	s.checkSize(n)

	v := string(s.runes[:n])
	t := token.New(
		v,
		token.FindWordKind(v),
		where.New(s.line, s.col, s.col+n),
	)

	s.runes = s.runes[n:]
	s.col = t.Where().End()

	return t
}
