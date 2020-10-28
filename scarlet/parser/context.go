package parser

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type TokenItr interface {
	More() bool
	Get() token.Lexeme
	Next() token.Lexeme
	Prev() token.Lexeme
	LookAhead() token.Lexeme
	Snippet() token.Snippet
}

type context struct {
	TokenItr
	parent *context
}
