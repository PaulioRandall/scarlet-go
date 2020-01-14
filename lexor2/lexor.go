package lexor

import (
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

// scanNewline attempts to scan a newline token. If successful a non-empty
// newline token is returned.
func (scn *Scanner) scanNewline() (_ token.Token) {

	size := len(scn.runes)

	if size < 1 {
		return
	}

	if scn.runes[0] == '\n' { // LF
		return scn.tokenize(1, token.NEWLINE, true)
	}

	if size > 1 && scn.runes[0] == '\r' && scn.runes[1] == '\n' { // CRLF
		return scn.tokenize(2, token.NEWLINE, true)
	}

	return
}

// scanComment attempts to scan a comment. If successful a non-empty comment
// token is returned.
func (scn *Scanner) scanComment() (_ token.Token) {

	var n int

	size := len(scn.runes)
	if size < 2 {
		return
	}

	if scn.runes[0] != '/' || scn.runes[1] != '/' {
		return
	}

	for n = 2; n < size; n++ { // Scan to the end of the line or file
		if scn.runes[n] == '\n' {
			if scn.runes[n-1] == '\r' {
				n--
			}
			break
		}
	}

	return scn.tokenize(n, token.COMMENT, false)
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
