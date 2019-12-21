package scanner

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token2"
)

var symbolSet map[string]token.Kind = map[string]token.Kind{
	":=": token.ASSIGN,
	"(":  token.OPEN_PAREN,
	")":  token.CLOSE_PAREN,
	",":  token.ID_DELIM,
	"@":  token.SPELL,
}

// findSymbol satisfies the source.TokenFinder function prototype. It attempts
// to match the next token to a symbol kind returning its length if matched.
func findSymbol(r []rune) (_ int, _ token.Kind, _ error) {

	src := string(r)

	for s, k := range symbolSet {
		if strings.HasPrefix(src, s) {
			return len([]rune(s)), k, nil
		}
	}

	return
}
