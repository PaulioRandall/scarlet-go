package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseAssign parses an assignment into a statement. Assumes that the next
// statement in the input channel is an assignment.
func (p *Parser) parseAssign(inline bool) Stat {

	ids := p.parseIDs()
	p.checkNoDuplicates(ids)

	ass := p.takeEnsure(token.LEXEME_ASSIGN)
	srcs := p.parseDelimExpr(true)

	if !inline {
		p.takeEnsure(token.LEXEME_TERMINATOR)
	}

	return assignStat{
		ass:  ass,
		ids:  ids,
		srcs: srcs,
	}
}

// parseIDs parses a delimitered list of ID tokens used for an assignment.
func (p *Parser) parseIDs() (ids []token.Token) {
	for {

		tk := p.takeEnsure(token.LEXEME_ID)
		ids = append(ids, tk)

		if p.peek().Lexeme == token.LEXEME_DELIM {
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
	ids  []token.Token
	srcs []Expr
}

// String satisfies the Expr interface.
func (ex assignStat) String() (s string) {

	s += "Assign (" + ex.ass.String() + ")"

	s += "\n\tIDs"
	for _, id := range ex.ids {
		s += "\n\t\t" + id.String()
	}

	s += "\n\tExprs"
	exprs := ""
	for _, src := range ex.srcs {
		exprs += "\n\t" + strings.ReplaceAll(src.String(), "\n", "\n\t")
	}

	return s + strings.ReplaceAll(exprs, "\n", "\n\t")
}

// Eval satisfies the Expr interface.
func (ex assignStat) Eval(ctx Context) (_ Value) {

	var (
		idCount = len(ex.ids)
		values  []Value
	)

	for _, src := range ex.srcs {
		val := src.Eval(ctx)

		if val.k == TUPLE {
			tuple := val.v.([]Value)
			values = append(values, tuple...)
			continue
		}

		values = append(values, val)
	}

	if idCount != len(values) {
		panic(bard.NewHorror(ex.ass, nil,
			"Expected the left-hand ID count and "+
				"right-hand value count to be equal",
		))
	}

	for i := 0; i < idCount; i++ {
		ctx.set(ex.ids[i].Value, values[i])
	}

	return
}
