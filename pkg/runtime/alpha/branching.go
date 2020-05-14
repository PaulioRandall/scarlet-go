package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func exeMatch(ctx *alphaContext, m st.Match) {
	for _, g := range m.Cases {
		if exeGuard(ctx, g) {
			break
		}
	}
}

func exeGuard(ctx *alphaContext, g st.Guard) bool {

	pass, ok := evalExpression(ctx, g.Condition).(boolLiteral)

	if !ok {
		panic(err("ExeGuard", g.Open, "Unexpected non-boolean result"))
	}

	if pass {
		guardCtx := ctx.Spawn()
		exeBlock(guardCtx, g.Block)
	}

	return bool(pass)
}
