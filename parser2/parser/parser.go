// Package parser converts a stream of Tokens into a parser tree...
package parser

import (
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

// Parse parses a series of Tokens into a series of parse trees.
func Parse(itr TokenItr) ([]Node, error) {
	ctx := newCtx(itr, nil)
	return statements(ctx)
}

func newCtx(itr TokenItr, parent *context) *context {
	return &context{
		TokenItr: itr,
		parent:   parent,
	}
}

func statements(ctx *context) ([]Node, error) {

	var (
		r = []Node{}
		n Node
		e error
	)

	for ctx.More() {
		switch l := ctx.LookAhead(); {
		case l.Token == token.IDENT:
			ctx.Next()
			n, e = indentLeads(ctx)

		case l.IsLiteral():
			n, e = expectExpr(ctx)

		default:
			return nil, errSnip(l.Snippet,
				"%s does not lead any known statement", l.Token.String())
		}

		if e != nil {
			return nil, e
		}

		r = append(r, n)
	}

	return r, nil
}

// indentLeads must only be used when the next Token is an IDENT and it begins
// a statement.
func indentLeads(ctx *context) (Node, error) {

	l := ctx.LookAhead()

	switch l.Token {
	case token.ASSIGN:
		return singleAssignment(ctx)

	case token.DELIM:
		return multiAssignment(ctx)
	}

	return nil, errSnip(l.Snippet,
		"%s does not follow an identifier in any known statement", l.Token.String())
}

// Assumes: IDENT ASSIGN ...
// Pattern: IDENT ASSIGN <expr>
func singleAssignment(ctx *context) (Node, error) {

	var e error
	var s SingleAssign

	s.Left, e = expectIdent(ctx.Get())
	if e != nil {
		return s, e
	}

	s.Infix = ctx.Next().Snippet
	s.Right, e = expectExpr(ctx)
	if e != nil {
		return s, e
	}

	s.Snippet = position.SuperSnippet(s.Left.Pos(), s.Right.Pos())
	return s, nil
}

// Assumes: IDENT DELIM ...
// Pattern: IDENT {DELIM IDENT} ASSIGN <expr> {DELIM <expr>}
func multiAssignment(ctx *context) (Node, error) {

	var (
		lSnip position.Snippet
		rSnip position.Snippet
		zero  MultiAssign
		m     MultiAssign
		e     error
	)

	if m.Left, lSnip, e = multiAssignLeft(ctx); e != nil {
		return zero, e
	}

	l := ctx.Next()
	if l.Token != token.ASSIGN {
		return zero, errSnip(l.Snippet, "Expected assignment symbol")
	}
	m.Infix = l.Snippet

	if m.Right, rSnip, e = multiAssignRight(ctx); e != nil {
		return zero, e
	}
	m.Snippet = position.SuperSnippet(lSnip, rSnip)

	lSize, rSize := len(m.Left), len(m.Right)
	if lSize < rSize {
		return zero, errSnip(m.Snippet,
			"Not enough expressions on left or too many on right of assignment")
	} else if lSize > rSize {
		return zero, errSnip(m.Snippet,
			"Too many expressions on left or not enough on right of assignment")
	}

	return m, nil
}

// Assumes: IDENT DELIM ...
// Pattern: IDENT {DELIM IDENT}
func multiAssignLeft(ctx *context) ([]Expr, position.Snippet, error) {

	var (
		zero position.Snippet
		snip position.Snippet
		l    lexeme.Lexeme
		r    []Expr
		id   Ident
		e    error
	)

	l = ctx.Get()
	snip = l.Snippet

	if id, e = expectIdent(l); e != nil {
		return nil, zero, e
	}
	r = append(r, id)

	for ctx.LookAhead().Token == token.DELIM {
		ctx.Next()
		l = ctx.Next()

		if id, e = expectIdent(l); e != nil {
			return nil, zero, e
		}
		r = append(r, id)
	}

	snip = position.SuperSnippet(snip, l.Snippet)
	return r, snip, nil
}

// Pattern: <expr> {DELIM <expr>}
func multiAssignRight(ctx *context) ([]Expr, position.Snippet, error) {

	var (
		zero position.Snippet
		snip position.Snippet
		r    []Expr
		ex   Expr
		e    error
	)

	if ex, e = expectExpr(ctx); e != nil {
		return nil, zero, e
	}
	r = append(r, ex)
	snip = ex.Pos()

	for ctx.LookAhead().Token == token.DELIM {
		ctx.Next()

		if ex, e = expectExpr(ctx); e != nil {
			return nil, zero, e
		}
		r = append(r, ex)
	}

	snip = position.SuperSnippet(snip, ex.Pos())
	return r, snip, nil
}

// Pattern: BOOL | NUMBER | STRING
func expectLiteral(ctx *context) (Expr, error) {

	if !ctx.More() {
		return nil, errPos(ctx.Snippet().End,
			"Expected expression but reached EOF")
	}

	l := ctx.Next()
	switch l.Token {
	case token.TRUE, token.FALSE:
		return boolLit(l), nil
	case token.NUMBER:
		return numLit(l), nil
	case token.STRING:
		return strLit(l), nil
	}

	return nil, errSnip(l.Snippet,
		"%s does not lead any known expression", l.Token.String())
}

// Pattern: <literal> {<operator> <literal>}
func expectExpr(ctx *context) (Expr, error) {
	return expectExprRight(ctx, 0)
}

// Pattern: <literal>
func expectExprLeft(ctx *context) (Expr, error) {
	return expectLiteral(ctx)
}

// Pattern: <literal> {<operator> <literal>}
func expectExprRight(ctx *context, leftOpPrec int) (Expr, error) {
	left, e := expectExprLeft(ctx)
	if e != nil {
		return nil, e
	}
	return maybeBinaryExpr(ctx, left, leftOpPrec)
}

// Pattern: {<operator> <literal>}
func maybeBinaryExpr(ctx *context, left Expr, leftOpPrec int) (Expr, error) {

	// TODO: Multi-operator & precedence needs testing!!

	if !ctx.LookAhead().IsOperator() {
		return left, nil
	}

	if leftOpPrec >= ctx.LookAhead().Precedence() {
		return left, nil
	}

	op := ctx.Next()
	right, e := expectExprRight(ctx, op.Precedence())
	if e != nil {
		return nil, e
	}

	binExpr := BinaryExpr{
		Left:    left,
		Op:      op.Token,
		OpPos:   op.Snippet,
		Right:   right,
		Snippet: position.SuperSnippet(left.Pos(), right.Pos()),
	}

	return maybeBinaryExpr(ctx, binExpr, leftOpPrec)
}
