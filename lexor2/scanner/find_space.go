package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token2"
)

// findSpace satisfies the source.TokenFinder function prototype.
func findSpace(r []rune) (n int, k token.Kind, _ error) {

	for _, ru := range r {

		_, k, _ = findNewline(r)
		if !unicode.IsSpace(ru) || k == token.WHITESPACE {
			break
		}

		n++
	}

	if n > 0 {
		k = token.WHITESPACE
	}

	return
}
