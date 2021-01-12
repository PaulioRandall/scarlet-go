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

func TestBinder_1(t *testing.T) {

	// x <- 1
	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "x"),
		token.MakeLex2(token.DEFINE, "<-"),
		token.MakeLex2(token.NUM, "1"),
	}

	exp := ast.MakeBinder(
		[]ast.Ident{ast.MakeIdent(in[0])},
		in[1],
		[]ast.Expr{ast.MakeLiteral(in[2])},
	)

	itr := NewIterator(in)
	act, e := binder(itr)

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

	var stmt ast.Stmt = ast.MakeBinder(
		[]ast.Ident{ast.MakeIdent(in[0])},
		in[1],
		[]ast.Expr{ast.MakeLiteral(in[2])},
	)
	exp := ast.Tree{Root: stmt}

	itr := NewIterator(in)
	act, e := parseNext(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}
