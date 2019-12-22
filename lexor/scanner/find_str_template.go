package scanner

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findStrTemplate satisfies the source.TokenFinder function prototype. It
// attempts to match the next token to the string template kind returning its
// length and kind if matched.
func findStrTemplate(r []rune) (_ int, _ token.Kind, e error) {

	prev := rune(0)

	for i, ru := range r {

		switch {
		case i == 0 && ru != '"':
			return
		case i == 0 && ru == '"':
			break
		case prev != '\\' && ru == '"':
			return i + 1, token.STR_TEMPLATE, nil
		case ru == '\n':
			goto ERROR
		}

		prev = ru
	}

ERROR:
	e = errors.New("Unterminated string template")
	return
}
