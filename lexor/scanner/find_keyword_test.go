package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindKeyword_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findKeyword
}

func TestFindKeyword_2(t *testing.T) {
	// Check it works when a keyword is the only input token.

	tokenFinderTest(t, findKeyword, "F", 1, token.FUNC)
	tokenFinderTest(t, findKeyword, "DO", 2, token.DO)
	tokenFinderTest(t, findKeyword, "WATCH", 5, token.WATCH)
	tokenFinderTest(t, findKeyword, "MATCH", 5, token.MATCH)
	tokenFinderTest(t, findKeyword, "END", 3, token.END)
	tokenFinderTest(t, findKeyword, "TRUE", 4, token.TRUE)
	tokenFinderTest(t, findKeyword, "FALSE", 5, token.FALSE)
}

func TestFindKeyword_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a keyword
	// is the first.

	in := "F END"
	expN, expK := 1, token.FUNC
	tokenFinderTest(t, findKeyword, in, expN, expK)
}

func TestFindKeyword_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a keyword.

	in := "   "
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findId, in, expN, expK)
}
