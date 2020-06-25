package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
)

func EvalExpression(ctx *Context, expr Expression) (Result, error) {

	switch expr.Kind() {
	case ST_LITERAL:
		return EvalLiteral(expr.(Literal)), nil

	case ST_IDENTIFIER:
		return EvalIdentifier(ctx, expr.(Identifier)), nil
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
