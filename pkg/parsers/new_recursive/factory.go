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
	return voidExpr{tk}
}

func newIdentifier(tk Token) Expression {
	return identifierExpr{tk}
}

func newLiteral(tk Token) Expression {
	return literalExpr{tk}
}

func newListAccessor(id, index Expression) Expression {
	return listAccessorExpr{
		id:    id,
		index: index,
	}
}

func newList(open, close Token, items []Expression) Expression {
	return listConstructorExpr{
		open:  open,
		close: close,
		items: items,
	}
}

func newNegation(expr Expression) Expression {
	return negationExpr{expr}
}

func newAssignment(target, source Expression) Statement {
	return assignmentStat{
		target: target,
		source: source,
	}
}

func newBlock(open, close Token, stats []Statement) Block {
	return blockExpr{
		open:  open,
		close: close,
		stats: stats,
	}
}

func newNonWrappedBlock(stats []Statement) Block {
	return blockExpr{
		stats: stats,
	}
}

func newParameters(open, close Token, inputs, outputs []Token) Parameters {
	return parametersDef{
		open:    open,
		close:   close,
		inputs:  inputs,
		outputs: outputs,
	}
}

func newFunction(key Token, params Parameters, body Block) Expression {
	return functionExpr{
		key:    key,
		params: params,
		body:   body,
	}
}

func newOperation(operator Token, left, right Expression) Operation {
	return Operation{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
