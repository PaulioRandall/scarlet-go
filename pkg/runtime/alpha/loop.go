package alpha

import (
	"github.com/shopspring/decimal"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
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

func exeForEach(ctx *alphaContext, f ForEach) {

	v := evalExpression(ctx, f.List)
	list, ok := v.(listLiteral)

	if !ok {
		err.Panic("Not a list", err.At(f.List.Token()))
	}

	loopCtx := ctx.Spawn(false)
	items := []result(list)
	size := len(items)

	for i, v := range items {

		// Set the iteration  index
		d := decimal.NewFromInt(int64(i))
		n := numberLiteral(d)
		loopCtx.SetLocal(f.IndexId, n)

		// Set the iteration value
		loopCtx.SetLocal(f.ValueId, v)

		// Set the more, i.e. false only if this is the last item
		m := i < size-1
		loopCtx.SetLocal(f.MoreId, boolLiteral(m))

		exeBlock(loopCtx, f.Block)
	}
}
