package symbol

import (
	"unicode"
)

// SymbolStream provides access to an ordered stream of terminal symbols (runes)
// representing a script. The stream also monitors the current cursor position
// in the form of line and column indexes.
type SymbolStream interface {

	// Empty returns true if the stream is empty.
	Empty() bool

	// Len returns the number of symbols remaining in the stream.
	Len() int

	// LineIndex returns the current line index within the text being read.
	LineIndex() int

	// ColIndex returns the current column index from the current line within the
	// text being read.
	ColIndex() int

	// IsMatch returns true if s matches the sequence of symbols starting from
	// start.
	IsMatch(start int, s string) bool

	// IsNewLine returns true if the sequence of symbols starting from start match
	// a line break, i.e. LF or CRLF.
	IsNewline(start int) bool

	// CountSymbolsWhile loops through the symbol stream, starting at start, while
	// f returns true returning the number of successful iterations. f is invoked
	// at the beginning of each iteration like a traditional while loop.
	CountSymbolsWhile(start int, f func(int, rune) bool) int

	// CountNewlineSymbols returns the number of symbols representing a line
	// break at the start index within the symbol stream. If no line break occurs
	// at start then 0 is returned.
	CountNewlineSymbols(start int) int

	// CountConsecutiveLetters returns the number of consecutive letters
	// (unicode category L) and underscores '_', starting at start, within the
	// symbol stream.
	CountConsecutiveLetters(start int) int

	// CountConsecutiveDigits returns the number of consecutive digits
	// (unicode category Nd) and underscores '_', starting at start, within the
	// symbol stream.
	CountConsecutiveDigits(start int) int

	// IndexOfNextNewline returns the index within the symbol stream, starting
	// at start, where the next line break occurs.
	IndexOfNextNewline(start int) int

	// Peek performs a read without eating up the symbols in the stream or
	// updating the line and column indexes.
	Peek(runeCount int) string

	// Read reads the specified number of symbols from the stream updating the
	// line and column indexes accordingly. If you want to record the line or
	// column index of the read symbols, get them before performing the read.
	Read(runeCount int, isNewline bool) string
}

// impl is the implementation of the SymbolStream interface.
type impl struct {
	runes []rune // Symbols representing a script.
	line  int
	col   int
}

// NewSymbolStream uses s as the contents of a new SymbolStream.
func NewSymbolStream(s string) SymbolStream {
	return &impl{
		runes: []rune(s),
	}
}

// Empty satisfies the SymbolStream interface.
func (ss *impl) Empty() bool {
	return len(ss.runes) == 0
}

// Len satisfies the SymbolStream interface.
func (ss *impl) Len() int {
	return len(ss.runes)
}

// LineIndex satisfies the SymbolStream interface.
func (ss *impl) LineIndex() int {
	return ss.line
}

// ColIndex satisfies the SymbolStream interface.
func (ss *impl) ColIndex() int {
	return ss.col
}

// IsMatch satisfies the SymbolStream interface.
func (ss *impl) IsMatch(start int, s string) bool {

	haystack := ss.runes[start:]

	if len(s) > len(haystack) {
		codingError("Bad argument, `s` is bigger than the `haystack`")
	}

	for i, ru := range s {
		if haystack[i] != ru {
			return false
		}
	}

	return true
}

// IsNewline satisfies the SymbolStream interface.
func (ss *impl) IsNewline(start int) bool {
	return ss.CountNewlineSymbols(start) > 0
}

// CountSymbolsWhile satisfies the SymbolStream interface.
func (ss *impl) CountSymbolsWhile(start int, f func(int, rune) bool) (i int) {

	var ru rune

	for i, ru = range ss.runes[start:] {
		if !f(i, ru) {
			break
		}
	}

	return i
}

// CountNewlineSymbols satisfies the SymbolStream interface.
func (ss *impl) CountNewlineSymbols(start int) int {

	const (
		LF        = "\n"
		CRLF      = "\r\n"
		NOT_FOUND = 0
	)

	size := ss.Len()

	if size > 0 && ss.IsMatch(start, LF) {
		return len(LF)
	}

	if size > 1 && ss.IsMatch(start, CRLF) {
		return len(CRLF)
	}

	return NOT_FOUND
}

// CountConsecutiveLetters satisfies the SymbolStream interface.
func (ss *impl) CountConsecutiveLetters(start int) int {
	return ss.CountSymbolsWhile(start, func(i int, ru rune) bool {
		return ru == '_' || unicode.IsLetter(ru)
	})
}

// CountConsecutiveDigits satisfies the SymbolStream interface.
func (ss *impl) CountConsecutiveDigits(start int) int {
	return ss.CountSymbolsWhile(start, func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

// IndexOfNextNewline satisfies the SymbolStream interface.
func (ss *impl) IndexOfNextNewline(start int) int {
	return ss.CountSymbolsWhile(start, func(i int, ru rune) bool {
		return !ss.IsNewline(i)
	})
}

// Peek satisfies the SymbolStream interface.
func (ss *impl) Peek(runeCount int) string {
	return string(ss.runes[:runeCount])
}

// Peek satisfies the SymbolStream interface.
func (ss *impl) Read(runeCount int, isNewline bool) string {

	if ss.Len() < runeCount {
		codingError("Bad argument, requested read amount is bigger than the number of remaining runes")
	}

	r := ss.Peek(runeCount)
	ss.runes = ss.runes[runeCount:]

	if isNewline {
		ss.line++
		ss.col = 0
	} else {
		ss.col += runeCount
	}

	return r
}

// codingError generates a panic to stop the program because a programmer has
// introduced an error.
func codingError(msg string) {
	panic("PROGRAMMERS ERROR! " + msg)
}
