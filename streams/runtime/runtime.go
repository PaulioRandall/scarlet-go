package runtime

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

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

	values := EvalExpressions(ctx, stat.Exprs)

	if stat.IDs != nil {
		if len(stat.IDs) > len(values) {
			runtimeError(stat.Assign, "Not enough expression results to populate variables")
		} else if len(stat.IDs) < len(values) {
			runtimeError(stat.Assign, "Not enough variables to contain expression results")
		}
	}

	for i, id := range stat.IDs {
		ctx.Set(id.Value, values[i])
	}
}

func EvalExpressions(ctx *Context, exprs []recursive.Expression) []Value {

	for _, _ = range exprs {

	}

	return nil
}

func EvalExpression(ctx *Context, expr recursive.Expression) Value {
	return nil
}

func runtimeError(tk lexeme.Token, msg string) {
	if tk == (lexeme.Token{}) {
		panic("[RUNTIME] " + msg)
	} else {
		panic("[RUNTIME] (" + tk.String() + ") " + msg)
	}
}
