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
			Kind: token.KIND_EOF,
			Line: s.line,
			Col:  s.col,
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
		return s.tokenize(n, token.KIND_NEWLINE, true)
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
	return s.tokenize(n, token.KIND_WHITESPACE, false)
}

func (s *scanner) scanComment() (_ token.Token) {

	const (
		COMMENT_PREFIX     = token.LEXEME_COMMENT_START
		COMMENT_PREFIX_LEN = len(COMMENT_PREFIX)
	)

	if s.doesNotMatchNonTerminal(0, COMMENT_PREFIX) {
		return
	}

	n := s.runesUntilNewline(COMMENT_PREFIX_LEN)
	return s.tokenize(n, token.KIND_COMMENT, false)
}

func (s *scanner) scanWord() (_ token.Token) {

	n := s.countWordRunes(0)

	if n == 0 {
		return
	}

	w := string(s.runes[:n])
	k := token.KeywordToKind(w)

	if k == token.KIND_UNDEFINED {
		k = token.KIND_ID
	}

	return s.tokenize(n, k, false)
}

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

func (s *scanner) scanNumberLiteral() (_ token.Token) {

	const (
		DELIM     = token.LEXEME_FRACTIONAL_DELIM
		DELIM_LEN = len(DELIM)
	)

	intLen := s.countDigitRunes(0)

	if intLen == 0 {
		return
	}

	if intLen == s.len() || s.doesNotMatchNonTerminal(intLen, DELIM) {
		return s.tokenize(intLen, token.INT, false)
	}

	fractionalLen := s.countDigitRunes(intLen + DELIM_LEN)

	if fractionalLen == 0 {
		panic(bard.NewTerror(s.line, s.col+intLen+DELIM_LEN, nil,
			"Expected digit after decimal point",
		))
	}

	n := intLen + DELIM_LEN + fractionalLen
	return s.tokenize(n, token.REAL, false)
}

func (s *scanner) scanStringLiteral() (_ token.Token) {

	const (
		PREFIX = token.LEXEME_STRING_START
		SUFFIX = token.LEXEME_STRING_END
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

	return s.tokenize(n+1, token.STR, false)
}

func (s *scanner) scanStringTemplate() (_ token.Token) {

	const (
		PREFIX        = token.LEXEME_TEMPLATE_START
		SUFFIX        = token.LEXEME_TEMPLATE_END
		SUFFIX_LEN    = len(SUFFIX)
		ESCAPE_SYMBOL = token.LEXEME_TEMPLATE_ESCAPE
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
	return s.tokenize(n, token.TEMPLATE, false)
}
