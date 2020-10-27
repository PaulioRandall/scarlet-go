package compiler

import (
	"github.com/PaulioRandall/scarlet-go/token/code"
	"github.com/PaulioRandall/scarlet-go/token/inst"
	"github.com/PaulioRandall/scarlet-go/token/token"
	"github.com/PaulioRandall/scarlet-go/token/tree"
	"github.com/PaulioRandall/scarlet-go/token/value"
)

// CompileAll converts the each parse tree into a slice of statements and
// aggregates them all into a slice of slices.
func CompileAll(trees []tree.Node) (r [][]inst.Inst, e error) {
	r = make([][]inst.Inst, len(trees))
	for i, t := range trees {
		if r[i], e = Compile(t); e != nil {
			return nil, e
		}
	}
	return r, nil
}

// Compile converts the parse tree 't' into a slice of instructions.
func Compile(t tree.Node) ([]inst.Inst, error) {
	switch v := t.(type) {
	case tree.SingleAssign:
		return singleAssign(v), nil
	case tree.MultiAssign:
		return multiAssign(v), nil
	case tree.Literal, tree.BinaryExpr:
		return nil, errSnip(t.Pos(), "Result of expression ignored")
	default:
		return nil, errSnip(t.Pos(), "Unknown node type")
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
		return value.Num{Number: v.Val}
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
		return code.BIN_OP_ADD
	case token.SUB:
		return code.BIN_OP_SUB
	case token.MUL:
		return code.BIN_OP_MUL
	case token.DIV:
		return code.BIN_OP_DIV
	case token.REM:
		return code.BIN_OP_REM
	case token.LESS:
		return code.BIN_OP_LESS
	case token.MORE:
		return code.BIN_OP_MORE
	case token.LESS_EQUAL:
		return code.BIN_OP_LEQU
	case token.MORE_EQUAL:
		return code.BIN_OP_MEQU
	case token.EQUAL:
		return code.BIN_OP_EQU
	case token.NOT_EQUAL:
		return code.BIN_OP_NEQU
	case token.AND:
		return code.BIN_OP_AND
	case token.OR:
		return code.BIN_OP_OR
	default:
		panic("[ERROR] Unknown operator token")
	}
}
