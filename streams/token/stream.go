// token package was created separate the TokenStream interface and its default
// implementation from its more advanced implementations. The API is responsible
// for providing basic token streaming functionality for its users.
//
// Key decisions:
// 1. The simple slice-based implementation was included here because it is tiny
// and has no additional dependencies. Moving it to its own file or package felt
// unnecessary.
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

// impl is a simple TokenStream implementation which reads from a slice.
type impl struct {
	tokens []lexeme.Token
	index  int
	prev   lexeme.Token
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

// ToStream creates a simple TokenStream which reads from the slice tks.
func ToStream(tks []lexeme.Token) TokenStream {
	return &impl{
		tokens: tks,
	}
}

// ToArray consumes all the tokens from ts into an array.
func ToArray(ts TokenStream) []lexeme.Token {

	var tk lexeme.Token
	var tokens []lexeme.Token

	for tk = ts.Read(); tk.Lexeme != lexeme.LEXEME_EOF; tk = ts.Read() {
		tokens = append(tokens, tk)
	}

	return append(tokens, tk)
}

// PrintAll pretty prints all tokens in tks.
func PrintAll(tks []lexeme.Token) {
	for _, tk := range tks {
		switch k := tk.Lexeme; k {
		case lexeme.LEXEME_TERMINATOR:
			println(k)
		case lexeme.LEXEME_EOF:
			println(k)
			println()
		default:
			print(k + " ")
		}
	}
}
