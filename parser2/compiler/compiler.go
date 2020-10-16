package compiler

import (
	"github.com/PaulioRandall/scarlet-go/token2/code"
	"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/token"
	"github.com/PaulioRandall/scarlet-go/token2/tree"
	"github.com/PaulioRandall/scarlet-go/token2/value"
)

func Compile(n tree.Node) ([]inst.Inst, error) {
	switch v := n.(type) {
	case tree.SingleAssign:
		return singleAssign(v), nil
	case tree.MultiAssign:
		return multiAssign(v), nil
	case tree.Literal, tree.BinaryExpr:
		return nil, errSnip(n.Pos(), "Result of expression ignored")
	default:
		return nil, errSnip(n.Pos(), "Unknown node type")
	}
}

func singleAssign(n tree.SingleAssign) []inst.Inst {
	ins := expression(n.Right)
	return append(ins, inst.Inst{
		Code: code.SCOPE_BIND,
		Data: createAssignData(n.Left),
	})
}

func multiAssign(n tree.MultiAssign) (ins []inst.Inst) {
	for i, v := range n.Right {
		ins = append(ins, expression(v)...)
		ins = append(ins, inst.Inst{
			Code: code.SCOPE_BIND,
			Data: createAssignData(n.Left[i]),
		})
	}
	return
}

func createAssignData(n tree.Assignee) value.Value {
	switch v := n.(type) {
	case tree.Ident:
		return value.Ident(v.Val)
	default:
		panic("[ERROR] Unknown assignee type")
	}
}

func expression(n tree.Expr) []inst.Inst {
	switch v := n.(type) {
	case tree.Literal:
		return []inst.Inst{inst.Inst{
			Code: code.STACK_PUSH,
			Data: createLitData(v),
		}}
	case tree.BinaryExpr:
		return binaryExpression(v)
	default:
		panic("[ERROR] Unknown expression type")
	}
}

func createLitData(n tree.Literal) value.Value {
	switch v := n.(type) {
	case tree.BoolLit:
		return value.Bool(v.Val)
	case tree.NumLit:
		return value.Num{v.Val}
	case tree.StrLit:
		return value.Str(v.Val)
	default:
		panic("[ERROR] Unknown literal type")
	}
}

func binaryExpression(n tree.BinaryExpr) []inst.Inst {
	l := expression(n.Left)
	r := expression(n.Right)
	ins := append(l, r...) // left associative
	return append(ins, inst.Inst{
		Code: findOpCode(n.Op),
	})
}

func findOpCode(tk token.Token) code.Code {
	switch tk {
	case token.ADD:
		return code.OP_ADD
	case token.SUB:
		return code.OP_SUB
	case token.MUL:
		return code.OP_MUL
	case token.DIV:
		return code.OP_DIV
	case token.REM:
		return code.OP_REM
	case token.LESS:
		return code.OP_LESS
	case token.MORE:
		return code.OP_MORE
	case token.LESS_EQUAL:
		return code.OP_LEQU
	case token.MORE_EQUAL:
		return code.OP_MEQU
	case token.EQUAL:
		return code.OP_EQU
	case token.NOT_EQUAL:
		return code.OP_NEQU
	case token.AND:
		return code.OP_AND
	case token.OR:
		return code.OP_OR
	default:
		panic("[ERROR] Unknown operator token")
	}
}
