package matching

import (
	"testing"

	e "github.com/PaulioRandall/scarlet-go/pkg/err"

	tests "github.com/PaulioRandall/scarlet-go/pkg/scanners/tests"
)

func TestScanErr(t *testing.T) {
	var _ e.Err = scanErr{}
}

func Test_T1_Newlines(t *testing.T) {
	tests.Run(t, ScanAll, tests.T1_Newlines)
}

func Test_T2_Whitespace(t *testing.T) {
	tests.Run(t, ScanAll, tests.T2_Whitespace)
}

func Test_T3_Comments(t *testing.T) {
	tests.Run(t, ScanAll, tests.T3_Comments)
}

func Test_T4_Match(t *testing.T) {
	tests.Run(t, ScanAll, tests.T4_Match)
}

func Test_T5_False(t *testing.T) {
	tests.Run(t, ScanAll, tests.T5_False)
}

func Test_T6_True(t *testing.T) {
	tests.Run(t, ScanAll, tests.T6_True)
}

func Test_T7_List(t *testing.T) {
	tests.Run(t, ScanAll, tests.T7_List)
}

func Test_T8_Fix(t *testing.T) {
	tests.Run(t, ScanAll, tests.T8_Fix)
}

func Test_T10_F(t *testing.T) {
	tests.Run(t, ScanAll, tests.T10_F)
}

func Test_T11_Identifiers(t *testing.T) {
	tests.Run(t, ScanAll, tests.T11_Identifiers)
}

func Test_T12_Assign(t *testing.T) {
	tests.Run(t, ScanAll, tests.T12_Assign)
}

func Test_T13_Output(t *testing.T) {
	tests.Run(t, ScanAll, tests.T13_Output)
}

func Test_T14_LessThanOrEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.T14_LessThanOrEqual)
}

func Test_T15_MoreThanOrEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.T15_MoreThanOrEqual)
}

func Test_T16_BlockOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.T16_BlockOpen)
}

func Test_T17_BlockOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.T17_BlockClose)
}

func Test_T18_ParenOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.T18_ParenOpen)
}

func Test_T19_ParenClose(t *testing.T) {
	tests.Run(t, ScanAll, tests.T19_ParenClose)
}

func Test_T20_GuardOpen(t *testing.T) {
	tests.Run(t, ScanAll, tests.T20_GuardOpen)
}

func Test_T21_GuardClose(t *testing.T) {
	tests.Run(t, ScanAll, tests.T21_GuardClose)
}

func Test_T22_Delim(t *testing.T) {
	tests.Run(t, ScanAll, tests.T22_Delim)
}

func Test_T23_Void(t *testing.T) {
	tests.Run(t, ScanAll, tests.T23_Void)
}

func Test_T24_Terminator(t *testing.T) {
	tests.Run(t, ScanAll, tests.T24_Terminator)
}

func Test_T25_Spell(t *testing.T) {
	tests.Run(t, ScanAll, tests.T25_Spell)
}

func Test_T26_Add(t *testing.T) {
	tests.Run(t, ScanAll, tests.T26_Add)
}

func Test_T27_Subtract(t *testing.T) {
	tests.Run(t, ScanAll, tests.T27_Subtract)
}

func Test_T28_Multiply(t *testing.T) {
	tests.Run(t, ScanAll, tests.T28_Multiply)
}

func Test_T29_Divide(t *testing.T) {
	tests.Run(t, ScanAll, tests.T29_Divide)
}

func Test_T30_Remainder(t *testing.T) {
	tests.Run(t, ScanAll, tests.T30_Remainder)
}

func Test_T31_And(t *testing.T) {
	tests.Run(t, ScanAll, tests.T31_And)
}

func Test_T32_Or(t *testing.T) {
	tests.Run(t, ScanAll, tests.T32_Or)
}

func Test_T33_Equal(t *testing.T) {
	tests.Run(t, ScanAll, tests.T33_Equal)
}

func Test_T34_NotEqual(t *testing.T) {
	tests.Run(t, ScanAll, tests.T34_NotEqual)
}

func Test_T35_LessThan(t *testing.T) {
	tests.Run(t, ScanAll, tests.T35_LessThan)
}

func Test_T36_MoreThan(t *testing.T) {
	tests.Run(t, ScanAll, tests.T36_MoreThan)
}

func Test_T37_String(t *testing.T) {
	tests.Run(t, ScanAll, tests.T37_String)
}

func Test_T38_Template(t *testing.T) {
	tests.Run(t, ScanAll, tests.T38_Template)
}

func Test_T39_Number(t *testing.T) {
	tests.Run(t, ScanAll, tests.T39_Number)
}

func Test_T40_Loop(t *testing.T) {
	tests.Run(t, ScanAll, tests.T40_Loop)
}

func Test_T41_Append(t *testing.T) {
	tests.Run(t, ScanAll, tests.T41_Append)
}

func Test_T42_Prepend(t *testing.T) {
	tests.Run(t, ScanAll, tests.T42_Prepend)
}

func Test_S1_Assignment(t *testing.T) {
	tests.Run(t, ScanAll, tests.S1_Assignment)
}

func Test_S2_MultiAssignment(t *testing.T) {
	tests.Run(t, ScanAll, tests.S2_MultiAssignment)
}

func Test_S3_GuardBlock(t *testing.T) {
	tests.Run(t, ScanAll, tests.S3_GuardBlock)
}

func Test_S4_MatchBlock(t *testing.T) {
	tests.Run(t, ScanAll, tests.S4_MatchBlock)
}

func Test_S5_FuncDef(t *testing.T) {
	tests.Run(t, ScanAll, tests.S5_FuncDef)
}

func Test_S6_FuncCall(t *testing.T) {
	tests.Run(t, ScanAll, tests.S6_FuncCall)
}

func Test_S7_Expression(t *testing.T) {
	tests.Run(t, ScanAll, tests.S7_Expression)
}

func Test_S8_Block(t *testing.T) {
	tests.Run(t, ScanAll, tests.S8_Block)
}

func Test_S9_List(t *testing.T) {
	tests.Run(t, ScanAll, tests.S9_List)
}

func Test_S10_Loop(t *testing.T) {
	tests.Run(t, ScanAll, tests.S10_Loop)
}

func Test_S11_ModifyList(t *testing.T) {
	tests.Run(t, ScanAll, tests.S11_ModifyList)
}
