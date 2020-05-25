package alpha

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
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
		err.Panic("Not a function", err.At(idExp.Token()))
	}

	return f
}

func checkFuncCallArgs(exp []Token, act []Expression, callTk Token) {

	a, b := len(exp), len(act)

	if a != b {
		m := fmt.Sprintf("Expected %d parameters, given %d", a, b)
		err.Panic(m, err.At(callTk))
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
		err.Panic("Need one result", err.At(tk))
	}

	return a[0]
}
