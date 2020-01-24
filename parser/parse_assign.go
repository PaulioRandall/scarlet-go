package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// assignStat represents a statement that creates or updates variables by
// evaluating expressions into values which are mapped to their identifier
// within a context.
type assignStat struct {
	tokenExpr
	ids  []token.Token
	srcs []Expr
}

// Eval satisfies the Expr interface.
func (ex assignStat) Eval(ctx Context) (_ Value) {

	size := len(ex.ids)
	for i := 0; i < size; i++ {
		v := ex.srcs[i].Eval(ctx)
		ctx.set(ex.ids[i].Value, v)
	}

	return
}

// String satisfies the Expr interface.
func (ex assignStat) String() (s string) {

	size := len(ex.ids)
	for i := 0; i < size; i++ {

		if i != 0 {
			s += "\n"
		}

		s += "Assign "
		s += "[" + ex.ids[i].String() + "] "
		s += "[" + ex.tk.String() + "] "
		s += "[" + ex.srcs[i].String() + "]"
	}

	return
}

// parseAssign parses an assignment into a statement. Assumes that the next
// statement in the input channel is an assignment.
func (p *Parser) parseAssign(first token.Token) Stat {

	ids, ass := p.parseAssignIDs(first)
	srcs := p.parseAssignSources()

	if len(ids) != len(srcs) {
		panic(ass.String() + ": Assignment requires the ID and expression count match")
	}

	return assignStat{
		tokenExpr: tokenExpr{ass},
		ids:       ids,
		srcs:      srcs,
	}
}

// parseAssignIDs parses a delimitered list of ID tokens used for an assignment.
func (p *Parser) parseAssignIDs(first token.Token) (ids []token.Token, tk token.Token) {

	ids = []token.Token{first}

	for {
		tk = p.takeEnsure(token.DELIM, token.ASSIGN)

		if tk.Kind == token.ASSIGN {
			break
		}

		tk = p.takeEnsure(token.ID)
		ids = append(ids, tk)
	}

	return
}

// parseAssignSources parses the sources of an assignment.
func (p *Parser) parseAssignSources() (srcs []Expr) {

	var next token.Token

	for next.Kind != token.TERMINATOR {

		next = p.take()
		ex := p.parseExpr(next)
		srcs = append(srcs, ex)

		next = p.takeEnsure(token.DELIM, token.TERMINATOR)
	}

	return
}
