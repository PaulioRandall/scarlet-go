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
			Kind: token.KIND_EOF,
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

	if n := s.howManyNewlineTerminals(0); n > 0 {
		return s.tokenize(n, token.KIND_NEWLINE, true)
	}

	return
}

// scanComment attempts to scan a comment. If successful a non-empty comment
// token is returned.
func (s *scanner) scanComment() (_ token.Token) {

	if s.doesNotMatch(0, token.TERMINAL_COMMENT_START) {
		return
	}

	const PREFIXES = 1 // Number of terminals that signify a comment start
	n := s.howManyRunesUntilNewline(PREFIXES)

	return s.tokenize(n, token.KIND_COMMENT, false)
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

	n := s.howManyRunesUntil(0, isSpace)
	return s.tokenize(n, token.KIND_WHITESPACE, false)
}

// scanNumLiteral attempts to scan a literal number. If successful a non-empty
// number literal token is returned.
func (s *scanner) scanNumLiteral() (_ token.Token) {

	const (
		DELIM_RUNE = token.TERMINAL_FRACTIONAL_DELIM
		DELIM_LEN  = 1
	)

	isNotDigit := func(_ int, ru rune) bool {
		return !unicode.IsDigit(ru)
	}

	intLen := s.howManyRunesUntil(0, isNotDigit)

	if intLen == 0 {
		return
	}

	if intLen == s.len() || s.doesNotMatch(intLen, DELIM_RUNE) {
		return s.tokenize(intLen, token.INT, false)
	}

	fractionalLen := s.howManyRunesUntil(intLen+DELIM_LEN, isNotDigit)

	if fractionalLen == 0 {
		panic(bard.NewTerror(s.line, s.col+intLen+DELIM_LEN, nil,
			"Expected digit after decimal point",
		))
	}

	n := intLen + DELIM_LEN + fractionalLen
	return s.tokenize(n, token.REAL, false)
}

// scanStrLiteral attempts to scan a string literal. If successful a non-empty
// string literal token is returned. Assumes that the scanners rune slice
// length is greater than 0.
func (s *scanner) scanStrLiteral() (_ token.Token) {

	const (
		START_RUNE = token.TERMINAL_STRING_START
		END_RUNE   = token.TERMINAL_STRING_END
	)

	panic_unterminatedString := func() {
		panic(bard.NewTerror(s.line, s.col, nil,
			"Unterminated string literal",
		))
	}

	n := s.howManyRunesUntil(0, func(i int, ru rune) bool {

		switch {
		case i == 0 && ru != START_RUNE:
			return true
		case i == 0:
			return false
		case ru == END_RUNE:
			return true
		case s.matchesNewline(i):
			panic_unterminatedString()
		}

		return false
	})

	if n == 0 {
		return
	}

	if n == s.len() {
		panic_unterminatedString()
	}

	return s.tokenize(n+1, token.STR, false)
}

// scanStrTemplate attempts to scan a string template. If successful a non-empty
// string template token is returned. Assumes that the scanners rune array
// length is greater than 0.
func (s *scanner) scanStrTemplate() (_ token.Token) {

	const (
		START_RUNE  = token.TERMINAL_TEMPLATE_START
		END_RUNE    = token.TERMINAL_TEMPLATE_END
		ESCAPE_RUNE = token.TERMINAL_TEMPLATE_ESCAPE
	)

	panic_unterminatedString := func() {
		panic(bard.NewTerror(s.line, s.col, nil,
			"Unterminated string template",
		))
	}

	var prev rune

	n := s.howManyRunesUntil(0, func(i int, ru rune) bool {

		switch {
		case i == 0 && ru != START_RUNE:
			return true
		case i == 0 && ru == END_RUNE:
			return false
		case prev != ESCAPE_RUNE && ru == END_RUNE:
			return true
		case s.matchesNewline(i):
			panic_unterminatedString()
		}

		return false
	})

	if n == 0 {
		return
	}

	if n == s.len() {
		panic_unterminatedString()
	}

	return s.tokenize(n+1, token.TEMPLATE, false)
}

// scanWord attempts to scan a keyword or identifier. If successful a non-empty
// keyword or identifier token is returned.
func (s *scanner) scanWord() (_ token.Token) {

	const UNDERSCORE_RUNE = token.TERMINAL_WORD_UNDERSCORE

	n := s.howManyRunesUntil(0, func(i int, ru rune) bool {

		if i == 0 && ru == UNDERSCORE_RUNE {
			return true
		}

		return ru != UNDERSCORE_RUNE && !unicode.IsLetter(ru)
	})

	if n == 0 {
		return
	}

	k := identifyAskeywordOrID(s.runes[:n])
	return s.tokenize(n, k, false)
}

// scanSymbol attempts to scan a symbol token. If successful a non-empty
// symbol token is returned. Assumes that the scanners rune array length is
// greater than 0.
func (s *scanner) scanSymbol() (_ token.Token) {

	var n int
	var k token.Kind

	if s.empty() {
		return
	}

	switch { // The order matters! This might be best moved to the token package.
	case s.matchesNonTerminal(0, token.NON_TERMINAL_ASSIGNMENT):
		n, k = 2, token.KIND_ASSIGN
	case s.matchesNonTerminal(0, token.NON_TERMINAL_RETURN_PARAMS):
		n, k = 2, token.KIND_RETURNS
	case s.matchesTerminal(0, token.TERMINAL_OPEN_PAREN):
		n, k = 1, token.KIND_OPEN_PAREN
	case s.matchesTerminal(0, token.TERMINAL_CLOSE_PAREN):
		n, k = 1, token.KIND_CLOSE_PAREN
	case s.matchesTerminal(0, token.TERMINAL_OPEN_GUARD):
		n, k = 1, token.KIND_OPEN_GUARD
	case s.matchesTerminal(0, token.TERMINAL_CLOSE_GUARD):
		n, k = 1, token.KIND_CLOSE_GUARD
	case s.matchesTerminal(0, token.TERMINAL_OPEN_LIST):
		n, k = 1, token.KIND_OPEN_LIST
	case s.matchesTerminal(0, token.TERMINAL_CLOSE_LIST):
		n, k = 1, token.KIND_CLOSE_LIST
	case s.matchesTerminal(0, token.TERMINAL_LIST_DELIM):
		n, k = 1, token.KIND_DELIM
	case s.matchesTerminal(0, token.TERMINAL_VOID_VALUE):
		n, k = 1, token.VOID
	case s.matchesTerminal(0, token.TERMINAL_STATEMENT_TERMINATOR):
		n, k = 1, token.TERMINATOR
	case s.matchesTerminal(0, token.TERMINAL_SPELL_PREFIX):
		n, k = 1, token.SPELL
	case s.matchesTerminal(0, token.TERMINAL_UNIVERSAL_NEGATION):
		n, k = 1, token.NOT
	case s.matchesTerminal(0, token.TERMINAL_TEA_DRINKING_NEGATION):
		n, k = 1, token.NOT
	case s.matchesTerminal(0, token.TERMINAL_MATH_ADDITION):
		n, k = 1, token.ADD
	case s.matchesTerminal(0, token.TERMINAL_MATH_SUBTRACTION):
		n, k = 1, token.SUBTRACT
	case s.matchesTerminal(0, token.TERMINAL_MATH_MULTIPLICATION):
		n, k = 1, token.MULTIPLY
	case s.matchesTerminal(0, token.TERMINAL_MATH_DIVISION):
		n, k = 1, token.DIVIDE
	case s.matchesTerminal(0, token.TERMINAL_MATH_REMAINDER):
		n, k = 1, token.MOD
	case s.matchesTerminal(0, token.TERMINAL_LOGICAL_AND):
		n, k = 1, token.AND
	case s.matchesTerminal(0, token.TERMINAL_LOGICAL_OR):
		n, k = 1, token.OR
	case s.matchesTerminal(0, token.TERMINAL_EQUALITY):
		n, k = 1, token.EQU
	case s.matchesTerminal(0, token.TERMINAL_UNEQUALITY):
		n, k = 1, token.NEQ
	case s.matchesNonTerminal(0, token.NON_TERMINAL_LESS_THAN_OR_EQUAL):
		n, k = 2, token.LT_OR_EQU
	case s.matchesNonTerminal(0, token.NON_TERMINAL_GREATER_THAN_OR_EQUAL):
		n, k = 2, token.MT_OR_EQU
	case s.matchesTerminal(0, token.TERMINAL_LESS_THAN):
		n, k = 1, token.LT
	case s.matchesTerminal(0, token.TERMINAL_MORE_THAN):
		n, k = 1, token.MT
	}

	if k == token.KIND_UNDEFINED {
		return
	}

	return s.tokenize(n, k, false)
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

	if k == token.KIND_UNDEFINED {
		return token.KIND_ID
	}

	return k
}
