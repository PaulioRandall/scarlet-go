package lexor

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// empty returns true if the scanners rune slice is empty.
func (s *scanner) empty() bool {
	return len(s.runes) == 0
}

// len returns the length of the scanners rune slice.
func (s *scanner) len() int {
	return len(s.runes)
}

// matches returns true if the scanners rune slice matches the specified
// sequence of runes.
func (s *scanner) matchesTerminal(i int, terminal rune) bool {

	if i >= len(s.runes) {
		panic("SANITY CHECK! Bad argument, start is bigger than the remaining source code")
	}

	return s.runes[i] == terminal
}

func (s *scanner) matchesNonTerminal(start int, needle string) bool {

	haystack := s.runes[start:]

	if len(needle) > len(haystack) {
		panic("SANITY CHECK! Bad argument, the `needle` is bigger than the `haystack`")
	}

	for i, ru := range needle {
		if haystack[i] != ru {
			return false
		}
	}

	return true
}

func (s *scanner) doesNotMatchNonTerminal(start int, needle string) bool {
	return !s.matchesNonTerminal(start, needle)
}

// matchesNewline returns true if the scanners rune slice begins with a sequence
// of newline terminals.
func (s *scanner) matchesNewline(start int) bool {
	return s.countNewlineRunes(start) > 0
}

// noMatchNewline returns false if the scanners rune slice begins with a
// sequence of newline terminals.
func (s *scanner) doesNotMatchNewline(start int) bool {
	return !s.matchesNewline(start)
}

// howManyRunesUntil iterates the scanners rune slice executing the function for
// on each rune. The number of runes counted before the function results in true
// is returned to the user. If the function never returns true then the length
// of the rune slice, from the start index, is returned.
func (s *scanner) howManyRunesUntil(start int, f func(int, rune) bool) (i int) {

	var ru rune

	for i, ru = range s.runes[start:] {
		if f(i, ru) {
			break
		}
	}

	return i
}

func (s *scanner) countNewlineRunes(start int) int {

	const (
		LF        = token.LEXEME_NEWLINE_LF
		CRLF      = token.LEXEME_NEWLINE_CRLF
		NOT_FOUND = 0
	)

	size := s.len()

	if size > 0 && s.matchesNonTerminal(start, LF) {
		return len(LF)
	}

	if size > 1 && s.matchesNonTerminal(start, CRLF) {
		return len(CRLF)
	}

	return NOT_FOUND
}

func (s *scanner) runesUntilNewline(start int) int {
	return s.howManyRunesUntil(start, func(i int, ru rune) bool {
		return s.matchesNewline(i)
	})
}

func (s *scanner) countWordRunes(start int) int {
	return s.howManyRunesUntil(start, func(i int, ru rune) bool {

		if i == 0 && ru == '_' {
			return true
		}

		return ru != '_' && !unicode.IsLetter(ru)
	})
}

func (s *scanner) countDigitRunes(start int) int {
	return s.howManyRunesUntil(start, func(_ int, ru rune) bool {
		return !unicode.IsDigit(ru)
	})
}

// tokenize slices off the next token from the scanners rune array and updates
// the line and column numbers accordingly.
func (s *scanner) tokenize(n int, k token.Kind, newline bool) (tk token.Token) {

	if s.len() < n {
		panic("SANITY CHECK! Bad argument, n is bigger than the remaining source code")
	}

	tk = token.Token{
		Kind:  k,
		Value: string(s.runes[:n]),
		Line:  s.line,
		Col:   s.col,
	}

	s.runes = s.runes[n:]

	if newline {
		s.line++
		s.col = 0
	} else {
		s.col += n
	}

	return
}
