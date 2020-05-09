package matching

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/scanners/tests"
)

func TestScanner(t *testing.T) {

	tester := tests.Tester{
		T:   t,
		Tag: "scanner.matching",
		F:   ScanAll,
	}

	tester.Run(tests.A1)
	tester.Run(tests.A2)
}
