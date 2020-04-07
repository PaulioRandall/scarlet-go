package lexor

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

/******************************************************************************
File:
	Contains functions used in parsing a written script into tokens.
******************************************************************************/

// scanFunc is the common signiture used by every scanning function that
// follows. If a concrete scanning function finds a match it must return a
// non-zero token and 'true' else it must return a zero token and 'false'.
type scanFunc func(*scanner) (token.Token, bool)

func scanNewline(s *scanner) (_ token.Token, _ bool) {

	if n := s.countNewlineRunes(0); n > 0 {
		tk := s.tokenize(n, token.LEXEME_NEWLINE, true)
		return tk, true
	}

	return
}

func scanWhitespace(s *scanner) (_ token.Token, _ bool) {

	isNotSpace := func(i int, ru rune) bool {
		return s.matchesNewline(i) || !unicode.IsSpace(ru)
	}

	if n := s.howManyRunesUntil(0, isNotSpace); n > 0 {
		tk := s.tokenize(n, token.LEXEME_WHITESPACE, false)
		return tk, true
	}

	return
}

func scanComment(s *scanner) (tk token.Token, match bool) {

	const (
		COMMENT_PREFIX     = token.SYMBOL_COMMENT_START
		COMMENT_PREFIX_LEN = len(COMMENT_PREFIX)
	)

	if s.matchesNonTerminal(0, COMMENT_PREFIX) {
		n := s.runesUntilNewline(COMMENT_PREFIX_LEN)
		tk = s.tokenize(n, token.LEXEME_COMMENT, false)
		match = true
	}

	return
}

func scanWord(s *scanner) (_ token.Token, _ bool) {

	n := s.countWordRunes(0)

	if n == 0 {
		return
	}

	w := string(s.runes[:n])

	for _, kw := range token.Keywords() {
		if kw.Symbol == w {
			tk := s.tokenize(n, kw.Lexeme, false)
			return tk, true
		}
	}

	tk := s.tokenize(n, token.LEXEME_ID, false)
	return tk, true
}

func scanSymbol(s *scanner) (_ token.Token, _ bool) {

	if s.empty() {
		return
	}

	size := s.len()

	for _, sym := range token.LoneSymbols() {

		if size < sym.Len {
			continue
		}

		if s.matchesNonTerminal(0, sym.Symbol) {
			tk := s.tokenize(sym.Len, sym.Lexeme, false)
			return tk, true
		}
	}

	return
}

func scanNumberLiteral(s *scanner) (_ token.Token, _ bool) {

	const (
		DELIM     = token.SYMBOL_FRACTIONAL_DELIM
		DELIM_LEN = len(DELIM)
	)

	intLen := s.countDigitRunes(0)

	if intLen == 0 {
		// If there are no digits then this is not a number.
		return
	}

	if intLen == s.len() || s.doesNotMatchNonTerminal(intLen, DELIM) {
		// If this is the last token in the scanner or the next terminal is not the
		// delimiter between a floats integral and fractional parts then it must be
		// an integral.
		tk := s.tokenize(intLen, token.LEXEME_INT, false)
		return tk, true
	}

	fractionalLen := s.countDigitRunes(intLen + DELIM_LEN)

	if fractionalLen == 0 {
		// One or many fractional digits must follow a delimiter. Zero following
		// digits is invalid syntax, so we must panic.
		panic(bard.NewTerror(s.line, s.col+intLen+DELIM_LEN, nil,
			"Invalid syntax, expected digit after decimal point",
		))
	}

	n := intLen + DELIM_LEN + fractionalLen
	tk := s.tokenize(n, token.LEXEME_FLOAT, false)
	return tk, true
}

func scanStringLiteral(s *scanner) (_ token.Token, _ bool) {

	const (
		PREFIX = token.STRING_SYMBOL_START
		SUFFIX = token.STRING_SYMBOL_END
	)

	n := s.howManyRunesUntil(0, func(i int, _ rune) bool {

		switch {
		case i == 0:
			// If the initial terminals are not signify a string literal then exit
			// straight away.
			return s.doesNotMatchNonTerminal(i, PREFIX)
		case s.matchesNonTerminal(i, SUFFIX):
			// If
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

	tk := s.tokenize(n+1, token.LEXEME_STRING, false)
	return tk, true
}

func scanStringTemplate(s *scanner) (_ token.Token, _ bool) {

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
	tk := s.tokenize(n, token.LEXEME_TEMPLATE, false)
	return tk, true
}
