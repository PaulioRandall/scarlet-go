package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindSpace_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findSpace
}

func TestFindSpace_2(t *testing.T) {
	// Check it works when whitespace is the only input token.

	in := " \t\v\f"
	expN, expK := 4, token.WHITESPACE
	tokenFinderTest(t, findSpace, in, expN, expK)
}

func TestFindSpace_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and whitespace
	// is the first.

	in := "  \b"
	expN, expK := 2, token.WHITESPACE
	tokenFinderTest(t, findSpace, in, expN, expK)
}

func TestFindSpace_4(t *testing.T) {
	// Check it works when `\n` is the only input token.

	in := "\n"
	expN, expK := 1, token.NEWLINE
	tokenFinderTest(t, findSpace, in, expN, expK)
}

func TestFindSpace_5(t *testing.T) {
	// Check it works when `\r\n` is the only input token.

	in := "\r\n"
	expN, expK := 2, token.NEWLINE
	tokenFinderTest(t, findSpace, in, expN, expK)
}

func TestFindSpace_6(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a newline
	// is the first.

	in := "\r\nabc"
	expN, expK := 2, token.NEWLINE
	tokenFinderTest(t, findSpace, in, expN, expK)
}

func TestFindSpace_7(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not whitespace.

	in := "abc"
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findSpace, in, expN, expK)
}
