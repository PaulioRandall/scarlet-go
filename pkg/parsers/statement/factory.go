package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Factory interface {
	NewVoid(tk Token) Void
	NewIdentifier(tk Token) Identifier
	NewLiteral(tk Token) Literal
	NewList(open, close Token, items []Expression) List
	NewListAccessor(list, index Expression) ListAccessor
	NewNegation(expr Expression) Negation
	NewAssignment(target, source Expression) Assignment
	NewBlock(start, end Token, stats []Statement) Block
	NewNonWrappedBlock(stats []Statement) Block
	NewParameters(open, close Token, inputs, outputs []Token) Parameters
	NewFunction(key Token, params Parameters, body Block) Function
	NewNumericOperation(operator Token, left, right Expression) NumericOperation
}

func NewFactory() Factory {
	return fac{}
}

func Precedence(expr Expression) int {
	if v, ok := expr.(Operation); ok {
		return v.Precedence()
	}

	return 0
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

func (fac) NewList(open, close Token, items []Expression) List {
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

func (fac) NewBlock(start, end Token, stats []Statement) Block {
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

func (fac) NewParameters(open, close Token, inputs, outputs []Token) Parameters {
	return Parameters{
		open:    open,
		close:   close,
		Inputs:  inputs,
		Outputs: outputs,
	}
}

func (fac) NewFunction(key Token, params Parameters, body Block) Function {
	return Function{
		key:    key,
		Params: params,
		Body:   body,
	}
}

func (fac) NewNumericOperation(operator Token, left, right Expression) NumericOperation {
	return NumericOperation{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
