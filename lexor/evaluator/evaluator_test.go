package evaluator

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func wrapErrTest(t *testing.T, f lexor.ScanToken, expAt int, exp lexor.ScanErr) {
	e := lexor.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(lexor.ScanErr).Where())
}

func TestWrap_1(t *testing.T) {
	lexor.ScanTokenTest(t,
		New("`abc`"),
		token.NewToken(token.STR_LITERAL, "abc", 0, 0, 5),
	)
}

func TestWrap_2(t *testing.T) {
	lexor.ScanTokenTest(t,
		New("abc\n F"),
		token.NewToken(token.ID, "abc", 0, 0, 3),
		token.NewToken(token.NEWLINE, "\n", 0, 3, 4),
		token.NewToken(token.FUNC, "F", 1, 1, 2),
	)
}

func TestWrap_3(t *testing.T) {
	wrapErrTest(t,
		New("~~~"),
		0,
		lexor.NewScanErr("", 0, 0, 0),
	)
}
