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

func Test_A4_Key_Match(t *testing.T) {
	tests.Run(t, ScanAll, tests.A4_Key_Match)
}

func Test_A5_Key_False(t *testing.T) {
	tests.Run(t, ScanAll, tests.A5_Key_False)
}

func Test_A6_Key_True(t *testing.T) {
	tests.Run(t, ScanAll, tests.A6_Key_True)
}

func Test_A7_Key_List(t *testing.T) {
	tests.Run(t, ScanAll, tests.A7_Key_List)
}

func Test_A8_Key_Fix(t *testing.T) {
	tests.Run(t, ScanAll, tests.A8_Key_Fix)
}

func Test_A9_Key_Eof(t *testing.T) {
	tests.Run(t, ScanAll, tests.A9_Key_Eof)
}

func Test_A10_Key_F(t *testing.T) {
	tests.Run(t, ScanAll, tests.A10_Key_F)
}

func Test_A11_Identifiers(t *testing.T) {
	tests.Run(t, ScanAll, tests.A11_Identifiers)
}

func Test_A12_Sym_Assign(t *testing.T) {
	tests.Run(t, ScanAll, tests.A12_Sym_Assign)
}

func Test_A13_Sym_Returns(t *testing.T) {
	tests.Run(t, ScanAll, tests.A13_Sym_Returns)
}

func Test_A14_Sym_LessThanOrEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.A14_Sym_LessThanOrEqual)
}

func Test_A15_Sym_MoreThanOrEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.A15_Sym_MoreThanOrEqual)
}
