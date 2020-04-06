package lexor

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// ****************************************************************************
// * evaluator
// ****************************************************************************

// evaluator is a structure for parsed source code tokens. This involves
// removing quotes from string literals and discarding whitespace tokens etc.
// It requires a TokenStream as a source for tokens and implements the
// TokenStream interface so it may be wrapped.
type evaluator struct {
	ts   TokenStream
	prev token.Lexeme
}

// NewEvaluator creates a new evaluator to evaluate tokens within a stream.
func NewEvaluator(delegate TokenStream) TokenStream {
	return &evaluator{
		ts: delegate,
	}
}

// Next satisfies the TokenStream interface.
func (ev *evaluator) Next() (_ token.Token) {

	var tk token.Token
	var lex token.Lexeme
	prev := ev.prev

	for tk = ev.ts.Next(); tk != (token.Token{}); tk = ev.ts.Next() {

		lex = tk.Lexeme

		if lex == token.LEXEME_WHITESPACE || lex == token.LEXEME_COMMENT {
			continue
		}

		switch prev {
		case token.LEXEME_OPEN_LIST, token.LEXEME_DELIM, token.LEXEME_TERMINATOR:
			fallthrough
		case token.LEXEME_DO, token.LEXEME_UNDEFINED:
			if lex == token.LEXEME_NEWLINE {
				continue
			}
		}

		if lex == token.LEXEME_NEWLINE {
			tk.Lexeme = token.LEXEME_TERMINATOR
		}

		if lex == token.LEXEME_STRING || lex == token.LEXEME_TEMPLATE {
			trimStrQuotes(&tk)
		}

		ev.prev = tk.Lexeme
		return tk
	}

	return
}

// trimStrQuotes removes the leading and trailing quotes from string literals
// and templates.
func trimStrQuotes(tk *token.Token) {
	s := tk.Value
	tk.Value = s[1 : len(s)-1]
}
