package compiler

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

// CompileAll converts the each AST into a slice of statements and
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

// Compile converts the AST 't' into a slice of instructions.
func Compile(t tree.Node) ([]inst.Inst, error) {
	switch v := t.(type) {
	case tree.SingleAssign:
		return singleAssign(v), nil

	case tree.MultiAssign:
		return multiAssign(v), nil

	case tree.SpellCall:
		return spellCall(v), nil

	case tree.Ident, tree.Literal, tree.BinaryExpr:
		return nil, errSnip(t.Pos(), "Result of expression ignored")

	default:
		return nil, errSnip(t.Pos(), "Unknown node type")
	}
}

func singleAssign(n tree.SingleAssign) []inst.Inst {
	ins := expression(n.Right)
	return append(ins, inst.Inst{
		Code: inst.SCOPE_BIND,
		Data: createAssignData(n.Left),
	})
}

func multiAssign(n tree.MultiAssign) (ins []inst.Inst) {
	for _, v := range n.Right {
		ins = append(ins, expression(v)...)
	}
	// Reverse because the last expr result will be at the top of the stack
	for i := len(n.Left) - 1; i >= 0; i-- {
		ins = append(ins, inst.Inst{
			Code: inst.SCOPE_BIND,
			Data: createAssignData(n.Left[i]),
		})
	}
	return
}

func spellCall(n tree.SpellCall) (ins []inst.Inst) {
	ins = append(ins, inst.Inst{
		Code: inst.STACK_PUSH,
	})
	for i := len(n.Args) - 1; i >= 0; i-- {
		ins = append(ins, expression(n.Args[i])...)
	}
	return append(ins, inst.Inst{
		Code: inst.SPELL_CALL,
		Data: value.Ident(n.Name),
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
		return []inst.Inst{inst.Inst{
			Code: inst.STACK_PUSH,
			Data: createLitData(v),
		}}

	case tree.Ident:
		return []inst.Inst{inst.Inst{
			Code: inst.FETCH_PUSH,
			Data: value.Ident(v.Val),
		}}

	case tree.BinaryExpr:
		return binaryExpression(v)

	case tree.SpellCall:
		return spellCall(v)

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
		return value.Str(v.Val[1 : len(v.Val)-1])
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

func findOpCode(tk token.Token) inst.Code {
	switch tk {
	case token.ADD:
		return inst.BIN_OP_ADD
	case token.SUB:
		return inst.BIN_OP_SUB
	case token.MUL:
		return inst.BIN_OP_MUL
	case token.DIV:
		return inst.BIN_OP_DIV
	case token.REM:
		return inst.BIN_OP_REM
	case token.LESS:
		return inst.BIN_OP_LESS
	case token.MORE:
		return inst.BIN_OP_MORE
	case token.LESS_EQUAL:
		return inst.BIN_OP_LEQU
	case token.MORE_EQUAL:
		return inst.BIN_OP_MEQU
	case token.EQUAL:
		return inst.BIN_OP_EQU
	case token.NOT_EQUAL:
		return inst.BIN_OP_NEQU
	case token.AND:
		return inst.BIN_OP_AND
	case token.OR:
		return inst.BIN_OP_OR
	default:
		panic("[ERROR] Unknown operator token")
	}
}
