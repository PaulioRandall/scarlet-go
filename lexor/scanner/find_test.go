package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFind__1(t *testing.T) {
	// Check the find... functions are a type of lexor.TokenFinder.
	var _ TokenFinder = findComment
	var _ TokenFinder = findSpace
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
}

func TestFind__4(t *testing.T) {

	// Check 0 and UNDEFINED are returned when the first token is not a comment.
	tokenFinderTest(t, findComment, "   // abc", 0, token.UNDEFINED)
	tokenFinderTest(t, findSpace, "abc   ", 0, token.UNDEFINED)
}
