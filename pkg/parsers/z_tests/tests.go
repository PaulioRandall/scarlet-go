package z_tests

import (
	"strconv"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"

	"github.com/stretchr/testify/require"
)

func tok(m Morpheme, v string) Token {
	return token{m, v}
}

type token struct {
	m Morpheme
	v string
}

func (tk token) Morpheme() Morpheme {
	return tk.m
}

func (tk token) Value() string {
	return tk.v
}

func (tk token) Line() int {
	return 0
}

func (tk token) Col() int {
	return 0
}

type ParseFunc func(in []Token) []Statement

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

func expectOneStat(t *testing.T, exp Statement, acts []Statement) {
	expectSize(t, 1, acts)
	expectStat(t, exp, acts[0])
}

func expectStat(t *testing.T, exp, act Statement) {
	require.Equal(t, exp, act,
		"Expect: "+exp.String(0)+"\n"+
			"Actual: "+act.String(0),
	)
}

func expectSize(t *testing.T, exp int, acts []Statement) {
	require.Equal(t, exp, len(acts),
		"Expected "+strconv.Itoa(exp)+
			" statement but got "+strconv.Itoa(len(acts)))
}

func expectPanic(t *testing.T, f func()) {
	require.Panics(t, f, "Expected a panic")
}

func stStr(sts []Statement, i int) string {
	if i < len(sts) {
		return sts[i].String(1)
	}
	return ""
}
