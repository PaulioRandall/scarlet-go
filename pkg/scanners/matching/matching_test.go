package matching

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	. "github.com/PaulioRandall/scarlet-go/pkg/scanners/new_tests"
)

type testFunc func() (string, []Token)

func okTest(t *testing.T, f testFunc) {
	in, exps := f()
	act, e := ScanAll_(in)
	AssertResults(t, exps, act, e)
}

func Test_T1_1(t *testing.T) {
	okTest(t, T1_1)
}

func Test_T1_2(t *testing.T) {
	okTest(t, T1_2)
}

func Test_T2_1(t *testing.T) {
	okTest(t, T2_1)
}

func Test_T3_1(t *testing.T) {
	okTest(t, T3_1)
}

func Test_T4_1(t *testing.T) {
	okTest(t, T4_1)
}

func Test_T5_1(t *testing.T) {
	okTest(t, T5_1)
}

func Test_T5_2(t *testing.T) {
	okTest(t, T5_2)
}

func Test_T6_1(t *testing.T) {
	okTest(t, T6_1)
}

func Test_T7_1(t *testing.T) {
	okTest(t, T7_1)
}

func Test_T8_1(t *testing.T) {
	okTest(t, T8_1)
}

func Test_T9_1(t *testing.T) {
	okTest(t, T9_1)
}

func Test_T9_2(t *testing.T) {
	okTest(t, T9_2)
}

func Test_T9_3(t *testing.T) {
	okTest(t, T9_3)
}

func Test_T9_4(t *testing.T) {
	okTest(t, T9_4)
}

func Test_T9_5(t *testing.T) {
	okTest(t, T9_5)
}

func Test_T9_6(t *testing.T) {
	okTest(t, T9_6)
}

func Test_T9_7(t *testing.T) {
	okTest(t, T9_7)
}

func Test_T10_1(t *testing.T) {
	okTest(t, T10_1)
}

func Test_T11_1(t *testing.T) {
	okTest(t, T11_1)
}

func Test_T12_1(t *testing.T) {
	okTest(t, T12_1)
}
