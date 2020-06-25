package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
)

func EvalExpression(ctx *Context, expr Expression) (Result, error) {

	switch expr.Kind() {
	case ST_LITERAL:
		return EvalLiteral(expr.(Literal)), nil
	}

	panic(err.NewBySnippet("Unknown expression type", expr))
}

func EvalLiteral(lit Literal) Result {
	return ResultOf(lit.Tk())
}
