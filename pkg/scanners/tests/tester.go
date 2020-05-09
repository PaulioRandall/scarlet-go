package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

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

func check(t *testing.T, exps, acts []Token) {

	tkStr := func(tks []Token, i int) (_ string) {
		if i < len(tks) {
			return tks[i].String()
		}
		return
	}

	expSize := len(exps)
	actSize := len(acts) - 1

	require.Equal(t, EOF, acts[actSize].Type,
		"Expected EOF but got ("+tkStr(acts, actSize)+")")

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+tkStr(exps, i)+") but no actual tokens remain")

		require.True(t, i < expSize,
			"Didn't expect any more tokens but got ("+tkStr(acts, i)+")")

		require.Equal(t, exps[i], acts[i],
			"Expected ("+tkStr(exps, i)+") but got ("+tkStr(acts, i)+")")
	}
}

func checkOne(t *testing.T, exp Token, acts []Token) {
	check(t, []Token{exp}, acts)
}

func checkPanic(t *testing.T, f func()) {
	require.Panics(t, f, "Expected a panic")
}
