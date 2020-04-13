package symbol

// impl is the one and only implementation of the SymbolStream interface.
type impl struct {
	runes []rune // Symbols representing a script.
	line  int
	col   int
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
func (ss *impl) CountSymbolsWhile(start int, f func(int, rune) bool) int {

	var ru rune
	var i int

	for i, ru = range ss.runes[start:] {
		if !f(i, ru) {
			break
		}
	}

	return i
}

// PeekTerminal satisfies the SymbolStream interface.
func (ss *impl) PeekTerminal(index int) rune {
	return ss.runes[index]
}

// PeekNonTerminal satisfies the SymbolStream interface.
func (ss *impl) PeekNonTerminal(runeCount int) string {
	return string(ss.runes[:runeCount])
}

// ReadNonTerminal satisfies the SymbolStream interface.
func (ss *impl) ReadNonTerminal(runeCount int) string {

	if runeCount > ss.Len() {
		panic("PROGRAMMERS ERROR! Bad argument, " +
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

	r := ss.PeekNonTerminal(runeCount)
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

	if ss.IsMatch(start, LF) {
		return len(LF)
	}

	if ss.IsMatch(start, CRLF) {
		return len(CRLF)
	}

	return NOT_FOUND
}

// IndexOfNextNewline satisfies the SymbolStream interface.
func (ss *impl) IndexOfNextNewline(start int) int {
	return ss.CountSymbolsWhile(start, func(i int, _ rune) bool {
		return !ss.IsNewline(i)
	})
}
