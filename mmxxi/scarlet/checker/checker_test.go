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
	token.MakeLex2(token.NUM, `"Scarlet"`), // 12
	token.MakeLex2(token.TERMINATOR, "\n"),
}

func TestBinding_1(t *testing.T) {

	// x B <- true
	in := ast.MakeBinding(
		[]ast.Ident{
			ast.MakeIdent(tks[0], ast.T_BOOL),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeLiteral(tks[8]),
		},
	)

	e := validateBinding(in)
	require.Nil(t, e, "Unexpected error: %+v", e)
	//require.NotNil(t, e, "Expected error")
}

func TestBinding_2(t *testing.T) {

	// x B, y N <- true, 1
	in := ast.MakeBinding(
		[]ast.Ident{
			ast.MakeIdent(tks[0], ast.T_BOOL),
			ast.MakeIdent(tks[3], ast.T_NUM),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeLiteral(tks[8]),
			ast.MakeLiteral(tks[10]),
		},
	)

	e := validateBinding(in)
	require.Nil(t, e, "Unexpected error: %+v", e)
	//require.NotNil(t, e, "Expected error")
}

func TestBinding_3(t *testing.T) {

	// x B, y N <- true
	in := ast.MakeBinding(
		[]ast.Ident{
			ast.MakeIdent(tks[0], ast.T_BOOL),
			ast.MakeIdent(tks[3], ast.T_NUM),
		},
		tks[7],
		[]ast.Expr{
			ast.MakeLiteral(tks[8]),
		},
	)

	e := validateBinding(in)
	//require.Nil(t, e, "Unexpected error: %+v", e)
	require.NotNil(t, e, "Expected error")
}
