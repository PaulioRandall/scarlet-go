package scanner

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findStrLiteral satisfies the source.TokenFinder function prototype. It
// attempts to match the next token to the string literal kind returning its
// length and kind if matched.
func findStrLiteral(r []rune) (_ int, _ token.Kind, e error) {

	for i, ru := range r {
		switch {
		case i == 0:
			if ru == '`' {
				continue
			}
			return
		case ru == '`':
			return i + 1, token.STR_LITERAL, nil
		case ru == '\n':
			goto ERROR
		}
	}

ERROR:
	e = errors.New("Unterminated string literal")
	return
}
