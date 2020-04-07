package lexor

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenStream provides access to an ordered stream of tokens.
type TokenStream interface {

	// Next returns the next token in the stream. An EOF token is always returned
	// if the token stream is empty.
	Next() token.Token
}

// NewScanner creates a new token scanner. The input string is parsed one token
// at a time on each call to TokenStream.Next.
func NewScanner(s string) TokenStream {
	return &scanner{
		runes: []rune(s),
		line:  0,
		col:   0,
	}
}

// NewEvaluator creates a new evaluator to evaluate tokens within a stream.
func NewEvaluator(delegate TokenStream) TokenStream {
	return &evaluator{
		ts: delegate,
	}
}
