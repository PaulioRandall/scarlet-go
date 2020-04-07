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

// symbolStream provides access to an ordered stream of terminal symbols (runes)
// representing a script. The stream also stores the current cursor position in
// the form of line and column indexes.
type symbolStream struct {
	runes []rune // Symbols representing a script.
	line  int
	col   int
}

// empty returns true if the stream is empty.
func (ss *symbolStream) empty() bool {
	return len(ss.runes) == 0
}

// len returns the length of the stream.
func (ss *symbolStream) len() int {
	return len(ss.runes)
}

// lineIndex returns the current line index the stream is at within the script.
func (ss *symbolStream) lineIndex() int {
	return ss.line
}

// colIndex returns the current column index of the current line.
func (ss *symbolStream) colIndex() int {
	return ss.col
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
	return ss.howManyNewlineSymbols(start) > 0
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

func (ss *symbolStream) howManyNewlineSymbols(start int) int {

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

func (ss *symbolStream) howManyConsecutiveLetters(start int) int {
	return ss.countRunesWhile(start, func(i int, ru rune) bool {

		if i == 0 && ru == '_' {
			return false
		}

		return ru == '_' || unicode.IsLetter(ru)
	})
}

func (ss *symbolStream) howManyConsecutiveDigits(start int) int {
	return ss.countRunesWhile(start, func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

func (ss *symbolStream) whensTheNextNewline(start int) int {
	return ss.countRunesWhile(start, func(i int, ru rune) bool {
		return !ss.isNewline(i)
	})
}

// peek performs a read without eating up the symbols in the stream or updating
// the line and column indexes.
func (ss *symbolStream) peek(runeCount int) string {
	return string(ss.runes[:runeCount])
}

// read reads the specified number of symbols from the stream updating the line
// and column indexes accordingly. If you want to record the line or column
// index of the read symbols, record them before performing the read.
func (ss *symbolStream) read(runeCount int, isNewline bool) string {

	if ss.len() < runeCount {
		codingError("Bad argument, requested read amount is bigger than the number of remaining runes")
	}

	r := ss.peek(runeCount)
	ss.runes = ss.runes[runeCount:]

	if isNewline {
		ss.line++
		ss.col = 0
	} else {
		ss.col += runeCount
	}

	return r
}
