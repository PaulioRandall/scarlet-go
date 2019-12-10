package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValueExpr_1(t *testing.T) {
	// Check it is a type of Expr.
	var _ Expr = ValueExpr{}
}

func TestIdExpr_1(t *testing.T) {
	// Check it is a type of Expr.
	var _ Expr = IdExpr{
		t:  nil,
		id: `abc`,
	}
}

func TestIdExpr_2(t *testing.T) {
	// Ensure Eval returns the correct Value from the context.

	var ex Expr = IdExpr{nil, `abc`}
	v := Value{`xyz`}

	ctx := NewRootCtx()
	ctx.Set(`abc`, v)

	act, e := ex.Eval(ctx)
	require.Nil(t, e)
	assert.Equal(t, v, act)
}

func TestFuncExpr_1(t *testing.T) {
	// Check it is a type of Expr.
	var _ Expr = FuncExpr{}
}

func TestSpellExpr_1(t *testing.T) {
	// Check it is a type of Expr.
	var _ Expr = SpellExpr{}
}

func TestAssignExpr_1(t *testing.T) {
	// Check it is a type of Expr.
	var _ Expr = AssignExpr{}
}

func TestAssignExpr_2(t *testing.T) {
	// Ensure Eval performs a one expression to one ID assignment and returns
	// an empty value.

	exp := Value{`abc`}
	a := Assign{
		Dst: []string{`s`},
		Src: []Expr{ValueExpr{nil, exp}},
	}

	as := AssignExpr{nil, a}

	ctx := NewRootCtx()
	r, e := as.Eval(ctx)

	require.Nil(t, e)
	assert.Empty(t, r)
	assert.Equal(t, exp, ctx.Get(`s`))
}
