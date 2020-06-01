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
	NewBlock(start Token, stats []Statement, end Token) Block
	NewNonWrappedBlock(stats []Statement) Block
	NewFunction(key Token, params Parameters, body Block) Function
}

func NewFactory() Factory {
	return fac{}
}

type fac struct{}

func (fac) NewVoid(tk Token) Void {
	return Void{tk}
}

func (fac) NewIdentifier(tk Token) Identifier {
	return Identifier{tk}
}

func (fac) NewLiteral(tk Token) Literal {
	return Literal{tk}
}

func (fac) NewList(open Token, items []Expression, close Token) List {
	return List{
		Open:  open,
		Items: items,
		Close: close,
	}
}

func (fac) NewListAccessor(list, index Expression) ListAccessor {
	return ListAccessor{
		List:  list,
		Index: index,
	}
}

func (fac) NewNegation(expr Expression) Negation {
	return Negation{expr}
}

func (fac) NewAssignment(target, source Expression) Assignment {
	return Assignment{
		Target: target,
		Source: source,
	}
}

func (fac) NewBlock(start Token, stats []Statement, end Token) Block {
	return Block{
		start: start,
		Stats: stats,
		end:   end,
	}
}

func (fac) NewNonWrappedBlock(stats []Statement) Block {
	return Block{
		Stats: stats,
	}
}

func (fac) NewFunction(key Token, params Parameters, body Block) Function {
	return Function{
		key:    key,
		Params: params,
		Body:   body,
	}
}
