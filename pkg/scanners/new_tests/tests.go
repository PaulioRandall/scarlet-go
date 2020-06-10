package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

func AssertError(t *testing.T, acts []Token, e error) {
	require.NotNil(t, e, "Expected an error for this test")
	require.Nil(t, acts, "Upon error, resultant token slice should be nil")
}

func AssertResults(t *testing.T, exps, acts []Token, e error) {
	requireNoError(t, e)
	require.NotNil(t, exps, "SANITY CHECK! What tokens were expected?")
	require.NotNil(t, acts, "Expected a non-nil token slice")
	assertMany(t, exps, acts)
}

func requireNoError(t *testing.T, e error) {
	if e != nil {
		require.FailNow(t, "%s", e)
	}
}

func assertMany(t *testing.T, exps, acts []Token) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+tkStr(exps, i)+")\nBut no actual tokens remain")

		require.True(t, i < expSize,
			"Did not expect any more tokens\nBut got ("+tkStr(acts, i)+")")

		assertToken(t, exps[i], acts[i])
	}
}

func assertToken(t *testing.T, exp, act Token) {
	require.NotNil(t, act, "Expected token ("+ToString(exp)+")\nBut got nil")

	m := "Expected (" + ToString(exp) + ")\nActual   (" + ToString(act) + ")"

	require.Equal(t, exp.Type(), act.Type(), m)
	require.Equal(t, exp.Value(), act.Value(), m)
	require.Equal(t, exp.Line(), act.Line(), m)
	require.Equal(t, exp.Col(), act.Col(), m)
}

func tkStr(tks []Token, i int) (_ string) {
	if i < len(tks) {
		return ToString(tks[i])
	}
	return
}
