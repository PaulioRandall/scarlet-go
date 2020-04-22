package runtime

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func Run(ss []st.Statement) Context {
	ctx := Context{
		make(map[string]Value),
		make(map[string]Value),
		nil,
	}

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

	case st.Match:
		ExeMatch(ctx, v)

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
		ctx.Set(id, values[i])
	}
}

func ExeMatch(ctx *Context, m st.Match) {
	for _, g := range m.Cases {
		if ExeGuard(ctx, g) {
			break
		}
	}
}

func ExeGuard(ctx *Context, g st.Guard) bool {

	pass, ok := EvalExpression(ctx, g.Cond).(Bool)

	if !ok {
		panic(err("ExeGuard", g.Open, "Unexpected non-boolean result"))
	}

	if pass {
		ExeBlock(ctx, g.Block)
	}

	return bool(pass)
}

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

	case st.FuncDef:
		return EvalFuncDef(ctx, v)

	case st.FuncCall:
		return EvalFuncCall(ctx, v)
	}

	panic(err("EvalExpression", expr.Token(), "Unknown expression type"))
}

func EvalList(ctx *Context, list st.List) Value {
	return List(EvalExpressions(ctx, list.Exprs))
}

func EvalFuncDef(ctx *Context, f st.FuncDef) Value {
	return Function(f)
}

func EvalFuncCall(ctx *Context, call st.FuncCall) Value {

	v := EvalExpression(ctx, call.ID)
	def, ok := v.(Function)

	if !ok {
		panic(err("EvalFuncCall", call.ID.Token(), "Expected function value"))
	}

	if len(def.Input) != len(call.Input) {
		panic(err("EvalFuncCall", call.ID.Token(),
			"Expected %d parameters, got %d", len(def.Input), len(call.Input),
		))
	}

	subCtx := ctx.Spawn()

	for i, paramExpr := range call.Input {

		v := EvalExpression(ctx, paramExpr)
		id := def.Input[i]
		v = single(v, call.ID.Token())

		ident := st.Identifier{false, id}
		subCtx.Set(ident, v)
	}

	ExeBlock(subCtx, def.Body)

	r := make([]Value, len(def.Output))

	for i, id := range def.Output {
		r[i] = subCtx.Get(id.Value)
	}

	return Tuple(r)
}

func single(v Value, tk token.Token) Value {

	t, ok := v.(Tuple)
	if !ok {
		return v
	}

	a := []Value(t)

	if t == nil || len(a) != 1 {
		panic(err("singletonTuple", tk, "Expected exactly one result"))
	}

	return a[0]
}
