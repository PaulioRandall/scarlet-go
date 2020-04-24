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

	case st.Expression:
		_ = EvalExpression(ctx, v)

	default:
		panic(err("ExeStatement", s.Token(), "Unknown statement type"))
	}
}

func ExeAssignment(ctx *Context, a st.Assignment) {

	values := EvalExpressions(ctx, a.Exprs)
	checkAssignments(a.IDs, values, a.Assign)

	for i, id := range a.IDs {
		ctx.Set(id, values[i])
	}
}

func checkAssignments(ids []st.Identifier, vals []Value, operator token.Token) {

	a, b := len(ids), len(vals)

	if a > b {
		panic(err("ExeStatement", operator,
			"Missing expression values to populate variables... have %d, want %d",
			len(ids), len(vals),
		))
	}

	if a < b {
		panic(err("ExeStatement", operator,
			"Too many expression values to populate variables... have %d, want %d",
			len(ids), len(vals),
		))
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

func EvalIdentifier(ctx *Context, id st.Identifier) Value {

	v := ctx.Get(id.Source.Value)

	if v == nil {
		panic(err("EvalExpression", id.Source, "Undefined identifier"))
	}

	return v
}

func EvalList(ctx *Context, list st.List) Value {
	return List(EvalExpressions(ctx, list.Exprs))
}

func EvalListAccess(ctx *Context, la st.ListAccess) Value {

	v := EvalIdentifier(ctx, la.ID)
	list, ok := v.(List)

	if !ok {
		panic(err("EvalListAccess", la.ID.Source, "Can't get item of a non-list"))
	}

	n := EvalExpression(ctx, la.Index)
	index, ok := n.(Number)

	if !ok {
		panic(err("EvalListAccess", la.Index.Token(), "Expected number as result"))
	}

	i := index.ToInt()
	if i < 0 {
		panic(err("EvalListAccess", la.ID.Source, "Index must be greater than zero"))
	}

	items := []Value(list)
	if i >= int64(len(items)) {
		panic(err("EvalListAccess", la.Index.Token(), "Index out of range"))
	}

	return items[i]
}
