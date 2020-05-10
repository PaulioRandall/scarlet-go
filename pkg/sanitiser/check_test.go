package sanitiser

import (
	"strconv"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

func checkIgnores(t *testing.T, tk Token) {
	out := SanitiseAll([]Token{tk})
	checkSize(t, 1, out)
	checkToken(t, tk, out[0])
}

func checkRemoves(t *testing.T, tk Token) {
	out := SanitiseAll([]Token{tk})
	checkSize(t, 0, out)
}

func checkMany(t *testing.T, exp, in []Token) {

	out := SanitiseAll(in)

	expSize := len(exp)
	outSize := len(out)

	for i := 0; i < expSize || i < outSize; i++ {

		require.True(t, i < outSize,
			"Expected ("+tkStr(exp, i)+") but no actual tokens remain")

		require.True(t, i < expSize,
			"Didn't expect any more tokens but got ("+tkStr(out, i)+")")

		checkToken(t, exp[i], out[i])
	}
}

func checkOneNot(t *testing.T, notExp Token, acts []Token) {
	checkSize(t, 2, acts)
	require.NotEqual(t, notExp, acts[0],
		"Expected any token except ("+notExp.String()+") but got it")
	checkEOF(t, acts)
}

func checkToken(t *testing.T, exp, act Token) {
	require.Equal(t, exp, act,
		"Expected ("+exp.String()+") but got ("+act.String()+")")
}

func checkSize(t *testing.T, exp int, acts []Token) {
	require.Equal(t, exp, len(acts),
		"Expected "+strconv.Itoa(exp)+
			" tokens (inc EOF) but got "+strconv.Itoa(len(acts)))
}

func checkEOF(t *testing.T, acts []Token) {
	i := len(acts) - 1
	require.Equal(t, EOF, acts[i].Type,
		"Expected EOF but got ("+tkStr(acts, i)+")")
}

func checkPanic(t *testing.T, f func()) {
	require.Panics(t, f, "Expected a panic")
}

func tkStr(tks []Token, i int) (_ string) {
	if i < len(tks) {
		return tks[i].String()
	}
	return
}
