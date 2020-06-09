package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
)

// TODO: Is this redundant?

func Test_Void(t *testing.T) {
	var _ Void = voidExpr{}
}

func Test_Identifier(t *testing.T) {
	var _ Identifier = identifierExpr{}
}

func Test_Literal(t *testing.T) {
	var _ Literal = literalExpr{}
}

func Test_ListAccessor(t *testing.T) {
	var _ ListAccessor = listAccessorExpr{}
}

func Test_Negation(t *testing.T) {
	var _ Negation = negationExpr{}
}

func Test_Operation(t *testing.T) {
	var _ Operation = operationExpr{}
}

func Test_Assignment(t *testing.T) {
	var _ Assignment = assignmentStat{}
}

func Test_AssignmentBlock(t *testing.T) {
	var _ AssignmentBlock = assignmentBlockStat{}
}

func Test_Block(t *testing.T) {
	var _ Block = blockExpr{}
}

func Test_Parameters(t *testing.T) {
	var _ Parameters = parametersDef{}
}

func Test_Function(t *testing.T) {
	var _ Function = functionExpr{}
}

func Test_ExpressionFunction(t *testing.T) {
	var _ ExpressionFunction = expressionFunctionExpr{}
}
