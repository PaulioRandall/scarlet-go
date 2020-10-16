package compiler

import (
	"github.com/PaulioRandall/scarlet-go/token2/code"
	"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/tree"
	"github.com/PaulioRandall/scarlet-go/token2/value"
)

// TODO: Need to test singleAssign!

func Compile(n tree.Node) []inst.Inst {

	var ins []inst.Inst

	switch v := n.(type) {
	case tree.SingleAssign:
		ins = singleAssign(v)

	default:
		panic("[ERROR] Unknown node type")
	}

	return ins
}

func singleAssign(n tree.SingleAssign) []inst.Inst {
	ins := expression(n.Right)
	return append(ins, inst.Inst{
		Code: code.SCOPE_BIND,
		Data: createAssignData(n.Left),
	})
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
		return []inst.Inst{
			inst.Inst{
				Code: code.STACK_PUSH,
				Data: createLitData(v),
			}}

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
