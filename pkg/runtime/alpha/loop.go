package alpha

import (
	"github.com/shopspring/decimal"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func exeLoop(ctx *alphaContext, l Loop) {

	loopCtx := ctx.Spawn(false)

	for i := 0; ; i++ {

		d := decimal.NewFromInt(int64(i))
		n := numberLiteral(d)
		loopCtx.SetLocal(l.IndexVar, n)

		if !exeGuard(loopCtx, l.Guard) {
			break
		}
	}
}
