package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// identifyWord identifies the next word in the rune slice returning the kind
// of the word token and its length.
func identifyWord(runes []rune) (k token.Kind, n int) {

	for _, ru := range runes {
		if !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	return identifyWordKind(runes, n)
}

// identifyWordKind identifies the kind of the word token.
func identifyWordKind(runes []rune, size int) (k token.Kind, n int) {

	n = size

	switch string(runes[:size]) {
	case `PROCEDURE`:
		k = token.PROCEDURE
	case `END`:
		k = token.END
	default:
		n = 0
	}

	return
}
