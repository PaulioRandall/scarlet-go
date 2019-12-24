package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findKeyword satisfies the source.TokenFinder function prototype.
func findKeyword(r []rune) (n int, k token.Kind, _ error) {

	for _, ru := range r {
		if !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	n, k = checkIfKeyword(r, n)
	return
}

// checkIfKeyword returns the the input `n` and the identified keyword kind if
// the word represts a keyword else returns zero and `UNDEFINED`.
func checkIfKeyword(r []rune, n int) (_ int, k token.Kind) {
	if n > 0 {
		s := string(r[:n])
		k = kindOfKeyword(s)
	}

	if k == token.UNDEFINED {
		n = 0
	}

	return n, k
}

// kindOfKeyword identifies the kind of the keyword or returns UNDEFINED.
func kindOfKeyword(s string) token.Kind {

	ks := map[string]token.Kind{
		`F`:     token.FUNC,
		`DO`:    token.DO,
		`WATCH`: token.WATCH,
		`MATCH`: token.MATCH,
		`END`:   token.END,
		`TRUE`:  token.TRUE,
		`FALSE`: token.FALSE,
	}

	for k, v := range ks {
		if k == s {
			return v
		}
	}

	return token.UNDEFINED
}
