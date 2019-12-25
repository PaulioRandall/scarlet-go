package scanner

import (
	"errors"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findComment satisfies the source.TokenFinder function prototype. It attempts
// to match the next token to the comment kind returning its length if matched.
func findComment(r []rune) (n int, k token.Kind, e error) {

	prev := rune(0)

	for _, ru := range r {

		switch {
		case n < 2 && ru != '/':
			goto NOT_COMMENT
		case n < 2:
			// First two runes must be `//`
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

// findSpace satisfies the source.TokenFinder function prototype. It attempts
// to find whitespace or a newline token.
func findSpace(r []rune) (n int, _ token.Kind, _ error) {

	prev := rune(0)

	if len(r) == 0 {
		goto NOT_SPACE
	}

	for _, ru := range r {
		switch {
		case n == 0 && ru == '\n': // First rune only
			goto FOUND_NEWLINE
		case n == 1 && prev == '\r' && ru == '\n': // Second rune only
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

FOUND_WHITESPACE:
	return n, token.WHITESPACE, nil

FOUND_NEWLINE:
	return n + 1, token.NEWLINE, nil

NOT_SPACE:
	return 0, token.UNDEFINED, nil
}

// findNumLiteral satisfies the source.TokenFinder function prototype. It
// attempts to match the next token to the number literal kind returning its
// length and kind if matched.
func findNumLiteral(r []rune) (_ int, _ token.Kind, e error) {

	size := len(r)
	n := countDigits(r, size, 0)

	if n == 0 {
		return
	}

	if n < size && r[n] == '.' {

		n++ // Decimal point
		n += countDigits(r, size, n)

		if n == 0 {
			e = errors.New("Expected digit after decimal point")
			return
		}
	}

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
