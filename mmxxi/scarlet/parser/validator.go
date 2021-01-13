package parser

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	//"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

func err(itr LexIterator, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", itr.Line(), m)
	return errors.New(m)
}

func errNode(n ast.Node, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", n.Snippet().Start.Line, m)
	return errors.New(m)
}

func validateStmt(stmt ast.Stmt) error {
	switch v := stmt.(type) {
	case nil:
		panic("Nil statement no allowed")
	case ast.Binding:
		return validateBinding(v)
	default:
		return errNode(stmt, "Unknown statement type")
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
		return errNode(b, "Invalid binding: too many items on left or too few on right")
	}

	if leftLen < rightLen {
		return errNode(b,
			"Invalid binding: too many items on right or too few on left")
	}

	return nil
}
