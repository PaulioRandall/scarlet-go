package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func exeLoop(ctx *alphaContext, l st.Loop) {

	loopCtx := ctx.Spawn()

	for i := 0; ; i++ {

		n := numberLiteral(float64(i))
		loopCtx.SetLocal(l.IndexVar, n)

		if !exeGuard(loopCtx, l.Guard) {
			break
		}
	}
}
