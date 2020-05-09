package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

type ScanFunc func(in string) []Token

func DoTests(t *testing.T, tag string, f ScanFunc) {

	var tName string
	tag = "scanner." + tag

	defer func() {
		if e := recover(); e != nil {
			v, _ := e.(error)
			require.Nil(t, e, tName+v.Error())
		}
	}()

	tName = tag + ".a1: "
	a1(t, tName, f)

	tName = tag + ".a2: "
	a2(t, tName, f)
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

func a1(t *testing.T, tag string, f ScanFunc) {

	in := "\n\r\n"

	exps := []Token{
		Token{NEWLINE, "\n", 0, 0},
		Token{NEWLINE, "\r\n", 1, 0},
	}

	acts := f(in)

	check(t, tag, exps, acts)
}

func a2(t *testing.T, tag string, f ScanFunc) {

	in := " \t\r\v\f"

	exps := []Token{
		Token{WHITESPACE, " \t\r\v\f", 0, 0},
	}

	acts := f(in)

	check(t, tag, exps, acts)
}
