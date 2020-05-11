package recursive

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/parsers/tests"
)

func Test_A1_Assignment(t *testing.T) {
	tests.Run(t, ParseAll, tests.A1_Assignment)
}

func Test_A2_MultiAssignment(t *testing.T) {
	tests.Run(t, ParseAll, tests.A2_MultiAssignment)
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

func Test_E12_WithFuncCall(t *testing.T) {
	tests.Run(t, ParseAll, tests.E12_WithFuncCall)
}

func Test_F1_FuncDefInline(t *testing.T) {
	tests.Run(t, ParseAll, tests.F1_FuncDefInline)
}

func Test_F2_FuncDef(t *testing.T) {
	tests.Run(t, ParseAll, tests.F2_FuncDef)
}

func Test_F3_FuncCallNoParams(t *testing.T) {
	tests.Run(t, ParseAll, tests.F3_FuncCallNoParams)
}
