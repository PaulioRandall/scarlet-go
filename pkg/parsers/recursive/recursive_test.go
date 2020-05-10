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

func Test_F1_FuncInline(t *testing.T) {
	tests.Run(t, ParseAll, tests.F1_FuncInline)
}

func Test_F2_Func(t *testing.T) {
	tests.Run(t, ParseAll, tests.F2_Func)
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
