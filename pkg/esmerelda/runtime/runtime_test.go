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

func Test_R1_7(t *testing.T) {

	// GIVEN a context with a const definition
	// EVAL an assignment block
	// WITH a const definition override assignment
	// THEN an error will be returned

	// def a := 1
	given := quickAssignBlock(true, "a", 1)

	ctx := NewCtx(nil, true)
	ctx.SetDefinition("a", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})

	e := EvalAssignBlock(ctx, given)
	require.NotNil(t, e, "Should not be able to reassign defintions")
}

func Test_R1_8(t *testing.T) {

	// GIVEN an empty context
	// EVAL an assignment block
	// WITH a void as target
	// THEN no error will be returned
	// AND no entry placed in the context

	// a := 1
	given := NewAssignBlock(
		false,
		[]Expr{NewVoid(tok(TK_VOID, "_"))},
		[]Expr{NewLiteral(tok(TK_NUMBER, "1"))},
		1,
	)

	ctx := NewCtx(nil, true)
	e := EvalAssignBlock(ctx, given)

	require.Nil(t, e)
	require.Empty(t, ctx.Locals(), "Void assignment targets should be ignored")
	require.Empty(t, ctx.Definitions(), "Void assignment targets should be ignored")
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

func Test_R5_1(t *testing.T) {

	// EVAL a func call
	// WITH no arguments
	// AND no statements in the body
	// THEN no errors are returned

	// f := F() {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{},
		[]Token{},
		Expr(quickBlock()),
	)

	// f()
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, "}"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	e := EvalStatement(ctx, given)

	require.Nil(t, e)
}

func Test_R5_2(t *testing.T) {

	// EVAL a func call
	// WITH multiple input arguments
	// AND no statements in the body
	// THEN no errors are returned

	// f := F(a, b, c) {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		[]Token{},
		Expr(quickBlock()),
	)

	// f(1, true, "abc")
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, ")"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_BOOL, "true")),
			NewLiteral(tok(TK_STRING, `"abc"`)),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	e := EvalStatement(ctx, given)

	require.Nil(t, e)
}

func Test_R5_3(t *testing.T) {

	// EVAL a func call
	// WITH multiple output arguments
	// AND no statements in the body
	// THEN no errors are returned

	// f := F(-> a, b, c) {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{},
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		Expr(quickBlock()),
	)

	// f()
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, ")"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	e := EvalStatement(ctx, given)

	require.Nil(t, e)
}

func Test_R5_4(t *testing.T) {

	// EVAL a func call
	// WITH multiple input and output arguments
	// AND no statements in the body
	// THEN no errors are returned

	// f := F(a, b, c-> a, b, c) {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		Expr(quickBlock()),
	)

	// f(1, true, "abc")
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, ")"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_BOOL, "true")),
			NewLiteral(tok(TK_STRING, `"abc"`)),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	e := EvalStatement(ctx, given)

	require.Nil(t, e)
}

func Test_R5_5(t *testing.T) {

	// EVAL a func call
	// WITH not enough input arguments
	// AND no statements in the body
	// THEN an error is returned

	// f := F(a, b, c) {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		[]Token{},
		Expr(quickBlock()),
	)

	// f(1, true)
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, ")"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_BOOL, "true")),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	e := EvalStatement(ctx, given)
	require.NotNil(t, e)
}

func Test_R5_6(t *testing.T) {

	// EVAL a func call
	// WITH too many input arguments
	// AND no statements in the body
	// THEN an error is returned

	// f := F(a, b) {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
		},
		[]Token{},
		Expr(quickBlock()),
	)

	// f(1, true, "abc")
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, ")"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_BOOL, "true")),
			NewLiteral(tok(TK_STRING, `"abc"`)),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	e := EvalStatement(ctx, given)
	require.NotNil(t, e)
}

func Test_R5_7(t *testing.T) {

	// EVAL a func call
	// WITH input arguments as output arguments
	// AND no statements in the body
	// THEN the output results equal the input arguments

	// f := F(a, b, c -> a, b, c) {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		Expr(quickBlock()),
	)

	// f(1, true, "abc")
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, ")"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_BOOL, "true")),
			NewLiteral(tok(TK_STRING, `"abc"`)),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	r, e := EvalExpr(ctx, given)
	require.Nil(t, e)

	require.True(t, r.Is(RT_TUPLE), "Expected tuple as output")
	results, _ := r.Tuple()

	require.Equal(t, 3, len(results), "Expected exactly 3 function outputs")

	exp := Result{typ: RT_NUMBER, val: number.New("1")}
	require.Equal(t, exp, results[0])

	exp = Result{typ: RT_BOOL, val: true}
	require.Equal(t, exp, results[1])

	exp = Result{typ: RT_STRING, val: "abc"}
	require.Equal(t, exp, results[2])
}

func Test_R5_8(t *testing.T) {

	// EVAL a func call
	// WITH input arguments
	// AND output arguments
	// AND a statement assigning inputs variables to outputs variables
	// THEN the output results equal the input arguments

	// f := F(a, b, c -> x, y, z) {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
			tok(TK_IDENTIFIER, "c"),
		},
		[]Token{
			tok(TK_IDENTIFIER, "x"),
			tok(TK_IDENTIFIER, "y"),
			tok(TK_IDENTIFIER, "z"),
		},
		quickBlock(NewAssignBlock(
			false,
			[]Expr{
				NewIdentifier(tok(TK_IDENTIFIER, "x")),
				NewIdentifier(tok(TK_IDENTIFIER, "y")),
				NewIdentifier(tok(TK_IDENTIFIER, "z")),
			},
			[]Expr{
				NewIdentifier(tok(TK_IDENTIFIER, "a")),
				NewIdentifier(tok(TK_IDENTIFIER, "b")),
				NewIdentifier(tok(TK_IDENTIFIER, "c")),
			},
			3,
		)),
	)

	// f(1, true, "abc")
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, ")"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{
			NewLiteral(tok(TK_NUMBER, "1")),
			NewLiteral(tok(TK_BOOL, "true")),
			NewLiteral(tok(TK_STRING, `"abc"`)),
		},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	r, e := EvalExpr(ctx, given)
	require.Nil(t, e)

	require.True(t, r.Is(RT_TUPLE), "Expected tuple as output")
	results, _ := r.Tuple()

	require.Equal(t, 3, len(results), "Expected exactly 3 function outputs")

	exp := Result{typ: RT_NUMBER, val: number.New("1")}
	require.Equal(t, exp, results[0])

	exp = Result{typ: RT_BOOL, val: true}
	require.Equal(t, exp, results[1])

	exp = Result{typ: RT_STRING, val: "abc"}
	require.Equal(t, exp, results[2])
}

func Test_R5_9(t *testing.T) {

	// EVAL a func call
	// WITH an unassigned output argument
	// AND no statements in the body
	// THEN a single void result is returned

	// f := F(-> a) {}
	setup := NewFuncDef(
		tok(TK_FUNCTION, "F"),
		[]Token{},
		[]Token{
			tok(TK_IDENTIFIER, "a"),
		},
		Expr(quickBlock()),
	)

	// f()
	given := NewFuncCall(
		tok(TK_PAREN_CLOSE, ")"),
		NewIdentifier(tok(TK_IDENTIFIER, "f")),
		[]Expr{},
	)

	ctx := NewCtx(nil, true)
	ctx.SetLocal("f", Result{
		typ: RT_FUNC_DEF,
		val: setup,
	})

	r, e := EvalExpr(ctx, given)
	require.Nil(t, e)

	require.True(t, r.Is(RT_TUPLE), "Expected tuple as output")
	results, _ := r.Tuple()

	require.Equal(t, 1, len(results), "Expected exactly 1 function output")

	exp := Result{typ: RT_VOID, val: VoidResult{}}
	require.Equal(t, exp, results[0])
}
