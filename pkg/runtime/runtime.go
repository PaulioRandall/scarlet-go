package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func Run(stats statement.Statements) Context {
	ctx := Context{make(map[string]Value), nil}
	ExeStatements(&ctx, stats)
	return ctx
}

func ExeStatements(ctx *Context, stats statement.Statements) {
	for _, s := range []statement.Statement(stats) {
		ExeStatement(ctx, s)
	}
}

func ExeStatement(ctx *Context, stat statement.Statement) {

	values := EvalExpressions(ctx, stat.Exprs)

	if stat.IDs != nil {

		if len(stat.IDs) > len(values) {
			panic(err("ExeStatement", stat.Assign,
				"Missing expression values to populate variables... have %d, want %d",
				len(stat.IDs), len(values),
			))

		} else if len(stat.IDs) < len(values) {
			panic(err("ExeStatement", stat.Assign,
				"Too many expression values to populate variables... have %d, want %d",
				len(stat.IDs), len(values),
			))
		}
	}

	for i, id := range stat.IDs {
		ctx.Set(id.Value, values[i])
	}
}

func EvalExpressions(ctx *Context, exprs []statement.Expression) []Value {

	var values []Value

	for _, expr := range exprs {
		v := EvalExpression(ctx, expr)
		values = append(values, v)
	}

	return values
}

func EvalExpression(ctx *Context, expr statement.Expression) Value {
	switch v := expr.(type) {
	case statement.Value:
		return valueOf(v.Source)
	}

	panic(err("EvalExpression", expr.Token(), "Unknown expression type"))
	return nil
}
