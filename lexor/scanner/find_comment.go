package scanner

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// findComment satisfies the source.TokenFinder function prototype. It attempts
// to match the next token to the comment kind returning its length if matched.
func findComment(r []rune) (n int, k token.Kind, e error) {

	prev := rune(0)

	for _, ru := range r {

		switch {
		case n < 2:
			if ru != '/' {
				goto NOT_COMMENT
			}
		case prev == '\r' && ru == '\n':
			n--
			goto FOUND
		case ru == '\n':
			goto FOUND
		}

		n++
		prev = ru
	}

	if n > 0 {
		goto FOUND
	}

NOT_COMMENT:
	n = 0
	return

FOUND:
	k = token.COMMENT
	return
}
