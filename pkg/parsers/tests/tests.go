package tests

import (
	"strconv"
	"testing"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

type ParseFunc func(in []token.Token) []st.Statement

type TestFunc func(t *testing.T, pf ParseFunc)

func Run(t *testing.T, pf ParseFunc, tf TestFunc) {

	defer func() {
		if e := recover(); e != nil {
			v, _ := e.(error)
			require.Nil(t, e, v.Error())
		}
	}()

	tf(t, pf)
}

/*
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
*/

func expectOneStat(t *testing.T, exp st.Statement, acts []st.Statement) {
	expectSize(t, 1, acts)
	expectStat(t, exp, acts[0])
}

func expectStat(t *testing.T, exp, act st.Statement) {
	require.Equal(t, exp, act,
		"Expected ("+exp.String(1)+") but got ("+act.String(1)+")")
}

func expectSize(t *testing.T, exp int, acts []st.Statement) {
	require.Equal(t, exp, len(acts),
		"Expected "+strconv.Itoa(exp)+
			" statement but got "+strconv.Itoa(len(acts)))
}

func expectPanic(t *testing.T, f func()) {
	require.Panics(t, f, "Expected a panic")
}

func stStr(sts []st.Statement, i int) (_ string) {
	if i < len(sts) {
		return sts[i].String(1)
	}
	return
}
