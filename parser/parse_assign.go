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
func (ex assignStat) String() string {
	return "[" + ex.id.String() + "] " +
		"[" + ex.tk.String() + "] " +
		"[" + ex.src.String() + "]"
}

// parseAssign parses an assignment into a statement. Assumes that the next
// statement in the input channel is an assignment.
func (p *Parser) parseAssign(dst token.Token) Stat {

	p.checkToken(dst, token.ID)

	ass := <-p.in
	p.checkToken(ass, token.ASSIGN)

	src := <-p.in
	p.checkToken(src, token.STR_LITERAL)

	srcEx := valueExpr{
		tokenExpr: tokenExpr{src},
		v:         Value{STR, src.Value},
	}

	return assignStat{
		tokenExpr: tokenExpr{ass},
		id:        dst,
		src:       srcEx,
	}
}
