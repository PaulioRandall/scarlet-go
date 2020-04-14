package runtime

import (
	"github.com/PaulioRandall/scarlet-go/streams/parser/recursive"
)

func Run(stats recursive.Statements) Context {
	ctx := Context{make(map[string]Value), nil}
	ExeStatements(&ctx, stats)
	return ctx
}

func ExeStatements(ctx *Context, stats recursive.Statements) {
	for _, s := range []recursive.Statement(stats) {
		ExeStatement(ctx, s)
	}
}

func ExeStatement(ctx *Context, stat recursive.Statement) {
	// 1. Evaluate expressions
	// 2. Do the number of results from the expressions match the number of ids?
	// 3. For each ID assign its expression result
}

func EvalExpressions(ctx *Context, exprs []recursive.Expression) []Value {
	return nil
}

func EvalExpression(ctx *Context, expr recursive.Expression) Value {
	return nil
}
