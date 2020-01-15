package lexor

import (
	//"strconv"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

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
	}

	for _, f := range fs {
		tk = f()

		if tk != token.ZERO() {
			return
		}
	}

	return
}

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
func newlineTerminals(runes []rune) (_ int) {

	size := len(runes)

	if size < 1 {
		return
	}

	if runes[0] == '\n' { // LF
		return 1
	}

	if size > 1 && runes[0] == '\r' && runes[1] == '\n' { // CRLF
		return 2
	}

	return
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

// scanComment attempts to scan a series of whitespace characters. If
// successful a non-empty whitespace token is returned.
func (scn *Scanner) scanSpace() (_ token.Token) {

	var i int
	var ru rune

	for i, ru = range scn.runes {
		if !unicode.IsSpace(ru) {
			break
		}

		if n := newlineTerminals(scn.runes[i:]); n > 0 {
			break
		}
	}

	if i == 0 {
		return
	}

	i++ // Convert from index to count
	return scn.tokenize(i, token.WHITESPACE, false)
}
