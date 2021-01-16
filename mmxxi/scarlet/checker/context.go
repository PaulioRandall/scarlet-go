package checker

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
)

type context struct {
	user    map[string]ast.ValType
	globals map[string]ast.ValType
	defs    map[string]ast.ValType
	vars    map[string]ast.ValType
}

func (ctx context) get(id string) ast.ValType {
	if t, ok := ctx.globals[id]; ok {
		return t
	}
	if t, ok := ctx.defs[id]; ok {
		return t
	}
	if t, ok := ctx.vars[id]; ok {
		return t
	}
	return ast.T_UNDEFINED
}
