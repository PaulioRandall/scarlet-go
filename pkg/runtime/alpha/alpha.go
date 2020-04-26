package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func Run(ss []st.Statement) alphaContext {
	ctx := alphaContext{
		make(map[string]result),
		make(map[string]result),
		nil,
	}

	exeStatements(&ctx, ss)
	return ctx
}
