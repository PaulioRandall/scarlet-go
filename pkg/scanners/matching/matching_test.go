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

func Test_T13_1(t *testing.T) {
	okTest(t, T13_1)
}

func Test_T14_1(t *testing.T) {
	okTest(t, T14_1)
}

func Test_T15_1(t *testing.T) {
	okTest(t, T15_1)
}

func Test_T16_1(t *testing.T) {
	okTest(t, T16_1)
}

func Test_T17_1(t *testing.T) {
	okTest(t, T17_1)
}

func Test_T18_1(t *testing.T) {
	okTest(t, T18_1)
}

func Test_T19_1(t *testing.T) {
	okTest(t, T19_1)
}

func Test_T20_1(t *testing.T) {
	okTest(t, T20_1)
}

func Test_T21_1(t *testing.T) {
	okTest(t, T21_1)
}

func Test_T22_1(t *testing.T) {
	okTest(t, T22_1)
}

func Test_T23_1(t *testing.T) {
	okTest(t, T23_1)
}

func Test_T24_1(t *testing.T) {
	okTest(t, T24_1)
}

func Test_T25_1(t *testing.T) {
	okTest(t, T25_1)
}

func Test_T26_1(t *testing.T) {
	okTest(t, T26_1)
}

func Test_T27_1(t *testing.T) {
	okTest(t, T27_1)
}

func Test_T28_1(t *testing.T) {
	okTest(t, T28_1)
}

func Test_T29_1(t *testing.T) {
	okTest(t, T29_1)
}

func Test_T30_1(t *testing.T) {
	okTest(t, T30_1)
}

func Test_T31_1(t *testing.T) {
	okTest(t, T31_1)
}

func Test_T32_1(t *testing.T) {
	okTest(t, T32_1)
}

func Test_T33_1(t *testing.T) {
	okTest(t, T33_1)
}

func Test_T34_1(t *testing.T) {
	okTest(t, T34_1)
}

func Test_T35_1(t *testing.T) {
	okTest(t, T35_1)
}

func Test_T35_2(t *testing.T) {
	okTest(t, T35_2)
}

func Test_T36_1(t *testing.T) {
	okTest(t, T36_1)
}

func Test_T36_2(t *testing.T) {
	okTest(t, T36_2)
}

func Test_T36_3(t *testing.T) {
	okTest(t, T36_3)
}

func Test_T36_4(t *testing.T) {
	okTest(t, T36_4)
}

func Test_T37_1(t *testing.T) {
	okTest(t, T37_1)
}

func Test_T38_1(t *testing.T) {
	okTest(t, T38_1)
}

func Test_T39_1(t *testing.T) {
	okTest(t, T39_1)
}

func Test_T40_1(t *testing.T) {
	okTest(t, T40_1)
}

func Test_T41_1(t *testing.T) {
	okTest(t, T41_1)
}

func Test_S1(t *testing.T) {
	okTest(t, S1)
}
