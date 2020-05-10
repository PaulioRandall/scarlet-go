package recursive

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/parsers/tests"
)

func Test_A1(t *testing.T) {
	tests.Run(t, ParseAll, tests.A1)
}
