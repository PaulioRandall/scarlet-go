package aggregator

import (
	"github.com/PaulioRandall/scarlet-go/stat"
	"github.com/PaulioRandall/scarlet-go/token"
)

// findParams satisfies the SequenceFinder function prototype. It attempts to
// match the next set of tokens to a set of parameters enclosed by parentheses.
func findParams(tok []token.Token) (n int, k stat.Kind, e token.Perror) {

	size := len(tok)

	if size < 1 || tok[0].Kind() != token.OPEN_PAREN {
		return
	}

	// TODO: create finder for all types of parameters
	switch n, k, e = findIds(tok[1:]); {
	case e != nil:
	case k == stat.UNDEFINED:
		e = token.PerrorBySnippet("Expected set of IDs", tok[n].Where())
	case size <= n || tok[n].Kind() != token.CLOSE_PAREN:
		n, k = 0, stat.UNDEFINED
		e = token.PerrorBySnippet("Expected a closing parentheses", tok[n].Where())
	default:
		n++
	}

	return
}
