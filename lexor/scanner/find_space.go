package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findSpace satisfies the source.TokenFinder function prototype. It attempts
// to find whitespace or a newline token.
func findSpace(r []rune) (n int, k token.Kind, _ error) {

	prev := rune(0)

	for _, ru := range r {
		switch {
		case n == 0 && ru == '\n':
			goto FOUND_NEWLINE
		case n == 1 && prev == '\r' && ru == '\n':
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

	if n == 0 {
		goto NOT_SPACE
	}

FOUND_WHITESPACE:
	k = token.WHITESPACE
	return

FOUND_NEWLINE:
	n++
	k = token.NEWLINE
	return

NOT_SPACE:
	return
}
