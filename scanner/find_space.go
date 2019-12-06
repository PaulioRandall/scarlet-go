package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findSpace satisfies the source.TokenFinder function prototype.
func findSpace(r []rune) (n int, k token.Kind) {

	for _, ru := range r {
		if _, kd := findNewline(r); !unicode.IsSpace(ru) || kd == token.WHITESPACE {
			break
		}

		n++
	}

	if n > 0 {
		k = token.WHITESPACE
	}

	return
}
