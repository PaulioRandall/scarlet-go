package runtime

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

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
