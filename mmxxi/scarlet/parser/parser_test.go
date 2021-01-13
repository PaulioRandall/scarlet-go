package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"

	"github.com/stretchr/testify/require"
)

func TestIdentList_1(t *testing.T) {

	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
	}

	exp := []ast.Ident{ast.MakeIdent(in[0])}

	itr := NewIterator(in)
	act, e := identList(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestIdentList_2(t *testing.T) {

	// a, b, c
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "b"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "c"),
	}

	exp := []ast.Ident{
		ast.MakeIdent(in[0]),
		ast.MakeIdent(in[2]),
		ast.MakeIdent(in[4]),
	}

	itr := NewIterator(in)
	act, e := identList(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestIdentList_3(t *testing.T) {

	in := []token.Lexeme{}

	itr := NewIterator(in)
	_, e := identList(itr)

	require.NotNil(t, e, "Expected error")
}

func TestIdentList_4(t *testing.T) {

	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
		token.MakeLex2(token.DELIM, ","),
	}

	itr := NewIterator(in)
	_, e := identList(itr)

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
		[]ast.Ident{ast.MakeIdent(in[0])},
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
		[]ast.Ident{
			ast.MakeIdent(in[0]),
			ast.MakeIdent(in[2]),
			ast.MakeIdent(in[4]),
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

	in := []token.Lexeme{
		token.MakeLex2(token.BOOL, "true"),
	}

	exp := ast.MakeLiteral(in[0])

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
			[]ast.Ident{ast.MakeIdent(in[0])},
			in[1],
			[]ast.Expr{ast.MakeLiteral(in[2])},
		),
	}

	itr := NewIterator(in)
	act, e := parseNext(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestParseNext_Fail_1(t *testing.T) {

	// x, y <- 1
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "x"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "y"),
		token.MakeLex2(token.ASSIGN, "<-"),
		token.MakeLex2(token.NUM, "1"),
		token.MakeLex2(token.TERMINATOR, "\n"),
	}

	itr := NewIterator(in)
	_, e := parseNext(itr)

	require.NotNil(t, e, "Expected error")
}

func TestParseAll_1(t *testing.T) {

	// pi := 3.14
	// x, y, z <- true, 1, "Scarlet"
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "pi"),
		token.MakeLex2(token.DEFINE, ":="),
		token.MakeLex2(token.NUM, "3.14"),
		token.MakeLex2(token.TERMINATOR, "\n"), // 3
		token.MakeLex2(token.IDENT, "x"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "y"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "z"),
		token.MakeLex2(token.ASSIGN, "<-"), // 9
		token.MakeLex2(token.BOOL, "true"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.NUM, "1"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.STR, `"Scarlet"`),
		token.MakeLex2(token.TERMINATOR, "\n"), // 15
	}

	exp := []ast.Tree{
		ast.Tree{
			Root: ast.MakeBinding(
				[]ast.Ident{ast.MakeIdent(in[0])},
				in[1],
				[]ast.Expr{ast.MakeLiteral(in[2])},
			),
		},
		ast.Tree{
			Root: ast.MakeBinding(
				[]ast.Ident{
					ast.MakeIdent(in[4]),
					ast.MakeIdent(in[6]),
					ast.MakeIdent(in[8]),
				},
				in[9],
				[]ast.Expr{
					ast.MakeLiteral(in[10]),
					ast.MakeLiteral(in[12]),
					ast.MakeLiteral(in[14]),
				},
			),
		},
	}

	itr := NewIterator(in)
	act, e := ParseAll(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}
