package alpha

import (
	errr "github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
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
		errr.Panic(
			"Expected boolean value",
			errr.At(g.Open),
		)
	}

	if pass {
		guardCtx := ctx.Spawn(false)
		exeBlock(guardCtx, g.Block)
	}

	return bool(pass)
}
