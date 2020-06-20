package vorbis

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/esmerelda"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

type stItr struct {
	StatementIterator
	stats []Expression
	index int
}

func (i *stItr) Next() (Expression, error) {

	if i.index >= len(i.stats) {
		return nil, nil
	}

	st := i.stats[i.index]
	i.index++
	return st, nil
}

func tok(ty TokenType, v string) Token {
	return NewToken(ty, v, 0, 0)
}

func Test_S1_1(t *testing.T) {

	stats := []Expression{
		NewGuard(
			tok(TK_GUARD_OPEN, "["),
			NewOperation(
				tok(TK_EQUAL, "=="),
				NewLiteral(tok(TK_NUMBER, "1")),
				NewLiteral(tok(TK_NUMBER, "1")),
			),
			NewBlock(
				tok(TK_BLOCK_OPEN, "{"),
				tok(TK_BLOCK_CLOSE, "}"),
				[]Expression{},
			),
		),
	}

	itr := &stItr{
		stats: stats,
	}

	errs := AuditStatements(itr)
	require.Nil(t, errs)
}
