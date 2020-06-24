package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/runtime/result"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
)

func evalStatements(ctx *Context, stats []Expression) error {

	for _, st := range stats {
		e := evalStatement(ctx, st)
		if e != nil {
			return e
		}
	}

	return nil
}

func evalStatement(ctx *Context, st Expression) error {

	switch st.Kind() {
	case ST_ASSIGNMENT:
		return evalAssignmentBlock(ctx, st.(AssignmentBlock))
	}

	return err.NewBySnippet("Unknown statement type", st)
}

func evalAssignmentBlock(ctx *Context, as AssignmentBlock) error {

	values, e := evalAssignmentSources(ctx, as.Sources(), as.Count())
	if e != nil {
		return e
	}

	return doAssignments(ctx, as.Const(), as.Targets(), values, as.Count())
}

func evalAssignmentSources(
	ctx *Context,
	sources []Expression,
	count int,
) ([]Result, error) {

	var e error
	r := make([]Result, count)

	for i, s := range sources {

		r[i], e = evalExpression(ctx, s)
		if e != nil {
			return nil, e
		}
	}

	return r, nil
}

func doAssignments(
	ctx *Context,
	final bool,
	targets []Expression,
	values []Result,
	count int,
) error {

	for i := 0; i < count; i++ {
		e := doAssignment(ctx, final, targets[i], values[i])
		if e != nil {
			return e
		}
	}

	return nil
}

func doAssignment(
	ctx *Context,
	final bool,
	target Expression,
	value Result,
) error {
	// TODO
	return nil
}

func evalExpression(ctx *Context, expr Expression) (Result, error) {
	// TODO
	return Result{}, nil
}
