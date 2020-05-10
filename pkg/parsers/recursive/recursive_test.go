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
