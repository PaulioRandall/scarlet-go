package alpha

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func evalIdentifier(ctx *alphaContext, id Identifier) result {

	v := ctx.Get(id.Token().Value())

	if v == nil {
		err.Panic("Undefined identifier", err.At(id.Token()))
	}

	return v
}

func evalExpressions(ctx *alphaContext, exprs []Expression) []result {

	var vs []result

	for _, expr := range exprs {

		v := evalExpression(ctx, expr)
		t, ok := v.(tuple)

		if !ok {
			vs = append(vs, v)
			continue
		}

		for _, v := range []result(t) {
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

	case ExprFuncDef:
		return evalExprFuncDef(ctx, v)

	case FuncCall:
		return evalFuncCall(ctx, v)
	}

	err.Panic("Unknown expression type", err.At(expr.Token()))
	return nil
}
