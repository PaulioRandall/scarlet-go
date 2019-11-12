package count

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/cookies"
)

// CountSpaces counts the number of whitespace runes that form the prefix of
// the input `runes` slice.
func CountSpaces(runes []rune) (n int) {

	for _, ru := range runes {
		if !unicode.IsSpace(ru) || cookies.NewlineRunes(runes, n) != 0 {
			return
		}

		n++
	}

	return
}
