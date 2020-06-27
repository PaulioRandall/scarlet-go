package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
)

func EvalExpr(ctx *Context, expr Expr) (Result, error) {

	switch expr.Kind() {
	case ST_LITERAL:
		return EvalLiteral(expr.(Literal)), nil

	case ST_IDENTIFIER:
		return EvalIdentifier(ctx, expr.(Identifier)), nil

	case ST_FUNC_DEF:
		return EvalFuncDef(ctx, expr.(FuncDef)), nil

	case ST_EXPR_FUNC:
		return EvalExprFunc(ctx, expr.(ExprFunc)), nil

	case ST_NEGATION:
		return EvalNegation(ctx, expr.(Negation))

	case ST_FUNC_CALL:
		return EvalFuncCall(ctx, expr.(FuncCall))
	}

	panic(err.NewBySnippet("Unknown expression type", expr))
}

func EvalLiteral(lit Literal) Result {
	return ResultOf(lit.Tk())
}

func EvalIdentifier(ctx *Context, id Identifier) Result {

	v, ok := ctx.Get(id.Tk().Value())
	if ok {
		return v
	}

	return Result{
		typ: RT_VOID,
		val: VoidResult{},
	}
}

func EvalFuncDef(ctx *Context, f FuncDef) Result {
	return Result{
		typ: RT_FUNC_DEF,
		val: f,
	}
}

func EvalExprFunc(ctx *Context, f ExprFunc) Result {
	return Result{
		typ: RT_EXPR_FUNC_DEF,
		val: f,
	}
}

func EvalNegation(ctx *Context, neg Negation) (Result, error) {

	subject, e := EvalExpr(ctx, neg.Expr())
	if e != nil {
		return Result{}, e
	}

	if subject.IsNot(RT_BOOL) && subject.IsNot(RT_NUMBER) {
		return Result{}, err.NewBySnippet(
			"Expected boolean or number result",
			neg.Expr(),
		)
	}

	subject.Negate()
	return subject, nil
}

func EvalFuncCall(ctx *Context, fc FuncCall) (Result, error) {

	fResult, e := EvalExpr(ctx, fc.Func())
	if e != nil {
		return Result{}, e
	}

	f, ok := fResult.Func()
	if !ok {
		return Result{}, err.NewBySnippet("Not a function", fc)
	}

	if len(f.Inputs()) != len(fc.Args()) {
		return Result{}, err.NewBySnippet("Wrong number of function arguments", fc)
	}

	inputs, e := evalFuncCallArgs(ctx, fc.Args())
	if e != nil {
		return Result{}, e
	}

	outputs, e := evalFuncBody(ctx, f, inputs)
	if e != nil {
		return Result{}, e
	}

	r := Result{
		typ: RT_TUPLE,
		val: outputs,
	}
	return r, nil
}

func evalFuncCallArgs(ctx *Context, args []Expr) ([]Result, error) {

	r := []Result{}

	for _, a := range args {

		v, e := EvalExpr(ctx, a)
		if e != nil {
			return nil, e
		}

		r = append(r, v)
	}

	return r, nil
}

func evalFuncBody(ctx *Context, f FuncDef, inputs []Result) ([]Result, error) {

	funcCtx := NewCtx(ctx, true)
	for i, in := range f.Inputs() {
		funcCtx.SetLocal(in.Value(), inputs[i])
	}

	e := EvalBlock(funcCtx, f.Body().(Block))
	if e != nil {
		return nil, e
	}

	r := []Result{}

	for _, out := range f.Outputs() {

		v, ok := funcCtx.GetVar(out.Value())
		if !ok {
			v = Result{
				typ: RT_VOID,
				val: VoidResult{},
			}
		}

		r = append(r, v)
	}

	return r, nil
}
