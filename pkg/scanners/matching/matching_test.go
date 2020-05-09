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

	tester.Run(tests.A1_Newlines)
	tester.Run(tests.A2_Whitespace)
	tester.Run(tests.A3_Comments)
	tester.Run(tests.A4_Match)
	tester.Run(tests.A5_Bool_False)
	tester.Run(tests.A6_Bool_True)
}
