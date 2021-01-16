package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"

	"github.com/stretchr/testify/require"
)

func TestVarList_1(t *testing.T) {

	// a
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
	}

	exp := []ast.Var{ast.MakeVar(in[0], ast.T_INFER)}

	itr := NewIterator(in)
	act, e := varList(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestVarList_2(t *testing.T) {

	// a N
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
		token.MakeLex2(token.T_NUM, "N"),
	}

	exp := []ast.Var{ast.MakeVar(in[0], ast.T_NUM)}

	itr := NewIterator(in)
	act, e := varList(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestVarList_3(t *testing.T) {

	// a, b, c
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "b"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "c"),
	}

	exp := []ast.Var{
		ast.MakeVar(in[0], ast.T_INFER),
		ast.MakeVar(in[2], ast.T_INFER),
		ast.MakeVar(in[4], ast.T_INFER),
	}

	itr := NewIterator(in)
	act, e := varList(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestVarList_4(t *testing.T) {

	// a B, b N, c S
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"), // 0
		token.MakeLex2(token.T_BOOL, "B"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "b"), // 3
		token.MakeLex2(token.T_NUM, "N"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "c"), // 6
		token.MakeLex2(token.T_STR, "S"),
	}

	exp := []ast.Var{
		ast.MakeVar(in[0], ast.T_BOOL),
		ast.MakeVar(in[3], ast.T_NUM),
		ast.MakeVar(in[6], ast.T_STR),
	}

	itr := NewIterator(in)
	act, e := varList(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestVarList_5(t *testing.T) {

	// a, b N, c S
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"), // 0
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "b"), // 2
		token.MakeLex2(token.T_NUM, "N"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "c"), // 5
		token.MakeLex2(token.T_STR, "S"),
	}

	exp := []ast.Var{
		ast.MakeVar(in[0], ast.T_NUM),
		ast.MakeVar(in[2], ast.T_NUM),
		ast.MakeVar(in[5], ast.T_STR),
	}

	itr := NewIterator(in)
	act, e := varList(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestVarList_6(t *testing.T) {

	in := []token.Lexeme{}

	itr := NewIterator(in)
	_, e := varList(itr)

	require.NotNil(t, e, "Expected error")
}

func TestVarList_7(t *testing.T) {

	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
		token.MakeLex2(token.DELIM, ","),
	}

	itr := NewIterator(in)
	_, e := varList(itr)

	require.NotNil(t, e, "Expected error")
}

func TestBinding_1(t *testing.T) {

	// x := 1
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "x"),
		token.MakeLex2(token.DEFINE, ":="),
		token.MakeLex2(token.NUM, "1"),
	}

	exp := ast.MakeBinding(
		[]ast.Var{ast.MakeVar(in[0], ast.T_INFER)},
		in[1],
		[]ast.Expr{ast.MakeLiteral(in[2])},
	)

	itr := NewIterator(in)
	act, e := binding(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestBinding_2(t *testing.T) {

	// a, b, c <- 1, 2, 3
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "b"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "c"),
		token.MakeLex2(token.ASSIGN, "<-"),
		token.MakeLex2(token.NUM, "1"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.NUM, "2"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.NUM, "3"),
	}

	exp := ast.MakeBinding(
		[]ast.Var{
			ast.MakeVar(in[0], ast.T_INFER),
			ast.MakeVar(in[2], ast.T_INFER),
			ast.MakeVar(in[4], ast.T_INFER),
		},
		in[5],
		[]ast.Expr{
			ast.MakeLiteral(in[6]),
			ast.MakeLiteral(in[8]),
			ast.MakeLiteral(in[10]),
		},
	)

	itr := NewIterator(in)
	act, e := binding(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestExpression_1(t *testing.T) {

	// true
	in := []token.Lexeme{
		token.MakeLex2(token.BOOL, "true"),
	}

	exp := ast.MakeLiteral(in[0])

	itr := NewIterator(in)
	act, e := expression(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestExpression_2(t *testing.T) {

	// abc
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "abc"),
	}

	exp := ast.MakeIdent(in[0], ast.T_INFER)

	itr := NewIterator(in)
	act, e := expression(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestParseNext_1(t *testing.T) {

	// x <- 1
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "x"),
		token.MakeLex2(token.DEFINE, "<-"),
		token.MakeLex2(token.NUM, "1"),
		token.MakeLex2(token.TERMINATOR, "\n"),
	}

	exp := ast.Tree{
		Root: ast.MakeBinding(
			[]ast.Var{ast.MakeVar(in[0], ast.T_INFER)},
			in[1],
			[]ast.Expr{ast.MakeLiteral(in[2])},
		),
	}

	itr := NewIterator(in)
	act, e := parseNext(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestParseAll_1(t *testing.T) {

	// pi N := 3.14
	// x, y B <- true, false
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "pi"), // 0
		token.MakeLex2(token.T_NUM, "N"),
		token.MakeLex2(token.DEFINE, ":="),
		token.MakeLex2(token.NUM, "3.14"), // 3
		token.MakeLex2(token.TERMINATOR, "\n"),
		token.MakeLex2(token.IDENT, "x"), // 5
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "y"), // 7
		token.MakeLex2(token.T_BOOL, "B"),
		token.MakeLex2(token.ASSIGN, "<-"),
		token.MakeLex2(token.BOOL, "true"), // 10
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.BOOL, "false"), // 12
		token.MakeLex2(token.TERMINATOR, "\n"),
	}

	exp := []ast.Tree{
		ast.Tree{
			Root: ast.MakeBinding(
				[]ast.Var{ast.MakeVar(in[0], ast.T_NUM)},
				in[2],
				[]ast.Expr{ast.MakeLiteral(in[3])},
			),
		},
		ast.Tree{
			Root: ast.MakeBinding(
				[]ast.Var{
					ast.MakeVar(in[5], ast.T_BOOL),
					ast.MakeVar(in[7], ast.T_BOOL),
				},
				in[9],
				[]ast.Expr{
					ast.MakeLiteral(in[10]),
					ast.MakeLiteral(in[12]),
				},
			),
		},
	}

	itr := NewIterator(in)
	act, e := ParseAll(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}
