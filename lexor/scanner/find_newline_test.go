package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindNewline_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findNewline
}

func TestFindNewline_2(t *testing.T) {
	// Check it works when `\n` is the only input token.

	in := "\n"
	expN, expK := 1, token.NEWLINE
	tokenFinderTest(t, findNewline, in, expN, expK)
}

func TestFindNewline_3(t *testing.T) {
	// Check it works when `\r\n` is the only input token.

	in := "\r\n"
	expN, expK := 2, token.NEWLINE
	tokenFinderTest(t, findNewline, in, expN, expK)
}

func TestFindNewline_4(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a newline
	// is the first.

	in := "\r\nabc"
	expN, expK := 2, token.NEWLINE
	tokenFinderTest(t, findNewline, in, expN, expK)
}

func TestFindNewline_5(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a newline.

	in := "   "
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findNewline, in, expN, expK)
}
