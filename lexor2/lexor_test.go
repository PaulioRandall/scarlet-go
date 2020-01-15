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
		"// comment" + "\n" +
		" \t\v\f")

	doTest(t, s,
		tok(token.NEWLINE, "\r\n", 0, 0),
		tok(token.COMMENT, "// comment", 1, 0),
		tok(token.NEWLINE, "\n", 1, 10),
		tok(token.WHITESPACE, " \t\v\f", 2, 0),
	)

	assert.Empty(t, s.Next())
}
