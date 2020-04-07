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

// scanner is a structure for parsing source code into tokens. It implements
// the TokenStream interface so it may be wrapped.
type scanner struct {
	symbolStream
}

// Next satisfies the TokenStream interface.
func (sc *scanner) Next() token.Token {

	if sc.empty() {
		return token.Token{
			Lexeme: token.LEXEME_EOF,
			Line:   sc.line,
			Col:    sc.col,
		}
	}

	fs := []scanFunc{
		sc.scanNewline, // LF & CRLF
		sc.scanWhitespace,
		sc.scanComment,
		sc.scanWord,   // Identifiers & keywords
		sc.scanSymbol, // :=, +, <, etc
		sc.scanNumberLiteral,
		sc.scanStringLiteral,  // `literal`
		sc.scanStringTemplate, // "Template: {identifier}"
	}

	for _, f := range fs {
		if tk, match := f(); match {
			return tk
		}
	}

	panic(bard.NewTerror(sc.line, sc.col, nil,
		"Could not identify next token",
	))
}

// scanFunc is the common signiture used by every scanning function that
// follows. If a concrete scanning function finds a match it must return a
// non-zero token and 'true' else it must return a zero token and 'false'.
type scanFunc func() (token.Token, bool)

func (sc *scanner) scanNewline() (_ token.Token, _ bool) {

	if n := sc.countNewlineRunes(0); n > 0 {
		tk := sc.tokenize(n, token.LEXEME_NEWLINE, true)
		return tk, true
	}

	return
}

func (sc *scanner) scanWhitespace() (_ token.Token, _ bool) {

	isNotSpace := func(i int, ru rune) bool {
		return sc.matchesNewline(i) || !unicode.IsSpace(ru)
	}

	if n := sc.howManyRunesUntil(0, isNotSpace); n > 0 {
		tk := sc.tokenize(n, token.LEXEME_WHITESPACE, false)
		return tk, true
	}

	return
}

func (sc *scanner) scanComment() (_ token.Token, _ bool) {

	const (
		COMMENT_PREFIX     = token.SYMBOL_COMMENT_START
		COMMENT_PREFIX_LEN = len(COMMENT_PREFIX)
	)

	if sc.matchesNonTerminal(0, COMMENT_PREFIX) {
		n := sc.runesUntilNewline(COMMENT_PREFIX_LEN)
		tk := sc.tokenize(n, token.LEXEME_COMMENT, false)
		return tk, true
	}

	return
}

func (sc *scanner) scanWord() (_ token.Token, _ bool) {

	n := sc.countWordRunes(0)

	if n == 0 {
		return
	}

	w := string(sc.runes[:n])

	for _, kw := range token.Keywords() {
		if kw.Symbol == w {
			tk := sc.tokenize(n, kw.Lexeme, false)
			return tk, true
		}
	}

	tk := sc.tokenize(n, token.LEXEME_ID, false)
	return tk, true
}

func (sc *scanner) scanSymbol() (_ token.Token, _ bool) {

	if sc.empty() {
		return
	}

	size := sc.len()

	for _, sym := range token.LoneSymbols() {

		if size < sym.Len {
			continue
		}

		if sc.matchesNonTerminal(0, sym.Symbol) {
			tk := sc.tokenize(sym.Len, sym.Lexeme, false)
			return tk, true
		}
	}

	return
}

func (sc *scanner) scanNumberLiteral() (_ token.Token, _ bool) {

	const (
		DELIM     = token.SYMBOL_FRACTIONAL_DELIM
		DELIM_LEN = len(DELIM)
	)

	intLen := sc.countDigitRunes(0)

	if intLen == 0 {
		// If there are no digits then this is not a number.
		return
	}

	if intLen == sc.len() || sc.doesNotMatchNonTerminal(intLen, DELIM) {
		// If this is the last token in the scanner or the next terminal is not the
		// delimiter between a floats integral and fractional parts then it must be
		// an integral.
		tk := sc.tokenize(intLen, token.LEXEME_INT, false)
		return tk, true
	}

	fractionalLen := sc.countDigitRunes(intLen + DELIM_LEN)

	if fractionalLen == 0 {
		// One or many fractional digits must follow a delimiter. Zero following
		// digits is invalid syntax, so we must panic.
		panic(bard.NewTerror(sc.line, sc.col+intLen+DELIM_LEN, nil,
			"Invalid syntax, expected digit after decimal point",
		))
	}

	n := intLen + DELIM_LEN + fractionalLen
	tk := sc.tokenize(n, token.LEXEME_FLOAT, false)
	return tk, true
}

func (sc *scanner) scanStringLiteral() (_ token.Token, _ bool) {

	const (
		PREFIX = token.STRING_SYMBOL_START
		SUFFIX = token.STRING_SYMBOL_END
	)

	n := sc.howManyRunesUntil(0, func(i int, _ rune) bool {

		switch {
		case i == 0:
			// If the initial terminals are not signify a string literal then exit
			// straight away.
			return sc.doesNotMatchNonTerminal(i, PREFIX)
		case sc.matchesNonTerminal(i, SUFFIX):
			// If
			return true
		case sc.matchesNewline(i):
			panic(bard.NewTerror(sc.line, sc.col, nil,
				"Newline encountered before a string literal was terminated",
			))
		case i+1 == sc.len():
			panic(bard.NewTerror(sc.line, sc.col, nil,
				"EOF encountered before a string literal was terminated",
			))
		}

		return false
	})

	if n == 0 {
		return
	}

	tk := sc.tokenize(n+1, token.LEXEME_STRING, false)
	return tk, true
}

func (sc *scanner) scanStringTemplate() (_ token.Token, _ bool) {

	const (
		PREFIX        = token.TEMPLATE_SYMBOL_START
		SUFFIX        = token.TEMPLATE_SYMBOL_END
		SUFFIX_LEN    = len(SUFFIX)
		ESCAPE_SYMBOL = token.TEMPLATE_SYMBOL_ESCAPE
	)

	var prevEscaped bool

	n := sc.howManyRunesUntil(0, func(i int, _ rune) bool {

		escaped := prevEscaped
		prevEscaped = false

		switch {
		case i == 0:
			return sc.doesNotMatchNonTerminal(i, PREFIX)
		case sc.matchesNonTerminal(i, ESCAPE_SYMBOL):
			prevEscaped = true
			return false
		case !escaped && sc.matchesNonTerminal(i, SUFFIX):
			return true
		case sc.matchesNewline(i):
			panic(bard.NewTerror(sc.line, sc.col, nil,
				"Newline encountered before a string template was terminated",
			))
		case i+1 == sc.len():
			panic(bard.NewTerror(sc.line, sc.col, nil,
				"EOF encountered before a string template was terminated",
			))
		}

		return false
	})

	if n == 0 {
		return
	}

	n += SUFFIX_LEN
	tk := sc.tokenize(n, token.LEXEME_TEMPLATE, false)
	return tk, true
}
