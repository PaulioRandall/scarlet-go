package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Factory interface {
	NewIdentifier(tk Token) Identifier
	NewLiteral(tk Token) Literal
	NewNegation(expr Expression) Negation
	NewAssignment(tk Token, expr Expression) Assignment
	NewAssignmentBlock(as []Assignment) AssignmentBlock
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

func (f fac) NewNegation(expr Expression) Negation {
	return Negation{expr}
}

func (f fac) NewAssignment(tk Token, expr Expression) Assignment {
	return Assignment{
		Target: tk,
		Source: expr,
	}
}

func (f fac) NewAssignmentBlock(as []Assignment) AssignmentBlock {
	return AssignmentBlock{as}
}
