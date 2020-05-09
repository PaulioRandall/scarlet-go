package matching

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/scanners/tests"
)

func Test_A1_Newlines(t *testing.T) {
	tests.Run(t, ScanAll, tests.A1_Newlines)
}

func Test_A2_Whitespace(t *testing.T) {
	tests.Run(t, ScanAll, tests.A2_Whitespace)
}

func Test_A3_Comments(t *testing.T) {
	tests.Run(t, ScanAll, tests.A3_Comments)
}

func Test_A4_Match(t *testing.T) {
	tests.Run(t, ScanAll, tests.A4_Match)
}

func Test_A5_Bool_False(t *testing.T) {
	tests.Run(t, ScanAll, tests.A5_Bool_False)
}

func Test_A6_Bool_True(t *testing.T) {
	tests.Run(t, ScanAll, tests.A6_Bool_True)
}

func Test_A7_List(t *testing.T) {
	tests.Run(t, ScanAll, tests.A7_List)
}

func Test_A8_Fix(t *testing.T) {
	tests.Run(t, ScanAll, tests.A8_Fix)
}

func Test_A9_Eof(t *testing.T) {
	tests.Run(t, ScanAll, tests.A9_Eof)
}
