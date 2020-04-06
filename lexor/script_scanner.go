package lexor

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenStream represents a stream of tokens.
type TokenStream interface {

	// Next returns the next token in the stream. An EOF token is always returned
	// if the token stream is empty.
	Next() token.Token
}

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
			Lexeme: token.LEXEME_EOF,
			Line:   s.line,
			Col:    s.col,
		}
	}

	type scanFunc func() token.Token

	fs := []scanFunc{
		s.scanNewline, // LF & CRLF
		s.scanWhitespace,
		s.scanComment,
		s.scanWord,   // Identifiers & keywords
		s.scanSymbol, // :=, +, <, etc
		s.scanNumberLiteral,
		s.scanStringLiteral,  // `literal`
		s.scanStringTemplate, // "Template: {identifier}"
	}

	for _, f := range fs {
		if tk = f(); tk != (token.Token{}) {
			return
		}
	}

	println(string(s.runes))

	panic(bard.NewTerror(s.line, s.col, nil,
		"Could not identify next token",
	))
}

func (s *scanner) scanNewline() (_ token.Token) {

	if n := s.countNewlineRunes(0); n > 0 {
		return s.tokenize(n, token.LEXEME_NEWLINE, true)
	}

	return
}

func (s *scanner) scanWhitespace() (_ token.Token) {

	isSpace := func(i int, ru rune) bool {
		return s.matchesNewline(i) || !unicode.IsSpace(ru)
	}

	if isSpace(0, s.runes[0]) {
		return
	}

	n := s.howManyRunesUntil(0, isSpace)
	return s.tokenize(n, token.LEXEME_WHITESPACE, false)
}

func (s *scanner) scanComment() (_ token.Token) {

	const (
		COMMENT_PREFIX     = token.SYMBOL_COMMENT_START
		COMMENT_PREFIX_LEN = len(COMMENT_PREFIX)
	)

	if s.doesNotMatchNonTerminal(0, COMMENT_PREFIX) {
		return
	}

	n := s.runesUntilNewline(COMMENT_PREFIX_LEN)
	return s.tokenize(n, token.LEXEME_COMMENT, false)
}

func (s *scanner) scanWord() (_ token.Token) {

	n := s.countWordRunes(0)

	if n == 0 {
		return
	}

	w := string(s.runes[:n])

	for _, kw := range token.Keywords() {
		if kw.Symbol == w {
			return s.tokenize(n, kw.Lexeme, false)
		}
	}

	return s.tokenize(n, token.LEXEME_ID, false)
}

func (s *scanner) scanSymbol() (_ token.Token) {

	if s.empty() {
		return
	}

	size := s.len()

	for _, sym := range token.LoneSymbols() {

		if size < sym.Len {
			continue
		}

		if s.matchesNonTerminal(0, sym.Symbol) {
			return s.tokenize(sym.Len, sym.Lexeme, false)
		}
	}

	return
}

func (s *scanner) scanNumberLiteral() (_ token.Token) {

	const (
		DELIM     = token.SYMBOL_FRACTIONAL_DELIM
		DELIM_LEN = len(DELIM)
	)

	intLen := s.countDigitRunes(0)

	if intLen == 0 {
		return
	}

	if intLen == s.len() || s.doesNotMatchNonTerminal(intLen, DELIM) {
		return s.tokenize(intLen, token.LEXEME_INT, false)
	}

	fractionalLen := s.countDigitRunes(intLen + DELIM_LEN)

	if fractionalLen == 0 {
		panic(bard.NewTerror(s.line, s.col+intLen+DELIM_LEN, nil,
			"Expected digit after decimal point",
		))
	}

	n := intLen + DELIM_LEN + fractionalLen
	return s.tokenize(n, token.LEXEME_FLOAT, false)
}

func (s *scanner) scanStringLiteral() (_ token.Token) {

	const (
		PREFIX = token.STRING_SYMBOL_START
		SUFFIX = token.STRING_SYMBOL_END
	)

	n := s.howManyRunesUntil(0, func(i int, _ rune) bool {

		switch {
		case i == 0:
			return s.doesNotMatchNonTerminal(i, PREFIX)
		case s.matchesNonTerminal(i, SUFFIX):
			return true
		case s.matchesNewline(i):
			panic(bard.NewTerror(s.line, s.col, nil,
				"Newline encountered before a string literal was terminated",
			))
		case i+1 == s.len():
			panic(bard.NewTerror(s.line, s.col, nil,
				"EOF encountered before a string literal was terminated",
			))
		}

		return false
	})

	if n == 0 {
		return
	}

	return s.tokenize(n+1, token.LEXEME_STRING, false)
}

func (s *scanner) scanStringTemplate() (_ token.Token) {

	const (
		PREFIX        = token.TEMPLATE_SYMBOL_START
		SUFFIX        = token.TEMPLATE_SYMBOL_END
		SUFFIX_LEN    = len(SUFFIX)
		ESCAPE_SYMBOL = token.TEMPLATE_SYMBOL_ESCAPE
	)

	var prevEscaped bool

	n := s.howManyRunesUntil(0, func(i int, _ rune) bool {

		escaped := prevEscaped
		prevEscaped = false

		switch {
		case i == 0:
			return s.doesNotMatchNonTerminal(i, PREFIX)
		case s.matchesNonTerminal(i, ESCAPE_SYMBOL):
			prevEscaped = true
			return false
		case !escaped && s.matchesNonTerminal(i, SUFFIX):
			return true
		case s.matchesNewline(i):
			panic(bard.NewTerror(s.line, s.col, nil,
				"Newline encountered before a string template was terminated",
			))
		case i+1 == s.len():
			panic(bard.NewTerror(s.line, s.col, nil,
				"EOF encountered before a string template was terminated",
			))
		}

		return false
	})

	if n == 0 {
		return
	}

	n += SUFFIX_LEN
	return s.tokenize(n, token.LEXEME_TEMPLATE, false)
}
