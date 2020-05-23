package z_tests

import (
	"strconv"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"

	"github.com/stretchr/testify/require"
)

type ScanFunc func(in string) []Token

type TestFunc func(t *testing.T, sf ScanFunc)

func Run(t *testing.T, sf ScanFunc, tf TestFunc) {

	defer func() {
		if e := recover(); e != nil {
			v, _ := e.(error)
			require.Nil(t, e, v.Error())
		}
	}()

	tf(t, sf)
}

func checkMany(t *testing.T, exps, acts []Token) {

	checkEOF(t, acts)

	expSize := len(exps)
	actSize := len(acts) - 1

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+tkStr(exps, i)+") but no actual tokens remain")

		require.True(t, i < expSize,
			"Didn't expect any more tokens but got ("+tkStr(acts, i)+")")

		checkToken(t, exps[i], acts[i])
	}
}

func checkOne(t *testing.T, exp Token, acts []Token) {
	checkSize(t, 2, acts)
	checkToken(t, exp, acts[0])
	checkEOF(t, acts)
}

func checkFirstNot(t *testing.T, notExp Token, acts []Token) {
	checkMinSize(t, 2, acts)
	require.NotEqual(t, notExp, acts[0],
		"Expected any token except ("+ToString(notExp)+") but got it")
	checkEOF(t, acts)
}

func checkEOF(t *testing.T, acts []Token) {
	i := len(acts) - 1
	tk := acts[i]
	m := "Expected EOF but got (" + tkStr(acts, i) + ")"

	require.Equal(t, K_EOF, tk.Kind(), m)
	require.Equal(t, M_EOF, tk.Morpheme(), m)
}

func checkSize(t *testing.T, exp int, acts []Token) {
	require.Equal(t, exp, len(acts),
		"Expected "+strconv.Itoa(exp)+
			" tokens (inc EOF) but got "+strconv.Itoa(len(acts)))
}

func checkMinSize(t *testing.T, min int, acts []Token) {
	require.True(t, min <= len(acts),
		"Expected minimum "+strconv.Itoa(min)+
			" tokens (inc EOF) but got "+strconv.Itoa(len(acts)))
}

func checkPanic(t *testing.T, f func()) {
	require.Panics(t, f, "Expected a panic")
}

func checkToken(t *testing.T, exp, act Token) {
	require.NotNil(t, act, "Expected token ("+ToString(exp)+") got nil")

	m := "Expected (" + ToString(exp) + ") but got (" + ToString(act) + ")"

	require.Equal(t, exp.Kind(), act.Kind(), m)
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
