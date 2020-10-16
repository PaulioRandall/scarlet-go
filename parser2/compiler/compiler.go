package compiler

import (
	"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/tree"
	"github.com/PaulioRandall/scarlet-go/token2/value"
)

// TODO: Need to test singleAssign!

func Compile(n tree.Node, ds *inst.DataSet) []inst.RiscInst {

	var in []inst.RiscInst

	switch v := n.(type) {
	case tree.SingleAssign:
		in = singleAssign(v, ds)

	default:
		panic("[ERROR] Unknown node type")
	}

	return in
}

func singleAssign(n tree.SingleAssign, ds *inst.DataSet) []inst.RiscInst {
	ins := expression(n.Right, ds)
	return append(ins, inst.RiscInst{
		Inst: inst.SCP_BIND,
		Data: createAssignData(n.Left, ds),
	})
}

func createAssignData(n tree.Assignee, ds *inst.DataSet) inst.DataRef {
	switch v := n.(type) {
	case tree.Ident:
		return ds.Insert(value.Ident(v.Val))

	default:
		panic("[ERROR] Unknown assignee type")
	}
}

func expression(n tree.Expr, ds *inst.DataSet) []inst.RiscInst {
	switch v := n.(type) {
	case tree.Literal:
		return []inst.RiscInst{
			inst.RiscInst{
				Inst: inst.STK_PUSH,
				Data: createLitData(v, ds),
			}}

	default:
		panic("[ERROR] Unknown expression type")
	}
}

func createLitData(n tree.Literal, ds *inst.DataSet) inst.DataRef {
	switch v := n.(type) {
	case tree.BoolLit:
		return ds.Insert(value.Bool(v.Val))
	case tree.NumLit:
		return ds.Insert(value.Num{v.Val})
	case tree.StrLit:
		return ds.Insert(value.Str(v.Val))

	default:
		panic("[ERROR] Unknown literal type")
	}
}
