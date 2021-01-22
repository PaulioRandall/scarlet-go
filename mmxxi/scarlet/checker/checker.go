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
		return checkExpr(ctx, v)
	case ast.Stmt:
		return checkStmt(ctx, v)
	default:
		return nil
	}
}

func checkVar(ctx rootCtx, n ast.Var) error {
	if n.ValType == ast.T_UNDEFINED {
		return errNode(n, "Invalid variable: undefined type")
	}
	return nil
}

func checkExpr(ctx rootCtx, n ast.Expr) error {
	switch v := n.(type) {
	case nil:
		panic("Nil expression not allowed")

	case ast.Ident:
		return checkIdent(ctx, v)

	case ast.Literal:
		return nil

	default:
		return errNode(v, "Invalid expression: unknown type")
	}
}

func checkStmt(ctx rootCtx, n ast.Stmt) error {
	switch v := n.(type) {
	case nil:
		panic("Nil statement not allowed")

	case ast.Binding:
		return checkBinding(ctx, v)

	default:
		return errNode(v, "Invalid statement: unknown type")
	}
}

func checkBinding(ctx rootCtx, n ast.Binding) error {

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
		if exp != ast.T_INFER && exp != resolveType(ctx, right[i]) {
			return badBind(right[i], "expression has wrong type, expected %s", exp)
		}
	}

	return nil
}

func checkIdent(ctx rootCtx, n ast.Ident) error {
	if n.ValType == ast.T_UNDEFINED {
		return errNode(n, "Invalid ident: Undefined variable type")
	}
	if !ctx.exists(n.Val) {
		return errNode(n, "Missing value: Undefined variable")
	}
	return nil
}

func resolveType(ctx rootCtx, n ast.TypedNode) ast.ValType {
	switch v := n.(type) {
	case ast.Ident:
		if n.ValueType() == ast.T_INFER {
			return ctx.get(v.Val)
		}
	}

	return n.ValueType()
}

func errNode(n ast.Node, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", n.Snippet().Start.Line, m)
	return errors.New(m)
}
