package matching

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/scanners/tests"
)

func Test_A1(t *testing.T) {
	tests.Run(t, ScanAll, tests.A1_Newlines)
}

func Test_A2(t *testing.T) {
	tests.Run(t, ScanAll, tests.A2_Whitespace)
}

func Test_A3(t *testing.T) {
	tests.Run(t, ScanAll, tests.A3_Comments)
}

func Test_A4(t *testing.T) {
	tests.Run(t, ScanAll, tests.A4_Match)
}

func Test_A5(t *testing.T) {
	tests.Run(t, ScanAll, tests.A5_Bool_False)
}

func Test_A6(t *testing.T) {
	tests.Run(t, ScanAll, tests.A6_Bool_True)
}
