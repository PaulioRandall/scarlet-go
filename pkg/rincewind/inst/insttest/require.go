package insttest

import (
	"fmt"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/number"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"

	"github.com/stretchr/testify/require"
)

func RequireSlice(t *testing.T, exps, acts []Instruction) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+exps[i].String()+")\nBut no actual instructions remain")

		require.True(t, i < expSize,
			"Did not expect any more instructions\nBut got ("+acts[i].String()+")")

		RequireInstruction(t, exps[i], acts[i])
	}
}

func RequireInstruction(t *testing.T, exp, act Instruction) {

	require.NotNil(t, act, "Expected instruction ("+exp.String()+")\nBut got nil")
	msg := "Expected (" + exp.String() + ")\nActual   (" + act.String() + ")"

	require.Equal(t, exp.Code(), act.Code(), msg)
	RequireData(t, exp.Data(), act.Data(), msg)
	requireSnippet(t, exp, act, msg)
}

func RequireData(t *testing.T, exp, act interface{}, msg string) {

	if v, ok := exp.([]interface{}); ok {
		require.Implements(t, ([]interface{})(nil), act)
		RequireDataSlice(t, v, act.([]interface{}), msg)
		return
	}

	RequireDataItem(t, exp, act, msg)
}

func RequireDataSlice(t *testing.T, exps, acts []interface{}, msg string) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected %q\nBut no actual data items remain", exps[i])

		require.True(t, i < expSize,
			"Did not expect any more data items\nBut got %q", acts[i])

		RequireDataItem(t, exps[i], acts[i], msg)
	}
}

func RequireDataItem(t *testing.T, exp, act interface{}, msg string) {

	switch v := exp.(type) {
	case number.Number:
		require.Implements(t, (*number.Number)(nil), act, msg)
		n := act.(number.Number)
		require.True(t, v.Equal(n), msg)

	default:
		require.Equal(t, exp, act, msg)
	}
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
