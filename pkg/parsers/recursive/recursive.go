package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// ParseAll parses all tokens in tks into Statements.
func ParseAll(tks []token.Token) []Statement {
	p := parser{itr: token.NewIterator(tks)}
	return p.script()
}

// parser stores a single read token to enable look ahead by one behaviour.
type parser struct {
	itr *token.TokenIterator
	tk  token.Token
}

// script parses all statements within the parsers iterator.
//
// Preconditions: None
func (p *parser) script() []Statement {

	var ss []Statement

	for !p.itr.Empty() && !p.accept(token.EOF) {
		s := p.statement()
		ss = append(ss, s)
	}

	return ss
}

// statement parses a single statement from the parsers iterator.
//
// Preconditions:
// - p.itr is not empty
func (p *parser) statement() Statement {

	s := Statement{}

	p.accept(token.ANY)

	if p.confirm(token.ID) {
		if p.identifiers(&s) {
			return s
		}
	}

	if !p.expressions(&s) {
		panic(unexpected("statement", p.tk, token.ANY))
	}

	return s
}

// identifiers parses a delimiter separated list of identifiers then invokes
// the approriate function to parse the remainder of the statement.
//
// Preconditions:
// - p.tk = ID
func (p *parser) identifiers(s *Statement) bool {

	switch {
	case p.inspect(token.DELIM):
	case p.inspect(token.ASSIGN):

	default:
		return false
	}

	ids := []token.Token{p.tk}

	for p.inspect(token.DELIM) { // a, b, c...
		p.expect(token.DELIM)   // Skip delimitier
		p.expect(token.ID)      // Must be an ID after a delimitier
		ids = append(ids, p.tk) // Append next ID
	}

	if p.accept(token.ASSIGN) {
		p.assignment(s, ids)
		return true
	}

	p.expect(token.ANOTHER)
	return true
}

// assignment?
//
// Preconditions:
// - p.tk = ASSIGN
func (p *parser) assignment(s *Statement, ids []token.Token) {
	s.IDs = ids
	s.Assign = p.tk
	p.accept(token.ANY)
	p.expressions(s)
}

// expressions?
//
// Preconditions:
// - p.tk = <Any>
func (p *parser) expressions(s *Statement) bool {

	var found bool

	for p.expression(s, true) {
		found = true

		if !p.accept(token.DELIM) {
			break
		}

		p.expect(token.ANY)
	}

	p.expect(token.TERMINATOR)
	return found
}

// expression?
//
// Preconditions:
// - p.tk = <Any>
func (p *parser) expression(s *Statement, required bool) bool {
	switch {
	case p.confirm(token.ID):
		s.Exprs = append(s.Exprs, Identifier{p.tk})

	case p.factor(s):
		s.Exprs = append(s.Exprs, Value{p.tk})

	default:
		if required {
			panic(unexpected("expression", p.tk, token.ANOTHER))
		}
		return false
	}

	return true
}

// factor to parse
//
// Preconditions:
// - p.tk = <Any>
func (p *parser) factor(s *Statement) bool {

	switch {
	case p.confirm(token.VOID),
		p.confirm(token.BOOL),
		p.confirm(token.INT),
		p.confirm(token.FLOAT),
		p.confirm(token.STRING),
		p.confirm(token.TEMPLATE):

	default:
		return false
	}

	return true
}
