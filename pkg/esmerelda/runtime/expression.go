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
