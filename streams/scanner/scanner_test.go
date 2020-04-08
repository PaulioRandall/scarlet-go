package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"

	"github.com/stretchr/testify/require"
)

func tok(lex lexeme.Lexeme, v string, l, c int) lexeme.Token {
	return lexeme.New(lex, v, l, c)
}

func doTest(t *testing.T, ts token.TokenStream, exp ...lexeme.Token) {

	for i := 0; i < len(exp); i++ {

		tk := ts.Next()

		if tk == (lexeme.Token{}) {
			require.Equal(t, len(exp), i, "Expected scanner to return more tokens")
			return
		}

		require.Equal(t, exp[i], tk)
	}
}

func TestScanner_Next_1(t *testing.T) {

	sc := New("\r\n" +
		" \t\r\v\f" + "// comment" + "\n" +
		"123" + " " + "123.456" + "\r\n" +
		"`abc`" + `"abc"` + "\n" +
		"abc_xyz" + "\r\n" +
		"F" + " " + "TRUE" + "\n" +
		"@" + "~" + ":=" + "*" + "->" + ")" + "\r\n")

	doTest(t, sc,
		// Line 0
		tok(lexeme.LEXEME_NEWLINE, "\r\n", 0, 0),
		// Line 1
		tok(lexeme.LEXEME_WHITESPACE, " \t\r\v\f", 1, 0),
		tok(lexeme.LEXEME_COMMENT, "// comment", 1, 5),
		tok(lexeme.LEXEME_NEWLINE, "\n", 1, 15),
		// Line 2
		tok(lexeme.LEXEME_INT, "123", 2, 0),
		tok(lexeme.LEXEME_WHITESPACE, " ", 2, 3),
		tok(lexeme.LEXEME_FLOAT, "123.456", 2, 4),
		tok(lexeme.LEXEME_NEWLINE, "\r\n", 2, 11),
		// Line 3
		tok(lexeme.LEXEME_STRING, "`abc`", 3, 0),
		tok(lexeme.LEXEME_TEMPLATE, `"abc"`, 3, 5),
		tok(lexeme.LEXEME_NEWLINE, "\n", 3, 10),
		// Line 4
		tok(lexeme.LEXEME_ID, "abc_xyz", 4, 0),
		tok(lexeme.LEXEME_NEWLINE, "\r\n", 4, 7),
		// Line 5
		tok(lexeme.LEXEME_FUNC, "F", 5, 0),
		tok(lexeme.LEXEME_WHITESPACE, " ", 5, 1),
		tok(lexeme.LEXEME_BOOL, "TRUE", 5, 2),
		tok(lexeme.LEXEME_NEWLINE, "\n", 5, 6),
		// Line 6
		tok(lexeme.LEXEME_SPELL, "@", 6, 0),
		tok(lexeme.LEXEME_NOT, "~", 6, 1),
		tok(lexeme.LEXEME_ASSIGN, ":=", 6, 2),
		tok(lexeme.LEXEME_MULTIPLY, "*", 6, 4),
		tok(lexeme.LEXEME_RETURNS, "->", 6, 5),
		tok(lexeme.LEXEME_CLOSE_PAREN, ")", 6, 7),
		tok(lexeme.LEXEME_NEWLINE, "\r\n", 6, 8),
		// Line 7
	)

	expEOF := lexeme.Token{
		Lexeme: lexeme.LEXEME_EOF,
		Line:   7,
		Col:    0,
	}

	require.Equal(t, expEOF, sc.Next())
}

func TestScanner_Next_2(t *testing.T) {
	require.Panics(t, func() {
		New("123.a").Next()
	})
}
