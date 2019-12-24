package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindNumLiteral_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findNumLiteral
}

func TestFindNumLiteral_2(t *testing.T) {
	// Check it works when an integer is the only input token.

	in := "123"
	expN, expK := 3, token.NUM_LITERAL
	tokenFinderTest(t, findNumLiteral, in, expN, expK)
}

func TestFindNumLiteral_3(t *testing.T) {
	// Check it works when a floating point number is the only input token.

	in := "123.456"
	expN, expK := 7, token.NUM_LITERAL
	tokenFinderTest(t, findNumLiteral, in, expN, expK)
}

func TestFindNumLiteral_4(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a number
	// is the first.

	in := "123.456)"
	expN, expK := 7, token.NUM_LITERAL
	tokenFinderTest(t, findNumLiteral, in, expN, expK)
}

func TestFindNumLiteral_5(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a number.

	in := "(123.456"
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findNumLiteral, in, expN, expK)
}
