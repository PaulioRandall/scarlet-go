package parser

import (
	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseAssign parses an assignment into a statement. Assumes that the next
// statement in the input channel is an assignment.
func (p *Parser) parseAssign(inline bool) Stat {

	var fix token.Token
	if p.peek().Kind == token.FIX {
		fix = p.take()
	}

	ids := p.parseIDs()
	p.checkNoDuplicates(ids)

	ass := p.takeEnsure(token.ASSIGN)
	srcs := p.parseDelimExpr(true)

	if !inline {
		p.takeEnsure(token.TERMINATOR)
	}

	if len(ids) != len(srcs) {
		panic(bard.NewHorror(ass, nil,
			"Assignment requires the ID and expression count match",
		))
	}

	return assignStat{
		ass:  ass,
		fix:  fix,
		ids:  ids,
		srcs: srcs,
	}
}

// parseIDs parses a delimitered list of ID tokens used for an assignment.
func (p *Parser) parseIDs() (ids []token.Token) {
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

// checkNoDuplicates checks that there are no duplicate IDs within a slice. A
// panic ensues otherwise.
func (p *Parser) checkNoDuplicates(ids []token.Token) {
	for i, sub := range ids {
		for j, obj := range ids {
			if i != j && sub.Value == obj.Value {
				panic(bard.NewHorror(obj, nil, "Duplicate IDs not allowed"))
			}
		}
	}
}

// assignStat represents a statement that creates or updates variables by
// evaluating expressions into values which are mapped to their identifier
// within a context.
type assignStat struct {
	ass  token.Token
	fix  token.Token
	ids  []token.Token
	srcs []Expr
}

// String satisfies the Expr interface.
func (ex assignStat) String() (s string) {

	var (
		isFixed = ex.fix != (token.Token{})
		size    = len(ex.ids)
	)

	for i := 0; i < size; i++ {

		if i != 0 {
			s += "\n"
		}

		s += "Assign "

		if isFixed {
			s += "[" + ex.fix.String() + "] "
		}

		s += "[" + ex.ids[i].String() + "] "
		s += "[" + ex.ass.String() + "] "
		s += "[" + ex.srcs[i].String() + "]"
	}

	return
}

// Eval satisfies the Expr interface.
func (ex assignStat) Eval(ctx Context) (_ Value) {

	var (
		size   = len(ex.ids)
		values = make([]Value, size)
	)

	for i := 0; i < size; i++ {
		values[i] = ex.srcs[i].Eval(ctx)
	}

	for i := 0; i < size; i++ {
		isFixed := ex.fix != (token.Token{})
		ctx.set(ex.ids[i].Value, values[i], isFixed)
	}

	return
}
