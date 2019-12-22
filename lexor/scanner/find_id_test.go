package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindId_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findId
}

func TestFindId_2(t *testing.T) {
	// Check it works when an ID is the only input token.

	in := "abc"
	expN, expK := 3, token.ID
	tokenFinderTest(t, findId, in, expN, expK)
}

func TestFindId_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and an ID is
	// the first.

	in := "abc efg"
	expN, expK := 3, token.ID
	tokenFinderTest(t, findId, in, expN, expK)
}

func TestFindId_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not an ID.

	in := "   "
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findId, in, expN, expK)
}
