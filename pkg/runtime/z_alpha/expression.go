package z_alpha

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
)

func evalIdentifier(ctx *alphaContext, id Identifier) result {

	v := ctx.Get(id.Token().Value())

	if v == nil {
		panic(err("evalIdentifier", id.Token(), "Undefined identifier"))
	}

	return v
}

func evalExpressions(ctx *alphaContext, exprs []Expression) []result {

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

func evalExpression(ctx *alphaContext, expr Expression) result {
	switch v := expr.(type) {
	case Identifier:
		return evalIdentifier(ctx, v)

	case Value:
		return valueOf(v.Token())

	case Operation:
		return evalOperation(ctx, v)

	case List:
		return evalList(ctx, v)

	case ListAccess:
		return evalListAccess(ctx, v)

	case FuncDef:
		return evalFuncDef(ctx, v)

	case FuncCall:
		return evalFuncCall(ctx, v)
	}

	panic(err("EvalExpression", expr.Token(), "Unknown expression type"))
}
