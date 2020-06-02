package statement

import (
	"testing"
)

func Test_Void(t *testing.T) {
	var _ Expression = Void{}
}

func Test_Identifier(t *testing.T) {
	var _ Expression = Identifier{}
}

func Test_Literal(t *testing.T) {
	var _ Expression = Literal{}
}

func Test_List(t *testing.T) {
	var _ Statement = List{}
}

func Test_ListAccessor(t *testing.T) {
	var _ Statement = ListAccessor{}
}

func Test_Negation(t *testing.T) {
	var _ Expression = Negation{}
}

func Test_Assignment(t *testing.T) {
	var _ Statement = Assignment{}
}

func Test_Block(t *testing.T) {
	var _ Statement = Block{}
}

func Test_Parameters(t *testing.T) {
	var _ Expression = Parameters{}
}

func Test_Function(t *testing.T) {
	var _ Expression = Function{}
}
