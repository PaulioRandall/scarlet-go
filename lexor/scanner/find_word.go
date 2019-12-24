package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findWord satisfies the source.TokenFinder function prototype. It attempts to
// match the next token to either the identifier kind or a keyword kind
// returning its length and kind if matched.
func findWord(r []rune) (n int, k token.Kind, _ error) {

	for _, ru := range r {
		if !unicode.IsLetter(ru) && ru != '_' {
			break
		}

		n++
	}

	if n > 0 {
		k = keywordOrID(r[:n])
	}

	return
}

// keywordOrID identifies the kind of the keyword or returns the ID kind.
func keywordOrID(r []rune) token.Kind {

	ks := map[string]token.Kind{
		`GLOBAL`: token.GLOBAL,
		`F`:      token.FUNC,
		`DO`:     token.DO,
		`WATCH`:  token.WATCH,
		`MATCH`:  token.MATCH,
		`END`:    token.END,
		`TRUE`:   token.TRUE,
		`FALSE`:  token.FALSE,
	}

	src := string(r)

	for k, v := range ks {
		if k == src {
			return v
		}
	}

	return token.ID
}
