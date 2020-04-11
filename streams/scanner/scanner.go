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
		sc.scanSymbol, // DO, END, :=, +, <, etc
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

	panic(terror(sc, 0, "Could not identify next token"))
}

// scanSymbol implements the scanFunc signiture to return a non-zero token if
// a simple math, logic, or control symbol (lone symbol, non-terminal) appears
// next within the script, i.e. :=, +, <, etc.
func (sc *Scanner) scanSymbol() (_ lexeme.Token, _ bool) {

	for _, nt := range nonTerminals() {

		matchLen := nt.Matcher(symbol.SymbolStream(sc))

		if matchLen > 0 {
			tk := sc.tokenize(matchLen, nt.Lexeme)
			return tk, true
		}
	}

	return
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
			panic(terror(sc, 0,
				"Newline encountered before a string literal was terminated",
			))
		case i+1 == sc.Len():
			panic(terror(sc, 0,
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
			panic(terror(sc, 0,
				"Newline encountered before a string template was terminated",
			))
		case i+1 == sc.Len():
			panic(terror(sc, 0,
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

	tk.Value = sc.SymbolStream.Slice(n, lex == lexeme.LEXEME_NEWLINE)

	return tk
}

// terror was created because I am lazy. It will probably be removed when I
// update the error handling.
func terror(ss symbol.SymbolStream, colOffset int, msg string) bard.Terror {
	return bard.NewTerror(
		ss.LineIndex(),
		ss.ColIndex()+colOffset,
		nil,
		msg,
	)
}
