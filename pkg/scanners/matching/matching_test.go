package matching

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/scanners/new_tests"
)

func Test_T1_1(t *testing.T) {
	in, exps := T1_1()
	act, e := ScanAll_(in)
	AssertResults(t, exps, act, e)
}
