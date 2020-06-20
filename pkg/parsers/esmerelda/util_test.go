package esmerelda

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

type tkStream struct {
	tks []Token
}

func (s *tkStream) get(i int) Token {

	if len(s.tks) > i {
		return s.tks[i]
	}

	return nil
}

func (s *tkStream) Next() Token {

	tk := s.get(0)
	if tk != nil {
		s.tks = s.tks[1:]
	}

	return tk
}

func (s *tkStream) Peek() Token {
	return s.get(0)
}

func (s *tkStream) PeekBeyond() Token {
	return s.get(1)
}

type ParseFunc func(in []Token) []Expression
type TestFunc func(t *testing.T, pf ParseFunc)

func tok(ty TokenType, v string) Token {
	return NewToken(ty, v, 0, 0)
}

func expectError(t *testing.T, acts []Expression, e error) {
	require.NotNil(t, e, "Expected error")
	require.Nil(t, acts, "Only expected an error, but result was not nil")
}

func expectOneStat(t *testing.T, exp Expression, acts []Expression, e error) {
	checkNoErr(t, e)
	expectSize(t, 1, acts)
	expectStat(t, exp, acts[0])
}

func expectStats(t *testing.T, exps []Expression, acts []Expression, e error) {
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

func expectStat(t *testing.T, exp, act Expression) {
	require.Equal(t, exp, act,
		"Expect: %s\nActual: %s", exp.String(), act.String())
}

func expectSize(t *testing.T, exp int, acts []Expression) {
	require.Equal(t, exp, len(acts),
		"Expected %d statements, got %d", exp, len(acts))
}

func checkNoErr(t *testing.T, e error) {
	if e != nil {
		require.FailNow(t, e.Error())
	}
}
