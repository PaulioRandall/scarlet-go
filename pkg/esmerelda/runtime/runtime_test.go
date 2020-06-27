package runtime

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/number"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"

	"github.com/stretchr/testify/require"
)

func tok(ty TokenType, val string) Token {
	return NewToken(ty, val, 0, 0)
}

func quickBlock(exprs ...Expr) Block {
	return NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "{"),
		exprs,
	)
}

func quickAssignBlock(final bool, target string, source interface{}) AssignBlock {

	var src Token

	switch v := source.(type) {
	case bool:
		src = tok(TK_BOOL, strconv.FormatBool(v))
	case int:
		src = tok(TK_NUMBER, strconv.Itoa(v))
	case float64:
		src = tok(TK_NUMBER, strconv.FormatFloat(v, 'f', -1, 64))
	case string:
		src = tok(TK_STRING, v)
	}

	return NewAssignBlock(
		final,
		[]Expr{NewIdentifier(tok(TK_IDENTIFIER, target))},
		[]Expr{NewLiteral(src)},
		1,
	)
}

func requireCtxValue(t *testing.T, ctx *Context, final bool, id string, exp Result) {

	actDef, okDef := ctx.GetDefined(id)
	actVar, okVar := ctx.GetVar(id)

	if final {

		require.True(t, okDef, "Context missing definintion %q", exp.String())
		require.False(t, okVar, "Unexpected context variable %q", exp.String())
		require.Empty(t, actVar)

		require.Equal(t, exp, actDef, "Wrong value in context, have %q, want %q",
			actVar.String(),
			exp.String(),
		)

	} else {

		require.True(t, okVar, "Context missing variable %q", exp.String())
		require.False(t, okDef, "Unexpected context definition %q", exp.String())
		require.Empty(t, actDef)

		require.Equal(t, exp, actVar, "Wrong value in context, have %q, want %q",
			actVar.String(),
			exp.String(),
		)
	}

}

func Test_R1_1(t *testing.T) {

	// GIVEN an empty context
	// EVAL an assignment block
	// WITH one literal assigned to one target
	// AND is a const definition
	// THEN the context will be updated to reflect the const assignment

	// a := 1
	given := quickAssignBlock(false, "a", 1)

	ctx := NewCtx(nil, true)
	e := EvalAssignBlock(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "a", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})
}

func Test_R1_2(t *testing.T) {

	// GIVEN an empty context
	// EVAL an assignment block
	// WITH a const definition assignment
	// THEN the context will be updated to reflect the assignment

	// def a := 1
	given := quickAssignBlock(true, "a", 1)

	ctx := NewCtx(nil, true)
	e := EvalAssignBlock(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, true, "a", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})
}

func Test_R1_3(t *testing.T) {

	// GIVEN an empty context
	// EVAL an assignment block
	// WITH multiple assignments
	// THEN the context will be updated to reflect the assignments

	// a, b, c := true, 1, "abc"
	given := NewAssignBlock(
		false,
		[]Expr{
			NewIdentifier(tok(TK_IDENTIFIER, "a")),
			NewIdentifier(tok(TK_IDENTIFIER, "b")),
			NewIdentifier(tok(TK_IDENTIFIER, "c")),
		},
		[]Expr{
			NewLiteral(tok(TK_BOOL, "true")),
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_STRING, `"abc"`)),
		},
		3,
	)

	ctx := NewCtx(nil, true)
	e := EvalAssignBlock(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "a", Result{
		typ: RT_BOOL,
		val: true,
	})

	requireCtxValue(t, ctx, false, "b", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})

	requireCtxValue(t, ctx, false, "c", Result{
		typ: RT_STRING,
		val: "abc",
	})
}

func Test_R1_4(t *testing.T) {

	// GIVEN a non-empty context
	// EVAL an assignment block
	// WITH an identifier assigned to another
	// THEN the context will be updated to reflect the assignment

	// a := b
	given := NewAssignBlock(
		false,
		[]Expr{NewIdentifier(tok(TK_IDENTIFIER, "a"))},
		[]Expr{NewIdentifier(tok(TK_IDENTIFIER, "b"))},
		1,
	)

	id := Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	}

	ctx := NewCtx(nil, true)
	ctx.SetLocal("b", id)

	e := EvalAssignBlock(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "a", id)
	requireCtxValue(t, ctx, false, "b", id)
}

func Test_R1_5(t *testing.T) {

	// GIVEN an empty context
	// EVAL an assignment block
	// WITH a function assigned to an identifier
	// THEN the context will be updated to reflect the assignment

	f := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{},
		[]Token{},
		NewBlock(
			tok(TK_BLOCK_OPEN, "{"),
			tok(TK_BLOCK_CLOSE, "}"),
			[]Expr{},
		),
	)

	// a := F() {}
	given := NewAssignBlock(
		false,
		[]Expr{NewIdentifier(tok(TK_IDENTIFIER, "a"))},
		[]Expr{f},
		1,
	)

	ctx := NewCtx(nil, true)
	e := EvalAssignBlock(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "a", Result{
		typ: RT_FUNC_DEF,
		val: f,
	})
}

func Test_R1_6(t *testing.T) {

	// GIVEN an empty context
	// EVAL an assignment block
	// WITH a function assigned to an identifier
	// THEN the context will be updated to reflect the assignment

	f := NewExprFunc(
		tok(TK_EXPR_FUNC, "E"),
		[]Token{},
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	// a := E() 1
	given := NewAssignBlock(
		false,
		[]Expr{NewIdentifier(tok(TK_IDENTIFIER, "a"))},
		[]Expr{f},
		1,
	)

	ctx := NewCtx(nil, true)
	e := EvalAssignBlock(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "a", Result{
		typ: RT_EXPR_FUNC_DEF,
		val: f,
	})
}

func Test_R2_1(t *testing.T) {

	// EVAL a negation expression
	// WITH a number
	// THEN the negated number is returned

	// -1
	given := NewNegation(
		NewLiteral(tok(TK_NUMBER, "1")),
	)

	exp := Result{
		typ: RT_NUMBER,
		val: number.New("-1"),
	}

	ctx := NewCtx(nil, true)
	act, e := EvalExpr(ctx, given)

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func Test_R2_2(t *testing.T) {

	// EVAL a negation expression
	// WITH a bool
	// THEN the negated bool is returned

	// -true
	given := NewNegation(
		NewLiteral(tok(TK_BOOL, "true")),
	)

	exp := Result{
		typ: RT_BOOL,
		val: false,
	}

	ctx := NewCtx(nil, true)
	act, e := EvalExpr(ctx, given)

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func Test_R3_1(t *testing.T) {

	// EVAL a guard statement
	// WITH a true condition
	// THEN the body is evaluated

	// [true] { a := 1 }
	given := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewLiteral(tok(TK_BOOL, "true")),
		quickBlock(quickAssignBlock(false, "a", 1)),
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("a", Result{
		typ: RT_NUMBER,
		val: number.New("0"),
	})

	e := EvalStatement(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "a", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})
}

func Test_R3_2(t *testing.T) {

	// EVAL a guard statement
	// WITH a false condition
	// THEN the body is not evaluated

	// [false] { a := 1 }
	given := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewLiteral(tok(TK_BOOL, "false")),
		quickBlock(quickAssignBlock(false, "a", 1)),
	)

	ctx := NewCtx(nil, true)
	e := EvalStatement(ctx, given)

	require.Nil(t, e)

	require.Empty(t, ctx.Definitions(), "Expected empty context")
	require.Empty(t, ctx.Locals(), "Expected empty context")
}

func Test_R3_3(t *testing.T) {

	// EVAL a guard statement
	// WITH a false condition
	// THEN the body is not evaluated

	// ["abc"] {}
	given := NewGuard(
		tok(TK_GUARD_OPEN, "["),
		NewLiteral(tok(TK_STRING, `"abc"`)),
		quickBlock(),
	)

	ctx := NewCtx(nil, true)
	e := EvalStatement(ctx, given)

	require.NotNil(t, e, "Expected error: invalid condition expression")
}

func Test_R4_1(t *testing.T) {

	// EVAL a when statement
	// WITH a true subject
	// AND no cases
	// THEN no case is evaluated

	// when x := true {
	//
	// }
	given := NewWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "{"),
		tok(TK_IDENTIFIER, "x"),
		NewLiteral(tok(TK_BOOL, `true`)),
		[]Expr{},
	)

	ctx := NewCtx(nil, true)
	e := EvalStatement(ctx, given)

	require.Nil(t, e)
}

func Test_R4_2(t *testing.T) {

	// GIVEN a context with an intiialised variable
	// EVAL a when statement
	// WITH a 'true' subject
	// AND one 'false' when case then one 'true' when case
	// THEN the 'true' case is evaluated

	// y := false
	// when x := true {
	//   false: y := 0
	//   true:  y := 1
	// }
	given := NewWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "{"),
		tok(TK_IDENTIFIER, "x"),
		NewLiteral(tok(TK_BOOL, `true`)),
		[]Expr{
			NewWhenCase(
				NewLiteral(tok(TK_BOOL, `false`)),
				quickBlock(quickAssignBlock(false, "y", 0)),
			),
			NewWhenCase(
				NewLiteral(tok(TK_BOOL, `true`)),
				quickBlock(quickAssignBlock(false, "y", 1)),
			),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("y", Result{
		typ: RT_BOOL,
		val: false,
	})

	e := EvalStatement(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "y", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})
}

func Test_R4_3(t *testing.T) {

	// GIVEN a context with an intiialised variable
	// EVAL a when statement
	// WITH a false evaluating guard then a true evaluating guard
	// THEN the true guard body is evaluated

	// y := false
	// when x := true {
	//   [false]: y := 0
	//   [true]:  y := 1
	// }
	given := NewWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "{"),
		tok(TK_IDENTIFIER, "x"),
		NewLiteral(tok(TK_BOOL, `true`)),
		[]Expr{
			NewGuard(
				tok(TK_GUARD_OPEN, "["),
				NewLiteral(tok(TK_BOOL, "false")),
				quickBlock(quickAssignBlock(false, "y", 0)),
			),
			NewGuard(
				tok(TK_GUARD_OPEN, "["),
				NewLiteral(tok(TK_BOOL, "true")),
				quickBlock(quickAssignBlock(false, "y", 1)),
			),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("y", Result{
		typ: RT_BOOL,
		val: false,
	})

	e := EvalStatement(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "y", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})
}

func Test_R4_4(t *testing.T) {

	// GIVEN a context with an intiialised variable
	// EVAL a when statement
	// WITH a 'true' subject
	// AND one 'true' when case with an inline block body
	// THEN the true guard body is evaluated

	// y := false
	// when x := true {
	//   [true]:  y := 1
	// }
	given := NewWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "{"),
		tok(TK_IDENTIFIER, "x"),
		NewLiteral(tok(TK_BOOL, `true`)),
		[]Expr{
			NewWhenCase(
				NewLiteral(tok(TK_BOOL, `true`)),
				NewUndelimBlock(
					[]Expr{quickAssignBlock(false, "y", 1)},
				),
			),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("y", Result{
		typ: RT_BOOL,
		val: false,
	})

	e := EvalStatement(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "y", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})
}

func Test_R4_5(t *testing.T) {

	// GIVEN a context with an intiialised variable
	// EVAL a when statement
	// WITH a several interlaced when and guard cases
	// THEN the correct case body is evaluated

	// y := 0
	// when x := 4 {
	//   0:       y := false
	//   1:       y := false
	//   [false]: y := false
	//   3:       y := false
	//   [-true]: y := false
	//   4:       y := true
	//   5:       y := false
	//   [true]:  y := false
	// }

	notEvaluated := quickBlock(quickAssignBlock(false, "y", false))

	given := NewWhen(
		tok(TK_WHEN, "when"),
		tok(TK_BLOCK_CLOSE, "{"),
		tok(TK_IDENTIFIER, "x"),
		NewLiteral(tok(TK_NUMBER, "4")),
		[]Expr{
			NewWhenCase(
				NewLiteral(tok(TK_NUMBER, "0")),
				notEvaluated,
			),
			NewWhenCase(
				NewLiteral(tok(TK_NUMBER, "1")),
				notEvaluated,
			),
			NewGuard(
				tok(TK_GUARD_OPEN, "["),
				NewLiteral(tok(TK_BOOL, "false")),
				notEvaluated,
			),
			NewWhenCase(
				NewLiteral(tok(TK_NUMBER, "3")),
				notEvaluated,
			),
			NewGuard(
				tok(TK_GUARD_OPEN, "["),
				NewNegation(NewLiteral(tok(TK_BOOL, "true"))),
				notEvaluated,
			),
			NewWhenCase(
				NewLiteral(tok(TK_NUMBER, "4")),
				quickBlock(quickAssignBlock(false, "y", true)),
			),
			NewWhenCase(
				NewLiteral(tok(TK_NUMBER, "5")),
				notEvaluated,
			),
			NewGuard(
				tok(TK_GUARD_OPEN, "["),
				NewLiteral(tok(TK_BOOL, "true")),
				notEvaluated,
			),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("y", Result{
		typ: RT_NUMBER,
		val: number.New("0"),
	})

	e := EvalStatement(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "y", Result{
		typ: RT_BOOL,
		val: true,
	})
}
