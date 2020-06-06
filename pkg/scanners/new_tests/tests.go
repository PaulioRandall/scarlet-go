package new_tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

func assertError(t *testing.T, toks []Token, e error) {
	require.NotNil(t, e, "Expected an error for this test")
	require.Nil(t, toks, "Upon error, resultant token slice should be nil")
}

func assertResults(t *testing.T, exp, toks []Token, e error) {
	require.Nil(t, e, "Did not expect an error for this test")
	require.NotNil(t, toks, "SANITY CHECK! What tokens were expected?")
	require.NotNil(t, toks, "Expected a non-nil token slice")
}

func checkMany(t *testing.T, exps, acts []Token) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+tkStr(exps, i)+"), but no actual tokens remain")

		require.True(t, i < expSize,
			"Did not expect any more tokens, but got ("+tkStr(acts, i)+")")

		checkToken(t, exps[i], acts[i])
	}
}

func checkToken(t *testing.T, exp, act Token) {
	require.NotNil(t, act, "Expected token ("+ToString(exp)+"), but got nil")

	m := "Expected (" + ToString(exp) + "), but got (" + ToString(act) + ")"

	require.Equal(t, exp.Morpheme(), act.Morpheme(), m)
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
