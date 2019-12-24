package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindWord_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findWord
}

func TestFindWord_2(t *testing.T) {
	// Check it works when a word is the only input token.

	tokenFinderTest(t, findWord, "GLOBAL", 6, token.GLOBAL)
	tokenFinderTest(t, findWord, "F", 1, token.FUNC)
	tokenFinderTest(t, findWord, "DO", 2, token.DO)
	tokenFinderTest(t, findWord, "WATCH", 5, token.WATCH)
	tokenFinderTest(t, findWord, "MATCH", 5, token.MATCH)
	tokenFinderTest(t, findWord, "END", 3, token.END)
	tokenFinderTest(t, findWord, "TRUE", 4, token.TRUE)
	tokenFinderTest(t, findWord, "FALSE", 5, token.FALSE)
	tokenFinderTest(t, findWord, "an_identifier", 13, token.ID)
}

func TestFindWord_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a word
	// is the first.

	in := "F END"
	expN, expK := 1, token.FUNC
	tokenFinderTest(t, findWord, in, expN, expK)
}

func TestFindWord_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a word.

	in := "   "
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findWord, in, expN, expK)
}
