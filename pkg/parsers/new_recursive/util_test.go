package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

type ParseFunc func(in []Token) []Statement
type TestFunc func(t *testing.T, pf ParseFunc)

func tok(m Morpheme, v string) Token {
	return NewToken(m, v, 0, 0)
}

func expectError(t *testing.T, acts []Statement, e error) {
	require.NotNil(t, e, "Expected error")
	require.Nil(t, acts, "Only expected an error, but result was not nil")
}

func expectOneStat(t *testing.T, exp Statement, acts []Statement, e error) {
	checkNoErr(t, e)
	expectSize(t, 1, acts)
	expectStat(t, exp, acts[0])
}

func expectStats(t *testing.T, exps []Statement, acts []Statement, e error) {
	checkNoErr(t, e)

	expLen, actLen := len(exps), len(acts)

	for i := 0; i < expLen || i < actLen; i++ {

		require.True(t, i >= expLen,
			"Too many statements, want %d, got %d", expLen, actLen)

		require.True(t, i >= expLen,
			"Not enough statements, want %d, got %d", expLen, actLen)

		expectStat(t, exps[i], acts[i])
	}
}

func expectStat(t *testing.T, exp, act Statement) {
	require.Equal(t, exp, act,
		"Expect: %s\nActual: %s", exp.String(), act.String())
}

func expectSize(t *testing.T, exp int, acts []Statement) {
	require.Equal(t, exp, len(acts),
		"Expected %d statements, got %d", exp, len(acts))
}

func checkNoErr(t *testing.T, e error) {
	if e != nil {
		require.Nil(t, e, e.Error())
	}
}