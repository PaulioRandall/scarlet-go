package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
)

func scanner_ScanErrTest(t *testing.T, f token.ScanToken, expAt int, exp perror.Perror) {
	e := token.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(perror.Perror).Where())
}

func TestScan_1(t *testing.T) {
	token.ScanTokenTest(t,
		New("abc"),
		token.New("abc", token.ID, 0, 0, 3),
	)
}

func TestScan_2(t *testing.T) {
	token.ScanTokenTest(t,
		New("FUNC abc\nEND"),
		token.New("FUNC", token.FUNC, 0, 0, 4),
		token.New(" ", token.WHITESPACE, 0, 4, 5),
		token.New("abc", token.ID, 0, 5, 8),
		token.New("\n", token.NEWLINE, 0, 8, 9),
		token.New("END", token.END, 1, 0, 3),
	)
}

func TestScan_3(t *testing.T) {
	scanner_ScanErrTest(t,
		New("~~~"),
		0,
		perror.New("", 0, 0, 0),
	)
}
