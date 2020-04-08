package evaluator

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"
)

type evaluator struct {
	ts   token.TokenStream
	prev lexeme.Lexeme
}

func New(delegate token.TokenStream) token.TokenStream {
	return &evaluator{
		ts: delegate,
	}
}

// Next satisfies the TokenStream interface.
func (ev *evaluator) Next() (_ lexeme.Token) {

	var tk lexeme.Token
	var lex lexeme.Lexeme
	prev := ev.prev

	for tk = ev.ts.Next(); tk != (lexeme.Token{}); tk = ev.ts.Next() {

		lex = tk.Lexeme

		if lex == lexeme.LEXEME_WHITESPACE || lex == lexeme.LEXEME_COMMENT {
			continue
		}

		switch prev {
		case lexeme.LEXEME_OPEN_LIST, lexeme.LEXEME_DELIM, lexeme.LEXEME_TERMINATOR:
			fallthrough
		case lexeme.LEXEME_DO, lexeme.LEXEME_UNDEFINED:
			if lex == lexeme.LEXEME_NEWLINE {
				continue
			}
		}

		if lex == lexeme.LEXEME_NEWLINE {
			tk.Lexeme = lexeme.LEXEME_TERMINATOR
		}

		if lex == lexeme.LEXEME_STRING || lex == lexeme.LEXEME_TEMPLATE {
			trimStrQuotes(&tk)
		}

		ev.prev = tk.Lexeme
		return tk
	}

	return
}

// trimStrQuotes removes the leading and trailing quotes from string literals
// and templates.
func trimStrQuotes(tk *lexeme.Token) {
	s := tk.Value
	tk.Value = s[1 : len(s)-1]
}
