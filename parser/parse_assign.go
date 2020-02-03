package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

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
		panic(bard.NewHorror(ass, nil,
			"Assignment requires the ID and expression count match",
		))
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

// assignStat represents a statement that creates or updates variables by
// evaluating expressions into values which are mapped to their identifier
// within a context.
type assignStat struct {
	tokenExpr
	sticky token.Token
	ids    []token.Token
	srcs   []Expr
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

// Eval satisfies the Expr interface.
func (ex assignStat) Eval(ctx Context) (_ Value) {

	var (
		size   int = len(ex.ids)
		setter func(string, Value)
		values = make([]Value, size)
	)

	if ex.sticky == (token.Token{}) {
		setter = ctx.set
	} else {
		setter = ctx.setSticky
	}

	for i := 0; i < size; i++ {
		values[i] = ex.srcs[i].Eval(ctx)
	}

	for i := 0; i < size; i++ {
		setter(ex.ids[i].Value, values[i])
	}

	return
}
