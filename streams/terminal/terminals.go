// terminal package was created to separate the concern of managing access to
// terminal symbols within a stream with the concern of scanning tokens from a
// script; this package manages access to the terminal symbols. Users of the
// SymbolStream interface are able to inspect and read off sequences of terminal
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
// made by those programming to the SymbolStream interface --or possible errors
// in the implementation itself--.
//
// This package does not hold any knowledge about the text being scanned except
// that it may be written over multiple lines. For simplicity and lack of
// requirement, no back tracking functionality is provided.
package terminal

// TerminalStream provides access to an ordered stream of terminal symbols
// (runes) representing a script. The stream also monitors the current cursor
// position in the form of line and column indexes.
type TerminalStream struct {
	runes []rune // Terminals representing a script.
	line  int
	col   int
}

// New uses s as the contents of a new TerminalStream.
func New(s string) *TerminalStream {
	return &TerminalStream{
		runes: []rune(s),
	}
}

// Empty returns true if the stream is empty.
func (ts *TerminalStream) Empty() bool {
	return len(ts.runes) == 0
}

// Len returns the number of symbols remaining in the stream.
func (ts *TerminalStream) Len() int {
	return len(ts.runes)
}

// IsMatch returns true if s matches the sequence of symbols starting from
// start.
func (ts *TerminalStream) IsMatch(start int, s string) bool {

	haystack := ts.runes[start:]

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

// CountSymbolsWhile loops through the symbol stream, starting at start, while
// f returns true returning the number of succetsful iterations. f is invoked
// at the beginning of each iteration like a traditional while loop.
func (ts *TerminalStream) CountSymbolsWhile(start int, f func(int, rune) bool) int {

	var ru rune
	var i int

	for i, ru = range ts.runes[start:] {
		if !f(i, ru) {
			break
		}
	}

	return i
}

// PeekTerminal performs a read for a single terminal symbol without eating it
// up from the stream or updating the line and column indexes.
func (ts *TerminalStream) PeekTerminal(index int) rune {
	return ts.runes[index]
}

// PeekNonTerminal performs a read for a non-terminal without eating up the
// terminal symbols from the stream or updating the line and column indexes.
func (ts *TerminalStream) PeekNonTerminal(runeCount int) string {
	return string(ts.runes[:runeCount])
}

// ReadNonTerminal reads the specified number of terminal symbols from the
// stream updating the line and column indexes accordingly. If you want to
// record the line or column index of the read terminals, get them before
// performing the slice.
func (ts *TerminalStream) ReadNonTerminal(runeCount int) string {

	if runeCount > ts.Len() {
		panic("PROGRAMMERS ERROR! Bad argument, " +
			"requested slice amount is bigger than the number of remaining runes")
	}

	for i := 0; i < runeCount; i++ {
		switch ts.CountNewlineSymbols(i) {
		case 2:
			i++
			fallthrough
		case 1:
			ts.line++
			ts.col = 0
		case 0:
			ts.col++
		}
	}

	r := ts.PeekNonTerminal(runeCount)
	ts.runes = ts.runes[runeCount:]

	return r
}

// LineIndex returns the current line index within the text being read.
func (ts *TerminalStream) LineIndex() int {
	return ts.line
}

// ColIndex returns the current column index from the current line within the
// text being read.
func (ts *TerminalStream) ColIndex() int {
	return ts.col
}

// IsNewLine returns true if the sequence of symbols starting from start match
// a line break, i.e. LF or CRLF.
func (ts *TerminalStream) IsNewline(start int) bool {
	return ts.CountNewlineSymbols(start) > 0
}

// CountNewlineSymbols returns the number of symbols representing a line
// break at the start index within the symbol stream. If no line break occurs
// at start then 0 is returned.
func (ts *TerminalStream) CountNewlineSymbols(start int) int {

	const (
		LF        = "\n"
		CRLF      = "\r\n"
		NOT_FOUND = 0
	)

	if ts.IsMatch(start, LF) {
		return len(LF)
	}

	if ts.IsMatch(start, CRLF) {
		return len(CRLF)
	}

	return NOT_FOUND
}

// IndexOfNextNewline returns the index within the symbol stream, starting
// at start, where the next line break occurs.
func (ts *TerminalStream) IndexOfNextNewline(start int) int {
	return ts.CountSymbolsWhile(start, func(i int, _ rune) bool {
		return !ts.IsNewline(i)
	})
}
