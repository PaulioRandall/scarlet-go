package statement

import (
	"testing"
)

func Test_Identifier(t *testing.T) {
	var _ Expression = Identifier{}
}

func Test_Literal(t *testing.T) {
	var _ Expression = Literal{}
}
