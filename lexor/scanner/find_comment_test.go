package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindComment_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findComment
}

func TestFindComment_2(t *testing.T) {
	// Check it works when a comment is the only input token.

	in := "// Die Hard is a Christmas movie"
	expN, expK := 32, token.COMMENT
	tokenFinderTest(t, findComment, in, expN, expK)
}

func TestFindComment_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a comment is
	// the first.

	in := "// abc\nefg"
	expN, expK := 6, token.COMMENT
	tokenFinderTest(t, findComment, in, expN, expK)
}

func TestFindComment_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a comment.

	in := "   "
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findComment, in, expN, expK)
}
