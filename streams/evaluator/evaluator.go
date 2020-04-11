// evaluator package was created to handle evaluation of scanned script tokens.
// Evaluation involves removing redundant tokens, such as comment and
// whitespace, and formatting values such as trimming off the quotes from
// string literals and templates.
//
// Key decisions: N/A
package evaluator

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"
)

// EvalAll creates an evaluator from s and evaluates all tokens from it into a
// new array.
func EvalAll(s []lexeme.Token) []lexeme.Token {
	ts := token.New(s)
	ev := New(ts)
	return token.ReadAll(ev)
}

// evaluator is a TokenStream providing functionality for evaluating a stream
// of tokens sourced from else where.
type evaluator struct {
	ts   token.TokenStream
	prev lexeme.Token // Storage for the previously returned token.
}

// New creates a new token evaluator as a TokenStream.
func New(delegate token.TokenStream) token.TokenStream {
	return &evaluator{
		ts: delegate,
	}
}

// Read satisfies the TokenStream interface.
func (ev *evaluator) Read() (_ lexeme.Token) {

	if ev.prev.Lexeme == lexeme.LEXEME_EOF {
		return ev.prev
	}

	tk := ev.readNextParsableToken()
	tk = formatToken(tk)

	ev.prev = tk
	return tk
}

// readNextParsableToken reads in tokens until a non-redundant one is found
// --The parser has no use for sugar and spice--.
func (ev *evaluator) readNextParsableToken() (tk lexeme.Token) {

	for tk = ev.ts.Read(); tk.Lexeme != lexeme.LEXEME_EOF; tk = ev.ts.Read() {
		if !isRedundantLexeme(tk.Lexeme, ev.prev.Lexeme) {
			break
		}
	}

	return tk
}

// isRedundantLexeme returns true if l is considered redundant to parsing.
func isRedundantLexeme(l, prev lexeme.Lexeme) bool {

	if l == lexeme.LEXEME_WHITESPACE || l == lexeme.LEXEME_COMMENT {
		return true
	}

	if l != lexeme.LEXEME_NEWLINE {
		return false
	}

	return prev == lexeme.LEXEME_OPEN_LIST ||
		prev == lexeme.LEXEME_DELIM ||
		prev == lexeme.LEXEME_TERMINATOR ||
		prev == lexeme.LEXEME_DO ||
		prev == lexeme.LEXEME_UNDEFINED
}

// Applies any special formatting to the token such as converting its lexeme
// type or trimming runes off its value.
func formatToken(tk lexeme.Token) lexeme.Token {

	switch tk.Lexeme {
	case lexeme.LEXEME_NEWLINE:
		// Non-redundant newline tokens are expression and statement terminators
		// in disguise.
		tk.Lexeme = lexeme.LEXEME_TERMINATOR

	case lexeme.LEXEME_STRING, lexeme.LEXEME_TEMPLATE:
		// Removes prefix and suffix from tk.Value
		s := tk.Value
		tk.Value = s[1 : len(s)-1]
	}

	return tk
}
