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
	tests.Run(t, ScanAll, tests.A4_Match)
}

func Test_A5_Key_False(t *testing.T) {
	tests.Run(t, ScanAll, tests.A5_False)
}

func Test_A6_Key_True(t *testing.T) {
	tests.Run(t, ScanAll, tests.A6_True)
}

func Test_A7_Key_List(t *testing.T) {
	tests.Run(t, ScanAll, tests.A7_List)
}

func Test_A8_Key_Fix(t *testing.T) {
	tests.Run(t, ScanAll, tests.A8_Fix)
}

func Test_A9_Key_Eof(t *testing.T) {
	tests.Run(t, ScanAll, tests.A9_Eof)
}

func Test_A10_Key_F(t *testing.T) {
	tests.Run(t, ScanAll, tests.A10_F)
}

func Test_A11_Identifiers(t *testing.T) {
	tests.Run(t, ScanAll, tests.A11_Identifiers)
}

func Test_A12_Sym_Assign(t *testing.T) {
	tests.Run(t, ScanAll, tests.A12_Assign)
}

func Test_A13_Sym_Returns(t *testing.T) {
	tests.Run(t, ScanAll, tests.A13_Returns)
}

func Test_A14_Sym_LessThanOrEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.A14_LessThanOrEqual)
}

func Test_A15_Sym_MoreThanOrEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.A15_MoreThanOrEqual)
}

func Test_A16_Sym_BlockOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.A16_BlockOpen)
}

func Test_A17_Sym_BlockOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.A17_BlockClose)
}

func Test_A18_Sym_ParenOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.A18_ParenOpen)
}

func Test_A19_Sym_ParenClose(t *testing.T) {
	tests.Run(t, ScanAll, tests.A19_ParenClose)
}

func Test_A20_Sym_GuardOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.A20_GuardOpen)
}

func Test_A21_Sym_GuardClose(t *testing.T) {
	tests.Run(t, ScanAll, tests.A21_GuardClose)
}

func Test_A22_Sym_Delim(t *testing.T) {
	tests.Run(t, ScanAll, tests.A22_Delim)
}

func Test_A23_Sym_Void(t *testing.T) {
	tests.Run(t, ScanAll, tests.A23_Void)
}

func Test_A24_Sym_Terminator(t *testing.T) {
	tests.Run(t, ScanAll, tests.A24_Terminator)
}

func Test_A25_Sym_Spell(t *testing.T) {
	tests.Run(t, ScanAll, tests.A25_Spell)
}
