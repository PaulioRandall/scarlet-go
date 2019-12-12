package strimmer

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/lexor/scanner"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) lexor.ScanToken {
	st := scanner.New(src)
	return wrap(st)
}

// wrap wraps a ScanToken function with one that iterates through the scanner to
// find and return the next significant token. It effectively filters all
// insgnificant tokens for the user.
func wrap(f lexor.ScanToken) lexor.ScanToken {

	if f == nil {
		return nil
	}

	return func() (t token.Token, st lexor.ScanToken, e lexor.ScanErr) {

		for st = f; st != nil; {

			t, st, e = st()

			if e != nil || t == nil {
				return
			}

			if isSignificant(t.Kind()) {
				st = wrap(st)
				return
			}
		}

		return
	}
}

// isSignificant returns true if the input `k` is an essential token to the
// parsing of a script, i.e. not whiespace or a comment.
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
