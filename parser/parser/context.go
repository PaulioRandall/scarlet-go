package parser

import (
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
	"github.com/PaulioRandall/scarlet-go/token/position"
)

type TokenItr interface {
	More() bool
	Get() lexeme.Lexeme
	Next() lexeme.Lexeme
	Prev() lexeme.Lexeme
	LookAhead() lexeme.Lexeme
	Snippet() position.Snippet
}

type context struct {
	TokenItr
	parent *context
}
