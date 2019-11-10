package scanner

import (
	"unicode"
)

// countSpaces counts the number of whitespace runes that form the prefix of
// the input `runes` slice.
func countSpaces(runes []rune) (n int) {
	for _, ru := range runes {
		if !unicode.IsSpace(ru) || isNewline(n, runes) {
			return
		}

		n++
	}

	return
}

// isNewline returns true if the next token is a newline token. This may be a
// single linefeed or carriage return followed by a linefeed.
func isNewline(i int, runes []rune) bool {
	if runes[i] == '\n' {
		return true
	}

	if i+1 < len(runes) {
		if runes[i] == '\r' && runes[i+1] == '\n' {
			return true
		}
	}

	return false
}
