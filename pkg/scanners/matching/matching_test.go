package matching

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/scanners/tests"
)

func TestScanner(t *testing.T) {
	tests.DoTests(t, ScanAll)
}
