package scanner

import (
	"github.com/PaulioRandall/scarlet-go/token2"
)

// findComment satisfies the source.TokenFinder function prototype. It attempts
// to match the next token to the comment kind returning its length if matched.
func findComment(r []rune) (n int, k token.Kind, e error) {
	for _, ru := range r {
		if (n == 0 || n == 1) && ru != '/' {
			goto NOT_FOUND
		}

		if _, k2, e2 := findNewline(r[n:]); e2 != nil {
			e = e2
			goto NOT_FOUND
		} else if k2 != token.UNDEFINED {
			goto FOUND
		}

		n++
	}

	if n > 2 {
		goto FOUND
	}

NOT_FOUND:
	n = 0
	return

FOUND:
	k = token.COMMENT
	return
}
