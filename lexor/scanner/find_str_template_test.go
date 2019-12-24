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

	in := `"abc \n@~\""`
	expN, expK := 12, token.STR_TEMPLATE
	tokenFinderTest(t, findStrTemplate, in, expN, expK)
}

func TestFindStrTemplate_3(t *testing.T) {
	// Check it works when there are multiple tokens in the input and a string
	// template is the first.

	in := `"abc" efg`
	expN, expK := 5, token.STR_TEMPLATE
	tokenFinderTest(t, findStrTemplate, in, expN, expK)
}

func TestFindStrTemplate_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a string
	// template.

	in := `   `
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findStrTemplate, in, expN, expK)
}
