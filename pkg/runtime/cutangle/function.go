package cutangle

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func evalFuncDef(ctx *alphaContext, f FuncDef) result {
	return funcLiteral(f)
}

func evalExprFuncDef(ctx *alphaContext, e ExprFuncDef) result {
	return exprFuncLiteral(e)
}

func evalFuncCall(ctx *alphaContext, call FuncCall) result {

	def := evalExpression(ctx, call.ID)

	switch v := def.(type) {
	case funcLiteral:
		return evalStdFuncCall(ctx, call, v)
	case exprFuncLiteral:
		return evalExprFuncCall(ctx, call, v)
	}

	err.Panic("Not a function or expression function", err.At(call.Token()))
	return nil
}

func evalStdFuncCall(ctx *alphaContext, call FuncCall, def funcLiteral) result {

	checkFuncCallArgs(def.Inputs, call.Inputs, call.ID.Token())

	funcCtx := evalFuncCallArgs(ctx, def.Inputs, call.Inputs)
	initFuncReturnArgs(funcCtx, def.Outputs)

	exeBlock(funcCtx, def.Body)
	results := collectFuncCallResults(funcCtx, def.Outputs)

	return tuple(results)
}

func evalExprFuncCall(ctx *alphaContext, call FuncCall, def exprFuncLiteral) result {
	checkFuncCallArgs(def.Inputs, call.Inputs, call.ID.Token())
	funcCtx := evalFuncCallArgs(ctx, def.Inputs, call.Inputs)
	return evalExpression(funcCtx, def.Expr)
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

func initFuncReturnArgs(ctx *alphaContext, outParams []OutputParam) {
	for _, p := range outParams {

		id := p.ID.Token()

		if p.Expr != nil {
			v := evalExpression(ctx, p.Expr)
			v = expectOneValue(v, p.Expr.Token())
			ctx.SetLocal(id, v)

		} else if v := ctx.GetLocal(id.Value()); v == nil {
			ctx.SetLocal(id, voidLiteral{})
		}
	}
}

func collectFuncCallResults(ctx *alphaContext, params []OutputParam) []result {

	r := make([]result, len(params))

	for i, p := range params {

		id := p.ID.Token()
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
