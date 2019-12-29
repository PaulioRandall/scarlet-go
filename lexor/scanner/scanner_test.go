package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

func scanErrTest(t *testing.T, f lexor.ScanToken, expAt int, exp lexor.ScanErr) {
	e := lexor.ScanTokenErrTest(t, f, expAt)
	lexor.AssertScanErr(t, exp, e.(lexor.ScanErr))
}

func TestScan_1(t *testing.T) {
	// Check it works when the input only contains one token.

	lexor.ScanTokenTest(t,
		New("// abc"),
		token.New(token.COMMENT, "// abc", 0, 0),
	)
}

func TestScan_2(t *testing.T) {
	// Check it works when the input contains multiple tokens.

	lexor.ScanTokenTest(t,
		New("F(x,y)// ^_^"),
		token.New(token.FUNC, "F", 0, 0),
		token.New(token.OPEN_PAREN, "(", 0, 1),
		token.New(token.ID, "x", 0, 2),
		token.New(token.DELIM, ",", 0, 3),
		token.New(token.ID, "y", 0, 4),
		token.New(token.CLOSE_PAREN, ")", 0, 5),
		token.New(token.COMMENT, "// ^_^", 0, 6),
	)
}

func TestScan_3(t *testing.T) {
	// Check it works when the input contains multiple tokens.

	lexor.ScanTokenTest(t,
		New("DO\nabc := `xyz`\nEND"),
		token.New(token.DO, "DO", 0, 0),
		token.New(token.NEWLINE, "\n", 0, 2),
		token.New(token.ID, "abc", 1, 0),
		token.New(token.WHITESPACE, " ", 1, 3),
		token.New(token.ASSIGN, ":=", 1, 4),
		token.New(token.WHITESPACE, " ", 1, 6),
		token.New(token.STR_LITERAL, "`xyz`", 1, 7),
		token.New(token.NEWLINE, "\n", 1, 12),
		token.New(token.END, "END", 2, 0),
	)
}

func TestScan_4(t *testing.T) {
	// Check it works when the input contains multiple tokens.

	lexor.ScanTokenTest(t,
		New("a:=_"),
		token.New(token.ID, "a", 0, 0),
		token.New(token.ASSIGN, ":=", 0, 1),
		token.New(token.VOID, "_", 0, 3),
	)
}

func TestScan_5(t *testing.T) {
	// Check an error occurrs when the input contains invalid tokens.

	scanErrTest(t,
		New("abc   £££"),
		2,
		lexor.NewScanErr("", nil, 0, 6),
	)
}
