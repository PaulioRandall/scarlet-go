package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Factory interface {
	Identifier(tk Token) Expression
	Literal(tk Token) Expression
}

func NewFactory() Factory {
	return fac{}
}

type fac struct{}

func (f fac) Identifier(tk Token) Expression {
	return Identifier{tk}
}

func (f fac) Literal(tk Token) Expression {
	return Literal{tk}
}
