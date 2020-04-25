package runtime

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func EvalFuncDef(ctx *Context, f st.FuncDef) Value {
	return Function(f)
}

func EvalFuncCall(ctx *Context, call st.FuncCall) Value {

	def := findFunction(ctx, call.ID)

	checkFuncCallArgs(def.Input, call.Input, call.ID.Token())
	subCtx := evalFuncCallArgs(ctx, def.Input, call.Input)

	ExeBlock(subCtx, def.Body)
	results := collectFuncCallResults(subCtx, def.Output)

	return Tuple(results)
}

func findFunction(ctx *Context, idExp st.Expression) Function {

	v := EvalExpression(ctx, idExp)
	f, ok := v.(Function)

	if !ok {
		panic(err("EvalFuncCall", idExp.Token(), "Expected function as result"))
	}

	return f
}

func checkFuncCallArgs(exp []token.Token, act []st.Expression, callTk token.Token) {

	a, b := len(exp), len(act)

	if a != b {
		panic(err("checkParamCount", callTk,
			"Expected %d parameters, got %d", a, b,
		))
	}
}

func evalFuncCallArgs(ctx *Context, ids []token.Token, params []st.Expression) *Context {

	subCtx := ctx.Spawn()

	for i, p := range params {

		v := EvalExpression(ctx, p)
		v = expectOneValue(v, p.Token())

		subCtx.Set(ids[i], v)
	}

	return subCtx
}

func collectFuncCallResults(ctx *Context, ids []token.Token) []Value {

	r := make([]Value, len(ids))

	for i, id := range ids {
		r[i] = ctx.Get(id.Value)
	}

	return r
}

func expectOneValue(v Value, tk token.Token) Value {

	t, ok := v.(Tuple)
	if !ok {
		return v
	}

	a := []Value(t)

	if t == nil || len(a) != 1 {
		panic(err("expectOneValue", tk, "Expected exactly one result"))
	}

	return a[0]
}
