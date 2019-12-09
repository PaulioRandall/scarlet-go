package parser

import (
	"testing"
)

func TestRootContext_1(t *testing.T) {
	// Check it is a type of Context.
	var _ Context = RootContext{}
}
