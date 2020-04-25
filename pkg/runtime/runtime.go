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
		if a.Fixed {
			ctx.SetFixed(id, values[i])
		} else {
			ctx.Set(id, values[i])
		}
	}
}

func checkAssignments(ids []token.Token, vals []Value, operator token.Token) {

	a, b := len(ids), len(vals)

	if a > b {
		panic(err("ExeStatement", operator,
			"Missing expression values to populate variables... have %d, want %d",
			a, b,
		))
	}

	if a < b {
		panic(err("ExeStatement", operator,
			"Too many expression values to populate variables... have %d, want %d",
			a, b,
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

func EvalIdentifier(ctx *Context, id st.Identifier) Value {

	v := ctx.Get(id.Source.Value)

	if v == nil {
		panic(err("EvalExpression", id.Source, "Undefined identifier"))
	}

	return v
}
