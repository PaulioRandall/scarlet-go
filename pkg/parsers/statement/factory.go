package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Factory interface {
	NewVoid(tk Token) Void
	NewIdentifier(tk Token) Identifier
	NewLiteral(tk Token) Literal
	NewList(open Token, items []Expression, close Token) List
	NewListAccessor(list, index Expression) ListAccessor
	NewNegation(expr Expression) Negation
	NewAssignment(target, source Expression) Assignment
	NewAssignmentBlock(as []Assignment) AssignmentBlock
}

func NewFactory() Factory {
	return fac{}
}

type fac struct{}

func (f fac) NewVoid(tk Token) Void {
	return Void{tk}
}

func (f fac) NewIdentifier(tk Token) Identifier {
	return Identifier{tk}
}

func (f fac) NewLiteral(tk Token) Literal {
	return Literal{tk}
}

func (f fac) NewList(open Token, items []Expression, close Token) List {
	return List{
		Open:  open,
		Items: items,
		Close: close,
	}
}

func (f fac) NewListAccessor(list, index Expression) ListAccessor {
	return ListAccessor{
		List:  list,
		Index: index,
	}
}

func (f fac) NewNegation(expr Expression) Negation {
	return Negation{expr}
}

func (f fac) NewAssignment(target, source Expression) Assignment {
	return Assignment{
		Target: target,
		Source: source,
	}
}

func (f fac) NewAssignmentBlock(as []Assignment) AssignmentBlock {
	return AssignmentBlock{as}
}
