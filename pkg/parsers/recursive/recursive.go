package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// ParseAll parses all tokens in tks into Statements.
func ParseAll(tks []token.Token) []st.Statement {
	p := &pipe{token.NewIterator(tks)}
	return statements(p)
}

// statements parses all statements within the parsers iterator.
//
// Preconditions: None
func statements(p *pipe) (ss []st.Statement) {

	for !p.itr.Empty() && !p.accept(token.EOF) && !p.inspect(token.BLOCK_CLOSE) {
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

	if a := assignment(p); a != nil {
		return *a
	}

	if g := guard(p); g != nil {
		return *g
	}

	if m := match(p); m != nil {
		return *m
	}

	panic(unexpected("statement", p.snoop(), token.ANY))
}

// assignment?
//
// Preconditions:
// - Next token is not empty
func assignment(p *pipe) *st.Assignment {

	if !p.accept(token.ID) {
		return nil
	}

	if !p.inspect(token.DELIM) && !p.inspect(token.ASSIGN) {
		p.retract()
		return nil
	}

	p.retract()
	ids := identifiers(p)

	if p.accept(token.ASSIGN) {

		tk := p.prior()
		exprs := expressions(p)

		if exprs == nil {
			panic(unexpected("assignment", p.snoop(), token.ANY))
		}

		p.expect(`assignment`, token.TERMINATOR)
		return &st.Assignment{ids, tk, exprs}
	}

	panic(unexpected("assignment", p.prior(), token.ANOTHER))
}

func match(p *pipe) *st.Match {

	if !p.accept(token.MATCH_OPEN) {
		return nil
	}

	m := st.Match{
		Open: p.prior(),
	}

	p.expect(`match`, token.TERMINATOR)

	m.Cases = guards(p)
	if m.Cases == nil {
		panic(unexpected("match", p.snoop(), token.GUARD_OPEN))
	}

	p.expect(`match`, token.BLOCK_CLOSE)
	m.Close = p.prior()
	p.expect(`block`, token.TERMINATOR)

	return &m
}

func guards(p *pipe) []st.Guard {

	var gs []st.Guard

	for g := guard(p); g != nil; g = guard(p) {
		gs = append(gs, *g)
	}

	return gs
}

// guard?
//
// Preconditions:
// - Next token is not empty
func guard(p *pipe) *st.Guard {

	if !p.accept(token.GUARD_OPEN) {
		return nil
	}

	g := &st.Guard{
		Open: p.prior(),
	}

	condition := expression(p)

	if condition == nil || !boolOperator(condition) {
		panic(err("guard", condition.Token(),
			`Expected expression with a bool result`,
		))
	}

	g.Cond = condition
	p.expect(`guard`, token.GUARD_CLOSE)
	g.Close = p.prior()

	if b := block(p); b != nil {
		g.Block = *b
	} else {
		g.Block = inlineBlock(p)
	}

	return g
}

func block(p *pipe) *st.Block {

	if !p.accept(token.BLOCK_OPEN) {
		return nil
	}

	b := &st.Block{
		Open:  p.prior(),
		Stats: statements(p),
	}

	p.expect(`block`, token.BLOCK_CLOSE)
	b.Close = p.prior()
	p.expect(`block`, token.TERMINATOR)

	return b
}

func inlineBlock(p *pipe) st.Block {
	return st.Block{
		Open:  p.snoop(),
		Stats: []st.Statement{statement(p)},
		Close: p.prior(),
	}
}

func boolOperator(ex st.Expression) bool {

	if _, ok := ex.(st.Identifier); ok {
		return true
	}

	if v, ok := ex.(st.Value); ok {
		return v.Source.Type == token.BOOL
	}

	if v, ok := ex.(st.Operation); ok {
		return st.IsBoolOperator(v.Operator.Type)
	}

	return false
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

	var left st.Expression

	switch {
	case isFuncCall(p):
		p.expect(`rightSide`, token.ID)
		left := st.NewValueExpression(p.prior())
		return funcCall(p, left)

	case literal(p):
		left = st.NewValueExpression(p.proceed())
		return operation(p, left, 0)

	case p.inspect(token.PAREN_OPEN):
		left = group(p)
		return operation(p, left, 0)

	case p.inspect(token.FUNC):
		return funcDef(p)

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
func literal(p *pipe) bool {

	switch {
	case p.inspect(token.ID),
		p.inspect(token.VOID),
		p.inspect(token.BOOL),
		p.inspect(token.NUMBER),
		p.inspect(token.STRING),
		p.inspect(token.TEMPLATE):

		return true
	}

	return false
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

	switch {
	case isFuncCall(p):
		p.expect(`rightSide`, token.ID)
		left := st.NewValueExpression(p.prior())
		return funcCall(p, left)

	case literal(p):
		return st.NewValueExpression(p.proceed())

	case p.inspect(token.PAREN_OPEN):
		return group(p)
	}

	panic(unexpected("rightSide", p.snoop(), `<literal> | PAREN_OPEN`))
}

func funcDef(p *pipe) st.Expression {

	p.expect(`funcDef`, token.FUNC)
	f := st.FuncDef{
		Open: p.prior(),
	}

	p.expect(`funcDef`, token.PAREN_OPEN)

	if p.inspect(token.ID) {
		f.Input = identifiers(p)
	}

	if p.accept(token.RETURNS) {
		f.Output = identifiers(p)
	}

	p.expect(`funcDef`, token.PAREN_CLOSE)

	if b := block(p); b != nil {
		f.Body = *b
	} else {
		f.Body = inlineBlock(p)
	}

	p.retract() // Put TERMINATOR back for the statement to end correctly

	return f
}

func isFuncCall(p *pipe) (is bool) {

	if p.accept(token.ID) {
		is = p.inspect(token.PAREN_OPEN)
		p.retract()
	}

	return is
}

func funcCall(p *pipe, left st.Expression) st.Expression {

	p.expect(`funcCall`, token.PAREN_OPEN)

	f := st.FuncCall{
		ID:    left,
		Input: expressions(p),
	}

	p.expect(`funcCall`, token.PAREN_CLOSE)

	return f
}

func list(p *pipe) st.Expression {
	start := p.proceed()
	exprs := expressions(p)
	p.expect(`list`, token.LIST_CLOSE)
	return st.List{start, exprs, p.prior()}
}
