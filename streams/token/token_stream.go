package token

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// TokenStream provides access to an ordered stream of tokens.
type TokenStream interface {

	// Next returns the next token in the stream. An EOF token is always returned
	// if the stream is empty.
	Next() lexeme.Token
}
