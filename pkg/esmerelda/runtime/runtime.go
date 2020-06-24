package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
)

func Run(stats []Expression) (*Context, error) {

	ctx := NewCtx(nil, true)
	e := evalStatements(ctx, stats)
	if e != nil {
		return nil, e
	}

	return ctx, nil
}
