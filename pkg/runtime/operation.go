package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func EvalOperation(ctx *Context, op statement.Operation) Value {

	tk := op.Operator

	switch tk.Type {
	case token.ADD:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Number(left + right)

	case token.SUBTRACT:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Number(left - right)

	case token.MULTIPLY:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Number(left * right)

	case token.DIVIDE:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Number(left / right)

	case token.REMAINDER:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Number(int64(left) % int64(right))

	case token.LESS_THAN:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Bool(left < right)

	case token.MORE_THAN:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Bool(left > right)

	case token.LESS_THAN_OR_EQUAL:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Bool(left <= right)

	case token.MORE_THAN_OR_EQUAL:
		left, right := EvalNumbers(ctx, op.Left, op.Right)
		return Bool(left >= right)

	case token.AND:
		left, right := EvalBools(ctx, op.Left, op.Right)
		return Bool(left && right)

	case token.OR:
		left, right := EvalBools(ctx, op.Left, op.Right)
		return Bool(left || right)

	case token.EQUAL:
		left, right := EvalValues(ctx, op.Left, op.Right)
		return Bool(left == right)

	case token.NOT_EQUAL:
		left, right := EvalValues(ctx, op.Left, op.Right)
		return Bool(left != right)
	}

	panic(err("EvalOperation", tk, "Unknown operation type"))
}

func EvalValues(ctx *Context, left, right statement.Expression) (Value, Value) {
	return EvalExpression(ctx, left), EvalExpression(ctx, right)
}

func EvalNumbers(ctx *Context, left, right statement.Expression) (float64, float64) {
	return EvalNumber(ctx, left), EvalNumber(ctx, right)
}

func EvalNumber(ctx *Context, ex statement.Expression) float64 {
	if v, ok := EvalExpression(ctx, ex).(Number); ok {
		return float64(v)
	}

	panic(err("EvalNumber", ex.Token(), "Expected Number as result"))
}

func EvalBools(ctx *Context, left, right statement.Expression) (bool, bool) {
	return EvalBool(ctx, left), EvalBool(ctx, right)
}

func EvalBool(ctx *Context, ex statement.Expression) bool {
	if v, ok := EvalExpression(ctx, ex).(Bool); ok {
		return bool(v)
	}

	panic(err("EvalBool", ex.Token(), "Expected Bool as result"))
}
