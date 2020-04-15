package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// ParseAll parses all tokens in tks into Statements.
func ParseAll(tks []lexeme.Token) []Statement {
	p := parser{itr: lexeme.NewIterator(tks)}
	return p.script()
}

// parser stores a single read token to enable look ahead by one behaviour.
type parser struct {
	itr *lexeme.TokenIterator
	tk  lexeme.Token
}

// script parses all statements within the parsers iterator.
//
// Preconditions: None
func (p *parser) script() []Statement {

	var ss []Statement

	for !p.itr.Empty() && !p.accept(lexeme.LEXEME_EOF) {
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

	p.accept(lexeme.LEXEME_ANY)

	if p.confirm(lexeme.LEXEME_ID) {
		if p.identifiers(&s) {
			return s
		}
	}

	if !p.expressions(&s) {
		panic(unexpected("statement", p.tk, lexeme.LEXEME_ANY))
	}

	return s
}

// identifiers parses a delimiter separated list of identifiers then invokes
// the approriate function to parse the remainder of the statement.
//
// Preconditions:
// - p.tk = LEXEME_ID
func (p *parser) identifiers(s *Statement) bool {

	switch {
	case p.inspect(lexeme.LEXEME_DELIM):
	case p.inspect(lexeme.LEXEME_ASSIGN):

	default:
		return false
	}

	ids := []lexeme.Token{p.tk}

	for p.inspect(lexeme.LEXEME_DELIM) { // a, b, c...
		p.expect(lexeme.LEXEME_DELIM) // Skip delimitier
		p.expect(lexeme.LEXEME_ID)    // Must be an ID after a delimitier
		ids = append(ids, p.tk)       // Append next ID
	}

	if p.accept(lexeme.LEXEME_ASSIGN) {
		p.assignment(s, ids)
		return true
	}

	p.expect(lexeme.LEXEME_ANOTHER)
	return true
}

// assignment?
//
// Preconditions:
// - p.tk = LEXEME_ASSIGN
func (p *parser) assignment(s *Statement, ids []lexeme.Token) {
	s.IDs = ids
	s.Assign = p.tk
	p.accept(lexeme.LEXEME_ANY)
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

		if !p.accept(lexeme.LEXEME_DELIM) {
			break
		}

		p.expect(lexeme.LEXEME_ANY)
	}

	p.expect(lexeme.LEXEME_TERMINATOR)
	return found
}

// expression?
//
// Preconditions:
// - p.tk = <Any>
func (p *parser) expression(s *Statement, required bool) bool {
	switch {
	case p.factor(s):
		s.Exprs = append(s.Exprs, Value{p.tk})

	default:
		if required {
			panic(unexpected("expression", p.tk, lexeme.LEXEME_ANOTHER))
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
	case p.confirm(lexeme.LEXEME_VOID),
		p.confirm(lexeme.LEXEME_BOOL),
		p.confirm(lexeme.LEXEME_INT),
		p.confirm(lexeme.LEXEME_FLOAT),
		p.confirm(lexeme.LEXEME_STRING),
		p.confirm(lexeme.LEXEME_TEMPLATE):

	default:
		return false
	}

	return true
}
