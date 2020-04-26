package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func evalExpressions(ctx *alphaContext, exprs []st.Expression) []Value {

	var values []Value

	for _, expr := range exprs {
		v := evalExpression(ctx, expr)

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

func evalExpression(ctx *alphaContext, expr st.Expression) Value {
	switch v := expr.(type) {
	case st.Identifier:
		return evalIdentifier(ctx, v)

	case st.Value:
		return valueOf(v.Token())

	case st.Operation:
		return evalOperation(ctx, v)

	case st.List:
		return evalList(ctx, v)

	case st.ListAccess:
		return evalListAccess(ctx, v)

	case st.FuncDef:
		return evalFuncDef(ctx, v)

	case st.FuncCall:
		return evalFuncCall(ctx, v)
	}

	panic(err("EvalExpression", expr.Token(), "Unknown expression type"))
}
