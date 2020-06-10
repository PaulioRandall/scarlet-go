package recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func newVoid(tk Token) Void {
	return voidExpr{tk}
}

func newIdentifier(tk Token) Identifier {
	return identifierExpr{tk}
}

func newLiteral(tk Token) Literal {
	return literalExpr{tk}
}

func newListAccessor(id, index Expression) ListAccessor {
	return listAccessorExpr{
		id:    id,
		index: index,
	}
}

func newNegation(expr Expression) Negation {
	return negationExpr{expr}
}

func newOperation(operator Token, left, right Expression) Operation {
	return operationExpr{
		operator: operator,
		left:     left,
		right:    right,
	}
}

func newAssignment(target, source Expression) Assignment {
	return assignmentStat{
		target: target,
		source: source,
	}
}

func newAssignmentBlock(assignments []Assignment) AssignmentBlock {
	return assignmentBlockStat{
		assignments: assignments,
	}
}

func newBlock(open, close Token, stats []Expression) Block {
	return blockExpr{
		open:  open,
		close: close,
		stats: stats,
	}
}

func newExpressionFunction(key Token, inputs []Token, expr Expression) ExpressionFunction {
	return expressionFunctionExpr{
		key:    key,
		inputs: inputs,
		expr:   expr,
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

func newFunction(key Token, params Parameters, body Block) Function {
	return functionExpr{
		key:    key,
		params: params,
		body:   body,
	}
}

func newWatch(key Token, ids []Token, body Block) Watch {
	return watchStat{
		key:  key,
		ids:  ids,
		body: body,
	}
}

func newGuard(open Token, condition Expression, body Block) Guard {
	return guardStat{
		open:      open,
		condition: condition,
		body:      body,
	}
}
