package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/cookies"
)

// countSpaces counts the number of whitespace runes that form the prefix of
// the input `runes` slice.
func countSpaces(runes []rune) (n int) {

	for _, ru := range runes {
		if !unicode.IsSpace(ru) || cookies.NewlineRunes(runes, n) != 0 {
			return
		}

		n++
	}

	return
}

// countWord counts the number of runes in the next word.
func countWord(runes []rune) (n int) {

	for _, ru := range runes {
		if !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	return
}
