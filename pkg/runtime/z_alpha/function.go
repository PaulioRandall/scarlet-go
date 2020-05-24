package z_alpha

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func evalFuncDef(ctx *alphaContext, f FuncDef) result {
	return functionLiteral(f)
}

func evalFuncCall(ctx *alphaContext, call FuncCall) result {

	def := findFunction(ctx, call.ID)

	checkFuncCallArgs(def.Inputs, call.Inputs, call.ID.Token())

	funcCtx := evalFuncCallArgs(ctx, def.Inputs, call.Inputs)
	initFuncReturnArgs(funcCtx, def.Outputs)

	exeBlock(funcCtx, def.Body)
	results := collectFuncCallResults(funcCtx, def.Outputs)

	return tuple(results)
}

func findFunction(ctx *alphaContext, idExp Expression) functionLiteral {

	v := evalExpression(ctx, idExp)
	f, ok := v.(functionLiteral)

	if !ok {
		panic(err("EvalFuncCall", idExp.Token(), "Expected function as result"))
	}

	return f
}

func checkFuncCallArgs(exp []Token, act []Expression, callTk Token) {

	a, b := len(exp), len(act)

	if a != b {
		panic(err("checkParamCount", callTk,
			"Expected %d parameters, got %d", a, b,
		))
	}
}

func evalFuncCallArgs(ctx *alphaContext, ids []Token, params []Expression) *alphaContext {

	funcCtx := ctx.Spawn(true)

	for i, p := range params {

		v := evalExpression(ctx, p)
		v = expectOneValue(v, p.Token())

		funcCtx.SetLocal(ids[i], v)
	}

	return funcCtx
}

func initFuncReturnArgs(ctx *alphaContext, outParams []Token) {
	for _, p := range outParams {
		if v := ctx.GetLocal(p.Value()); v == nil {
			ctx.SetLocal(p, voidLiteral{})
		}
	}
}

func collectFuncCallResults(ctx *alphaContext, ids []Token) []result {

	r := make([]result, len(ids))

	for i, id := range ids {

		v := ctx.GetLocal(id.Value())

		if v != nil {
			r[i] = v
		} else {
			r[i] = voidLiteral{}
		}
	}

	return r
}

func expectOneValue(v result, tk Token) result {

	t, ok := v.(tuple)
	if !ok {
		return v
	}

	a := []result(t)

	if t == nil || len(a) != 1 {
		panic(err("expectOneValue", tk, "Expected exactly one result"))
	}

	return a[0]
}
