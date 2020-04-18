package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// ParseAll parses all tokens in tks into Statements.
func ParseAll(tks []token.Token) []st.Statement {
	p := &pipe{token.NewIterator(tks)}
	return script(p)
}

// script parses all statements within the parsers iterator.
//
// Preconditions: None
func script(p *pipe) (ss []st.Statement) {

	for !p.itr.Empty() && !p.accept(token.EOF) {
		s := statement(p)
		ss = append(ss, s)
	}

	return
}

// statement parses a single statement from the parsers iterator.
//
// Preconditions:
// - Next token is not empty
func statement(p *pipe) (s st.Statement) {

	switch {
	case assignment(p, s):

	default:
		panic(unexpected("statement", p.snoop(), token.ANY))
	}

	p.expect(`statement`, token.TERMINATOR)
	return s
}

// assignment?
//
// Preconditions:
// - Next token is not empty
func assignment(p *pipe, s st.Statement) bool {

	if !p.accept(token.ID) {
		return false
	}

	if !p.inspect(token.DELIM) && !p.inspect(token.ASSIGN) {
		p.retract()
		return false
	}

	p.retract()
	ids := identifiers(p)

	if p.accept(token.ASSIGN) {

		tk := p.prior()
		exprs := expressions(p)

		if exprs == nil {
			panic(unexpected("statement", p.snoop(), token.ANY))
		}

		s = st.Assignment{ids, tk, exprs}
		return true
	}

	panic(unexpected("assignment", p.prior(), token.ANOTHER))
}

// E.g. `a, b, c`
//
// Preconditions:
// - next = token.ID
func identifiers(p *pipe) []token.Token {

	p.expect(`identifiers`, token.ID)
	ids := []token.Token{p.prior()}

	for p.accept(token.DELIM) {
		p.expect(`identifiers`, token.ID)
		ids = append(ids, p.prior())
	}

	return ids
}

// expressions?
//
// Preconditions:
// - p.prior() = <Any>
func expressions(p *pipe) []st.Expression {

	var exprs []st.Expression

	for ex := expression(p); ex != nil; ex = expression(p) {
		exprs = append(exprs, ex)

		if !p.accept(token.DELIM) {
			break
		}
	}

	return exprs
}

// expression?
//
// Preconditions: NONE
func expression(p *pipe) st.Expression {

	switch left := literal(p); {
	case left != nil:
		return operation(p, left, 0)

	case p.inspect(token.PAREN_OPEN):
		left = group(p)
		return operation(p, left, 0)

	case p.inspect(token.LIST_OPEN):
		return list(p)
	}

	return nil
}

// literal is used to determine if p.prior() is a literal value.
// E.g.identifier, bool, int, etc.
//
// Preconditions:
// - p.prior() = <Any>
func literal(p *pipe) st.Expression {

	switch {
	case p.accept(token.ID),
		p.accept(token.VOID),
		p.accept(token.BOOL),
		p.accept(token.NUMBER),
		p.accept(token.STRING),
		p.accept(token.TEMPLATE):

		return st.NewValueExpression(p.prior())
	}

	return nil
}

// Preconditions:
// - next = token.PAREN_OPEN
func group(p *pipe) st.Expression {
	p.expect(`group`, token.PAREN_OPEN)

	left := expression(p)
	if left == nil {
		panic(unexpected("group", p.prior(), token.ANOTHER))
	}

	p.expect(`group`, token.PAREN_CLOSE)
	return left
}

// operation?
//
// Preconditions: NONE
func operation(p *pipe, left st.Expression, leftPriority int) st.Expression {

	op := p.snoop()

	if leftPriority >= st.Precedence(op.Type) {
		return left
	}

	p.expect(`operation`, op.Type) // Because we only snooped at it previously

	right := rightSide(p)
	right = operation(p, right, st.Precedence(op.Type))

	left = st.Operation{left, op, right}
	left = operation(p, left, leftPriority)

	return left
}

func rightSide(p *pipe) st.Expression {

	switch left := literal(p); {
	case left != nil:
		return left

	case p.inspect(token.PAREN_OPEN):
		return group(p)
	}

	panic(unexpected("rightSide", p.snoop(), `<literal> | PAREN_OPEN`))
}

func list(p *pipe) st.Expression {
	start := p.proceed()
	exprs := expressions(p)
	p.expect(`list`, token.LIST_CLOSE)
	return st.List{start, exprs, p.prior()}
}
