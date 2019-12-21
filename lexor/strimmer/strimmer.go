// strimmer removes insignificant tokens, i.e. removing tokens that are not
// required for parsing such as whitespace. It wraps another ScanToken function.
package strimmer

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New makes a new strimmer function.
func New(src lexor.ScanToken) lexor.ScanToken {
	return wrap(src)
}

// wrap wraps a ScanToken with one that removes insignificant tokens. This means
// the underlying recursive function maybe called multiple times before a value
// is returned.
func wrap(f lexor.ScanToken) lexor.ScanToken {

	if f == nil {
		return nil
	}

	return func() (t token.Token, st lexor.ScanToken, e lexor.ScanErr) {

		for st = f; st != nil; {

			t, st, e = st()

			if e != nil || t == (token.Token{}) {
				return
			}

			if isSignificant(t.Kind) {
				st = wrap(st)
				return
			}
		}

		return
	}
}

// isSignificant returns true if the input `k` is an significant token to the
// parsing of a script, i.e. not whiespace or a comment etc.
func isSignificant(k token.Kind) bool {

	ks := []token.Kind{
		token.UNDEFINED,
		token.WHITESPACE,
		token.COMMENT,
	}

	for _, j := range ks {
		if k == j {
			return false
		}
	}

	return true
}
