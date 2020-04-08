package token

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// TokenStream provides access to an ordered stream of tokens.
type TokenStream interface {

	// Read returns the next token in the stream. An EOF token is always returned
	// if the stream is empty.
	Read() lexeme.Token
}

// impl is a simple TokenStream implementation which reads from a fixed slice.
type impl struct {
	tokens []lexeme.Token
	index  int
	prev   lexeme.Token
}

// New creates a simple token stream which reads from a fixed slice.
func New(tokens []lexeme.Token) TokenStream {
	return &impl{
		tokens: tokens,
	}
}

// Read satisfies the TokenStream interface.
func (ts *impl) Read() lexeme.Token {

	if ts.prev.Lexeme == lexeme.LEXEME_EOF {
		return ts.prev
	}

	if ts.index >= len(ts.tokens) {
		ts.prev = lexeme.Token{
			Lexeme: lexeme.LEXEME_EOF,
			Line:   -1,
			Col:    -1,
		}

		return ts.prev
	}

	ts.prev = ts.tokens[ts.index]
	ts.index++

	return ts.prev
}

// ReadAll reads all tokens from ts into an array.
func ReadAll(ts TokenStream) []lexeme.Token {

	var tk lexeme.Token
	var tokens []lexeme.Token

	for tk = ts.Read(); tk.Lexeme != lexeme.LEXEME_EOF; tk = ts.Read() {
		tokens = append(tokens, tk)
	}

	return append(tokens, tk)
}
