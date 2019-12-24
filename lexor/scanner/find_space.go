package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findSpace satisfies the source.TokenFinder function prototype. It attempts
// to find whitespace or a newline token.
func findSpace(r []rune) (n int, _ token.Kind, _ error) {

	prev := rune(0)

	if len(r) == 0 {
		goto NOT_SPACE
	}

	for _, ru := range r {
		switch {
		case n == 0 && ru == '\n': // First rune only
			goto FOUND_NEWLINE
		case n == 1 && prev == '\r' && ru == '\n': // Second rune only
			goto FOUND_NEWLINE
		case ru == '\n':
			if prev == '\r' {
				n--
			}
			goto FOUND_WHITESPACE
		case unicode.IsSpace(ru):
			n++
			prev = ru
		case n > 0:
			goto FOUND_WHITESPACE
		default:
			goto NOT_SPACE
		}
	}

FOUND_WHITESPACE:
	return n, token.WHITESPACE, nil

FOUND_NEWLINE:
	return n + 1, token.NEWLINE, nil

NOT_SPACE:
	return 0, token.UNDEFINED, nil
}
