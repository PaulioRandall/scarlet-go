package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestFindSymbol_1(t *testing.T) {
	// Check it is a type of lexor.TokenFinder.
	var _ TokenFinder = findSymbol
}

func TestFindSymbol_2(t *testing.T) {
	// Check it works on a range of valid inputs.

	f := func(in string, expN int, expK token.Kind) {
		tokenFinderTest(t, findSymbol, in, expN, expK)
	}

	// When input contains only one token, a symbol token
	f(":=", 2, token.ASSIGN)
	f("(", 1, token.OPEN_PAREN)
	f(")", 1, token.CLOSE_PAREN)
	f(",", 1, token.ID_DELIM)
	f("@", 1, token.SPELL)
	f("{", 1, token.OPEN_LIST)
	f("}", 1, token.CLOSE_LIST)

	f("+", 1, token.ADD)
	f("-", 1, token.SUBTRACT)
	f("/", 1, token.DIVIDE)
	f("*", 1, token.MULTIPLY)
	f("%", 1, token.MODULO)
	f("|", 1, token.OR)
	f("&", 1, token.AND)
	f("=", 1, token.EQUAL)
	f("#", 1, token.NOT_EQUAL)
	f("<=", 2, token.LT_OR_EQUAL)
	f(">=", 2, token.GT_OR_EQUAL)
	f("<", 1, token.LT)
	f(">", 1, token.GT)

	// When input contains multiple tokens, but the first is a symbol token
	f(":= 123.456", 2, token.ASSIGN)
	f("@Abc", 1, token.SPELL)
}

func TestFindSymbol_4(t *testing.T) {
	// Check 0 and UNDEFINED are returned when the first token is not a symbol.

	in := `  :=`
	expN, expK := 0, token.UNDEFINED
	tokenFinderTest(t, findSymbol, in, expN, expK)
}
