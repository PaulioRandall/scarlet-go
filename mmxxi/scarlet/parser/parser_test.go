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

	exp := []ast.Ident{

		ast.Ident{Snip: in[0].Snippet, Lex: in[0]},
	}

	itr := NewIterator(in)
	act, e := identList(itr)

	require.Nil(t, e, "Unexpected error: %+v", e)
	require.Equal(t, exp, act)
}

func TestIdentList_2(t *testing.T) {

	in := []token.Lexeme{
		token.MakeLex2(token.IDENT, "a"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "b"),
		token.MakeLex2(token.DELIM, ","),
		token.MakeLex2(token.IDENT, "c"),
	}

	exp := []ast.Ident{
		ast.Ident{Snip: in[0].Snippet, Lex: in[0]},
		ast.Ident{Snip: in[2].Snippet, Lex: in[2]},
		ast.Ident{Snip: in[4].Snippet, Lex: in[4]},
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
