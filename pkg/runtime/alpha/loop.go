package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func exeLoop(ctx *alphaContext, l st.Loop) {

	for i := 0; ; i++ {

		n := numberLiteral(float64(i))
		ctx.Set(l.IndexVar, n)

		if !exeGuard(ctx, l.Guard) {
			break
		}
	}
}
