package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindStrTemplate_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findStrTemplate
}

func TestFindStrTemplate_2(t *testing.T) {
	// Check it works when a string template is the only input token.

	in := `"abc @~\""`
	tokenFinderTest(t, findStrTemplate, in, 10, token.STR_TEMPLATE)
}

func TestFindStrTemplate_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a string
	// template is the first.

	in := `"abc" efg`
	tokenFinderTest(t, findStrTemplate, in, 5, token.STR_TEMPLATE)
}

func TestFindStrTemplate_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a string
	// template.

	in := `   `
	tokenFinderTest(t, findStrTemplate, in, 0, token.UNDEFINED)
}
