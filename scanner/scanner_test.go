package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func scanErrTest(t *testing.T, f token.ScanToken, expAt int, exp perror.Perror) {
	e := token.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(perror.Perror).Where())
}

func TestScan_1(t *testing.T) {
	// Check it works when the input only contains one token.

	token.ScanTokenTest(t,
		New("abc"),
		token.New("abc", token.ID, 0, 0, 3),
	)
}

func TestScan_2(t *testing.T) {
	// Check it works when the input contains multiple tokens.

	token.ScanTokenTest(t,
		New("FUNC\nabc :=\nEND"),
		token.New("FUNC", token.FUNC, 0, 0, 4),
		token.New("\n", token.NEWLINE, 0, 4, 5),
		token.New("abc", token.ID, 1, 0, 3),
		token.New(" ", token.WHITESPACE, 1, 3, 4),
		token.New(":=", token.ASSIGN, 1, 4, 6),
		token.New("\n", token.NEWLINE, 1, 6, 7),
		token.New("END", token.END, 2, 0, 3),
	)
}

func TestScan_3(t *testing.T) {
	// Check an error occurrs when the input contains invalid tokens.

	scanErrTest(t,
		New("abc   ~~~"),
		2,
		perror.New("", 0, 6, 6),
	)
}
