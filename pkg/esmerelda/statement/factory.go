package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func NewExit(tk Token, code Expression) exitStat {
	return exitStat{
		tk:   tk,
		code: code,
	}
}

func NewVoid(tk Token) voidExpr {
	return voidExpr{tk}
}

func NewIdentifier(tk Token) identifierExpr {
	return identifierExpr{tk}
}

func NewLiteral(tk Token) literalExpr {
	return literalExpr{tk}
}

func NewCollectionAccessor(collection, key Expression) collectionAccessorExpr {
	return collectionAccessorExpr{
		collection: collection,
		key:        key,
	}
}

func NewNegation(expr Expression) negationExpr {
	return negationExpr{expr}
}

func NewOperation(operator Token, left, right Expression) operationExpr {
	return operationExpr{
		operator: operator,
		left:     left,
		right:    right,
	}
}

func NewAssignment(target, source Expression) assignmentStat {
	return assignmentStat{
		target: target,
		source: source,
	}
}

func NewAssignmentBlock(final bool, targets, sources []Expression, count int) assignmentBlockStat {
	return assignmentBlockStat{
		final:   final,
		targets: targets,
		sources: sources,
		count:   count,
	}
}

func NewBlock(open, close Token, stats []Expression) blockExpr {
	return blockExpr{
		open:  open,
		close: close,
		stats: stats,
	}
}

func NewUnDelimiteredBlock(stats []Expression) unDelimiteredBlockExpr {
	return unDelimiteredBlockExpr{
		stats: stats,
	}
}

func NewExpressionFunction(key Token, inputs []Token, expr Expression) expressionFunctionExpr {
	return expressionFunctionExpr{
		key:    key,
		inputs: inputs,
		expr:   expr,
	}
}

func NewParameters(open, close Token, inputs, outputs []Token) parametersDef {
	return parametersDef{
		open:    open,
		close:   close,
		inputs:  inputs,
		outputs: outputs,
	}
}

func NewFunction(key Token, params Parameters, body Expression) functionExpr {
	return functionExpr{
		key:    key,
		params: params,
		body:   body,
	}
}

func NewFunctionCall(close Token, f Expression, args []Expression) functionCallExpr {
	return functionCallExpr{
		close:    close,
		function: f,
		args:     args,
	}
}

func NewWatch(key Token, ids []Token, body Block) watchStat {
	return watchStat{
		key:  key,
		ids:  ids,
		body: body,
	}
}

func NewGuard(open Token, condition Expression, body Block) guardStat {
	return guardStat{
		open:      open,
		condition: condition,
		body:      body,
	}
}

func NewWhenCase(object Expression, body Block) whenCaseStat {
	return whenCaseStat{
		object: object,
		body:   body,
	}
}

func NewWhen(key, close Token, init Assignment, cases []WhenCase) whenStat {
	return whenStat{
		key:   key,
		close: close,
		init:  init,
		cases: cases,
	}
}

func NewLoop(key Token, init Assignment, guard Guard) loopStat {
	return loopStat{
		key:   key,
		init:  init,
		guard: guard,
	}
}

func NewSpellCall(spell, close Token, args []Expression) spellCallExpr {
	return spellCallExpr{
		spell: spell,
		close: close,
		args:  args,
	}
}

func NewExists(close Token, subject Expression) existsExpr {
	return existsExpr{
		close:   close,
		subject: subject,
	}
}
