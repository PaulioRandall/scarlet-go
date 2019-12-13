package context

import (
	"testing"
)

func TestRootCtx_1(t *testing.T) {
	// Check it is a type of Context.
	var _ Context = rootCtx{}
}
