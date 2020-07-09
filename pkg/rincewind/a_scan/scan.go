package scan

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type ScanFunc func() (token.Tok, ScanFunc, error)

type SymbolItr interface {
	Next() (rune, bool)
}

func New(itr SymbolItr) ScanFunc {

	if itr == nil {
		failNow("Non-nil SymbolItr required")
	}

	scn := &scanner{
		itr: itr,
		col: -1, // -1 so index is before first symbol
	}
	scn.bufferNext()

	if scn.empty() {
		return nil
	}

	return scn.scan
}
