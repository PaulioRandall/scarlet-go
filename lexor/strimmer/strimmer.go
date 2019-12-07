package strimmer

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/lexor/scanner"
	"github.com/PaulioRandall/scarlet-go/perror"
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

	return func() (t token.Token, st lexor.ScanToken, e perror.Perror) {

		for st = f; st != nil; {

			t, st, e = st()

			if e != nil || t == nil {
				break
			}

			if !ignore(t.Kind()) {
				st = wrap(st)
				break
			}
		}

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
