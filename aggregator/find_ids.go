package aggregator

import (
	"github.com/PaulioRandall/scarlet-go/stat"
	"github.com/PaulioRandall/scarlet-go/token"
)

// findIds satisfies the SequenceFinder function prototype. It attempts to
// match the next set of tokens to a delimitered set of IDs returning its
// length if matched.
func findIds(tok []token.Token) (n int, k stat.Kind, e token.Perror) {

	nextIsID := true

	for i, t := range tok {

		switch {
		case i == 0:
			if t.Kind() != token.ID {
				return
			}
		case nextIsID:
			if e = expectId(t); e != nil {
				return
			}
		case t.Kind() != token.ID_DELIM:
			goto DONE
		}

		nextIsID = !nextIsID
		n++
	}

DONE:
	if n > 0 {
		k = stat.OK
	}

	return
}

// expectId accepts a token kind and returns an error if it is not the ID kind.
func expectId(t token.Token) (e token.Perror) {
	if t.Kind() != token.ID {
		e = token.PerrorBySnippet(
			"Expected an ID",
			t.Where(),
		)
	}
	return
}
