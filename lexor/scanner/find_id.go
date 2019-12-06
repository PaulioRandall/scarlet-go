package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findId satisfies the source.TokenFinder function prototype. It attempts to
// match the next token to the identifier kind returning its length if matched.
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
