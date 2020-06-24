package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
)

func Run(stats []Expression) *Context {
	ctx := NewCtx(nil, true)
	evalStatements(ctx, stats)
	return ctx
}
