package scanner

import (
	"unicode"
)

// countSpaces counts the number of whitespace runes that form the prefix of
// the input `runes` slice.
func countSpaces(runes []rune) (n int) {
	for _, ru := range runes {
		if !unicode.IsSpace(ru) || newlineRunes(runes, n) != 0 {
			return
		}

		n++
	}

	return
}

// newlineRunes returns the number runes representing the next new line if the
// next token is a newline token else zero is returned. I.e. if the next rune
// is `\n` then `1` is returned, if the the next rune are `\r\n` then `2` is
// returned else 0 is returned.
func newlineRunes(runes []rune, i int) int {

	if runes[i] == '\n' {
		return 1
	}

	if i+1 < len(runes) {
		if runes[i] == '\r' && runes[i+1] == '\n' {
			return 2
		}
	}

	return 0
}
