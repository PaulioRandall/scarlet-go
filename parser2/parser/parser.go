// Package parser converts a stream of Tokens into a parser tree...
package parser

import (
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

// Parse parses a series of Tokens into a series of parse trees.
func Parse(itr TokenItr) ([]Stat, error) {
	ctx := newCtx(itr, nil)
	return statements(ctx)
}

func newCtx(itr TokenItr, parent *context) *context {
	return &context{
		TokenItr: itr,
		parent:   parent,
	}
}

func statements(ctx *context) ([]Stat, error) {

	var (
		r = []Stat{}
		s Stat
		e error
	)

	for ctx.More() {
		switch l := ctx.LookAhead(); {
		case l.Token == token.IDENT:
			ctx.Next()
			s, e = indentLeads(ctx)

		default:
			return nil, errSnip(l.Snippet,
				"%s does not lead any known statement", l.Token.String())
		}

		if e != nil {
			return nil, e
		}

		r = append(r, s)
	}

	return r, nil
}

// indentLeads must only be used when the next Token is an IDENT and it begins
// a statement.
func indentLeads(ctx *context) (Stat, error) {

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
func singleAssignment(ctx *context) (Stat, error) {

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
func multiAssignment(ctx *context) (Stat, error) {

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
func expectExpr(ctx *context) (Expr, error) {

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
