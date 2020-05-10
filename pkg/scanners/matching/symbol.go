package matching

// symbols provides access to an ordered stream of terminal symbols (runes)
// representing a script. The stream also monitors the current cursor position
// in the form of line and column indexes.
type symbols struct {
	runes []rune
	line  int
	col   int
}

func (s *symbols) empty() bool {
	return len(s.runes) == 0
}

func (s *symbols) len() int {
	return len(s.runes)
}

func (s *symbols) drain() {
	s.runes = []rune{}
}

func (s *symbols) isMatch(start int, str string) bool {

	haystack := s.runes[start:]

	if len(str) > len(haystack) {
		return false
	}

	for i, ru := range str {
		if haystack[i] != ru {
			return false
		}
	}

	return true
}

func (s *symbols) countSymbolsWhile(start int, f func(int, rune) bool) int {

	runes := s.runes[start:]
	size := len(runes)

	for i := 0; i < size; i++ {
		if !f(i, runes[i]) {
			return i
		}
	}

	return size
}

func (s *symbols) peekTerminal(index int) rune {
	return s.runes[index]
}

func (s *symbols) peekNonTerminal(runeCount int) string {
	return string(s.runes[:runeCount])
}

func (s *symbols) readNonTerminal(runeCount int) string {

	if runeCount > s.len() {
		panic("PROGRAMMING ERROR! Bad argument, " +
			"requested slice amount is bigger than the number of remaining runes")
	}

	for i := 0; i < runeCount; i++ {
		switch s.countNewlineSymbols(i) {
		case 2:
			i++
			fallthrough

		case 1:
			s.line++
			s.col = 0

		case 0:
			s.col++
		}
	}

	r := s.peekNonTerminal(runeCount)
	s.runes = s.runes[runeCount:]

	return r
}

func (s *symbols) isNewline(start int) bool {
	return s.countNewlineSymbols(start) > 0
}

func (s *symbols) countNewlineSymbols(start int) int {
	switch {
	case s.isMatch(start, "\n"):
		return 1

	case s.isMatch(start, "\r\n"):
		return 2

	default:
		return 0
	}
}

func (s *symbols) indexOfNextNewline(start int) int {
	return s.countSymbolsWhile(start, func(i int, _ rune) bool {
		return !s.isNewline(i)
	})
}
