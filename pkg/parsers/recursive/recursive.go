package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// ParseAll parses all tokens in tks into Statements.
func ParseAll(tks []token.Token) []st.Statement {
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
func (p *parser) script() (ss []st.Statement) {

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
func (p *parser) statement() (s st.Statement) {

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

func (p *parser) assignment(s *st.Statement) {

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
func (p *parser) expressions(required bool) (exprs []st.Expression, found bool) {

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
func (p *parser) expression(required bool) (st.Expression, bool) {

	switch left := p.term(); {
	case left != nil:
		return p.operation(left, 0), true

	case p.inspect(token.PAREN_OPEN):
		left = p.grouping()
		return p.operation(left, 0), true

	case p.inspect(token.LIST_OPEN):
		return p.list(), true

	case p.inspect(token.LIST_CLOSE):
		return nil, false

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
func (p *parser) term() st.Expression {

	switch {
	case p.accept(token.ID),
		p.accept(token.VOID),
		p.accept(token.BOOL),
		p.accept(token.INT),
		p.accept(token.FLOAT),
		p.accept(token.STRING),
		p.accept(token.TEMPLATE):

		return st.NewValueExpression(p.tk)
	}

	return nil
}

// Preconditions:
// - next = token.PAREN_OPEN
func (p *parser) grouping() st.Expression {
	p.expect(`grouping`, token.PAREN_OPEN)
	left, _ := p.expression(true)
	p.expect(`grouping`, token.PAREN_CLOSE)
	return left
}

// operation?
//
// Preconditions: NONE
func (p *parser) operation(left st.Expression, leftPriority int) st.Expression {

	op := p.snoop()

	if leftPriority >= st.Precedence(op.Type) {
		return left
	}

	p.expect(`operation`, op.Type) // Because we only snooped at it previously

	right := p.rightExpression(true)
	right = p.operation(right, st.Precedence(op.Type))

	left = st.NewOperation(left, op, right)
	left = p.operation(left, leftPriority)

	return left
}

func (p *parser) rightExpression(required bool) st.Expression {

	switch left := p.term(); {
	case left != nil:
		return left

	case p.inspect(token.PAREN_OPEN):
		return p.grouping()

	case required:
		panic(unexpected("rightExpression", p.snoop(), `<term> | PAREN_OPEN`))
	}

	return nil
}

func (p *parser) list() st.Expression {
	start := p.proceed()
	exprs, _ := p.expressions(false)
	p.expect(`list`, token.LIST_CLOSE)
	return st.List{start, exprs, p.tk}
}
