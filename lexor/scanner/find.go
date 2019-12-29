package scanner

import (
	"errors"
	"strings"
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
	n := _countDigits(r, size, 0)

	if n == 0 {
		return
	}

	if n >= size || r[n] != '.' {
		return n, token.INT_LITERAL, nil
	}

	n++ // Decimal point
	d := _countDigits(r, size, n)

	if d == 0 {
		e = errors.New("Expected digit after decimal point")
		return
	} else {
		n += d
	}

	return n, token.REAL_LITERAL, nil
}

// countDigits counts an uninterupted series of digits in the rune slice
// starting from the specified index.
func _countDigits(r []rune, size int, start int) (n int) {
	for n = start; n < size; n++ {
		if !unicode.IsDigit(r[n]) {
			break
		}
	}

	return n - start
}

// findStrLiteral satisfies the source.TokenFinder function prototype. It
// attempts to match the next token to the string literal kind returning its
// length and kind if matched.
func findStrLiteral(r []rune) (_ int, _ token.Kind, e error) {

	for i, ru := range r {
		switch {
		case i == 0 && ru != '`':
			return
		case i == 0:
			continue
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

// findWord satisfies the source.TokenFinder function prototype. It attempts to
// match the next token to either the identifier kind or a keyword kind
// returning its length and kind if matched.
func findWord(r []rune) (n int, k token.Kind, _ error) {

	for _, ru := range r {
		if ru != '_' && !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	if n > 1 || (n == 1 && r[0] != '_') {
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

// findSymbol satisfies the source.TokenFinder function prototype. It attempts
// to match the next token to a symbol kind returning its length if matched.
func findSymbol(r []rune) (_ int, _ token.Kind, _ error) {

	type sym struct {
		v string
		n int
		k token.Kind
	}

	symbols := []sym{
		sym{":=", 2, token.ASSIGN},
		sym{"->", 2, token.RETURNS}, // Order matters! Must be before `-`
		sym{"(", 1, token.OPEN_PAREN},
		sym{")", 1, token.CLOSE_PAREN},
		sym{"[", 1, token.OPEN_GUARD},
		sym{"]", 1, token.CLOSE_GUARD},
		sym{"{", 1, token.OPEN_LIST},
		sym{"}", 1, token.CLOSE_LIST},
		sym{",", 1, token.ID_DELIM},
		sym{"@", 1, token.SPELL},
		sym{"+", 1, token.ADD},
		sym{"-", 1, token.SUBTRACT},
		sym{"/", 1, token.DIVIDE},
		sym{"*", 1, token.MULTIPLY},
		sym{"%", 1, token.MODULO},
		sym{"|", 1, token.OR},
		sym{"&", 1, token.AND},
		sym{"~", 1, token.NOT},
		sym{"¬", 1, token.NOT},
		sym{"=", 1, token.CMP_OPERATOR},
		sym{"#", 1, token.CMP_OPERATOR},
		sym{"<=", 2, token.CMP_OPERATOR}, // Order matters! Must be before `<`
		sym{">=", 2, token.CMP_OPERATOR}, // Order matters! Must be before `>`
		sym{"<", 1, token.CMP_OPERATOR},
		sym{">", 1, token.CMP_OPERATOR},
		sym{"_", 1, token.VOID},
	}

	src := string(r)

	for _, s := range symbols {
		if strings.HasPrefix(src, s.v) {
			return s.n, s.k, nil
		}
	}

	return
}
