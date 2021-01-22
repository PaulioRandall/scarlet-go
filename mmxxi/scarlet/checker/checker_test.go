package checker

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"

	"github.com/stretchr/testify/require"
)

var tks = []token.Lexeme{
	token.MakeLex2(token.IDENT, "x"), // 0
	token.MakeLex2(token.T_BOOL, "B"),
	token.MakeLex2(token.DELIM, ","),
	token.MakeLex2(token.IDENT, "y"), // 3
	token.MakeLex2(token.T_NUM, "N"),
	token.MakeLex2(token.DELIM, ","),
	token.MakeLex2(token.IDENT, "z"), // 6
	token.MakeLex2(token.ASSIGN, "<-"),
	token.MakeLex2(token.BOOL, "true"), // 8
	token.MakeLex2(token.DELIM, ","),
	token.MakeLex2(token.NUM, "1"), // 10
	token.MakeLex2(token.DELIM, ","),
	token.MakeLex2(token.IDENT, "abc"), // 12
	token.MakeLex2(token.TERMINATOR, "\n"),
}

// TODO: Test CheckDefine

func TestCheckAssign_1(t *testing.T) {

	// x B <- true
	in := ast.MakeBinding(
		[]ast.Var{
			ast.MakeVar(tks[0], ast.T_BOOL),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeLiteral(tks[8]),
		},
	).(ast.Assign)

	ctx := makeRootCtx()
	e := checkAssign(ctx, in)
	require.Nil(t, e, "Unexpected error: %+v", e)
	//require.NotNil(t, e, "Expected error")
}

func TestCheckAssign_2(t *testing.T) {

	// x B, y N <- true, 1
	in := ast.MakeBinding(
		[]ast.Var{
			ast.MakeVar(tks[0], ast.T_BOOL),
			ast.MakeVar(tks[3], ast.T_NUM),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeLiteral(tks[8]),
			ast.MakeLiteral(tks[10]),
		},
	).(ast.Assign)

	ctx := makeRootCtx()
	e := checkAssign(ctx, in)
	require.Nil(t, e, "Unexpected error: %+v", e)
	//require.NotNil(t, e, "Expected error")
}

func TestCheckAssign_3(t *testing.T) {

	// x B <- y
	in := ast.MakeBinding(
		[]ast.Var{
			ast.MakeVar(tks[0], ast.T_BOOL),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeIdent(tks[12], ast.T_RESOLVE),
		},
	).(ast.Assign)

	ctx := makeRootCtx()
	ctx.setVar(tks[12].Text, ast.T_BOOL)

	e := checkAssign(ctx, in)
	require.Nil(t, e, "Unexpected error: %+v", e)
	//require.NotNil(t, e, "Expected error")
}

func TestCheckAssign_fail_1(t *testing.T) {

	// x, y N <- true, 1
	in := ast.MakeBinding(
		[]ast.Var{
			ast.MakeVar(tks[0], ast.T_NUM),
			ast.MakeVar(tks[3], ast.T_NUM),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeLiteral(tks[8]), // Bool!
			ast.MakeLiteral(tks[10]),
		},
	).(ast.Assign)

	ctx := makeRootCtx()
	e := checkAssign(ctx, in)
	//require.Nil(t, e, "Unexpected error: %+v", e)
	require.NotNil(t, e, "Expected error")
}

func TestCheckAssign_fail_2(t *testing.T) {

	// Missing expression
	// x B, y N <- true
	in := ast.MakeBinding(
		[]ast.Var{
			ast.MakeVar(tks[0], ast.T_BOOL),
			ast.MakeVar(tks[3], ast.T_NUM),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeLiteral(tks[8]),
		},
	).(ast.Assign)

	ctx := makeRootCtx()
	e := checkAssign(ctx, in)
	//require.Nil(t, e, "Unexpected error: %+v", e)
	require.NotNil(t, e, "Expected error")
}

func TestCheckAssign_fail_3(t *testing.T) {

	// Wrong type being assigned
	// x B <- 1
	in := ast.MakeBinding(
		[]ast.Var{
			ast.MakeVar(tks[0], ast.T_NUM),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeLiteral(tks[8]),
		},
	).(ast.Assign)

	ctx := makeRootCtx()
	e := checkAssign(ctx, in)
	//require.Nil(t, e, "Unexpected error: %+v", e)
	require.NotNil(t, e, "Expected error")
}
