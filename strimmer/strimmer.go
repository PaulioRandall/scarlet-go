package strimmer

import (
	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/scanner/source"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) token.ScanToken {
	st := source.New(src)
	return wrap(st)
}

// wrap
func wrap(f token.ScanToken) token.ScanToken {
	return func() (t token.Token, st token.ScanToken, e perror.Perror) {

		st = f

		for t, st, e = st(); e == nil && t != token.Empty(); {
			if t.IsSignificant() {
				st = wrap(st)
				return
			}
		}

		return
	}
}
