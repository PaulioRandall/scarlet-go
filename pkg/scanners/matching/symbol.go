package matching

// symbolStream provides access to an ordered stream of terminal symbols
// (runes) representing a script. The stream also monitors the current cursor
// position in the form of line and column indexes.
type symbolStream struct {
	runes []rune
	line  int
	col   int
}

func (ss *symbolStream) empty() bool {
	return len(ss.runes) == 0
}

func (ss *symbolStream) len() int {
	return len(ss.runes)
}

func (ss *symbolStream) drain() {
	ss.runes = []rune{}
}

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

func (ss *symbolStream) peekTerminal(index int) rune {
	return ss.runes[index]
}

func (ss *symbolStream) peekNonTerminal(runeCount int) string {
	return string(ss.runes[:runeCount])
}

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

func (ss *symbolStream) lineIndex() int {
	return ss.line
}

func (ss *symbolStream) colIndex() int {
	return ss.col
}

func (ss *symbolStream) isNewline(start int) bool {
	return ss.countNewlineSymbols(start) > 0
}

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

func (ss *symbolStream) indexOfNextNewline(start int) int {
	return ss.countSymbolsWhile(start, func(i int, _ rune) bool {
		return !ss.isNewline(i)
	})
}
