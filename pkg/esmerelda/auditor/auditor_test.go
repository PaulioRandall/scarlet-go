package auditor

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"

	"github.com/stretchr/testify/require"
)

type stItr struct {
	StatementIterator
	sts   []Expression
	index int
}

func (i *stItr) Next() (Expression, error) {

	if i.index >= len(i.sts) {
		return nil, nil
	}

	st := i.sts[i.index]
	i.index++
	return st, nil
}

func tok(ty TokenType, v string) Token {
	return NewToken(ty, v, 0, 0)
}

func emptyBlock() Block {
	return NewBlock(
		tok(TK_BLOCK_OPEN, "{"),
		tok(TK_BLOCK_CLOSE, "}"),
		[]Expression{},
	)
}

func Test_S1_1(t *testing.T) {

	sts := []Expression{
		NewGuard(
			tok(TK_GUARD_OPEN, "["),
			NewOperation(
				tok(TK_EQUAL, "=="),
				NewLiteral(tok(TK_NUMBER, "1")),
				NewLiteral(tok(TK_NUMBER, "1")),
			),
			emptyBlock(),
		),
	}

	itr := &stItr{
		sts: sts,
	}

	errs := AuditStatements(itr)
	require.Nil(t, errs)
}

func Test_S1_2(t *testing.T) {

	sts := []Expression{
		NewGuard(
			tok(TK_GUARD_OPEN, "["),
			NewOperation(
				tok(TK_AND, "&"),
				NewLiteral(tok(TK_BOOL, "true")),
				NewLiteral(tok(TK_BOOL, "false")),
			),
			emptyBlock(),
		),
	}

	itr := &stItr{
		sts: sts,
	}

	errs := AuditStatements(itr)
	require.Nil(t, errs)
}

func Test_S1_3(t *testing.T) {

	sts := []Expression{
		NewGuard(
			tok(TK_GUARD_OPEN, "["),
			NewOperation(
				tok(TK_PLUS, "+"),
				NewLiteral(tok(TK_NUMBER, "1")),
				NewLiteral(tok(TK_NUMBER, "1")),
			),
			emptyBlock(),
		),
	}

	itr := &stItr{
		sts: sts,
	}

	errs := AuditStatements(itr)
	require.NotNil(t, errs)
	require.Equal(t, 1, len(errs), "Wrong number of errors")
}
