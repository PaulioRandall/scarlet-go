package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

func scanner_ScanErrTest(t *testing.T, f token.ScanToken, expAt int, exp perror.Perror) {
	e := token.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(perror.Perror).Where())
}

func TestScanner_Scan_1(t *testing.T) {
	token.ScanTokenTest(t,
		New("FUNC"),
		token.New("FUNC", token.FUNC, 0, 0, 4),
	)
}

func TestScanner_Scan_2(t *testing.T) {
	token.ScanTokenTest(t,
		New("FUNC\nEND"),
		token.New("FUNC", token.FUNC, 0, 0, 4),
		token.New("\n", token.NEWLINE, 0, 4, 5),
		token.New("END", token.END, 1, 0, 3),
	)
}

func TestScanner_Scan_3(t *testing.T) {
	token.ScanTokenTest(t,
		New("\t\t\t"),
		token.New("\t\t\t", token.WHITESPACE, 0, 0, 3),
	)
}

func TestScanner_Scan_4(t *testing.T) {
	token.ScanTokenTest(t,
		New("FUNC\t\tEND"),
		token.New("FUNC", token.FUNC, 0, 0, 4),
		token.New("\t\t", token.WHITESPACE, 0, 4, 6),
		token.New("END", token.END, 0, 6, 9),
	)
}

/*
func TestScanner_Scan_5(t *testing.T) {
	scanner_ScanErrTest(t,
		New("~~~"),
		0,
		perror.New("", 0, 0, 0),
	)
}

func TestScanner_Scan_6(t *testing.T) {
	scanner_ScanErrTest(t,
		New("FUNC\n  ~~~\nEND"),
		3,
		perror.New("", 1, 2, 2),
	)
}
*/
