package stats

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

func NewExit(tk Token, code Expr) exitStat {
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

func NewContainerItem(container, key Expr) containerItemExpr {
	return containerItemExpr{
		container: container,
		key:       key,
	}
}

func NewNegation(expr Expr) negationExpr {
	return negationExpr{expr}
}

func NewOperation(operator Token, left, right Expr) operationExpr {
	return operationExpr{
		operator: operator,
		left:     left,
		right:    right,
	}
}

func NewAssign(target, source Expr) assignStat {
	return assignStat{
		target: target,
		source: source,
	}
}

func NewAssignBlock(final bool, targets, sources []Expr, count int) assignBlockStat {
	return assignBlockStat{
		final:   final,
		targets: targets,
		sources: sources,
		count:   count,
	}
}

func NewBlock(open, close Token, stats []Expr) blockExpr {
	return blockExpr{
		open:  open,
		close: close,
		stats: stats,
	}
}

func NewUndelimBlock(stats []Expr) undelimBlockExpr {
	return undelimBlockExpr{
		stats: stats,
	}
}

func NewExprFunc(key Token, inputs []Token, expr Expr) exprFuncExpr {
	return exprFuncExpr{
		key:    key,
		inputs: inputs,
		expr:   expr,
	}
}

func NewFuncDef(key Token, inputs, outputs []Token, body Expr) funcDefExpr {
	return funcDefExpr{
		key:     key,
		inputs:  inputs,
		outputs: outputs,
		body:    body,
	}
}

func NewFuncCall(close Token, f Expr, args []Expr) funcCallExpr {
	return funcCallExpr{
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

func NewGuard(open Token, condition Expr, body Block) guardStat {
	return guardStat{
		open:      open,
		condition: condition,
		body:      body,
	}
}

func NewWhenCase(object Expr, body Block) whenCaseStat {
	return whenCaseStat{
		object: object,
		body:   body,
	}
}

func NewWhen(key, close Token, init Assign, cases []WhenCase) whenStat {
	return whenStat{
		key:   key,
		close: close,
		init:  init,
		cases: cases,
	}
}

func NewLoop(key Token, init Assign, guard Guard) loopStat {
	return loopStat{
		key:   key,
		init:  init,
		guard: guard,
	}
}

func NewSpellCall(spell, close Token, args []Expr) spellCallExpr {
	return spellCallExpr{
		spell: spell,
		close: close,
		args:  args,
	}
}

func NewExists(close Token, subject Expr) existsExpr {
	return existsExpr{
		close:   close,
		subject: subject,
	}
}
