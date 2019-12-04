package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestScanner_Scan_OK_1(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
	)
}

func TestScanner_Scan_OK_2(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE\nEND"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("\n", token.NEWLINE, 0, 9, 10),
		token.NewFlat("END", token.END, 1, 0, 3),
	)
}

func TestScanner_Scan_OK_3(t *testing.T) {
	token.ScanTokenTest(t,
		New("\t\t\t"),
		token.NewFlat("\t\t\t", token.WHITESPACE, 0, 0, 3),
	)
}

func TestScanner_Scan_OK_4(t *testing.T) {
	token.ScanTokenTest(t,
		New("PROCEDURE\t\tEND"),
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("\t\t", token.WHITESPACE, 0, 9, 11),
		token.NewFlat("END", token.END, 0, 11, 14),
	)
}
