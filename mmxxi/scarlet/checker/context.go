package checker

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
)

type (
	bindings map[string]ast.ValType

	// # new rootCtx, pushMajorCtx, pushMinorCtx
	// x <- 1 + 2
	//
	// # pushMajorCtx, pushMinorCtx, add a & b to minorCtx
	// div, err <- F(a N, b N -> c N, e S) {
	//    [b == 0] { # pushMinorCtx
	//      e <- "Can't divide by zero"
	//      <~
	//    } # popMinorCtx
	//    c <- a / b
	// } # set div=c, err=e, popMajorCtx

	// rootCtx (level 0) encapsulates a whole scroll.
	rootCtx struct {
		major *majorCtx
		defs  bindings
	}

	// majorCtx (level 1) represents constructs with a body of statements but with
	// isolated variables. No majorCtx can access the variables of another, unless
	// they are explicitly passed. I.e. the root of each scroll and function
	// bodies.
	majorCtx struct {
		next  *majorCtx
		minor *minorCtx
	}

	// minorCtx (level 2) represents constructs with a body of statements that
	// have access to all variables in other minorCtxs it is nested within but the
	// parent doesn't have access to theirs, i.e. conditionals and loops.
	minorCtx struct {
		next *minorCtx
		vars bindings
	}
)

const not_found = ast.T_UNDEFINED

func makeRootCtx() {
	c := rootCtx{}
	c.pushMajorCtx()
}

func (r rootCtx) pushMajorCtx() {
	r.major = &majorCtx{next: r.major}
	r.pushMinorCtx()
}

func (r rootCtx) popMajorCtx() {
	r.major = r.major.next
}

func (r rootCtx) pushMinorCtx() {
	r.major.minor = &minorCtx{
		next: r.major.minor,
		vars: bindings{},
	}
}

func (r rootCtx) popMinorCtx() {
	r.major.minor = r.major.minor.next
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
	for s := r.major.minor; s != nil; s = s.next {
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
