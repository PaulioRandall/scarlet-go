package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindStrLiteral_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findStrLiteral
}

func TestFindStrLiteral_2(t *testing.T) {
	// Check it works when a string literal is the only input token.

	in := "`abc @~\"`"
	expN, expK := 9, token.STR_LITERAL
	tokenFinderTest(t, findStrLiteral, in, expN, expK)
}

func TestFindStrLiteral_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a string
	// literal is the first.

	in := "`abc` efg"
	expN, expK := 5, token.STR_LITERAL
	tokenFinderTest(t, findStrLiteral, in, expN, expK)
}

func TestFindStrLiteral_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a string
	// literal.

	in := "   "
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findStrLiteral, in, expN, expK)
}
