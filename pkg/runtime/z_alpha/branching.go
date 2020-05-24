package z_alpha

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
)

func exeMatch(ctx *alphaContext, m Match) {
	for _, g := range m.Cases {
		if exeGuard(ctx, g) {
			break
		}
	}
}

func exeGuard(ctx *alphaContext, g Guard) bool {

	pass, ok := evalExpression(ctx, g.Condition).(boolLiteral)

	if !ok {
		panic(err("ExeGuard", g.Open, "Unexpected non-boolean result"))
	}

	if pass {
		guardCtx := ctx.Spawn(false)
		exeBlock(guardCtx, g.Block)
	}

	return bool(pass)
}
