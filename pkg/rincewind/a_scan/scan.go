package scan

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
)

type ScanFunc func() (token.Tok, ScanFunc, error)

type RuneItr interface {
	Next() (rune, bool)
}

func New(itr RuneItr) ScanFunc {

	if itr == nil {
		failNow("Non-nil RuneItr required")
	}

	scn := &scanner{
		itr: itr,
		col: -1, // -1 so index is before the first symbol
	}
	scn.bufferNext()

	if scn.empty() {
		return nil
	}

	return scn.scan
}

func ScanAll(itr RuneItr) ([]token.Token, error) {

	var (
		e   error
		tk  token.Token
		tks = []token.Token{}
	)

	for f := New(itr); f != nil; {
		if tk, f, e = f(); e != nil {
			return nil, e
		}

		tks = append(tks, tk)
	}

	return tks, nil
}
