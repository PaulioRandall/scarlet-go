package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

func scanner_ScanErrTest(t *testing.T, f token.ScanToken, expAt int, exp perror.Perror) {
	e := token.ScanTokenErrTest(t, f, expAt)
	assert.Equal(t, exp.Where(), e.(perror.Perror).Where())
}

func TestScanner_Scan_1(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
	)
}

func TestScanner_Scan_2(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE\nEND"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("\n", token.NEWLINE, 0, 9, 10),
		token.NewFlat("END", token.END, 1, 0, 3),
	)
}

func TestScanner_Scan_3(t *testing.T) {
	token.ScanTokenTest(t,
		New("\t\t\t"),
		token.NewFlat("\t\t\t", token.WHITESPACE, 0, 0, 3),
	)
}

func TestScanner_Scan_4(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE\t\tEND"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("\t\t", token.WHITESPACE, 0, 9, 11),
		token.NewFlat("END", token.END, 0, 11, 14),
	)
}

func TestScanner_Scan_5(t *testing.T) {
	scanner_ScanErrTest(t,
		New("~~~"),
		0,
		perror.New("", 0, 0, 0),
	)
}

func TestScanner_Scan_6(t *testing.T) {
	scanner_ScanErrTest(t,
		New("PROCEDURE\n  ~~~\nEND"),
		3,
		perror.New("", 1, 2, 2),
	)
}
