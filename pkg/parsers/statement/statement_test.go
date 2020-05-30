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

func Test_Assignment(t *testing.T) {
	var _ Expression = Assignment{}
}

func Test_AssignmentBlock(t *testing.T) {
	var _ Expression = AssignmentBlock{}
}
