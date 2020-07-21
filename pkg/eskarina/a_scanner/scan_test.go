package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme/lextest"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in string, exp *lexeme.Lexeme) {
	act, e := ScanStr(in)
	require.Nil(t, e, "%+v", e)
	lextest.Equal(t, exp, act)
}

func Test1_1(t *testing.T) {
	doTest(t, "\n", lextest.Feign(
		lextest.Lex(0, 0, "\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
	))
}

func Test1_2(t *testing.T) {
	doTest(t, "\r\n", lextest.Feign(
		lextest.Lex(0, 0, "\r\n", prop.PR_TERMINATOR, prop.PR_NEWLINE),
	))
}

func Test2_1(t *testing.T) {
	doTest(t, "# Comment", lextest.Feign(
		lextest.Lex(0, 0, "# Comment", prop.PR_REDUNDANT, prop.PR_COMMENT),
	))
}

func Test3_1(t *testing.T) {
	doTest(t, "   \t\v\f", lextest.Feign(
		lextest.Lex(0, 0, "   \t\v\f", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
	))
}
