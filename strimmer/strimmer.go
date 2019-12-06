package strimmer

import (
	"github.com/PaulioRandall/scarlet-go/perror"
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

	if f == nil {
		return nil
	}

	return func() (t token.Token, st token.ScanToken, e perror.Perror) {

		for {

			t, st, e = f()

			if e == nil || t == token.Empty() {
				break
			}

			if !ignore(t.Kind()) {
				st = wrap(st)
				break
			}

			if st == nil {
				break
			}
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

// ignore returns true if the input `k` is a kind for a token that should be
// discarded.
func ignore(k token.Kind) bool {

	ks := []token.Kind{
		token.UNDEFINED,
		token.WHITESPACE,
	}

	for _, j := range ks {
		if k == j {
			return true
		}
	}

	return false
}
