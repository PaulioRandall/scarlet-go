package matching

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/symbol"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, ts *symbol.TerminalStream, exp ...Token) {

	for i := 0; i < len(exp); i++ {

		tk := Read(ts)

		if tk == (Token{}) {
			require.Equal(t, len(exp), i, "Expected scanning to return more tokens")
			return
		}

		require.Equal(t, exp[i], tk)
	}
}

func TestScanner_Next_1(t *testing.T) {

	ts := symbol.New(
		"\r\n" +
			" \t\r\v\f" + "// comment" + "\n" +
			"123" + " " + "123.456" + "\r\n" +
			"`abc`" + `"abc"` + "\n" +
			"abc_xyz" + "\r\n" +
			"F" + " " + "TRUE" + "\n" +
			"@" + ":=" + "*" + "->" + ")" + "\r\n",
	)

	doTest(t, ts,
		// Line 0
		Token{LEXEME_NEWLINE, "\r\n", 0, 0},
		// Line 1
		Token{LEXEME_WHITESPACE, " \t\r\v\f", 1, 0},
		Token{LEXEME_COMMENT, "// comment", 1, 5},
		Token{LEXEME_NEWLINE, "\n", 1, 15},
		// Line 2
		Token{LEXEME_INT, "123", 2, 0},
		Token{LEXEME_WHITESPACE, " ", 2, 3},
		Token{LEXEME_FLOAT, "123.456", 2, 4},
		Token{LEXEME_NEWLINE, "\r\n", 2, 11},
		// Line 3
		Token{LEXEME_STRING, "`abc`", 3, 0},
		Token{LEXEME_TEMPLATE, `"abc"`, 3, 5},
		Token{LEXEME_NEWLINE, "\n", 3, 10},
		// Line 4
		Token{LEXEME_ID, "abc_xyz", 4, 0},
		Token{LEXEME_NEWLINE, "\r\n", 4, 7},
		// Line 5
		Token{LEXEME_FUNC, "F", 5, 0},
		Token{LEXEME_WHITESPACE, " ", 5, 1},
		Token{LEXEME_BOOL, "TRUE", 5, 2},
		Token{LEXEME_NEWLINE, "\n", 5, 6},
		// Line 6
		Token{LEXEME_SPELL, "@", 6, 0},
		Token{LEXEME_ASSIGN, ":=", 6, 1},
		Token{LEXEME_MULTIPLY, "*", 6, 3},
		Token{LEXEME_RETURNS, "->", 6, 4},
		Token{LEXEME_PAREN_CLOSE, ")", 6, 6},
		Token{LEXEME_NEWLINE, "\r\n", 6, 7},
		// Line 7
	)

	expEOF := Token{
		Lexeme: LEXEME_EOF,
		Line:   7,
		Col:    0,
	}

	require.Equal(t, expEOF, Read(ts))
}

func TestScanner_Next_2(t *testing.T) {
	require.Panics(t, func() {
		ts := symbol.New("123.a")
		Read(ts)
	})
}
