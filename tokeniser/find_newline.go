package tokeniser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// findNewline
func findNewline(r []rune) (int, token.Kind) {
	switch max := len(r); {
	case max == 0:
	case r[0] == '\n':
		return 1, token.NEWLINE
	case max > 1 && r[0] == '\r' && r[1] == '\n':
		return 2, token.NEWLINE
	}

	return 0, token.UNDEFINED
}
