package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

type Tester struct {
	T   *testing.T
	Tag string
	F   ScanFunc
}

type ScanFunc func(in string) []Token

type TestFunc func(t *testing.T, tag string, sf ScanFunc)

func (t *Tester) Run(f TestFunc) {

	defer func() {
		if e := recover(); e != nil {
			v, _ := e.(error)
			require.Nil(t.T, e, t.Tag+v.Error())
		}
	}()

	f(t.T, t.Tag, t.F)
}

func check(t *testing.T, tag string, exps, acts []Token) {

	tkStr := func(tks []Token, i int) (_ string) {
		if i < len(tks) {
			return tks[i].String()
		}
		return
	}

	expSize := len(exps)
	actSize := len(acts) - 1

	require.Equal(t, EOF, acts[actSize].Type, tag+
		" Expected EOF but got ("+tkStr(acts, actSize)+")")

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize, tag+
			" Expected ("+tkStr(exps, i)+") but no actual tokens remain")

		require.True(t, i < expSize, tag+
			" Didn't expect any more tokens but got ("+tkStr(acts, i)+")")

		require.Equal(t, exps[i], acts[i], tag+
			" Expected ("+tkStr(exps, i)+") but got ("+tkStr(acts, i)+")")
	}
}
