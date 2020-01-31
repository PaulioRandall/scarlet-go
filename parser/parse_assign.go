package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

// assignStat represents a statement that creates or updates variables by
// evaluating expressions into values which are mapped to their identifier
// within a context.
type assignStat struct {
	tokenExpr
	sticky token.Token
	ids    []token.Token
	srcs   []Expr
}

// Eval satisfies the Expr interface.
func (ex assignStat) Eval(ctx Context) (_ Value) {

	var (
		isSticky bool
		size     int = len(ex.ids)
	)

	if ex.sticky != (token.Token{}) {
		isSticky = true
	}

	for i := 0; i < size; i++ {
		v := ex.srcs[i].Eval(ctx)

		if isSticky {
			ctx.setSticky(ex.ids[i].Value, v)
		} else {
			ctx.set(ex.ids[i].Value, v)
		}
	}

	return
}

// String satisfies the Expr interface.
func (ex assignStat) String() (s string) {
	return ex.TabString(0)
}

// TabString satisfies the Expr interface.
func (ex assignStat) TabString(tabs int) (s string) {

	var (
		isSticky bool   = ex.sticky != (token.Token{})
		size     int    = len(ex.ids)
		pre      string = strings.Repeat("\t", tabs)
	)

	for i := 0; i < size; i++ {

		if i != 0 {
			s += "\n"
		}

		s += pre + "Assign "

		if isSticky {
			s += "[" + ex.sticky.String() + "] "
		}

		s += "[" + ex.ids[i].String() + "] "
		s += "[" + ex.tk.String() + "] "
		s += "[" + ex.srcs[i].String() + "]"
	}

	return
}

// parseAssign parses an assignment into a statement. Assumes that the next
// statement in the input channel is an assignment.
func (p *Parser) parseAssign() Stat {

	var sticky token.Token
	if p.peek().Kind == token.STICKY {
		sticky = p.take()
	}

	ids := p.parseAssignIDs()
	ass := p.takeEnsure(token.ASSIGN)
	srcs := p.parseDelimExpr()
	p.takeEnsure(token.TERMINATOR)

	if len(ids) != len(srcs) {
		panic(ass.String() + ": Assignment requires the ID and expression count match")
	}

	return assignStat{
		tokenExpr: tokenExpr{ass},
		sticky:    sticky,
		ids:       ids,
		srcs:      srcs,
	}
}

// parseAssignIDs parses a delimitered list of ID tokens used for an assignment.
func (p *Parser) parseAssignIDs() (ids []token.Token) {
	for {

		tk := p.takeEnsure(token.ID)
		ids = append(ids, tk)

		if p.peek().Kind == token.DELIM {
			p.take()
			continue
		}

		return
	}
}
