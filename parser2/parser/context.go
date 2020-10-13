package parser

import (
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
)

type TokenItr interface {
	More() bool
	Get() lexeme.Lexeme
	Next() lexeme.Lexeme
	Prev() lexeme.Lexeme
	LookAhead() lexeme.Lexeme
}

type context struct {
	TokenItr
	parent *context
}
