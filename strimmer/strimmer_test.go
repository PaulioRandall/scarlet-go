package strimmer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

func wrapErrTest(t *testing.T, f token.ScanToken, expAt int, exp perror.Perror) {
	e := token.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(perror.Perror).Where())
}

func TestWrap_1(t *testing.T) {
	token.ScanTokenTest(t,
		New("FUNC"),
		token.NewFlat("FUNC", token.FUNC, 0, 0, 4),
	)
}

func TestWrap_2(t *testing.T) {
	token.ScanTokenTest(t,
		New("FUNC\nEND"),
		token.NewFlat("FUNC", token.FUNC, 0, 0, 4),
		token.NewFlat("\n", token.NEWLINE, 0, 4, 5),
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
		New("FUNC\t\tEND"),
		token.NewFlat("FUNC", token.FUNC, 0, 0, 4),
		token.NewFlat("END", token.END, 0, 6, 9),
	)
}

/*
func TestWrap_5(t *testing.T) {
	wrapErrTest(t,
		New("~~~"),
		0,
		perror.New("", 0, 0, 0),
	)
}

func TestWrap_6(t *testing.T) {
	wrapErrTest(t,
		New("FUNC\n  ~~~\nEND"),
		2,
		perror.New("", 1, 2, 2),
	)
}
*/
