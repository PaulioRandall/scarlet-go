package z_alpha

import (
	"testing"

	e "github.com/PaulioRandall/scarlet-go/pkg/err"
)

func TestScanErr(t *testing.T) {
	var _ e.Err = runtimeErr{}
}
