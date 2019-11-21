package strimmer

import (
	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/scanner"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) token.ScanToken {
	st := scanner.New(src)
	return wrap(st)
}

// wrap wraps a ScanToken function with one that iterates through the scanner to
// find and return the next significant token. It effectively filters all
// insgnificant tokens for the user.
func wrap(f token.ScanToken) token.ScanToken {
	return func() (t token.Token, st token.ScanToken, e perror.Perror) {

		t, st, e = f()

		for e == nil && t != token.Empty() {
			if t.IsSignificant() {
				st = wrap(st)
				return
			}

			t, st, e = st()
		}

		/* Why does this result in an infinite loop?
		for t, st, e = f(); e == nil && t != token.Empty(); {
			if t.IsSignificant() {
				st = wrap(st)
				return
			}
		}
		*/

		return
	}
}
