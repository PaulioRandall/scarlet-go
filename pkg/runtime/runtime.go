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

func EvalIdentifier(ctx *Context, id st.Identifier) Value {

	v := ctx.Get(id.Value)

	if v == nil {
		panic(err("EvalExpression", id.Token(), "Undefined identifier"))
	}

	return v
}
