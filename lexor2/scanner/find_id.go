package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token2"
)

// findId satisfies the source.TokenFinder function prototype. It attempts to
// match the next token to the identifier kind returning its length if matched.
func findId(r []rune) (n int, k token.Kind, _ error) {
	for _, ru := range r {
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
