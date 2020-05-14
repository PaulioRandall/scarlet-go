package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func evalIdentifier(ctx *alphaContext, id st.Identifier) result {

	v := ctx.Get(id.Value)

	if v == nil {
		panic(err("evalIdentifier", id.Token(), "Undefined identifier"))
	}

	return v
}

func evalExpressions(ctx *alphaContext, exprs []st.Expression) []result {

	var vs []result

	for _, expr := range exprs {
		v := evalExpression(ctx, expr)

		if t, ok := v.(tuple); ok {
			for _, v := range []result(t) {
				vs = append(vs, v)
			}

		} else {
			vs = append(vs, v)
		}
	}

	return vs
}

func evalExpression(ctx *alphaContext, expr st.Expression) result {
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
