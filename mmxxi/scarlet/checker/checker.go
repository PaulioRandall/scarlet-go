package checker

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
)

func validateScroll(trees []ast.Tree) error {

	ctx := makeRootCtx()

	// Process definitions first
	for _, t := range trees {
		if def, ok := t.Root.(ast.Define); ok {
			if e := checkDefine(ctx, def); e != nil {
				return e
			}
		}
	}

	// Check everything
	for _, t := range trees {
		if _, ok := t.Root.(ast.Define); !ok {
			if e := checkNode(ctx, t.Root); e != nil {
				return e
			}
		}
	}

	return nil
}

func checkNode(ctx rootCtx, n ast.Node) error {
	switch v := n.(type) {
	case ast.Expr:
		return checkExpr(ctx, v)
	case ast.Stmt:
		return checkStmt(ctx, v)
	default:
		return errNode(v, "Invalid node: this shouldn't be here")
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

	case ast.Assign:
		return checkAssign(ctx, v)

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

	return nil
}

func checkDefine(ctx rootCtx, n ast.Define) error {
	if e := checkBinding(ctx, n); e != nil {
		return e
	}

	for i, v := range n.Base().Left {

		right := n.Base().Right[i]
		t := resolveType(ctx, right)

		if ctx.defExists(v.Val) {
			return errNode(v, "Invalid definition: already defined '"+v.Val+"'")
		}

		expT := v.ValueType()
		if expT != ast.T_INFER && expT != t {
			return errNode(right,
				"Invalid definition: expression has wrong type, expected %s", expT)
		}

		ctx.setDef(v.Val, t)
	}

	return nil
}

func checkAssign(ctx rootCtx, n ast.Assign) error {
	if e := checkBinding(ctx, n); e != nil {
		return e
	}

	vars := bindings{}
	for i, v := range n.Base().Left {

		if _, ok := vars[v.Val]; ok { // e.g. x, x <- 1, 2
			return errNode(v,
				"Invalid assignment: can't multi-assign to the same variable")
		}

		right := n.Base().Right[i]
		t := resolveType(ctx, right)
		expT := v.ValueType()

		if expT != ast.T_INFER && expT != t { // e.g. x B <- 1
			return errNode(right,
				"Invalid assignment: expression has wrong type, expected %s", expT)
		}

		vars[v.Val] = t
		ctx.setVar(v.Val, t)
	}

	return nil
}

func checkIdent(ctx rootCtx, n ast.Ident) error {
	if n.ValType == ast.T_UNDEFINED {
		return errNode(n, "Invalid ident: undefined variable type")
	}
	if !ctx.exists(n.Val) {
		return errNode(n, "Missing value: undefined variable")
	}
	return nil
}

func resolveType(ctx rootCtx, n ast.TypedNode) ast.ValType {
	if v, ok := n.(ast.Ident); ok && v.ValueType() == ast.T_RESOLVE {
		return ctx.get(v.Val)
	}
	return n.ValueType()
}

func errNode(n ast.Node, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", n.Snippet().Start.Line, m)
	return errors.New(m)
}
