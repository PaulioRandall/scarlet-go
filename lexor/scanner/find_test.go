package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func tokenFinderTest(
	t *testing.T,
	f TokenFinder,
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
	f TokenFinder,
	in string,
) {

	n, k, e := f([]rune(in))

	require.NotNil(t, e, "Expected error")
	assert.Empty(t, n, "Expected `n` to be 0")
	assert.Empty(t, k, "Expected token.UNDEFINED")
}

func TestFind__1(t *testing.T) {
	// Check the find... functions are a type of lexor.TokenFinder.
	var _ TokenFinder = findComment
	var _ TokenFinder = findSpace
	var _ TokenFinder = findNumLiteral
	var _ TokenFinder = findWord
	var _ TokenFinder = findStrLiteral
	var _ TokenFinder = findStrTemplate
	var _ TokenFinder = findSymbol
}

func TestFind__2(t *testing.T) {

	// 1. testing.T
	// 2. TokenFinder
	// 3. Input text
	// 4. Expected token length
	// 5. Expected token kind

	// Check it works when the input text only contains one token.
	tokenFinderTest(t, findComment, "// Die Hard", 11, token.COMMENT)
	tokenFinderTest(t, findSpace, " \t\v\f", 4, token.WHITESPACE)
	tokenFinderTest(t, findSpace, "\n", 1, token.NEWLINE)
	tokenFinderTest(t, findSpace, "\r\n", 2, token.NEWLINE)
	tokenFinderTest(t, findNumLiteral, "123", 3, token.NUM_LITERAL)
	tokenFinderTest(t, findNumLiteral, "123.456", 7, token.NUM_LITERAL)
	tokenFinderTest(t, findStrLiteral, "`abc @~\"`", 9, token.STR_LITERAL)
	tokenFinderTest(t, findStrTemplate, `"abc \n@~\""`, 12, token.STR_TEMPLATE)
	tokenFinderTest(t, findWord, "GLOBAL", 6, token.GLOBAL)
	tokenFinderTest(t, findWord, "F", 1, token.FUNC)
	tokenFinderTest(t, findWord, "DO", 2, token.DO)
	tokenFinderTest(t, findWord, "WATCH", 5, token.WATCH)
	tokenFinderTest(t, findWord, "MATCH", 5, token.MATCH)
	tokenFinderTest(t, findWord, "END", 3, token.END)
	tokenFinderTest(t, findWord, "TRUE", 4, token.TRUE)
	tokenFinderTest(t, findWord, "FALSE", 5, token.FALSE)
	tokenFinderTest(t, findWord, "an_identifier", 13, token.ID)
	tokenFinderTest(t, findSymbol, ":=", 2, token.ASSIGN)
	tokenFinderTest(t, findSymbol, "->", 2, token.RETURNS)
	tokenFinderTest(t, findSymbol, "(", 1, token.OPEN_PAREN)
	tokenFinderTest(t, findSymbol, ")", 1, token.CLOSE_PAREN)
	tokenFinderTest(t, findSymbol, "[", 1, token.OPEN_GUARD)
	tokenFinderTest(t, findSymbol, "]", 1, token.CLOSE_GUARD)
	tokenFinderTest(t, findSymbol, "{", 1, token.OPEN_LIST)
	tokenFinderTest(t, findSymbol, "}", 1, token.CLOSE_LIST)
	tokenFinderTest(t, findSymbol, ",", 1, token.ID_DELIM)
	tokenFinderTest(t, findSymbol, "@", 1, token.SPELL)
	tokenFinderTest(t, findSymbol, "+", 1, token.ADD)
	tokenFinderTest(t, findSymbol, "-", 1, token.SUBTRACT)
	tokenFinderTest(t, findSymbol, "/", 1, token.DIVIDE)
	tokenFinderTest(t, findSymbol, "*", 1, token.MULTIPLY)
	tokenFinderTest(t, findSymbol, "%", 1, token.MODULO)
	tokenFinderTest(t, findSymbol, "|", 1, token.OR)
	tokenFinderTest(t, findSymbol, "&", 1, token.AND)
	tokenFinderTest(t, findSymbol, "~", 1, token.NOT)
	tokenFinderTest(t, findSymbol, "¬", 1, token.NOT)
	tokenFinderTest(t, findSymbol, "=", 1, token.EQUAL)
	tokenFinderTest(t, findSymbol, "#", 1, token.NOT_EQUAL)
	tokenFinderTest(t, findSymbol, "<=", 2, token.LT_OR_EQUAL)
	tokenFinderTest(t, findSymbol, ">=", 2, token.GT_OR_EQUAL)
	tokenFinderTest(t, findSymbol, "<", 1, token.LT)
	tokenFinderTest(t, findSymbol, ">", 1, token.GT)
	tokenFinderTest(t, findSymbol, "_", 1, token.VOID)
}

func TestFind__3(t *testing.T) {

	// 1. testing.T
	// 2. TokenFinder
	// 3. Input text
	// 4. Expected token length
	// 5. Expected token kind

	// Check it works when there are multiple tokens in the input and the token
	// under test is first.
	tokenFinderTest(t, findComment, "// Die\nHard", 6, token.COMMENT)
	tokenFinderTest(t, findSpace, "  \b", 2, token.WHITESPACE)
	tokenFinderTest(t, findSpace, "\nabc", 1, token.NEWLINE)
	tokenFinderTest(t, findSpace, "\r\nabc", 2, token.NEWLINE)
	tokenFinderTest(t, findNumLiteral, "123.456abc", 7, token.NUM_LITERAL)
	tokenFinderTest(t, findStrLiteral, "`abc` efg", 5, token.STR_LITERAL)
	tokenFinderTest(t, findStrTemplate, `"abc" efg`, 5, token.STR_TEMPLATE)
	tokenFinderTest(t, findWord, "F()", 1, token.FUNC)
	tokenFinderTest(t, findSymbol, ":= 123.456", 2, token.ASSIGN)
}

func TestFind__4(t *testing.T) {

	// 1. testing.T
	// 2. TokenFinder
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
	// 2. TokenFinder
	// 3. Input text
	// 4. Expected token length
	// 5. Expected token kind

	// Check an error is returned if the text is a malformed instance of the
	// token under test.
	tokenFinderErrTest(t, findNumLiteral, "123.")
	tokenFinderErrTest(t, findStrLiteral, "`abc")
	tokenFinderErrTest(t, findStrTemplate, `"abc`)
}
