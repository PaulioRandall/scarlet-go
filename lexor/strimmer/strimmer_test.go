package strimmer

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func wrapErrTest(t *testing.T, f lexor.ScanToken, expAt int, exp perror.Perror) {
	e := lexor.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(perror.Perror).Where())
}

func TestWrap_1(t *testing.T) {
	lexor.ScanTokenTest(t,
		New("abc"),
		token.New(token.ID, "abc", 0, 0, 3),
	)
}

func TestWrap_2(t *testing.T) {
	lexor.ScanTokenTest(t,
		New("abc\n   efg"),
		token.New(token.ID, "abc", 0, 0, 3),
		token.New(token.NEWLINE, "\n", 0, 3, 4),
		token.New(token.ID, "efg", 1, 3, 6),
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
		perror.New("", 0, 0, 0),
	)
}