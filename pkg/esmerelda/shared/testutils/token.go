package testutils

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"

	"github.com/stretchr/testify/require"
)

func RequireTokenSlice(t *testing.T, exps, acts []token.Token) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+exps[i].String()+")\nBut no actual tokens remain")

		require.True(t, i < expSize,
			"Did not expect any more tokens\nBut got ("+acts[i].String()+")")

		requireToken(t, exps[i], acts[i])
	}
}

func requireToken(t *testing.T, exp, act token.Token) {

	require.NotNil(t, act, "Expected token ("+exp.String()+")\nBut got nil")
	msg := "Expected (" + exp.String() + ")\nActual   (" + act.String() + ")"

	require.Equal(t, exp.GenType(), act.GenType(), msg)
	require.Equal(t, exp.SubType(), act.SubType(), msg)
	require.Equal(t, exp.Props(), act.Props(), msg)
	require.Equal(t, exp.Raw(), act.Raw(), msg)
	require.Equal(t, exp.Value(), act.Value(), msg)
	requireSnippet(t, exp, act, msg)
}
