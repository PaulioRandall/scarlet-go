package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
)

type Token interface {
	Value() string
	Begin() (int, int)
	End() (int, int)
}

func Run(stats []Expression) (*Context, error) {

	ctx := NewCtx(nil, true)
	e := EvalStatements(ctx, stats)
	if e != nil {
		return nil, e
	}

	return ctx, nil
}
