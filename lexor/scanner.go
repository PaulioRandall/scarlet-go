package lexor

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/lexor/symbol"
	"github.com/PaulioRandall/scarlet-go/token"
)

/******************************************************************************
File:
	Contains functions used in parsing a written script into tokens.
******************************************************************************/

// scanner is a structure for parsing source code into tokens. It implements
// the TokenStream interface so it may be wrapped.
type scanner struct {
	symbol.SymbolStream
}

// Next satisfies the TokenStream interface.
func (sc *scanner) Next() token.Token {

	if sc.Empty() {
		return token.Token{
			Lexeme: token.LEXEME_EOF,
			Line:   sc.LineIndex(),
			Col:    sc.ColIndex(),
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

	panic(sc.terror(0, "Could not identify next token"))
}

// scanFunc is the common signiture used by every scanning function that
// follows. If a concrete scanning function finds a match it must return a
// non-zero token and 'true' else it must return a zero token and 'false'.
type scanFunc func() (token.Token, bool)

func (sc *scanner) scanNewline() (_ token.Token, _ bool) {

	if n := sc.CountNewlineSymbols(0); n > 0 {
		tk := sc.tokenize(n, token.LEXEME_NEWLINE)
		return tk, true
	}

	return
}

func (sc *scanner) scanWhitespace() (_ token.Token, _ bool) {

	isSpace := func(i int, ru rune) bool {
		return !sc.IsNewline(i) && unicode.IsSpace(ru)
	}

	if n := sc.CountSymbolsWhile(0, isSpace); n > 0 {
		tk := sc.tokenize(n, token.LEXEME_WHITESPACE)
		return tk, true
	}

	return
}

func (sc *scanner) scanComment() (_ token.Token, _ bool) {

	const (
		COMMENT_PREFIX     = token.SYMBOL_COMMENT_START
		COMMENT_PREFIX_LEN = len(COMMENT_PREFIX)
	)

	if sc.IsMatch(0, COMMENT_PREFIX) {
		n := sc.IndexOfNextNewline(COMMENT_PREFIX_LEN)
		tk := sc.tokenize(n, token.LEXEME_COMMENT)
		return tk, true
	}

	return
}

func (sc *scanner) scanWord() (_ token.Token, _ bool) {

	n := sc.CountSymbolsWhile(0, func(i int, ru rune) bool {
		return ru == '_' || unicode.IsLetter(ru)
	})

	if n == 0 {
		return
	}

	w := sc.Peek(n)

	if w[0] == '_' {
		if len(w) == 1 {
			return
		}

		panic(sc.terror(0, `Identifiers may not start with an underscore`))
	}

	for _, kw := range token.Keywords() {
		if kw.Symbol == w {
			tk := sc.tokenize(n, kw.Lexeme)
			return tk, true
		}
	}

	tk := sc.tokenize(n, token.LEXEME_ID)
	return tk, true
}

func (sc *scanner) scanSymbol() (_ token.Token, _ bool) {

	if sc.Empty() {
		return
	}

	size := sc.Len()

	for _, sym := range token.LoneSymbols() {

		if size < sym.Len {
			continue
		}

		if sc.IsMatch(0, sym.Symbol) {
			tk := sc.tokenize(sym.Len, sym.Lexeme)
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

	isDigit := func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	}

	intLen := sc.CountSymbolsWhile(0, isDigit)

	if intLen == 0 {
		// If there are no digits then this is not a number.
		return
	}

	if intLen == sc.Len() || !sc.IsMatch(intLen, DELIM) {
		// If this is the last token in the scanner or the next terminal is not the
		// delimiter between a floats integral and fractional parts then it must be
		// an integral.
		tk := sc.tokenize(intLen, token.LEXEME_INT)
		return tk, true
	}

	fractionalLen := sc.CountSymbolsWhile(intLen+DELIM_LEN, isDigit)

	if fractionalLen == 0 {
		// One or many fractional digits must follow a delimiter. Zero following
		// digits is invalid syntax, so we must panic.
		panic(sc.terror(
			intLen+DELIM_LEN,
			"Invalid syntax, expected digit after decimal point",
		))
	}

	n := intLen + DELIM_LEN + fractionalLen
	tk := sc.tokenize(n, token.LEXEME_FLOAT)
	return tk, true
}

func (sc *scanner) scanStringLiteral() (_ token.Token, _ bool) {

	const (
		PREFIX = token.STRING_SYMBOL_START
		SUFFIX = token.STRING_SYMBOL_END
	)

	n := sc.CountSymbolsWhile(0, func(i int, _ rune) bool {

		switch {
		case i == 0:
			// If the initial terminals are not signify a string literal then exit
			// straight away.
			return sc.IsMatch(i, PREFIX)
		case sc.IsMatch(i, SUFFIX):
			// If
			return false
		case sc.IsNewline(i):
			panic(sc.terror(0,
				"Newline encountered before a string literal was terminated",
			))
		case i+1 == sc.Len():
			panic(sc.terror(0,
				"EOF encountered before a string literal was terminated",
			))
		}

		return true
	})

	if n == 0 {
		return
	}

	tk := sc.tokenize(n+1, token.LEXEME_STRING)
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

	n := sc.CountSymbolsWhile(0, func(i int, _ rune) bool {

		escaped := prevEscaped
		prevEscaped = false

		switch {
		case i == 0:
			return sc.IsMatch(i, PREFIX)
		case sc.IsMatch(i, ESCAPE_SYMBOL):
			prevEscaped = true
			return true
		case !escaped && sc.IsMatch(i, SUFFIX):
			return false
		case sc.IsNewline(i):
			panic(sc.terror(0,
				"Newline encountered before a string template was terminated",
			))
		case i+1 == sc.Len():
			panic(sc.terror(0,
				"EOF encountered before a string template was terminated",
			))
		}

		return true
	})

	if n == 0 {
		return
	}

	n += SUFFIX_LEN
	tk := sc.tokenize(n, token.LEXEME_TEMPLATE)
	return tk, true
}

func (sc *scanner) tokenize(runeCount int, lex token.Lexeme) token.Token {

	tk := token.Token{
		Lexeme: lex,
		Line:   sc.LineIndex(),
		Col:    sc.ColIndex(),
	}

	tk.Value = sc.Read(runeCount, lex == token.LEXEME_NEWLINE)

	return tk
}

func (sc *scanner) terror(colOffset int, msg string) bard.Terror {
	return bard.NewTerror(
		sc.LineIndex(),
		sc.ColIndex()+colOffset,
		nil,
		msg,
	)
}
