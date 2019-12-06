package scanner

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

var symbolSet map[string]token.Kind = map[string]token.Kind{
	":=": token.ASSIGN,
}

// findSymbol satisfies the source.TokenFinder function prototype. It attempts
// to match the next token to a symbol kind returning its length if matched.
func findSymbol(runes []rune) (int, token.Kind) {

	src := string(runes)

	for s, k := range symbolSet {
		if strings.HasPrefix(src, s) {
			return len([]rune(s)), k
		}
	}

	return 0, token.UNDEFINED
}
