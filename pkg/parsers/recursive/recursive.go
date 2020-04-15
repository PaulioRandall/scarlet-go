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

	for expr, ok := p.expression(s, true); ok; expr, ok = p.expression(s, true) {

		found = true
		s.Exprs = append(s.Exprs, expr)

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
func (p *parser) expression(s *Statement, required bool) (Expression, bool) {

	switch {
	case p.term():
		if expr, ok := p.arithmetic(); ok {
			return expr, true
		}

		return ExpressionOf(p.tk), true

	default:
		if required {
			panic(unexpected("expression", p.tk, token.ANOTHER))
		}
	}

	return nil, false
}

// term is used to determine if p.tk is a term, e.g. identifier, bool, int, etc.
//
// Preconditions:
// - p.tk = <Any>
func (p *parser) term() bool {

	switch {
	case p.confirm(token.ID),
		p.confirm(token.VOID),
		p.confirm(token.BOOL),
		p.confirm(token.INT),
		p.confirm(token.FLOAT),
		p.confirm(token.STRING),
		p.confirm(token.TEMPLATE):

		return true
	}

	return false
}

// arithmetic?
//
// Preconditions:
// - p.tk = term
func (p *parser) arithmetic() (Expression, bool) {

	expr := ExpressionOf(p.tk)
	expr = p.highArithmetic(expr)
	expr = p.lowArithmetic(expr)

	return expr, true
}

// highArithmetic?
//
// Preconditions: NONE
func (p *parser) highArithmetic(left Expression) Expression {

	if p.inspect(token.MULTIPLY) || p.inspect(token.DIVIDE) {
		op := p.proceed() // * or /

		if p.expect(token.ANY); !p.term() { // Right side must be a term
			panic(unexpected("highArithmetic", p.itr.Peek(), token.ANOTHER))
		}

		right := ExpressionOf(p.tk)
		left = Arithmetic{left, op, right}
		left = p.highArithmetic(left)
	}

	return left
}

// lowArithmetic?
//
// Preconditions: NONE
func (p *parser) lowArithmetic(left Expression) Expression {

	if p.inspect(token.ADD) || p.inspect(token.SUBTRACT) {
		op := p.proceed() // + or -

		if p.expect(token.ANY); !p.term() { // Right side must be a term
			panic(unexpected("lowArithmetic", p.itr.Peek(), token.ANOTHER))
		}

		right := ExpressionOf(p.tk)
		right = p.highArithmetic(right)
		left = Arithmetic{left, op, right}
		left = p.lowArithmetic(left)
	}

	return left
}
