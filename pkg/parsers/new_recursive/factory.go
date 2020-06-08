package recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func Precedence(expr Expression) int {
	if v, ok := expr.(Operation); ok {
		return v.Precedence()
	}

	return 0
}

func newVoid(tk Token) Expression {
	return void{tk}
}

func newIdentifier(tk Token) Identifier {
	return Identifier{tk}
}

func newLiteral(tk Token) Literal {
	return Literal{tk}
}

func newList(open, close Token, items []Expression) List {
	return List{
		Open:  open,
		Items: items,
		Close: close,
	}
}

func newListAccessor(list, index Expression) ListAccessor {
	return ListAccessor{
		List:  list,
		Index: index,
	}
}

func newNegation(expr Expression) Negation {
	return Negation{expr}
}

func newAssignment(target, source Expression) Assignment {
	return Assignment{
		Target: target,
		Source: source,
	}
}

func newBlock(start, end Token, stats []Statement) Block {
	return Block{
		start: start,
		Stats: stats,
		end:   end,
	}
}

func newNonWrappedBlock(stats []Statement) Block {
	return Block{
		Stats: stats,
	}
}

func newParameters(open, close Token, inputs, outputs []Token) Parameters {
	return Parameters{
		open:    open,
		close:   close,
		Inputs:  inputs,
		Outputs: outputs,
	}
}

func newFunction(key Token, params Parameters, body Block) Function {
	return Function{
		key:    key,
		Params: params,
		Body:   body,
	}
}

func newOperation(operator Token, left, right Expression) Operation {
	return Operation{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
