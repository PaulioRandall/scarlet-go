package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/cookies"
	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/cookies/where"
	"github.com/PaulioRandall/scarlet-go/token"
)

// scanner holds the yet to be scanned source code and where that code begins
// within its source file.
type scanner struct {
	runes []rune // The source code
	line  int    // The line index
	col   int    // The column index of the line
}

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) token.ScanToken {
	s := scanner{
		runes: []rune(src),
	}
	return s.scan
}

// scan identifies and returns the next token in the source. The token must
// be one that appears at the start of a statement within the top level of a
// source file.
func (s *scanner) scan() (token.Token, token.ScanToken, perror.Perror) {

	var k token.Kind
	var n int

	if len(s.runes) == 0 {
		return token.Empty(), nil, nil
	}

	switch ru := s.runes[0]; {
	case cookies.NewlineRunes(s.runes, 0) != 0:
		return s.scanNewline(), s.scan, nil

	case unicode.IsSpace(ru):
		k = token.WHITESPACE
		n = countSpaces(s.runes)
		return s.scanRunes(n, k), s.scan, nil

	case unicode.IsLetter(ru):
		n = countWord(s.runes)
		return s.scanWord(n), s.scan, nil
	}

	return token.Empty(), nil, perror.New(
		"Unknown token",
		s.line,
		s.col,
		s.col,
	)
}

// checkSize validates that `n` is greater than zero and less than the number of
// remaining runes. If either case is false then a panic ensues.
func (s *scanner) checkSize(n int) {
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
func (s *scanner) scanRunes(n int, k token.Kind) token.Token {

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
func (s *scanner) scanNewline() token.Token {

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
func (s *scanner) scanWord(n int) token.Token {

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
