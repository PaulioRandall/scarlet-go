package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func tokenFinderTest(
	t *testing.T,
	f tokenFinder,
	in string,
	expN int,
	expK token.Kind,
) {

	n, k, e := f([]rune(in))

	require.Nil(t, e)
	assert.Equal(t, expN, n)
	assert.Equal(t, expK, k)
}

func tokenFinderErrTest(
	t *testing.T,
	f tokenFinder,
	in string,
) {

	n, k, e := f([]rune(in))

	require.NotNil(t, e, "Expected error")
	assert.Empty(t, n, "Expected `n` to be 0")
	assert.Empty(t, k, "Expected token.UNDEFINED")
}

func TestFind__1(t *testing.T) {
	// Check the find... functions are a type of `tokenFinder`.
	var _ tokenFinder = findComment
	var _ tokenFinder = findSpace
	var _ tokenFinder = findNumLiteral
	var _ tokenFinder = findWord
	var _ tokenFinder = findStrLiteral
	var _ tokenFinder = findStrTemplate
	var _ tokenFinder = findSymbol
}

func TestFind__2(t *testing.T) {

	// 1. testing.T
	// 2. tokenFinder function
	// 3. Input text
	// 4. Expected token length
	// 5. Expected token kind

	// Check it works when the input text only contains one token.
	tokenFinderTest(t, findComment, "// Die Hard", 11, token.COMMENT)
	tokenFinderTest(t, findSpace, " \t\v\f", 4, token.WHITESPACE)
	tokenFinderTest(t, findSpace, "\n", 1, token.NEWLINE)
	tokenFinderTest(t, findSpace, "\r\n", 2, token.NEWLINE)
	tokenFinderTest(t, findNumLiteral, "123", 3, token.INT_LITERAL)
	tokenFinderTest(t, findNumLiteral, "123.456", 7, token.REAL_LITERAL)
	tokenFinderTest(t, findStrLiteral, "`abc @~\"`", 9, token.STR_LITERAL)
	tokenFinderTest(t, findStrTemplate, `"abc \n@~\""`, 12, token.STR_TEMPLATE)
	tokenFinderTest(t, findWord, "GLOBAL", 6, token.GLOBAL)
	tokenFinderTest(t, findWord, "F", 1, token.FUNC)
	tokenFinderTest(t, findWord, "DO", 2, token.DO)
	tokenFinderTest(t, findWord, "WATCH", 5, token.WATCH)
	tokenFinderTest(t, findWord, "MATCH", 5, token.MATCH)
	tokenFinderTest(t, findWord, "END", 3, token.END)
	tokenFinderTest(t, findWord, "TRUE", 4, token.BOOL_LITERAL)
	tokenFinderTest(t, findWord, "FALSE", 5, token.BOOL_LITERAL)
	tokenFinderTest(t, findWord, "an_identifier", 13, token.ID)
	tokenFinderTest(t, findSymbol, ":=", 2, token.ASSIGN)
	tokenFinderTest(t, findSymbol, "->", 2, token.RETURNS)
	tokenFinderTest(t, findSymbol, "(", 1, token.OPEN_PAREN)
	tokenFinderTest(t, findSymbol, ")", 1, token.CLOSE_PAREN)
	tokenFinderTest(t, findSymbol, "[", 1, token.OPEN_GUARD)
	tokenFinderTest(t, findSymbol, "]", 1, token.CLOSE_GUARD)
	tokenFinderTest(t, findSymbol, "{", 1, token.OPEN_LIST)
	tokenFinderTest(t, findSymbol, "}", 1, token.CLOSE_LIST)
	tokenFinderTest(t, findSymbol, ",", 1, token.DELIM)
	tokenFinderTest(t, findSymbol, "@", 1, token.SPELL)
	tokenFinderTest(t, findSymbol, "+", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "-", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "/", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "*", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "%", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "|", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "&", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "~", 1, token.NOT)
	tokenFinderTest(t, findSymbol, "¬", 1, token.NOT)
	tokenFinderTest(t, findSymbol, "=", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "#", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "<=", 2, token.OPERATOR)
	tokenFinderTest(t, findSymbol, ">=", 2, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "<", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, ">", 1, token.OPERATOR)
	tokenFinderTest(t, findSymbol, "_", 1, token.VOID)
}

func TestFind__3(t *testing.T) {

	// 1. testing.T
	// 2. tokenFinder function
	// 3. Input text
	// 4. Expected token length
	// 5. Expected token kind

	// Check it works when there are multiple tokens in the input and the token
	// under test is first.
	tokenFinderTest(t, findComment, "// Die\nHard", 6, token.COMMENT)
	tokenFinderTest(t, findSpace, "  \b", 2, token.WHITESPACE)
	tokenFinderTest(t, findSpace, "\nabc", 1, token.NEWLINE)
	tokenFinderTest(t, findSpace, "\r\nabc", 2, token.NEWLINE)
	tokenFinderTest(t, findNumLiteral, "123.456abc", 7, token.REAL_LITERAL)
	tokenFinderTest(t, findStrLiteral, "`abc` efg", 5, token.STR_LITERAL)
	tokenFinderTest(t, findStrTemplate, `"abc" efg`, 5, token.STR_TEMPLATE)
	tokenFinderTest(t, findWord, "F()", 1, token.FUNC)
	tokenFinderTest(t, findSymbol, ":= 123.456", 2, token.ASSIGN)
}

func TestFind__4(t *testing.T) {

	// 1. testing.T
	// 2. tokenFinder function
	// 3. Input text
	// 4. Expected token length
	// 5. Expected token kind

	// Check 0 and UNDEFINED are returned when the first token is not a the
	// token under test.
	tokenFinderTest(t, findComment, "   // abc", 0, token.UNDEFINED)
	tokenFinderTest(t, findSpace, "abc   ", 0, token.UNDEFINED)
	tokenFinderTest(t, findNumLiteral, "   123", 0, token.UNDEFINED)
	tokenFinderTest(t, findWord, "   F", 0, token.UNDEFINED)
	tokenFinderTest(t, findStrLiteral, "   `abc`", 0, token.UNDEFINED)
	tokenFinderTest(t, findStrTemplate, `   "abc"`, 0, token.UNDEFINED)
	tokenFinderTest(t, findSymbol, `   :=`, 0, token.UNDEFINED)
}

func TestFind__5(t *testing.T) {

	// 1. testing.T
	// 2. tokenFinder function
	// 3. Input text
	// 4. Expected token length
	// 5. Expected token kind

	// Check an error is returned if the text is a malformed instance of the
	// token under test.
	tokenFinderErrTest(t, findNumLiteral, "123.")
	tokenFinderErrTest(t, findStrLiteral, "`abc")
	tokenFinderErrTest(t, findStrTemplate, `"abc`)
}
