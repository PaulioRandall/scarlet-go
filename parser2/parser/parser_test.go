package parser

/*
import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
	"github.com/PaulioRandall/scarlet-go/token2/tokentest"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {

	tm := &position.TextMarker{}
	genLex := func(v string, tk token.Token) lexeme.Lexeme {
		snip := tm.Snippet(v)
		tm.Advance(v)
		return lexeme.Make(v, tk, snip)
	}

	// x := 1
	tokenItr := tokentest.FeignSeries(
		genLex("x", token.IDENT),
		genLex(":=", token.ASSIGN),
		genLex("1", token.NUMBER),
	)

	exp := []Stat{}

	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	require.Equal(t, 1, len(act))

	for i, s := range act {
		require.True(t, i < len(exp), "Expected more statements")
		require.Equal(t, exp[i], s)
	}
}
*/
