package tokeniser

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findId
func findId(runes []rune) (n int, k token.Kind) {
	for _, ru := range runes {
		if !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	if n > 0 {
		k = token.ID
	}

	return
}
