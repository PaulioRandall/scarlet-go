// symbol package was created to separate the concern of managing access to
// terminal symbols within a stream with the concern of scanning tokens from a
// script; this package manages access to the terminal symbols. Users of the
// SymbolStream interface are able to inspect and read off sequences of terminal
// symbols from a stream, while the implementation keeps track of lines and
// columns within the streamed text.
//
// The API combines three responsibilities:
// 1. The base functions exposes simple stream functionality such as Len, Empty,
// IsMatch, CountSymbolsWhile, Peek, and Read.
// 2. The tracking functions, LineIndex and ColIndex, return the position of the
// stream relative to the text being streamed.
// 3. The remaining functions build upon the base functions to provide slightly
// higher level capabilities dealing with line breaks.
//
// Key decisions:
// 1. The three responsibilities seemed small and simple enough that a single
// interface combining them would be straight forward to create and maintain as
// well as being easier for package users to use.
// 2. The line separator terminals are hardcoded into the package because this
// is the simplist approach and the program is only expected to work on
// platforms using those line separators.
// 3. Any error results in a panic because all errors that can occur are those
// made by those programming to the SymbolStream interface --or possible errors
// in the implementation itself--.
//
// This package does not hold any knowledge about the text being scanned except
// that it may be written over multiple lines. For simplicity and lack of
// requirement, no back tracking functionality is provided.
package symbol

// SymbolStream provides access to an ordered stream of terminal symbols (runes)
// representing a script. The stream also monitors the current cursor position
// in the form of line and column indexes.
type SymbolStream interface {

	// Empty returns true if the stream is empty.
	Empty() bool

	// Len returns the number of symbols remaining in the stream.
	Len() int

	// IsMatch returns true if s matches the sequence of symbols starting from
	// start.
	IsMatch(start int, s string) bool

	// CountSymbolsWhile loops through the symbol stream, starting at start, while
	// f returns true returning the number of successful iterations. f is invoked
	// at the beginning of each iteration like a traditional while loop.
	CountSymbolsWhile(start int, f func(int, rune) bool) int

	// Peek performs a read for a single symbol without eating it up from the
	// stream or updating the line and column indexes.
	PeekSymbol(index int) rune

	// Peek performs a read without eating up the symbols in the stream or
	// updating the line and column indexes.
	Peek(runeCount int) string

	// Slice reads the specified number of symbols from the stream updating the
	// line and column indexes accordingly. If you want to record the line or
	// column index of the read symbols, get them before performing the slice.
	Slice(runeCount int) string

	// LineIndex returns the current line index within the text being read.
	LineIndex() int

	// ColIndex returns the current column index from the current line within the
	// text being read.
	ColIndex() int

	// IsNewLine returns true if the sequence of symbols starting from start match
	// a line break, i.e. LF or CRLF.
	IsNewline(start int) bool

	// CountNewlineSymbols returns the number of symbols representing a line
	// break at the start index within the symbol stream. If no line break occurs
	// at start then 0 is returned.
	CountNewlineSymbols(start int) int

	// IndexOfNextNewline returns the index within the symbol stream, starting
	// at start, where the next line break occurs.
	IndexOfNextNewline(start int) int
}

// impl is the one and only implementation of the SymbolStream interface.
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

// IsMatch satisfies the SymbolStream interface.
func (ss *impl) IsMatch(start int, s string) bool {

	haystack := ss.runes[start:]

	if len(s) > len(haystack) {
		return false
	}

	for i, ru := range s {
		if haystack[i] != ru {
			return false
		}
	}

	return true
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

// Peek satisfies the SymbolStream interface.
func (ss *impl) PeekSymbol(index int) rune {
	return ss.runes[index]
}

// Peek satisfies the SymbolStream interface.
func (ss *impl) Peek(runeCount int) string {
	return string(ss.runes[:runeCount])
}

// Slice satisfies the SymbolStream interface.
func (ss *impl) Slice(runeCount int) string {

	if ss.Len() < runeCount {
		codingError("Bad argument, " +
			"requested slice amount is bigger than the number of remaining runes")
	}

	for i := 0; i < runeCount; i++ {
		switch ss.CountNewlineSymbols(i) {
		case 2:
			i++
			fallthrough
		case 1:
			ss.line++
			ss.col = 0
		case 0:
			ss.col++
		}
	}

	r := ss.Peek(runeCount)
	ss.runes = ss.runes[runeCount:]

	return r
}

// LineIndex satisfies the SymbolStream interface.
func (ss *impl) LineIndex() int {
	return ss.line
}

// ColIndex satisfies the SymbolStream interface.
func (ss *impl) ColIndex() int {
	return ss.col
}

// IsNewline satisfies the SymbolStream interface.
func (ss *impl) IsNewline(start int) bool {
	return ss.CountNewlineSymbols(start) > 0
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

// IndexOfNextNewline satisfies the SymbolStream interface.
func (ss *impl) IndexOfNextNewline(start int) int {
	return ss.CountSymbolsWhile(start, func(i int, ru rune) bool {
		return !ss.IsNewline(i)
	})
}

// codingError generates a panic. It should be used when invalid API usage
// is detected.
func codingError(msg string) {
	panic("PROGRAMMERS ERROR! " + msg)
}
