package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueExpr_1(t *testing.T) {
	// Check it is a type of Expr.
	var _ Expr = ValueExpr{}
}

func TestIdExpr_1(t *testing.T) {
	// Check it is a type of Expr.
	var _ Expr = IdExpr(`abc`)
}

func TestIdExpr_2(t *testing.T) {
	// Ensure Eval returns the correct Value from the context.

	var ex Expr = IdExpr(`abc`)
	v := NewValue(STR, `xyz`)

	ctx := NewRootCtx()
	ctx.Set(`abc`, v)

	act := ex.Eval(ctx)
	assert.Equal(t, v, act)
}
