package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/number"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
	"github.com/PaulioRandall/scarlet-go/token2/tokentest"

	"github.com/stretchr/testify/require"
)

func positionLexemes(lexs ...lexeme.Lexeme) []lexeme.Lexeme {
	tm := position.TextMarker{}
	for i := 0; i < len(lexs); i++ {
		v := lexs[i].Val
		lexs[i].Snippet = tm.Snippet(v)
		tm.Advance(v)
	}
	return lexs
}

func requireNodes(t *testing.T, exp, act []Node) {
	require.Equal(t, len(exp), len(act))
	for i, n := range act {
		require.Equal(t, exp[i], n)
	}
}

func TestParse(t *testing.T) {

	// x := 1
	in := positionLexemes(
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("1", token.NUMBER),
	)

	exp := []Node{
		SingleAssign{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[2].Snippet.End,
			},
			Left: Ident{
				Snippet: in[0].Snippet,
				Val:     in[0].Val,
			},
			Infix: in[1].Snippet,
			Right: NumLit{
				Snippet: in[2].Snippet,
				Val:     number.New(in[2].Val),
			},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}
