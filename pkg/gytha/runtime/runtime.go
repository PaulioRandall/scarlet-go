package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
)

func Run(ss []Statement) alphaContext {
	ctx := alphaContext{
		true,
		make(map[string]result),
		make(map[string]result),
		nil,
	}

	exeStatements(&ctx, ss)
	return ctx
}
