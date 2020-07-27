package formatter

/*
import (
	"strings"
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"

	"github.com/stretchr/testify/require"
)

func Test_scanLines_1(t *testing.T) {

	given := "@Set(\"a\", 1)\n"

	exp := lextest.Feign(
		lextest.Lex(0, 0, "@Set", lexeme.PR_SPELL),
		lextest.Lex(0, 4, "(", lexeme.PR_DELIMITER, lexeme.PR_PARENTHESIS, lexeme.PR_OPENER),
		lextest.Lex(0, 5, `"a"`, lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_STRING),
		lextest.Lex(0, 8, ",", lexeme.PR_DELIMITER, lexeme.PR_SEPARATOR),
		lextest.Lex(0, 9, " ", lexeme.PR_REDUNDANT, lexeme.PR_WHITESPACE),
		lextest.Lex(0, 10, "1", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Lex(0, 11, ")", lexeme.PR_DELIMITER, lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER),
		lextest.Lex(0, 12, "\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
	)

	r := strings.NewReader(given)
	act, e := scanLines(r)

	require.Nil(t, e, "Unexpected error:\n%+v", e)
	require.Equal(t, given, act.raw)
	lextest.Equal(t, exp, act.head)
	require.Nil(t, act.next)
}

func Test_scanLines_2(t *testing.T) {

	r := strings.NewReader("1\n2\n3\n")
	act, e := scanLines(r)

	require.Nil(t, e, "Unexpected error:\n%+v", e)

	exp := lextest.Feign(
		lextest.Lex(0, 0, "1", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Lex(0, 1, "\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
	)
	require.Equal(t, "1\n", act.raw)
	lextest.Equal(t, exp, act.head)
	require.NotNil(t, act.next)

	act = act.next
	exp = lextest.Feign(
		lextest.Lex(1, 0, "2", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Lex(1, 1, "\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
	)
	require.Equal(t, "2\n", act.raw)
	lextest.Equal(t, exp, act.head)
	require.NotNil(t, act.next)

	act = act.next
	exp = lextest.Feign(
		lextest.Lex(2, 0, "3", lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Lex(2, 1, "\n", lexeme.PR_TERMINATOR, lexeme.PR_NEWLINE),
	)
	require.Equal(t, "3\n", act.raw)
	lextest.Equal(t, exp, act.head)
	require.Nil(t, act.next)
}
*/
