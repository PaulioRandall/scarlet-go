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
		"`abc`" + `"abc"`)

	doTest(t, s,
		// Line 0
		tok(token.NEWLINE, "\r\n", 0, 0),
		// Line 1
		tok(token.WHITESPACE, " \t\r\v\f", 1, 0),
		tok(token.COMMENT, "// comment", 1, 5),
		tok(token.NEWLINE, "\n", 1, 15),
		// Line 2
		tok(token.INT_LITERAL, "123", 2, 0), //
		tok(token.WHITESPACE, " ", 2, 3),
		tok(token.REAL_LITERAL, "123.456", 2, 4),
		tok(token.NEWLINE, "\r\n", 2, 11),
		// Line 3
		tok(token.STR_LITERAL, "`abc`", 3, 0),
		tok(token.STR_TEMPLATE, `"abc"`, 3, 5),
	)

	assert.Empty(t, s.Next())
}

func TestScanner_Next_2(t *testing.T) {
	require.Panics(t, func() {
		New("123.a").Next()
	})
}
