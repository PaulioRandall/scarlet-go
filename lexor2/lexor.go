package lexor

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// ****************************************************************************
// * Package API
// ****************************************************************************

// Scanner is a structure for parsing source code into tokens.
type Scanner struct {
	runes []rune // Source code
	line  int    // Line index
	col   int    // Column index
}

// New creates a new scanner to parse the input string.
func New(s string) *Scanner {
	return &Scanner{
		runes: []rune(s),
		line:  0,
		col:   0,
	}
}

// Next returns the next token in source code.
func (scn *Scanner) Next() (tk token.Token) {

	type scanFunc func() token.Token

	fs := []scanFunc{
		scn.scanNewline,
		scn.scanSpace,
		scn.scanComment,
		scn.scanNumLiteral,
		scn.scanStrLiteral,
		scn.scanStrTemplate,
	}

	if len(scn.runes) == 0 {
		return
	}

	for _, f := range fs {
		tk = f()

		if tk != token.ZERO() {
			return
		}
	}

	return
}

// ****************************************************************************
// * Helper functions
// ****************************************************************************

// tokenize slices off the next token from the scanners rune array and updates
// the line and column numbers accordingly.
func (scn *Scanner) tokenize(n int, k token.Kind, newline bool) (tk token.Token) {

	if len(scn.runes) < n {
		panic("Bad function argument, n is bigger than the source code")
	}

	tk = token.Token{
		Kind:  k,
		Value: string(scn.runes[:n]),
		Line:  scn.line,
		Col:   scn.col,
	}

	scn.runes = scn.runes[n:]

	if newline {
		scn.line++
		scn.col = 0
	} else {
		scn.col += n
	}

	return
}

// newlineTerminals returns the number of terminal symbols that make up the next
// newline token in the slice. If the next token is not a newline token then 0
// is returned.
func newlineTerminals(r []rune) (_ int) {

	size := len(r)

	if size < 1 {
		return
	}

	if r[0] == '\n' { // LF
		return 1
	}

	if size > 1 && r[0] == '\r' && r[1] == '\n' { // CRLF
		return 2
	}

	return
}

// countDigits counts an uninterupted series of digits in the rune slice
// starting from the specified index.
func countDigits(r []rune, start int) (n int) {

	size := len(r)

	for n = start; n < size; n++ {
		if !unicode.IsDigit(r[n]) {
			break
		}
	}

	return n - start
}

// ****************************************************************************
// * Scanning functions: func() token.Token
// ****************************************************************************

// scanNewline attempts to scan a newline token. If successful a non-empty
// newline token is returned.
func (scn *Scanner) scanNewline() (_ token.Token) {

	if n := newlineTerminals(scn.runes); n > 0 {
		return scn.tokenize(n, token.NEWLINE, true)
	}

	return
}

// scanComment attempts to scan a comment. If successful a non-empty comment
// token is returned.
func (scn *Scanner) scanComment() (_ token.Token) {

	const COMMENT_PREFIX int = 2
	var i int
	size := len(scn.runes)

	if size < COMMENT_PREFIX {
		return
	}

	if scn.runes[0] != '/' || scn.runes[1] != '/' {
		return
	}

	for i = COMMENT_PREFIX; i < size; i++ {
		if n := newlineTerminals(scn.runes[i:]); n > 0 {
			break
		}
	}

	return scn.tokenize(i, token.COMMENT, false)
}

// scanSpace attempts to scan a series of whitespace characters. If successful
// a non-empty whitespace token is returned. Assumes that the scanners rune
// array length is greater than 0.
func (scn *Scanner) scanSpace() (_ token.Token) {

	var n int

	if newlineTerminals(scn.runes) != 0 {
		return
	}

	if !unicode.IsSpace(scn.runes[0]) {
		return
	}

	for i, ru := range scn.runes {

		if !unicode.IsSpace(ru) {
			break
		}

		if n := newlineTerminals(scn.runes[i:]); n > 0 {
			break
		}

		n++
	}

	return scn.tokenize(n, token.WHITESPACE, false)
}

// scanNumLiteral attempts to scan a literal number. If successful a non-empty
// number literal token is returned.
func (scn *Scanner) scanNumLiteral() (_ token.Token) {

	r := scn.runes
	n := countDigits(r, 0)

	if n == 0 {
		return
	}

	if n == len(r) || r[n] != '.' {
		return scn.tokenize(n, token.INT_LITERAL, false)
	}

	n++ // +1 for decimal point
	d := countDigits(r, n)
	if d == 0 {
		panic("Expected digit after decimal point")
	}

	n += d
	return scn.tokenize(n, token.REAL_LITERAL, false)
}

// scanStrLiteral attempts to scan a string literal. If successful a non-empty
// string literal token is returned. Assumes that the scanners rune array
// length is greater than 0.
func (scn *Scanner) scanStrLiteral() (_ token.Token) {

	for i, ru := range scn.runes {
		switch {
		case i == 0 && ru != '`':
			return
		case i == 0:
			continue
		case ru == '`':
			return scn.tokenize(i+1, token.STR_LITERAL, false)
		case ru == '\n':
			goto ERROR
		}
	}

ERROR:
	panic("Unterminated string literal")
}

// scanStrTemplate attempts to scan a string template. If successful a non-empty
// string template token is returned. Assumes that the scanners rune array
// length is greater than 0.
func (scn *Scanner) scanStrTemplate() (_ token.Token) {

	prev := rune(0)

	for i, ru := range scn.runes {

		switch {
		case i == 0 && ru != '"':
			return
		case i == 0 && ru == '"':
			break
		case prev != '\\' && ru == '"':
			return scn.tokenize(i+1, token.STR_TEMPLATE, false)
		case ru == '\n':
			goto ERROR
		}

		prev = ru
	}

ERROR:
	panic("Unterminated string template")
}
