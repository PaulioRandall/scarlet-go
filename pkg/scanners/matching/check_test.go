package matching

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

func checkFormats(t *testing.T, exp, in Token) {
	out := SanitiseAll([]Token{in})
	checkSize(t, 1, out)
	checkToken(t, exp, out[0])
}

func checkRemoves(t *testing.T, tk Token) {
	out := SanitiseAll([]Token{tk})
	checkSize(t, 0, out)
}

func checkRemovesTerminators(t *testing.T, prev Token) {

	in := []Token{
		prev,
		Token{TERMINATOR, "", 0, 0},
		Token{TERMINATOR, "", 0, 0},
	}

	exp := []Token{prev}

	checkMany(t, exp, in)
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

func checkToken(t *testing.T, exp, act Token) {
	require.Equal(t, exp, act,
		"Expected ("+exp.String()+") but got ("+act.String()+")")
}

func checkSize(t *testing.T, exp int, acts []Token) {
	require.Equal(t, exp, len(acts),
		"Expected "+strconv.Itoa(exp)+
			" tokens but got "+strconv.Itoa(len(acts)))
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
