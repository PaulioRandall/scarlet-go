package tests

import (
	"strconv"
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

func checkOneNot(t *testing.T, notExp Token, acts []Token) {
	checkSize(t, 2, acts)
	require.NotEqual(t, notExp, acts[0],
		"Expected any token except ("+notExp.String()+") but got it anyway")
	checkEOF(t, acts)
}

func checkToken(t *testing.T, exp, act Token) {
	require.Equal(t, exp, act,
		"Expected ("+exp.String()+") but got ("+act.String()+")")
}

func checkSize(t *testing.T, exp int, acts []Token) {
	require.Equal(t, exp, len(acts),
		"Expected "+strconv.Itoa(2)+
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
