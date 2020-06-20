package esmerelda

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func newExit(tk Token, code Expression) exitStat {
	return exitStat{
		tk:   tk,
		code: code,
	}
}

func newVoid(tk Token) voidExpr {
	return voidExpr{tk}
}

func newIdentifier(tk Token) identifierExpr {
	return identifierExpr{tk}
}

func newLiteral(tk Token) literalExpr {
	return literalExpr{tk}
}

func newCollectionAccessor(collection, key Expression) collectionAccessorExpr {
	return collectionAccessorExpr{
		collection: collection,
		key:        key,
	}
}

func newNegation(expr Expression) negationExpr {
	return negationExpr{expr}
}

func newOperation(operator Token, left, right Expression) operationExpr {
	return operationExpr{
		operator: operator,
		left:     left,
		right:    right,
	}
}

func newAssignment(target, source Expression) assignmentStat {
	return assignmentStat{
		target: target,
		source: source,
	}
}

func newAssignmentBlock(targets, sources []Expression, count int) assignmentBlockStat {
	return assignmentBlockStat{
		targets: targets,
		sources: sources,
		count:   count,
	}
}

func newBlock(open, close Token, stats []Expression) blockExpr {
	return blockExpr{
		open:  open,
		close: close,
		stats: stats,
	}
}

func newUnDelimiteredBlock(stats []Expression) unDelimiteredBlockExpr {
	return unDelimiteredBlockExpr{
		stats: stats,
	}
}

func newExpressionFunction(key Token, inputs []Token, expr Expression) expressionFunctionExpr {
	return expressionFunctionExpr{
		key:    key,
		inputs: inputs,
		expr:   expr,
	}
}

func newParameters(open, close Token, inputs, outputs []Token) parametersDef {
	return parametersDef{
		open:    open,
		close:   close,
		inputs:  inputs,
		outputs: outputs,
	}
}

func newFunction(key Token, params Parameters, body Expression) functionExpr {
	return functionExpr{
		key:    key,
		params: params,
		body:   body,
	}
}

func newFunctionCall(close Token, f Expression, args []Expression) functionCallExpr {
	return functionCallExpr{
		close:    close,
		function: f,
		args:     args,
	}
}

func newWatch(key Token, ids []Token, body Block) watchStat {
	return watchStat{
		key:  key,
		ids:  ids,
		body: body,
	}
}

func newGuard(open Token, condition Expression, body Block) guardStat {
	return guardStat{
		open:      open,
		condition: condition,
		body:      body,
	}
}

func newWhenCase(object Expression, body Block) whenCaseStat {
	return whenCaseStat{
		object: object,
		body:   body,
	}
}

func newWhen(key, close Token, init Assignment, cases []WhenCase) whenStat {
	return whenStat{
		key:   key,
		close: close,
		init:  init,
		cases: cases,
	}
}

func newLoop(key Token, init Assignment, guard Guard) loopStat {
	return loopStat{
		key:   key,
		init:  init,
		guard: guard,
	}
}

func newSpellCall(spell, close Token, args []Expression) spellCallExpr {
	return spellCallExpr{
		spell: spell,
		close: close,
		args:  args,
	}
}

func newExists(close Token, subject Expression) existsExpr {
	return existsExpr{
		close:   close,
		subject: subject,
	}
}
