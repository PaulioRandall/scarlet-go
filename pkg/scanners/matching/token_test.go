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

func Test_A5_False(t *testing.T) {
	tests.Run(t, ScanAll, tests.A5_False)
}

func Test_A6_True(t *testing.T) {
	tests.Run(t, ScanAll, tests.A6_True)
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

func Test_A10_F(t *testing.T) {
	tests.Run(t, ScanAll, tests.A10_F)
}

func Test_A11_Identifiers(t *testing.T) {
	tests.Run(t, ScanAll, tests.A11_Identifiers)
}

func Test_A12_Assign(t *testing.T) {
	tests.Run(t, ScanAll, tests.A12_Assign)
}

func Test_A13_Returns(t *testing.T) {
	tests.Run(t, ScanAll, tests.A13_Returns)
}

func Test_A14_LessThanOrEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.A14_LessThanOrEqual)
}

func Test_A15_MoreThanOrEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.A15_MoreThanOrEqual)
}

func Test_A16_BlockOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.A16_BlockOpen)
}

func Test_A17_BlockOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.A17_BlockClose)
}

func Test_A18_ParenOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.A18_ParenOpen)
}

func Test_A19_ParenClose(t *testing.T) {
	tests.Run(t, ScanAll, tests.A19_ParenClose)
}

func Test_A20_GuardOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.A20_GuardOpen)
}

func Test_A21_GuardClose(t *testing.T) {
	tests.Run(t, ScanAll, tests.A21_GuardClose)
}

func Test_A22_Delim(t *testing.T) {
	tests.Run(t, ScanAll, tests.A22_Delim)
}

func Test_A23_Void(t *testing.T) {
	tests.Run(t, ScanAll, tests.A23_Void)
}

func Test_A24_Terminator(t *testing.T) {
	tests.Run(t, ScanAll, tests.A24_Terminator)
}

func Test_A25_Spell(t *testing.T) {
	tests.Run(t, ScanAll, tests.A25_Spell)
}

func Test_A26_Add(t *testing.T) {
	tests.Run(t, ScanAll, tests.A26_Add)
}

func Test_A27_Subtract(t *testing.T) {
	tests.Run(t, ScanAll, tests.A27_Subtract)
}

func Test_A28_Multiply(t *testing.T) {
	tests.Run(t, ScanAll, tests.A28_Multiply)
}

func Test_A29_Divide(t *testing.T) {
	tests.Run(t, ScanAll, tests.A29_Divide)
}

func Test_A30_Remainder(t *testing.T) {
	tests.Run(t, ScanAll, tests.A30_Remainder)
}

func Test_A31_And(t *testing.T) {
	tests.Run(t, ScanAll, tests.A31_And)
}

func Test_A32_Or(t *testing.T) {
	tests.Run(t, ScanAll, tests.A32_Or)
}

func Test_A33_Equal(t *testing.T) {
	tests.Run(t, ScanAll, tests.A33_Equal)
}

func Test_A34_NotEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.A34_NotEqual)
}

func Test_A35_LessThan(t *testing.T) {
	tests.Run(t, ScanAll, tests.A35_LessThan)
}

func Test_A36_MoreThan(t *testing.T) {
	tests.Run(t, ScanAll, tests.A36_MoreThan)
}

func Test_A37_String(t *testing.T) {
	tests.Run(t, ScanAll, tests.A37_String)
}

func Test_A38_Template(t *testing.T) {
	tests.Run(t, ScanAll, tests.A38_Template)
}

func Test_A39_Number(t *testing.T) {
	tests.Run(t, ScanAll, tests.A39_Number)
}
