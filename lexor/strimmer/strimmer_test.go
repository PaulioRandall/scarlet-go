package strimmer

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func wrapErrTest(t *testing.T, f lexor.ScanToken, expAt int, exp token.Perror) {
	e := lexor.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(token.Perror).Where())
}

func TestWrap_1(t *testing.T) {
	lexor.ScanTokenTest(t,
		New("abc"),
		token.NewToken(token.ID, "abc", 0, 0, 3),
	)
}

func TestWrap_2(t *testing.T) {
	lexor.ScanTokenTest(t,
		New("abc\n   efg"),
		token.NewToken(token.ID, "abc", 0, 0, 3),
		token.NewToken(token.NEWLINE, "\n", 0, 3, 4),
		token.NewToken(token.ID, "efg", 1, 3, 6),
	)
}

func TestWrap_3(t *testing.T) {
	lexor.ScanTokenTest(t,
		New("\t\t\t"),
	)
}

func TestWrap_4(t *testing.T) {
	wrapErrTest(t,
		New("~~~"),
		0,
		token.NewPerror("", 0, 0, 0),
	)
}
