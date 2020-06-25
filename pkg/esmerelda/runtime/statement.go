package runtime

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

func EvalStatements(ctx *Context, sts []Expression) error {

	for _, st := range sts {
		e := EvalStatement(ctx, st)
		if e != nil {
			return e
		}
	}

	return nil
}

func EvalStatement(ctx *Context, st Expression) error {

	switch st.Kind() {
	case ST_ASSIGNMENT:
		return EvalAssignmentBlock(ctx, st.(AssignmentBlock))
	}

	panic(err.NewBySnippet("Unknown statement type", st))
}

func EvalAssignmentBlock(ctx *Context, as AssignmentBlock) error {

	values, e := EvalAssignmentSources(ctx, as.Sources(), as.Count())
	if e != nil {
		return e
	}

	return doAssignments(ctx, as.Const(), as.Targets(), values, as.Count())
}

func EvalAssignmentSources(
	ctx *Context,
	sources []Expression,
	count int,
) ([]Result, error) {

	var e error
	r := make([]Result, count)

	for i, s := range sources {

		r[i], e = EvalExpression(ctx, s)
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

	switch target.Kind() {
	case ST_IDENTIFIER:

		id := target.(Identifier).Tk()
		e := checkNotDefined(ctx, id)
		if e != nil {
			return e
		}

		ctx.Set(final, id.Value(), value)
	}

	return nil
}

func checkNotDefined(ctx *Context, tk Token) error {

	v := tk.Value()
	if _, ok := ctx.GetDefined(v); ok {
		msg := fmt.Sprintf("%q cannot be changed, it was defined as constant", v)
		return err.NewBySnippet(msg, tk)
	}

	return nil
}
