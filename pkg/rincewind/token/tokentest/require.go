package tokentest

import (
	"fmt"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"

	"github.com/stretchr/testify/require"
)

func Tok(ge GenType, su SubType, raw string, line, colBegin, colEnd int) tok {
	return tok{
		ge:       ge,
		su:       su,
		raw:      raw,
		line:     line,
		colBegin: colBegin,
		colEnd:   colEnd,
	}
}

func HalfTok(ge GenType, su SubType, raw string) tok {
	return tok{
		ge:     ge,
		su:     su,
		raw:    raw,
		colEnd: len(raw),
	}
}

func RequireSlice(t *testing.T, exps, acts []Token) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+exps[i].String()+")\nBut no actual tokens remain")

		require.True(t, i < expSize,
			"Did not expect any more tokens\nBut got ("+acts[i].String()+")")

		RequireToken(t, exps[i], acts[i])
	}
}

func RequireToken(t *testing.T, exp, act Token) {

	require.NotNil(t, act, "Expected token ("+exp.String()+")\nBut got nil")
	msg := "Expected (" + exp.String() + ")\nActual   (" + act.String() + ")"

	require.Equal(t, exp.GenType(), act.GenType(), msg)
	require.Equal(t, exp.SubType(), act.SubType(), msg)
	require.Equal(t, exp.Raw(), act.Raw(), msg)
	require.Equal(t, exp.Value(), act.Value(), msg)
	requireSnippet(t, exp, act, msg)
}

func RequireSnippet(t *testing.T, exp, act Snippet) {

	var msg string
	if v, ok := exp.(fmt.Stringer); ok {
		msg = v.String()
	} else {
		line, col := exp.Begin()
		endLine, endCol := exp.End()
		msg = fmt.Sprintf("Begin %d:%d, End %d:%d", line, col, endLine, endCol)
	}

	requireSnippet(t, exp, act, msg)
}

func requireSnippet(t *testing.T, exp, act Snippet, msg string) {
	requirePos(t, exp.Begin, act.Begin, msg)
	requirePos(t, exp.End, act.End, msg)
}

func requirePos(t *testing.T, exp, act func() (int, int), msg string) {
	expLine, expCol := exp()
	actLine, actCol := act()
	require.Equal(t, expLine, actLine, msg)
	require.Equal(t, expCol, actCol, msg)
}
