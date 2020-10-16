package compiler

import (
	"github.com/PaulioRandall/scarlet-go/token2/code"
	"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/tree"
	"github.com/PaulioRandall/scarlet-go/token2/value"
)

func Compile(n tree.Node) []inst.Inst {
	switch v := n.(type) {
	case tree.SingleAssign:
		return singleAssign(v)
	case tree.MultiAssign:
		return multiAssign(v)

	default:
		panic("[ERROR] Unknown node type")
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
