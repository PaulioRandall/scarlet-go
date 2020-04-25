package runtime

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func EvalExpressions(ctx *Context, exprs []st.Expression) []Value {

	var values []Value

	for _, expr := range exprs {
		v := EvalExpression(ctx, expr)

		if t, ok := v.(Tuple); ok {
			for _, v := range []Value(t) {
				values = append(values, v)
			}

		} else {
			values = append(values, v)
		}
	}

	return values
}

func EvalExpression(ctx *Context, expr st.Expression) Value {
	switch v := expr.(type) {
	case st.Identifier:
		return EvalIdentifier(ctx, v)

	case st.Value:
		return valueOf(v.Source)

	case st.Operation:
		return EvalOperation(ctx, v)

	case st.List:
		return EvalList(ctx, v)

	case st.ListAccess:
		return EvalListAccess(ctx, v)

	case st.FuncDef:
		return EvalFuncDef(ctx, v)

	case st.FuncCall:
		return EvalFuncCall(ctx, v)
	}

	panic(err("EvalExpression", expr.Token(), "Unknown expression type"))
}
