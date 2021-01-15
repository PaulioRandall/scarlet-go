package checker

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
)

type context struct {
	defs []ast.Ident
	vars []ast.Ident
}

// TODO: For each routine, including main scroll:
// TODO: first pass identifies all defined identifiers
// TODO: second pass checks everything
