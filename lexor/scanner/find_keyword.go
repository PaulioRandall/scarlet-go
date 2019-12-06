package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findKeyword satisfies the source.TokenFinder function prototype.
func findKeyword(r []rune) (n int, _ token.Kind) {

	for _, ru := range r {
		if !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	return checkIfKeyword(r, n)
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
		`FUNC`: token.FUNC,
		`END`:  token.END,
	}

	for k, v := range ks {
		if k == s {
			return v
		}
	}

	return token.UNDEFINED
}
