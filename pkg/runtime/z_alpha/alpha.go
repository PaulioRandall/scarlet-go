package z_alpha

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
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
