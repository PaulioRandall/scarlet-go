package count

import (
	"unicode"
)

// CountWord counts the number of runes in the next word.
func CountWord(runes []rune) (n int) {

	for _, ru := range runes {
		if !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	return
}
