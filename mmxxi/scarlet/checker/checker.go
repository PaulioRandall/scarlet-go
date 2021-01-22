package checker

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
)

func validateRoutine(ctx rootCtx, trees []ast.Tree) error {

	for _, t := range trees {
		if e := checkNode(ctx, t.Root); e != nil {
			return e
		}
	}

	return nil
	// 	TODO: For each routine, including main scroll:
	// 				1. identify all defined identifiers
	// 				2. check everything
}

func checkNode(ctx rootCtx, n ast.Node) error {
	switch v := n.(type) {
	case ast.Expr:
		return checkExpr(v)
	case ast.Stmt:
		return checkStmt(v)
	default:
		return nil
	}
}

func checkVar(n ast.Var) error {
	if n.ValType == ast.T_UNDEFINED {
		return errNode(n, "Invalid variable: undefined type")
	}
	return nil
}

func checkExpr(n ast.Expr) error {
	switch v := n.(type) {
	case nil:
		panic("Nil expression not allowed")

	case ast.Ident:
		return checkIdent(v)

	case ast.Literal:
		return nil

	default:
		return errNode(v, "Invalid expression: unknown type")
	}
}

func checkStmt(n ast.Stmt) error {
	switch v := n.(type) {
	case nil:
		panic("Nil statement not allowed")

	case ast.Binding:
		return checkBinding(v)

	default:
		return errNode(v, "Invalid statement: unknown type")
	}
}

func checkBinding(n ast.Binding) error {

	badBind := func(n ast.Node, m string, args ...interface{}) error {
		return errNode(n, "Invalid binding: "+m, args...)
	}

	left, right := n.Base().Left, n.Base().Right

	if left == nil {
		return badBind(n, "left side is nil")
	}

	if right == nil {
		return badBind(n, "right side is nil")
	}

	leftLen, rightLen := len(left), len(right)

	if leftLen == 0 {
		return badBind(n, "left side is empty")
	}

	if leftLen > rightLen {
		return badBind(n, "too many items on left or too few on right")
	}

	if leftLen < rightLen {
		return badBind(n, "too few items on left or too many on right")
	}

	for i, _ := range left {
		exp := left[i].ValueType()
		if exp != ast.T_INFER && exp != right[i].ValueType() {
			return badBind(right[i], "expression has wrong type, expected %s", exp)
		}
	}

	return nil
}

func checkIdent(n ast.Ident) error {
	if n.ValType == ast.T_UNDEFINED {
		return errNode(n, "Invalid ident: Undefined variable type")
	}
	return nil
}

func errNode(n ast.Node, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", n.Snippet().Start.Line, m)
	return errors.New(m)
}
