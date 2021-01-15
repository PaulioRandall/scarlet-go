package checker

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"

	"github.com/stretchr/testify/require"
)

var tks = []token.Lexeme{
	token.MakeLex2(token.IDENT, "x"), // 0
	token.MakeLex2(token.DELIM, ","),
	token.MakeLex2(token.IDENT, "y"), // 2
	token.MakeLex2(token.DELIM, ","),
	token.MakeLex2(token.IDENT, "z"), // 4
	token.MakeLex2(token.ASSIGN, "<-"),
	token.MakeLex2(token.BOOL, "true"), // 6
	token.MakeLex2(token.DELIM, ","),
	token.MakeLex2(token.NUM, "1"), // 8
	token.MakeLex2(token.DELIM, ","),
	token.MakeLex2(token.NUM, `"Scarlet"`), // 10
	token.MakeLex2(token.TERMINATOR, "\n"), // 11
}

func TestBinding_1(t *testing.T) {

	// x, y <- 1, 2
	in := ast.MakeBinding(
		[]ast.Ident{
			ast.MakeIdent(tks[0]),
			ast.MakeIdent(tks[2]),
		},
		tks[5],
		[]ast.Expr{
			ast.MakeLiteral(tks[6]),
			ast.MakeLiteral(tks[8]),
		},
	)

	e := validateBinding(in)
	require.Nil(t, e, "Unexpected error: %+v", e)
	//require.NotNil(t, e, "Expected error")
}

func TestBinding_2(t *testing.T) {

	// x, y <- 1
	in := ast.MakeBinding(
		[]ast.Ident{
			ast.MakeIdent(tks[0]),
			ast.MakeIdent(tks[2]),
		},
		tks[5],
		[]ast.Expr{
			ast.MakeLiteral(tks[6]),
		},
	)

	e := validateBinding(in)
	//require.Nil(t, e, "Unexpected error: %+v", e)
	require.NotNil(t, e, "Expected error")
}
