// scanner package was created to handle scanning of tokens from a script at a
// high level; low level aspects live in the streams/symbol package.
//
// Key decisions:
// 1. This could be rewritten to be much more performant, but I decided that
// a focus on readability was more important. Also, each script is only scanned
// once per execution so optimisation will probably not have any meaningful
// effect.
//
// This package is responsible for scanning scripts only, evaluation is
// performed by the streams/evaluator package.
//
// TODO: Error handling needs to be simplified and the logger renamed to
// 			 something more meaningful.
package scanner

import (
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

	tk := sc.parseNextToken()

	if tk == (lexeme.Token{}) {
		panic(terror(sc, 0, "Could not identify next token"))
	}

	return tk
}

// parseNextToken attempts to match one of the non-terminal patterns to the next
// set of terminals in the script. If found, the terminals are removed and used
// to create a token.
func (sc *Scanner) parseNextToken() (_ lexeme.Token) {

	matchers := nonTerminals()
	size := len(matchers)

	for i := 0; i < size; i++ {

		nonTerminal := matchers[i]
		n := nonTerminal.Matcher(symbol.SymbolStream(sc))

		if n > 0 {
			return sc.tokenize(n, nonTerminal.Lexeme)
		}
	}

	return
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
