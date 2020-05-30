package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Factory interface {
	NewIdentifier(tk Token) Identifier
	NewLiteral(tk Token) Literal
}

func NewFactory() Factory {
	return fac{}
}

type fac struct{}

func (f fac) NewIdentifier(tk Token) Identifier {
	return Identifier{tk}
}

func (f fac) NewLiteral(tk Token) Literal {
	return Literal{tk}
}
