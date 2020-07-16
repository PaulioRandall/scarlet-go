package program

import (
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

func writeTokenPhaseFile(filename string, tks []token.Token) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return token.PrintAll(f, tks)
}
