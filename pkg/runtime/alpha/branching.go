package alpha

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func exeWhen(ctx *alphaContext, m When) {
	for _, g := range m.Cases {
		if exeGuard(ctx, g) {
			break
		}
	}
}

func exeGuard(ctx *alphaContext, g Guard) bool {

	pass, ok := evalExpression(ctx, g.Condition).(boolLiteral)

	if !ok {
		err.Panic("Expected boolean value", err.At(g.Open))
	}

	if pass {
		guardCtx := ctx.Spawn(false)
		exeBlock(guardCtx, g.Block)
	}

	return bool(pass)
}
