// symbol package was created to separate the concern of managing access to
// terminal symbols within a stream with the concern of scanning tokens from a
// script; this package manages access to the terminal symbols. Users of the
// symbolStream interface are able to inspect and read off sequences of terminal
// symbols from a stream, while the implementation keeps track of lines and
// columns within the streamed text.
//
// The API combines three responsibilities:
// 1. The base functions expose simple stream functionality such as Len, Empty,
// IsMatch, CountSymbolsWhile, PeekTerminal, PeekNonTerminal, and
// ReadNonTerminal.
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
// made by those programming to the symbolStream interface --or possible errors
// in the implementation itself--.
//
// This package does not hold any knowledge about the text being scanned except
// that it may be written over multiple lines. For simplicity and lack of
// requirement, no back tracking functionality is provided.
package matching

// symbolStream provides access to an ordered stream of terminal symbols
// (runes) representing a script. The stream also monitors the current cursor
// position in the form of line and column indexes.
type symbolStream struct {
	runes []rune // Terminals representing a script.
	line  int
	col   int
}

// empty returns true if the stream is empty.
func (ss *symbolStream) empty() bool {
	return len(ss.runes) == 0
}

// len returns the number of symbols remaining in the stream.
func (ss *symbolStream) len() int {
	return len(ss.runes)
}

// drain removes all remaining symbols in the stream.
func (ss *symbolStream) drain() {
	ss.runes = []rune{}
}

// isMatch returns true if s matches the sequence of symbols starting from
// start.
func (ss *symbolStream) isMatch(start int, s string) bool {

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

// countSymbolsWhile loops through the symbol stream, starting at start, while
// f returns true returning the number of succetsful iterations. f is invoked
// at the beginning of each iteration like a traditional while loop.
func (ss *symbolStream) countSymbolsWhile(start int, f func(int, rune) bool) int {

	var ru rune
	var i int

	for i, ru = range ss.runes[start:] {
		if !f(i, ru) {
			break
		}
	}

	return i
}

// peekTerminal performs a read for a single terminal symbol without eating it
// up from the stream or updating the line and column indexes.
func (ss *symbolStream) peekTerminal(index int) rune {
	return ss.runes[index]
}

// peekNonTerminal performs a read for a non-terminal without eating up the
// terminal symbols from the stream or updating the line and column indexes.
func (ss *symbolStream) peekNonTerminal(runeCount int) string {
	return string(ss.runes[:runeCount])
}

// readNonTerminal reads the specified number of terminal symbols from the
// stream updating the line and column indexes accordingly. If you want to
// record the line or column index of the read terminals, get them before
// performing the slice.
func (ss *symbolStream) readNonTerminal(runeCount int) string {

	if runeCount > ss.len() {
		panic("PROGRAMMERS ERROR! Bad argument, " +
			"requested slice amount is bigger than the number of remaining runes")
	}

	for i := 0; i < runeCount; i++ {
		switch ss.countNewlineSymbols(i) {
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

	r := ss.peekNonTerminal(runeCount)
	ss.runes = ss.runes[runeCount:]

	return r
}

// lineIndex returns the current line index within the text being read.
func (ss *symbolStream) lineIndex() int {
	return ss.line
}

// colIndex returns the current column index from the current line within the
// text being read.
func (ss *symbolStream) colIndex() int {
	return ss.col
}

// isNewline returns true if the sequence of symbols starting from start match
// a line break, i.e. LF or CRLF.
func (ss *symbolStream) isNewline(start int) bool {
	return ss.countNewlineSymbols(start) > 0
}

// countNewlineSymbols returns the number of symbols representing a line
// break at the start index within the symbol stream. If no line break occurs
// at start then 0 is returned.
func (ss *symbolStream) countNewlineSymbols(start int) int {

	const (
		LF        = "\n"
		CRLF      = "\r\n"
		NOT_FOUND = 0
	)

	if ss.isMatch(start, LF) {
		return len(LF)
	}

	if ss.isMatch(start, CRLF) {
		return len(CRLF)
	}

	return NOT_FOUND
}

// indexOfNextNewline returns the index within the symbol stream, starting
// at start, where the next line break occurs.
func (ss *symbolStream) indexOfNextNewline(start int) int {
	return ss.countSymbolsWhile(start, func(i int, _ rune) bool {
		return !ss.isNewline(i)
	})
}
