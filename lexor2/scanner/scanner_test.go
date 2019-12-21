package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor2"
	"github.com/PaulioRandall/scarlet-go/token2"
)

func scanErrTest(t *testing.T, f lexor.ScanToken, expAt int, exp lexor.ScanErr) {
	e := lexor.ScanTokenErrTest(t, f, expAt)
	lexor.AssertScanErr(t, exp, e.(lexor.ScanErr))
}

func TestScan_1(t *testing.T) {
	// Check it works when the input only contains one token.

	lexor.ScanTokenTest(t,
		New("// abc"),
		token.NewToken(token.COMMENT, "// abc", 0, 0),
	)
}

func TestScan_2(t *testing.T) {
	// Check it works when the input contains multiple tokens.

	lexor.ScanTokenTest(t,
		New("F(x,y)// ^_^"),
		token.NewToken(token.FUNC, "F", 0, 0),
		token.NewToken(token.OPEN_PAREN, "(", 0, 1),
		token.NewToken(token.ID, "x", 0, 2),
		token.NewToken(token.ID_DELIM, ",", 0, 3),
		token.NewToken(token.ID, "y", 0, 4),
		token.NewToken(token.CLOSE_PAREN, ")", 0, 5),
		token.NewToken(token.COMMENT, "// ^_^", 0, 6),
	)
}

func TestScan_3(t *testing.T) {
	// Check it works when the input contains multiple tokens.

	lexor.ScanTokenTest(t,
		New("DO\nabc := `xyz`\nEND"),
		token.NewToken(token.DO, "DO", 0, 0),
		token.NewToken(token.NEWLINE, "\n", 0, 2),
		token.NewToken(token.ID, "abc", 1, 0),
		token.NewToken(token.WHITESPACE, " ", 1, 3),
		token.NewToken(token.ASSIGN, ":=", 1, 4),
		token.NewToken(token.WHITESPACE, " ", 1, 6),
		token.NewToken(token.STR_LITERAL, "`xyz`", 1, 7),
		token.NewToken(token.NEWLINE, "\n", 1, 12),
		token.NewToken(token.END, "END", 2, 0),
	)
}

func TestScan_4(t *testing.T) {
	// Check an error occurrs when the input contains invalid tokens.

	scanErrTest(t,
		New("abc   ~~~"),
		2,
		lexor.NewScanErr("", nil, 0, 6),
	)
}