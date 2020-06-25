package runtime

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/number"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"

	"github.com/stretchr/testify/require"
)

func tok(ty token.TokenType, val string) token.Token {
	return token.NewToken(ty, val, 0, 0)
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

	// EVAL an assignment block
	// WITH only one identifier and one expression
	// AND is a const definition
	// THEN the context will be updated to reflect the const assignment

	// a := 1
	given := NewAssignmentBlock(
		false,
		[]Expression{NewIdentifier(tok(token.TK_IDENTIFIER, "a"))},
		[]Expression{NewLiteral(tok(token.TK_NUMBER, "1"))},
		1,
	)

	ctx := NewCtx(nil, true)
	e := EvalAssignmentBlock(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, false, "a", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})
}

func Test_R1_2(t *testing.T) {

	// EVAL an assignment block
	// WITH only one identifier and one expression
	// THEN the context will be updated to reflect the assignment

	// a := 1
	given := NewAssignmentBlock(
		true,
		[]Expression{NewIdentifier(tok(token.TK_IDENTIFIER, "a"))},
		[]Expression{NewLiteral(tok(token.TK_NUMBER, "1"))},
		1,
	)

	ctx := NewCtx(nil, true)
	e := EvalAssignmentBlock(ctx, given)
	require.Nil(t, e)

	requireCtxValue(t, ctx, true, "a", Result{
		typ: RT_NUMBER,
		val: number.New("1"),
	})
}

func Test_R1_3(t *testing.T) {

	// EVAL an assignment block
	// WITH multiple assignments
	// THEN the context will be updated to reflect the assignments

	// a := 1
	given := NewAssignmentBlock(
		false,
		[]Expression{
			NewIdentifier(tok(token.TK_IDENTIFIER, "a")),
			NewIdentifier(tok(token.TK_IDENTIFIER, "b")),
			NewIdentifier(tok(token.TK_IDENTIFIER, "c")),
		},
		[]Expression{
			NewLiteral(tok(token.TK_BOOL, "true")),
			NewLiteral(tok(token.TK_NUMBER, "1")),
			NewLiteral(tok(token.TK_STRING, `"abc"`)),
		},
		3,
	)

	ctx := NewCtx(nil, true)
	e := EvalAssignmentBlock(ctx, given)
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
