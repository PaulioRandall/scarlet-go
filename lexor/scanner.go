package lexor

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// ****************************************************************************
// * TokenStream
// ****************************************************************************

// TokenStream represents a stream of tokens.
type TokenStream interface {

	// Next returns the next token in the stream. An EOF token is always returned
	// if the token stream is empty.
	Next() token.Token
}

// ****************************************************************************
// * scanner
// ****************************************************************************

// scanner is a structure for parsing source code into tokens. It implements
// the TokenStream interface so it may be wrapped.
type scanner struct {
	runes []rune // Source code
	line  int    // Line index
	col   int    // Column index
}

// NewScanner creates a new scanner to parse the input string.
func NewScanner(s string) TokenStream {
	return &scanner{
		runes: []rune(s),
		line:  0,
		col:   0,
	}
}

// Next satisfies the TokenStream interface.
func (s *scanner) Next() (tk token.Token) {

	if s.empty() {
		return token.Token{
			Kind: token.EOF,
			Line: s.line,
			Col:  s.col,
		}
	}

	type scanFunc func() token.Token

	fs := []scanFunc{
		s.scanNewline,
		s.scanSpace,
		s.scanComment,
		s.scanWord,
		s.scanSymbol,
		s.scanNumLiteral,
		s.scanStrLiteral,
		s.scanStrTemplate,
	}

	for _, f := range fs {
		tk = f()

		if tk != token.ZERO() {
			return
		}
	}

	panic(bard.NewTerror(s.line, s.col, nil,
		"Could not identify next token",
	))
}

// scanNewline attempts to scan a newline token. If successful a non-empty
// newline token is returned.
func (s *scanner) scanNewline() (_ token.Token) {

	if n := s.countNewlineTerminals(0); n > 0 {
		return s.tokenize(n, token.NEWLINE, true)
	}

	return
}

// scanComment attempts to scan a comment. If successful a non-empty comment
// token is returned.
func (s *scanner) scanComment() (_ token.Token) {

	if s.doesNotMatch(terminal_commentStart) {
		return
	}

	const PREFIXES = 1 // Number of terminals that signify a comment start

	n := s.matchUntilNewline(PREFIXES)
	return s.tokenize(n+PREFIXES, token.COMMENT, false)
}

// scanSpace attempts to scan a series of whitespace characters. If successful
// a non-empty whitespace token is returned. Assumes that the scanners rune
// array length is greater than 0.
func (s *scanner) scanSpace() (_ token.Token) {

	isSpace := func(i int, ru rune) bool {
		return s.matchesNewline(i) || !unicode.IsSpace(ru)
	}

	if isSpace(0, s.runes[0]) {
		return
	}

	n := s.matchUntil(0, isSpace)
	return s.tokenize(n, token.WHITESPACE, false)
}

// scanNumLiteral attempts to scan a literal number. If successful a non-empty
// number literal token is returned.
func (scn *scanner) scanNumLiteral() (_ token.Token) {

	r := scn.runes
	n := countDigits(r, 0)

	if n == 0 {
		return
	}

	if n == len(r) || r[n] != terminal_fractionalDelim {
		return scn.tokenize(n, token.INT, false)
	}

	n++ // +1 for decimal point
	d := countDigits(r, n)
	if d == 0 {
		panic(bard.NewTerror(scn.line, scn.col+n, nil,
			"Expected digit after decimal point",
		))
	}

	n += d
	return scn.tokenize(n, token.REAL, false)
}

// scanStrLiteral attempts to scan a string literal. If successful a non-empty
// string literal token is returned. Assumes that the scanners rune array
// length is greater than 0.
func (scn *scanner) scanStrLiteral() (_ token.Token) {

	for i, ru := range scn.runes {
		switch {
		case i == 0 && ru != terminal_stringStart:
			return
		case i == 0:
			continue
		case ru == terminal_stringEnd:
			return scn.tokenize(i+1, token.STR, false)
		case ru == terminal_lineFeed:
			goto ERROR
		}
	}

ERROR:
	panic(bard.NewTerror(scn.line, scn.col, nil,
		"Unterminated string literal",
	))
}

// scanStrTemplate attempts to scan a string template. If successful a non-empty
// string template token is returned. Assumes that the scanners rune array
// length is greater than 0.
func (scn *scanner) scanStrTemplate() (_ token.Token) {

	prev := rune(0)

	for i, ru := range scn.runes {

		switch {
		case i == 0 && ru != terminal_templateStart:
			return
		case i == 0 && ru == terminal_templateStart:
			break
		case prev != terminal_templateEscape && ru == terminal_templateEnd:
			return scn.tokenize(i+1, token.TEMPLATE, false)
		case ru == terminal_lineFeed:
			goto ERROR
		}

		prev = ru
	}

ERROR:
	panic(bard.NewTerror(scn.line, scn.col, nil,
		"Unterminated string template",
	))
}

// scanWord attempts to scan a keyword or identifier. If successful a non-empty
// keyword or identifier token is returned.
func (scn *scanner) scanWord() (_ token.Token) {

	var n int
	r := scn.runes

	for _, ru := range r {
		if ru != terminal_wordUnderscore && !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	if n == 0 || (n == 1 && r[0] == terminal_wordUnderscore) {
		return
	}

	k := identifyAskeywordOrID(r[:n])
	return scn.tokenize(n, k, false)
}

// scanSymbol attempts to scan a symbol token. If successful a non-empty
// symbol token is returned. Assumes that the scanners rune array length is
// greater than 0.
func (scn *scanner) scanSymbol() (_ token.Token) {

	var (
		a rune
		b rune
		n int
		k token.Kind
	)

	if size := len(scn.runes); size == 0 {
		return
	} else if size > 1 {
		b = scn.runes[1]
	}

	a = scn.runes[0]

	switch {
	case a == ':' && b == '=':
		n, k = 2, token.ASSIGN
	case a == '-' && b == '>': // Order matters! Must be before `-`
		n, k = 2, token.RETURNS
	case a == '(':
		n, k = 1, token.OPEN_PAREN
	case a == ')':
		n, k = 1, token.CLOSE_PAREN
	case a == '[':
		n, k = 1, token.OPEN_GUARD
	case a == ']':
		n, k = 1, token.CLOSE_GUARD
	case a == '{':
		n, k = 1, token.OPEN_LIST
	case a == '}':
		n, k = 1, token.CLOSE_LIST
	case a == ',':
		n, k = 1, token.DELIM
	case a == '_':
		n, k = 1, token.VOID
	case a == ';':
		n, k = 1, token.TERMINATOR
	case a == '@':
		n, k = 1, token.SPELL
	case a == '~' || a == '¬': // Negation
		n, k = 1, token.NOT
	case a == '+':
		n, k = 1, token.ADD
	case a == '-':
		n, k = 1, token.SUBTRACT
	case a == '*':
		n, k = 1, token.MULTIPLY
	case a == '/':
		n, k = 1, token.DIVIDE
	case a == '%':
		n, k = 1, token.MOD
	case a == '&':
		n, k = 1, token.AND
	case a == '|':
		n, k = 1, token.OR
	case a == '=':
		n, k = 1, token.EQU
	case a == '#':
		n, k = 1, token.NEQ
	case a == '<' && b == '=': // Order matters! Must be before `<`
		n, k = 2, token.LT_OR_EQU
	case a == '>' && b == '=': // Order matters! Must be before `<`
		n, k = 2, token.MT_OR_EQU
	case a == '<':
		n, k = 1, token.LT
	case a == '>':
		n, k = 1, token.MT
	}

	if k == token.UNDEFINED {
		return
	}

	return scn.tokenize(n, k, false)
}

// tokenize slices off the next token from the scanners rune array and updates
// the line and column numbers accordingly.
func (scn *scanner) tokenize(n int, k token.Kind, newline bool) (tk token.Token) {

	if len(scn.runes) < n {
		panic("SANITY CHECK! Bad function argument, n is bigger than the source code")
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

// ****************************************************************************
// * Helper functions
// ****************************************************************************

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

// identifyAskeywordOrID identifies the keyword kind represented by the input
// rune slice. If no keyword can be found then the identifier kind is returned.
func identifyAskeywordOrID(r []rune) token.Kind {

	s := string(r)
	k := identifyKeyword(s)

	if k == token.UNDEFINED {
		return token.ID
	}

	return k
}
