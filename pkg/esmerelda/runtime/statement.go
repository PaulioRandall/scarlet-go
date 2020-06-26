package runtime

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

func EvalStatements(ctx *Context, sts []Expr) error {

	for _, st := range sts {
		e := EvalStatement(ctx, st)
		if e != nil {
			return e
		}
	}

	return nil
}

func EvalStatement(ctx *Context, st Expr) error {

	switch st.Kind() {
	case ST_ASSIGN_BLOCK:
		return EvalAssignBlock(ctx, st.(AssignBlock))

	case ST_ASSIGN:
		return EvalAssign(ctx, st.(Assign))

	case ST_GUARD:
		_, e := EvalGuard(ctx, st.(Guard))
		return e

	case ST_WHEN:
		return EvalWhen(ctx, st.(When))
	}

	panic(err.NewBySnippet("Unknown statement type", st))
}

func EvalAssignBlock(ctx *Context, as AssignBlock) error {

	values, e := EvalAssignSources(ctx, as.Sources(), as.Count())
	if e != nil {
		return e
	}

	return doAssignments(ctx, as.Const(), as.Targets(), values, as.Count())
}

func EvalAssign(ctx *Context, as Assign) error {

	v, e := EvalExpr(ctx, as.Source())
	if e != nil {
		return e
	}

	return doAssignment(ctx, as.Const(), as.Target(), v)
}

func EvalAssignSources(
	ctx *Context,
	sources []Expr,
	count int,
) ([]Result, error) {

	var e error
	r := make([]Result, count)

	for i, s := range sources {

		r[i], e = EvalExpr(ctx, s)
		if e != nil {
			return nil, e
		}
	}

	return r, nil
}

func doAssignments(
	ctx *Context,
	final bool,
	targets []Expr,
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
	target Expr,
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

func EvalBlock(ctx *Context, b Block) error {
	return EvalStatements(ctx, b.Stats())
}

func EvalGuard(ctx *Context, g Guard) (bool, error) {

	r, e := EvalExpr(ctx, g.Condition())
	if e != nil {
		return false, e
	}

	conditionMeet, ok := r.Bool()
	if !ok {
		msg := "Guard condition requires a boolean result"
		return false, err.NewBySnippet(msg, g.Condition())
	}

	if !conditionMeet {
		return false, nil
	}

	ctx = NewCtx(ctx, false)
	e = EvalBlock(ctx, g.Body())
	if e != nil {
		return false, e
	}

	return true, nil
}

func EvalWhen(ctx *Context, w When) error {

	subject, e := EvalExpr(ctx, w.Init())
	if e != nil {
		return e
	}

	ctx = NewCtx(ctx, false)
	id := w.Subject().Value()
	ctx.SetVar(id, subject)

	for _, wc := range w.Cases() {

		var match bool

		switch wc.Kind() {
		case ST_GUARD:
			match, e = EvalGuard(ctx, wc.(Guard))

		case ST_WHEN_CASE:
			match, e = EvalWhenCase(ctx, wc.(WhenCase), subject)

		default:
			return err.NewBySnippet("Unknown when case type", wc)
		}

		if match || e != nil {
			return e
		}
	}

	return nil
}

func EvalWhenCase(ctx *Context, wc WhenCase, subject Result) (bool, error) {

	r, e := EvalExpr(ctx, wc.Object())
	if e != nil {
		return false, e
	}

	if r.NotEqual(subject) {
		return false, nil
	}

	e = EvalBlock(ctx, wc.Body())
	if e != nil {
		return false, e
	}

	return true, nil
}
