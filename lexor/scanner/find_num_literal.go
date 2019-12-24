package scanner

import (
	"errors"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findNumLiteral satisfies the source.TokenFinder function prototype. It
// attempts to match the next token to the number literal kind returning its
// length and kind if matched.
func findNumLiteral(r []rune) (_ int, _ token.Kind, e error) {

	size := len(r)
	n := countDigits(r, size, 0)

	if n == 0 {
		return
	}

	if n == size || r[n] != '.' {
		goto FOUND
	}

	n++ // Decimal point
	n += countDigits(r, size, n)

	if n == 0 {
		e = errors.New("Expected digit after decimal point")
		return
	}

FOUND:
	return n, token.NUM_LITERAL, nil
}

// countDigits counts an uninterupted series of digits in the rune slice
// starting from the specified index.
func countDigits(r []rune, size int, start int) (n int) {
	for n = start; n < size; n++ {
		if !unicode.IsDigit(r[n]) {
			break
		}
	}

	return n - start
}
