package lexor

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// codingError generates a panic to stop the program because a programmer has
// introduced an error.
func codingError(msg string) {
	panic("PROGRAMMERS ERROR! " + msg)
}

// symbolStream provides access to an ordered stream of terminal symbols
// (runes).
type symbolStream struct {
	runes []rune // Symbols representing a script.
	line  int    // Current line index within the script.
	col   int    // Current column index within the line.
}

// empty returns true if the stream is empty.
func (ss *symbolStream) empty() bool {
	return len(ss.runes) == 0
}

// len returns the length of the stream.
func (ss *symbolStream) len() int {
	return len(ss.runes)
}

func (ss *symbolStream) isMatch(start int, needle string) bool {

	haystack := ss.runes[start:]

	if len(needle) > len(haystack) {
		codingError("Bad argument, the `needle` is bigger than the `haystack`")
	}

	for i, ru := range needle {
		if haystack[i] != ru {
			return false
		}
	}

	return true
}

func (ss *symbolStream) isNewline(start int) bool {
	return ss.countNewlineRunes(start) > 0
}

func (ss *symbolStream) countRunesWhile(start int, f func(int, rune) bool) (i int) {

	var ru rune

	for i, ru = range ss.runes[start:] {
		if !f(i, ru) {
			break
		}
	}

	return i
}

func (ss *symbolStream) countNewlineRunes(start int) int {

	const (
		LF        = token.NEWLINE_LF
		CRLF      = token.NEWLINE_CRLF
		NOT_FOUND = 0
	)

	size := ss.len()

	if size > 0 && ss.isMatch(start, LF) {
		return len(LF)
	}

	if size > 1 && ss.isMatch(start, CRLF) {
		return len(CRLF)
	}

	return NOT_FOUND
}

func (ss *symbolStream) runesUntilNewline(start int) int {
	return ss.countRunesWhile(start, func(i int, ru rune) bool {
		return !ss.isNewline(i)
	})
}

func (ss *symbolStream) countWordRunes(start int) int {
	return ss.countRunesWhile(start, func(i int, ru rune) bool {

		if i == 0 && ru == '_' {
			return false
		}

		return ru == '_' || unicode.IsLetter(ru)
	})
}

func (ss *symbolStream) countDigitRunes(start int) int {
	return ss.countRunesWhile(start, func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

func (ss *symbolStream) read(runeCount int, isNewline bool) string {

	if ss.len() < runeCount {
		codingError("Bad argument, requested read amount is bigger than the number of remaining runes")
	}

	r := string(ss.runes[:runeCount])
	ss.runes = ss.runes[runeCount:]

	if isNewline {
		ss.line++
		ss.col = 0
	} else {
		ss.col += runeCount
	}

	return r
}
