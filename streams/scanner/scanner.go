// scanner package was created to handle scanning of tokens from a script at a
// high level; low level aspects live in the streams/symbol package.
//
// Key decisions:
// 1. This could be rewritten to be much more performant, but I decided that
// a focus on readability was more important. Also, each script is only scanned
// once per execution so optimisation will probably not have any meaningful
// effect.
// 2. The terminal symbols used to represent various tokens have been separated
// into the lexeme package, even though the ordering of scanning functions
// depend on lexeme representations. This was conscious. It may not be possible
// to effectivly separate them.
//
// This package is responsible for scanning scripts only, evaluation is
// performed by the streams/evaluator package.
//
// TODO: Error handling needs to be simplified and the logger renamed to
// 			 something more meaningful.
// TODO: Some of these functions could probably do with being rewritten for
//       greater clarity.
package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/symbol"
	"github.com/PaulioRandall/scarlet-go/streams/token"
)

// ScanAll creates a scanner from s and reads all tokens from it into an array.
func ScanAll(s string) []lexeme.Token {
	sc := New(s)
	return token.ReadAll(sc)
}

// Scanner is a TokenStream providing functionality for parsing written scripts
// into a sequence of tokens. It provides the high level scanning code whilst
// the low level has been embedded.
type Scanner struct {
	symbol.SymbolStream
}

// New creates a new token Scanner as a TokenStream.
func New(s string) token.TokenStream {
	return &Scanner{
		symbol.NewSymbolStream(s),
	}
}

// scanFunc is the common signiture used by every scanning function used by
// the Read function. If a concrete scanning function finds a match it must
// return a non-zero token and true as the match else it must return a zero
// token along with false.
type scanFunc func() (lexeme.Token, bool)

// Read satisfies the TokenStream interface.
func (sc *Scanner) Read() lexeme.Token {

	if sc.Empty() {
		// TokenStream.Read requires an EOF token be returned upon an empty stream.
		return lexeme.Token{
			Lexeme: lexeme.LEXEME_EOF,
			Line:   sc.LineIndex(),
			Col:    sc.ColIndex(),
		}
	}

	// For proper parsing, the correct ordering of scanning functions may be
	// important depending on the lexeme representations. These can be found in
	// the lexeme package. Currently the followingering is required:
	// 1. scanComment before scanSymbol
	// 2. scanSymbol before scanWord
	fs := []scanFunc{
		sc.scanNewline,    // LF & CRLF
		sc.scanWhitespace, // Any whitespace except newlines
		sc.scanComment,
		sc.scanSymbol,         // :=, +, <, etc
		sc.scanWord,           // Identifiers & keywords
		sc.scanNumberLiteral,  // Ints & floats
		sc.scanStringLiteral,  // `literal`
		sc.scanStringTemplate, // "Template: {identifier}"
	}

	// Iterate each function list calling one at a time until one of them
	// identifies and parses the token.
	for _, f := range fs {
		if tk, match := f(); match {
			return tk
		}
	}

	panic(sc.terror(0, "Could not identify next token"))
}

// scanNewline implements the scanFunc signiture to return a non-zero token if
// a newline appears next within the script, i.e. LF or CRLF.
func (sc *Scanner) scanNewline() (_ lexeme.Token, _ bool) {

	if n := sc.CountNewlineSymbols(0); n > 0 {
		tk := sc.tokenize(n, lexeme.LEXEME_NEWLINE)
		return tk, true
	}

	return
}

// scanWhitespace implements the scanFunc signiture to return a non-zero token
// if one or many consecutive whitespace terminals appear next within the
// script. Newlines are not counted as whitespace, instead they are scanned by
// the scanNewline function.
func (sc *Scanner) scanWhitespace() (_ lexeme.Token, _ bool) {

	isSpace := func(i int, ru rune) bool {
		return !sc.IsNewline(i) && unicode.IsSpace(ru)
	}

	if n := sc.CountSymbolsWhile(0, isSpace); n > 0 {
		tk := sc.tokenize(n, lexeme.LEXEME_WHITESPACE)
		return tk, true
	}

	return
}

// scanComment implements the scanFunc signiture to return a non-zero token if
// an inline comment appears next within the script. Inline comments always
// include all terminals within the remainder of the line.
func (sc *Scanner) scanComment() (_ lexeme.Token, _ bool) {

	const (
		COMMENT_PREFIX     = lexeme.SYMBOL_COMMENT_START
		COMMENT_PREFIX_LEN = len(COMMENT_PREFIX)
	)

	if sc.IsMatch(0, COMMENT_PREFIX) {
		n := sc.IndexOfNextNewline(COMMENT_PREFIX_LEN)
		tk := sc.tokenize(n, lexeme.LEXEME_COMMENT)
		return tk, true
	}

	return
}

// scanSymbol implements the scanFunc signiture to return a non-zero token if
// a simple math, logic, or control symbol (lone symbol, non-terminal) appears
// next within the script, i.e. :=, +, <, etc.
func (sc *Scanner) scanSymbol() (_ lexeme.Token, _ bool) {

	for _, sym := range lexeme.Symbols() {

		if sc.Len() < sym.Len {
			// Ignore symbols that are shorter than the stream as they won't match
			// but will cause an out of range panic.
			continue
		}

		if sc.IsMatch(0, sym.Symbol) {
			tk := sc.tokenize(sym.Len, sym.Lexeme)
			return tk, true
		}
	}

	return
}

// scanWord implements the scanFunc signiture to return a non-zero token if
// either a keyword or identifier appear next within the script.
func (sc *Scanner) scanWord() (tk lexeme.Token, _ bool) {

	n := sc.CountSymbolsWhile(0, func(_ int, ru rune) bool {
		return lexeme.IsWordTerminal(ru)
	})

	if n == 0 {
		return
	}

	w := sc.Peek(n)
	lex := lexeme.FindKeywordLexeme(w)

	if lex != lexeme.LEXEME_UNDEFINED {
		tk = sc.tokenize(n, lex)
	} else {
		// Any word that is not a keyword (reserved word) must be an identifier.
		tk = sc.tokenize(n, lexeme.LEXEME_ID)
	}

	return tk, true
}

// scanNumberLiteral implements the scanFunc signiture to return a non-zero
// token if an int of float appear next within the script. Both integers and
// floats are parsed by this function.
func (sc *Scanner) scanNumberLiteral() (_ lexeme.Token, _ bool) {

	const (
		DELIM     = lexeme.SYMBOL_FRACTIONAL_DELIM
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
		// This must be an int if this is the last token in the scanner or the next
		// terminal is not the delimiter between a floats integral and fractional
		// parts.
		tk := sc.tokenize(intLen, lexeme.LEXEME_INT)
		return tk, true
	}

	fractionalLen := sc.CountSymbolsWhile(intLen+DELIM_LEN, isDigit)

	if fractionalLen == 0 {
		// One or many fractional digits must follow a delimiter. Zero following
		// digits is invalid syntax, so we must panic.
		panic(sc.terror(intLen+DELIM_LEN,
			"Invalid syntax, expected digit after decimal point",
		))
	}

	n := intLen + DELIM_LEN + fractionalLen
	tk := sc.tokenize(n, lexeme.LEXEME_FLOAT)
	return tk, true
}

// scanStringLiteral implements the scanFunc signiture to return a non-zero
// token if a literal string appears next within the script.
func (sc *Scanner) scanStringLiteral() (_ lexeme.Token, _ bool) {

	const (
		PREFIX = lexeme.STRING_SYMBOL_START
		SUFFIX = lexeme.STRING_SYMBOL_END
	)

	n := sc.CountSymbolsWhile(0, func(i int, _ rune) bool {

		if i == 0 {
			// If the initial terminal does not signify a string literal then exit
			// straight away, n will be 0.
			return sc.IsMatch(i, PREFIX)
		}

		switch { // We must panic if we can't find the end of the string.
		case sc.IsMatch(i, SUFFIX):
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

	tk := sc.tokenize(n+1, lexeme.LEXEME_STRING)
	return tk, true
}

// scanStringTemplate implements the scanFunc signiture to return a non-zero
// token if a template string appears next within the script. As the name
// suggests, templates can be populated with the value of identifiers, but the
// scanner is not concerned with parsing these. It does need to watch out for
// escaped terminals that also represent the string closer (suffix).
func (sc *Scanner) scanStringTemplate() (_ lexeme.Token, _ bool) {

	const (
		PREFIX        = lexeme.TEMPLATE_SYMBOL_START
		SUFFIX        = lexeme.TEMPLATE_SYMBOL_END
		SUFFIX_LEN    = len(SUFFIX)
		ESCAPE_SYMBOL = lexeme.TEMPLATE_SYMBOL_ESCAPE
	)

	var prevEscaped bool

	n := sc.CountSymbolsWhile(0, func(i int, _ rune) bool {

		escaped := prevEscaped
		prevEscaped = false

		if i == 0 {
			// If the initial terminal does not signify a string template then exit
			// straight away, n will be 0.
			return sc.IsMatch(i, PREFIX)
		}

		switch { // We must panic if we can't find the end of the template.
		case sc.IsMatch(i, ESCAPE_SYMBOL):
			prevEscaped = true
			return true
		case !escaped && sc.IsMatch(i, SUFFIX): // Ensure the suffix is not escaped
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
	tk := sc.tokenize(n, lexeme.LEXEME_TEMPLATE)
	return tk, true
}

// tokenize creates a new token from the next non-terminal. It reads off n
// symbols from the embedded SymbolStream ready for scanning the next token.
func (sc *Scanner) tokenize(n int, lex lexeme.Lexeme) lexeme.Token {

	tk := lexeme.Token{
		Lexeme: lex,
		Line:   sc.LineIndex(),
		Col:    sc.ColIndex(),
	}

	tk.Value = sc.SymbolStream.Read(n, lex == lexeme.LEXEME_NEWLINE)

	return tk
}

// terror was created because I am lazy. It will probably be removed when I
// update the error handling.
func (sc *Scanner) terror(colOffset int, msg string) bard.Terror {
	return bard.NewTerror(
		sc.LineIndex(),
		sc.ColIndex()+colOffset,
		nil,
		msg,
	)
}
