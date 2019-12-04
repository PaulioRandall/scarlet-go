package strimmer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

func wrapErrTest(t *testing.T, f token.ScanToken, expAt int, exp perror.Perror) {
	e := token.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(perror.Perror).Where())
}

func TestWrap_1(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
	)
}

func TestWrap_2(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE\nEND"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("\n", token.NEWLINE, 0, 9, 10),
		token.NewFlat("END", token.END, 1, 0, 3),
	)
}

func TestWrap_3(t *testing.T) {
	token.ScanTokenTest(t,
		New("\t\t\t"),
	)
}

func TestWrap_4(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE\t\tEND"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("END", token.END, 0, 11, 14),
	)
}

func TestWrap_5(t *testing.T) {
	wrapErrTest(t,
		New("~~~"),
		0,
		perror.New("", 0, 0, 0),
	)
}

func TestWrap_6(t *testing.T) {
	wrapErrTest(t,
		New("PROCEDURE\n  ~~~\nEND"),
		2,
		perror.New("", 1, 2, 2),
	)
}
