package scanner

import (
	"github.com/PaulioRandall/scarlet-go/token2"
)

// findNewline satisfies the source.TokenFinder function prototype.
func findNewline(r []rune) (n int, k token.Kind, _ error) {
	switch size := len(r); {
	case size < 1:
	case r[0] == '\n':
		n, k = 1, token.NEWLINE
	case size > 1 && r[0] == '\r' && r[1] == '\n':
		n, k = 2, token.NEWLINE
	}

	return
}
