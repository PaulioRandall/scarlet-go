package runtime

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
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

func EvalIdentifier(ctx *Context, id st.Identifier) Value {

	v := ctx.Get(id.Source.Value)

	if v == nil {
		panic(err("EvalExpression", id.Source, "Undefined identifier"))
	}

	return v
}
