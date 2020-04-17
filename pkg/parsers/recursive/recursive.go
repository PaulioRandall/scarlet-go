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
func (p *parser) script() (ss []Statement) {

	for !p.itr.Empty() && !p.accept(token.EOF) {
		s := p.statement()
		ss = append(ss, s)
	}

	return
}

// statement parses a single statement from the parsers iterator.
//
// Preconditions:
// - p.itr is not empty
func (p *parser) statement() (s Statement) {

	p.assignment(&s)

	exprs, ok := p.expressions(true)

	if ok {
		p.expect(`statement`, token.TERMINATOR)
		s.Exprs = exprs
		return s
	}

	panic(unexpected("statement", p.tk, token.ANY))
}

// E.g. `a, b, c`
//
// Preconditions:
// - p.tk = token.ID
func (p *parser) multipleIdentifiers() []token.Token {

	ids := []token.Token{p.tk}

	for p.accept(token.DELIM) {
		p.expect(`identifiers`, token.ID)
		ids = append(ids, p.tk)
	}

	return ids
}

func (p *parser) assignment(s *Statement) {

	if !p.accept(token.ID) {
		return
	}

	if !p.inspect(token.DELIM) && !p.inspect(token.ASSIGN) {
		p.retract()
		return
	}

	ids := p.multipleIdentifiers()

	if p.accept(token.ASSIGN) {
		s.IDs = ids
		s.Assign = p.tk
		return
	}

	panic(unexpected("assignment", p.tk, token.ANOTHER))
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
	}

	return
}

// expression?
//
// Preconditions:
// - p.tk = <Any>
func (p *parser) expression(required bool) (Expression, bool) {

	switch left, isOperation := p.operationInit(false); {
	case isOperation:
		return p.operation(left, 0), true

	case p.inspect(token.LIST_OPEN):
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
	case p.accept(token.ID),
		p.accept(token.VOID),
		p.accept(token.BOOL),
		p.accept(token.INT),
		p.accept(token.FLOAT),
		p.accept(token.STRING),
		p.accept(token.TEMPLATE):

		return true
	}

	return false
}

func (p *parser) operationInit(required bool) (Expression, bool) {

	switch {
	case p.term():
		return NewValueExpression(p.tk), true

	case p.accept(token.PAREN_OPEN):
		left, _ := p.expression(true)
		p.expect(`newOperation`, token.PAREN_CLOSE)

		if op, ok := left.(Operation); ok {
			return p.operation(left, op.Precedence()), true
		}

		return p.operation(left, 0), true
	}

	if required {
		panic(unexpected("operationInit", p.snoop(), `<term> | PAREN_OPEN`))
	}

	return nil, false
}

// operation?
//
// Preconditions: NONE
func (p *parser) operation(left Expression, leftPriority int) Expression {

	op := p.snoop()

	if leftPriority >= Precedence(op.Type) {
		return left
	}

	p.expect(`operation`, op.Type) // Because we only snooped at it previously

	right, _ := p.operationInit(true)
	right = p.operation(right, Precedence(op.Type))

	left = NewOperation(left, op, right)
	left = p.operation(left, leftPriority)

	return left
}

func (p *parser) list() Expression {

	start := p.proceed()
	exprs, _ := p.expressions(false)
	p.expect(`list`, token.LIST_CLOSE)

	return List{start, exprs, p.tk}
}
