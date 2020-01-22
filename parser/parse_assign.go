package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// assignStat represents a statement that creates or updates variables by
// evaluating expressions into values which are mapped to their identifier
// within a context.
type assignStat struct {
	tokenExpr
	id  token.Token
	src Expr
}

// Eval satisfies the Expr interface.
func (ex assignStat) Eval(ctx Context) (_ Value) {
	v := ex.src.Eval(ctx)
	ctx.set(ex.id.Value, v)
	return
}

// String satisfies the Expr interface.
func (ex assignStat) String() (s string) {
	s += "Assign "
	s += "[" + ex.id.String() + "] "
	s += "[" + ex.tk.String() + "] "
	s += "[" + ex.src.String() + "]"
	return
}

// parseAssign parses an assignment into a statement. Assumes that the next
// statement in the input channel is an assignment.
func (p *Parser) parseAssign(dst token.Token) Stat {

	p.ensure(dst, token.ID)
	ass := p.takeEnsure(token.ASSIGN)
	src := p.takeEnsure(token.STR_LITERAL, token.BOOL_LITERAL)
	p.takeEnsure(token.TERMINATOR)

	srcEx := valueExpr{
		tokenExpr: tokenExpr{src},
		v:         NewValue(src),
	}

	return assignStat{
		tokenExpr: tokenExpr{ass},
		id:        dst,
		src:       srcEx,
	}
}
