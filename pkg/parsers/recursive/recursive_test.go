package recursive

import (
	"testing"

	tests "github.com/PaulioRandall/scarlet-go/pkg/parsers/tests"
)

func Test_A1_Assignment(t *testing.T) {
	tests.Run(t, ParseAll, tests.A1_Assignment)
}

func Test_A2_MultiAssignment(t *testing.T) {
	tests.Run(t, ParseAll, tests.A2_MultiAssignment)
}

func Test_A3_Assignment(t *testing.T) {
	tests.Run(t, ParseAll, tests.A3_Assignment)
}

func Test_A4_MultiAssignment(t *testing.T) {
	tests.Run(t, ParseAll, tests.A4_MultiAssignment)
}

func Test_A5_Panics(t *testing.T) {
	tests.Run(t, ParseAll, tests.A5_Panics)
}

func Test_A6_Panics(t *testing.T) {
	tests.Run(t, ParseAll, tests.A6_Panics)
}

func Test_A7_Panics(t *testing.T) {
	tests.Run(t, ParseAll, tests.A7_Panics)
}

func Test_A8_ListItem(t *testing.T) {
	tests.Run(t, ParseAll, tests.A8_ListItem)
}

func Test_A9_ListItem(t *testing.T) {
	tests.Run(t, ParseAll, tests.A9_ListItem)
}

func Test_A10_ListItem(t *testing.T) {
	tests.Run(t, ParseAll, tests.A10_ListItem)
}

func Test_A11_ListItems(t *testing.T) {
	tests.Run(t, ParseAll, tests.A11_ListItems)
}

func Test_E1_Add(t *testing.T) {
	tests.Run(t, ParseAll, tests.E1_Add)
}

func Test_E2_Subtract(t *testing.T) {
	tests.Run(t, ParseAll, tests.E2_Subtract)
}

func Test_E3_Multiply(t *testing.T) {
	tests.Run(t, ParseAll, tests.E3_Multiply)
}

func Test_E4_Divide(t *testing.T) {
	tests.Run(t, ParseAll, tests.E4_Divide)
}

func Test_E5_AdditiveOrdering(t *testing.T) {
	tests.Run(t, ParseAll, tests.E5_AdditiveOrdering)
}

func Test_E6_AdditiveOrdering(t *testing.T) {
	tests.Run(t, ParseAll, tests.E6_AdditiveOrdering)
}

func Test_E7_MultiplicativeOrdering(t *testing.T) {
	tests.Run(t, ParseAll, tests.E7_MultiplicativeOrdering)
}

func Test_E8_MultiplicativeOrdering(t *testing.T) {
	tests.Run(t, ParseAll, tests.E8_MultiplicativeOrdering)
}

func Test_E9_OperationOrdering(t *testing.T) {
	tests.Run(t, ParseAll, tests.E9_OperationOrdering)
}

func Test_E10_OperationOrdering(t *testing.T) {
	tests.Run(t, ParseAll, tests.E10_OperationOrdering)
}

func Test_E11_OperationOrdering(t *testing.T) {
	tests.Run(t, ParseAll, tests.E11_OperationOrdering)
}

func Test_E12_FuncCall(t *testing.T) {
	tests.Run(t, ParseAll, tests.E12_FuncCall)
}

func Test_E13_Panics(t *testing.T) {
	tests.Run(t, ParseAll, tests.E13_Panics)
}

func Test_E14_Panics(t *testing.T) {
	tests.Run(t, ParseAll, tests.E14_Panics)
}

func Test_F1_FuncDef(t *testing.T) {
	tests.Run(t, ParseAll, tests.F1_FuncDef)
}

func Test_F2_FuncDef(t *testing.T) {
	tests.Run(t, ParseAll, tests.F2_FuncDef)
}

func Test_F3_FuncCall(t *testing.T) {
	tests.Run(t, ParseAll, tests.F3_FuncCall)
}

func Test_F4_FuncCall(t *testing.T) {
	tests.Run(t, ParseAll, tests.F4_FuncCall)
}

func Test_F5_FuncCall(t *testing.T) {
	tests.Run(t, ParseAll, tests.F5_FuncCall)
}

func Test_F6_FuncCall(t *testing.T) {
	tests.Run(t, ParseAll, tests.F6_FuncCall)
}

func Test_F7_FuncCallPanics(t *testing.T) {
	tests.Run(t, ParseAll, tests.F7_FuncCallPanics)
}

func Test_F8_FuncCallPanics(t *testing.T) {
	tests.Run(t, ParseAll, tests.F8_FuncCallPanics)
}

func Test_F9_FuncCallPanics(t *testing.T) {
	tests.Run(t, ParseAll, tests.F9_FuncCallPanics)
}

func Test_F10_FuncCallPanics(t *testing.T) {
	tests.Run(t, ParseAll, tests.F10_FuncCallPanics)
}

func Test_F11_FuncDef(t *testing.T) {
	tests.Run(t, ParseAll, tests.F11_FuncDef)
}

func Test_L1_ListDef(t *testing.T) {
	tests.Run(t, ParseAll, tests.L1_ListDef)
}

func Test_L2_ListDef(t *testing.T) {
	tests.Run(t, ParseAll, tests.L2_ListDef)
}

func Test_L3_ListAccess(t *testing.T) {
	tests.Run(t, ParseAll, tests.L3_ListAccess)
}

func Test_L4_ListAccess(t *testing.T) {
	tests.Run(t, ParseAll, tests.L4_ListAccess)
}

func Test_L5_ListAccess(t *testing.T) {
	tests.Run(t, ParseAll, tests.L5_ListAccess)
}

func Test_LP1_Conditional(t *testing.T) {
	tests.Run(t, ParseAll, tests.LP1_Conditional)
}

func Test_LP2_ForEach(t *testing.T) {
	tests.Run(t, ParseAll, tests.LP2_ForEach)
}
