package lexor

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func tok(k token.Kind, v string, l, c int) token.Token {
	return token.New(k, v, l, c)
}

func doTest(t *testing.T, scn *Scanner, exp ...token.Token) {

	for i := 0; i < len(exp); i++ {

		tk := scn.Next()

		if tk == token.ZERO() {
			require.Equal(t, len(exp), i, "Expected scanner to return more tokens")
			return
		}

		require.Equal(t, exp[i], tk)
	}
}

func TestScanner_Next_1(t *testing.T) {

	s := New("\r\n" +
		" \t\r\v\f" + "// comment" + "\n" +
		"123" + " " + "123.456" + "\r\n" +
		"`abc`" + `"abc"` + "\n" +
		"abc_xyz" + "\r\n" +
		"F" + " " + "MATCH" + " " + "TRUE" + "\n")

	doTest(t, s,
		// Line 0
		tok(token.NEWLINE, "\r\n", 0, 0),
		// Line 1
		tok(token.WHITESPACE, " \t\r\v\f", 1, 0),
		tok(token.COMMENT, "// comment", 1, 5),
		tok(token.NEWLINE, "\n", 1, 15),
		// Line 2
		tok(token.INT_LITERAL, "123", 2, 0),
		tok(token.WHITESPACE, " ", 2, 3),
		tok(token.REAL_LITERAL, "123.456", 2, 4),
		tok(token.NEWLINE, "\r\n", 2, 11),
		// Line 3
		tok(token.STR_LITERAL, "`abc`", 3, 0),
		tok(token.STR_TEMPLATE, `"abc"`, 3, 5),
		tok(token.NEWLINE, "\n", 3, 10),
		// Line 4
		tok(token.ID, "abc_xyz", 4, 0),
		tok(token.NEWLINE, "\r\n", 4, 7),
		// Line 5
		tok(token.FUNC, "F", 5, 0),
		tok(token.WHITESPACE, " ", 5, 1),
		tok(token.MATCH, "MATCH", 5, 2),
		tok(token.WHITESPACE, " ", 5, 7),
		tok(token.BOOL_LITERAL, "TRUE", 5, 8),
		tok(token.NEWLINE, "\n", 5, 12),
		// Line 6
	)

	assert.Empty(t, s.Next())
}

func TestScanner_Next_2(t *testing.T) {
	require.Panics(t, func() {
		New("123.a").Next()
	})
}
