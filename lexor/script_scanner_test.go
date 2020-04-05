package lexor

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func tok(k token.Kind, v string, l, c int) token.Token {
	return token.New(k, v, l, c)
}

func doTest(t *testing.T, ts TokenStream, exp ...token.Token) {

	for i := 0; i < len(exp); i++ {

		tk := ts.Next()

		if tk == (token.Token{}) {
			require.Equal(t, len(exp), i, "Expected scanner to return more tokens")
			return
		}

		require.Equal(t, exp[i], tk)
	}
}

func TestScanner_Next_1(t *testing.T) {

	sc := NewScanner("\r\n" +
		" \t\r\v\f" + "// comment" + "\n" +
		"123" + " " + "123.456" + "\r\n" +
		"`abc`" + `"abc"` + "\n" +
		"abc_xyz" + "\r\n" +
		"F" + " " + "MATCH" + " " + "TRUE" + "\n" +
		"@" + "~" + ":=" + "*" + "->" + ")" + "\r\n")

	doTest(t, sc,
		// Line 0
		tok(token.NEWLINE, "\r\n", 0, 0),
		// Line 1
		tok(token.WHITESPACE, " \t\r\v\f", 1, 0),
		tok(token.COMMENT, "// comment", 1, 5),
		tok(token.NEWLINE, "\n", 1, 15),
		// Line 2
		tok(token.INT, "123", 2, 0),
		tok(token.WHITESPACE, " ", 2, 3),
		tok(token.REAL, "123.456", 2, 4),
		tok(token.NEWLINE, "\r\n", 2, 11),
		// Line 3
		tok(token.STR, "`abc`", 3, 0),
		tok(token.TEMPLATE, `"abc"`, 3, 5),
		tok(token.NEWLINE, "\n", 3, 10),
		// Line 4
		tok(token.ID, "abc_xyz", 4, 0),
		tok(token.NEWLINE, "\r\n", 4, 7),
		// Line 5
		tok(token.FUNC, "F", 5, 0),
		tok(token.WHITESPACE, " ", 5, 1),
		tok(token.MATCH, "MATCH", 5, 2),
		tok(token.WHITESPACE, " ", 5, 7),
		tok(token.BOOL, "TRUE", 5, 8),
		tok(token.NEWLINE, "\n", 5, 12),
		// Line 6
		tok(token.SPELL, "@", 6, 0),
		tok(token.NOT, "~", 6, 1),
		tok(token.ASSIGN, ":=", 6, 2),
		tok(token.MULTIPLY, "*", 6, 4),
		tok(token.RETURNS, "->", 6, 5),
		tok(token.CLOSE_PAREN, ")", 6, 7),
		tok(token.NEWLINE, "\r\n", 6, 8),
		// Line 7
	)

	expEOF := token.Token{
		Kind: token.EOF,
		Line: 7,
		Col:  0,
	}

	require.Equal(t, expEOF, sc.Next())
}

func TestScanner_Next_2(t *testing.T) {
	require.Panics(t, func() {
		NewScanner("123.a").Next()
	})
}
