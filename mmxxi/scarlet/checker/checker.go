package checker

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
)

// func validateRoutine(ctx Context, trees []ast.Tree) error {
// 	TODO: For each routine, including main scroll:
// 				1. identify all defined identifiers
// 				2. check everything
// }

func validateNode(n ast.Node) error {
	switch v := n.(type) {
	case ast.Stmt:
		return validateStmt(v)
	default:
		return nil
	}
}

func validateStmt(stmt ast.Stmt) error {
	switch v := stmt.(type) {
	case nil:
		panic("Nil statement not allowed")

	case ast.Binding:
		return validateBinding(v)

	default:
		return errNode(v, "Unknown statement type")
	}
}

func validateBinding(b ast.Binding) error {

	left, right := b.Base().Left, b.Base().Right

	if left == nil {
		return errNode(b, "Invalid binding: left side is nil")
	}

	if right == nil {
		return errNode(b, "Invalid binding: right side is nil")
	}

	leftLen, rightLen := len(left), len(right)

	if leftLen == 0 {
		return errNode(b, "Invalid binding: left side is empty")
	}

	if leftLen > rightLen {
		return errNode(b,
			"Invalid binding: too many items on left or too few on right")
	}

	if leftLen < rightLen {
		return errNode(b,
			"Invalid binding: too few items on left or too many on right")
	}

	return nil
}

func validateExpr(expr ast.Expr) error {
	switch v := expr.(type) {
	case nil:
		panic("Nil expression not allowed")

	case ast.Ident, ast.Literal:
		return nil

	default:
		return errNode(v, "Unknown expression type")
	}
}

func errNode(n ast.Node, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", n.Snippet().Start.Line, m)
	return errors.New(m)
}
