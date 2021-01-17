package checker

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
)

type (
	bindings map[string]ast.ValType

	rootCtx struct {
		stack *baseCtx
		defs  bindings
		// usrTypes    bindings
	}

	baseCtx struct {
		next  *baseCtx
		stack *subCtx
	}

	subCtx struct {
		next *subCtx
		vars bindings
	}
)

const not_found = ast.T_UNDEFINED

func (r rootCtx) newBase() {
	r.stack = &baseCtx{next: r.stack}
	r.newSub()
}

func (r rootCtx) newSub() {
	r.stack.stack = &subCtx{
		next: r.stack.stack,
		vars: bindings{},
	}
}

func (r rootCtx) defExists(id string) bool {
	return r.getDef(id) != ast.T_UNDEFINED
}

func (r rootCtx) getDef(id string) ast.ValType {
	if t, ok := r.defs[id]; ok {
		return t
	}
	return not_found
}

func (r rootCtx) varExists(id string) bool {
	return r.getVar(id) != ast.T_UNDEFINED
}

func (r rootCtx) getVar(id string) ast.ValType {
	for s := r.stack.stack; s != nil; s = s.next {
		if t, ok := s.vars[id]; ok {
			return t
		}
	}
	return not_found
}

func (r rootCtx) exists(id string) bool {
	return r.get(id) != ast.T_UNDEFINED
}

func (r rootCtx) get(id string) ast.ValType {
	if t := r.getVar(id); t != not_found {
		return t
	}
	return r.getDef(id)
}
