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
	return ex.TabString(0)
}

// TabString satisfies the Expr interface.
func (ex assignStat) TabString(tabs int) (s string) {

	size := len(ex.ids)
	pre := strings.Repeat("\t", tabs)

	for i := 0; i < size; i++ {

		if i != 0 {
			s += "\n"
		}

		s += pre + "Assign "
		s += "[" + ex.ids[i].String() + "] "
		s += "[" + ex.tk.String() + "] "
		s += "[" + ex.srcs[i].String() + "]"
	}

	return
}

// parseAssign parses an assignment into a statement. Assumes that the next
// statement in the input channel is an assignment.
func (p *Parser) parseAssign() Stat {

	ids := p.parseAssignIDs()
	ass := p.takeEnsure(token.ASSIGN)
	srcs := p.parseAssignSources()
	p.takeEnsure(token.TERMINATOR)

	print(len(ids))
	print(":")
	println(len(srcs))

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

// parseAssignSources parses the sources of an assignment.
func (p *Parser) parseAssignSources() (srcs []Expr) {
	for {

		ex := p.parseExpr()
		srcs = append(srcs, ex)

		if p.peek().Kind == token.DELIM {
			p.take()
			continue
		}

		return
	}
}
