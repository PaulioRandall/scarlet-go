package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func Run(ss []st.Statement) alphaContext {
	ctx := alphaContext{
		make(map[string]Value),
		make(map[string]Value),
		nil,
	}

	exeStatements(&ctx, ss)
	return ctx
}

func evalIdentifier(ctx *alphaContext, id st.Identifier) Value {

	v := ctx.Get(id.Value)

	if v == nil {
		panic(err("evalExpression", id.Token(), "Undefined identifier"))
	}

	return v
}
