package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
)

func Run(sts []Expression) (*Context, error) {

	ctx := NewCtx(nil, true)
	e := EvalStatements(ctx, sts)
	if e != nil {
		return nil, e
	}

	return ctx, nil
}
