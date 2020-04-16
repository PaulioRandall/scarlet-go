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

	if exprs, ok := p.expressions(true); ok {
		p.expect(`statement`, token.TERMINATOR)
		s.Exprs = exprs
		return s
	}

	panic(unexpected("statement", p.tk, token.ANY))
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
		p.expect(`identifiers`, token.DELIM) // Skip delimitier
		p.expect(`identifiers`, token.ID)    // Must be an ID after a delimitier
		ids = append(ids, p.tk)              // Append next ID
	}

	if p.accept(token.ASSIGN) {
		p.assignment(s, ids)
		return true
	}

	p.expect(`identifiers`, token.ANOTHER)
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
	exprs, _ := p.expressions(true)
	p.expect(`assignment`, token.TERMINATOR)
	s.Exprs = exprs
}

// expressions?
//
// Preconditions:
// - p.tk = <Any>
func (p *parser) expressions(required bool) (exprs []Expression, found bool) {

	for expr, ok := p.expression(required); ok; expr, ok = p.expression(true) {

		found = true
		exprs = append(exprs, expr)

		if !p.accept(token.DELIM) {
			break
		}

		p.expect(`expressions`, token.ANY)
	}

	return
}

// expression?
//
// Preconditions:
// - p.tk = <Any>
func (p *parser) expression(required bool) (Expression, bool) {

	switch {
	case p.term():
		fallthrough

	case p.confirm(token.PAREN_OPEN):
		left := p.newOperation()
		return p.operation(left, 0), true

	case p.confirm(token.LIST_OPEN):
		return p.list(), true

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

// operation?
//
// Preconditions:
// - p.tk = <ANY>
func (p *parser) newOperation() Expression {
	switch {
	case p.confirm(token.PAREN_OPEN):
		p.expect(`newOperation`, token.ANY)
		expr, _ := p.expression(true)
		p.expect(`newOperation`, token.PAREN_CLOSE)
		return expr

	case p.term():
		return NewValueExpression(p.tk)

	default:
		panic(unexpected("newOperation", p.tk, token.ANOTHER))
	}
}

// operation?
//
// Preconditions: NONE
func (p *parser) operation(left Expression, leftPriority int) Expression {

	op := p.snoop()
	opPriority := precedence(op)

	if leftPriority >= opPriority {
		return left
	}

	p.expect(`operation`, op.Type)
	p.expect(`operation`, token.ANY)

	right := p.newOperation()
	right = p.operation(right, opPriority)

	left = NewMathExpression(left, op, right)
	left = p.operation(left, leftPriority)

	return left
}

func (p *parser) list() Expression {

	start := p.tk
	p.expect(`list`, token.ANY)
	exprs, ok := p.expressions(false)

	if ok {
		p.expect(`list`, token.LIST_CLOSE)
	} else {
		p.affirm(`list`, token.LIST_CLOSE)
	}

	return List{start, exprs, p.tk}
}

func precedence(tk token.Token) int {
	switch tk.Type {
	case token.MULTIPLY, token.DIVIDE, token.REMAINDER: // Multiplicative
		return 6
	case token.ADD, token.SUBTRACT: // Additive
		return 5
	case token.LESS_THAN, token.LESS_THAN_OR_EQUAL: // Relational
		fallthrough
	case token.MORE_THAN, token.MORE_THAN_OR_EQUAL: // Relational
		return 4
	case token.EQUAL, token.NOT_EQUAL: // Equality
		return 3
	case token.AND:
		return 2
	case token.OR:
		return 1
	default:
		return 0
	}
}
