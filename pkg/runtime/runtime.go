package runtime

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func Run(ss []st.Statement) Context {
	ctx := Context{make(map[string]Value), nil}
	ExeStatements(&ctx, ss)
	return ctx
}

func ExeBlock(ctx *Context, b st.Block) {
	ExeStatements(ctx, b.Stats)
}

func ExeStatements(ctx *Context, ss []st.Statement) {
	for _, s := range ss {
		ExeStatement(ctx, s)
	}
}

func ExeStatement(ctx *Context, s st.Statement) {
	switch v := s.(type) {
	case st.Assignment:
		ExeAssignment(ctx, v)

	case st.Guard:
		ExeGuard(ctx, v)

	default:
		panic(err("ExeStatement", s.Token(), "Unknown statement type"))
	}
}

func ExeAssignment(ctx *Context, a st.Assignment) {

	values := EvalExpressions(ctx, a.Exprs)

	if a.IDs != nil {

		if len(a.IDs) > len(values) {
			panic(err("ExeStatement", a.Assign,
				"Missing expression values to populate variables... have %d, want %d",
				len(a.IDs), len(values),
			))

		} else if len(a.IDs) < len(values) {
			panic(err("ExeStatement", a.Assign,
				"Too many expression values to populate variables... have %d, want %d",
				len(a.IDs), len(values),
			))
		}
	}

	for i, id := range a.IDs {
		ctx.Set(id.Value, values[i])
	}
}

func ExeGuard(ctx *Context, g st.Guard) {
	if pass, ok := EvalExpression(ctx, g.Cond).(Bool); !ok {
		panic(err("ExeGuard", g.Open, "Unexpected non-boolean result"))
	} else if pass {
		ExeBlock(ctx, g.Block)
	}
}

func EvalExpressions(ctx *Context, exprs []st.Expression) []Value {

	var values []Value

	for _, expr := range exprs {
		v := EvalExpression(ctx, expr)
		values = append(values, v)
	}

	return values
}

func EvalExpression(ctx *Context, expr st.Expression) Value {
	switch v := expr.(type) {
	case st.Identifier:
		val := ctx.Get(v.Source.Value)
		if val == nil {
			panic(err("EvalExpression", v.Source, "Undefined identifier"))
		}
		return val

	case st.Value:
		return valueOf(v.Source)

	case st.Operation:
		return EvalOperation(ctx, v)

	case st.List:
		return EvalList(ctx, v)
	}

	panic(err("EvalExpression", expr.Token(), "Unknown expression type"))
}

func EvalList(ctx *Context, list st.List) Value {
	return List(EvalExpressions(ctx, list.Exprs))
}
