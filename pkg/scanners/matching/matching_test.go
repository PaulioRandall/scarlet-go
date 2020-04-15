package matching

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, ss *symbolStream, exp ...Token) {

	for i := 0; i < len(exp); i++ {

		tk := read(ss)

		if tk == (Token{}) {
			require.Equal(t, len(exp), i, "Expected scanning to return more tokens")
			return
		}

		require.Equal(t, exp[i], tk)
	}
}

func TestScanner_Next_1(t *testing.T) {

	s := "\r\n" +
		" \t\r\v\f" + "// comment" + "\n" +
		"123" + " " + "123.456" + "\r\n" +
		"`abc`" + `"abc"` + "\n" +
		"abc_xyz" + "\r\n" +
		"F" + " " + "TRUE" + "\n" +
		"@" + ":=" + "*" + "->" + ")" + "\r\n"

	ss := &symbolStream{
		runes: []rune(s),
	}

	doTest(t, ss,
		// Line 0
		Token{NEWLINE, "\r\n", 0, 0},
		// Line 1
		Token{WHITESPACE, " \t\r\v\f", 1, 0},
		Token{COMMENT, "// comment", 1, 5},
		Token{NEWLINE, "\n", 1, 15},
		// Line 2
		Token{INT, "123", 2, 0},
		Token{WHITESPACE, " ", 2, 3},
		Token{FLOAT, "123.456", 2, 4},
		Token{NEWLINE, "\r\n", 2, 11},
		// Line 3
		Token{STRING, "`abc`", 3, 0},
		Token{TEMPLATE, `"abc"`, 3, 5},
		Token{NEWLINE, "\n", 3, 10},
		// Line 4
		Token{ID, "abc_xyz", 4, 0},
		Token{NEWLINE, "\r\n", 4, 7},
		// Line 5
		Token{FUNC, "F", 5, 0},
		Token{WHITESPACE, " ", 5, 1},
		Token{BOOL, "TRUE", 5, 2},
		Token{NEWLINE, "\n", 5, 6},
		// Line 6
		Token{SPELL, "@", 6, 0},
		Token{ASSIGN, ":=", 6, 1},
		Token{MULTIPLY, "*", 6, 3},
		Token{RETURNS, "->", 6, 4},
		Token{PAREN_CLOSE, ")", 6, 6},
		Token{NEWLINE, "\r\n", 6, 7},
		// Line 7
	)

	expEOF := Token{
		Type: EOF,
		Line: 7,
		Col:  0,
	}

	require.Equal(t, expEOF, read(ss))
}

func TestScanner_Next_2(t *testing.T) {
	require.Panics(t, func() {
		ss := &symbolStream{
			runes: []rune("123.a"),
		}
		read(ss)
	})
}
