// Package parser converts a stream of Tokens into a parser tree...
package parser

import (
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

// NEXT:
// Store Lexemes in Nodes instead of Snippets, the positions and snippets can
// still be accessed but Node consumers have more flexibility.
// NEXT:
// Update perror to store Lexemes or positions and snippets.

type LexemeIterator interface {
	More() bool
	Next() lexeme.Lexeme
	Prev() lexeme.Lexeme
	LookAhead() lexeme.Lexeme
}

func Parse(itr LexemeIterator) ([]Stat, error) {
	ctx := newCtx(itr, nil)
	return statements(ctx)
}

func newCtx(itr LexemeIterator, parent *context) *context {
	return &context{
		LexemeIterator: itr,
		parent:         parent,
	}
}

func statements(ctx *context) ([]Stat, error) {

	var (
		stats = []Stat{}
		s     Stat
		e     error
	)

	for ctx.More() {
		switch l := ctx.LookAhead(); {
		case l.Token == token.IDENT:
			s, e = leadingIdentStat(ctx)

		default:
			return nil, newErr("Unparsable statement")
		}

		if e != nil {
			return nil, e
		}

		stats = append(stats, s)
	}

	return stats, nil
}

// leadingIdentStat must only be used when the next Token is an IDENT and it
// begins a statement.
func leadingIdentStat(ctx *context) (Stat, error) {

	switch ctx.Next(); ctx.LookAhead().Token {
	case token.ASSIGN:
		ctx.Prev()
		return singleAssignment(ctx)

	case token.DELIM:
		ctx.Prev()
		return multiAssignment(ctx)
	}

	return nil, newErr("Unknown statement")
}

// Assumes: IDENT ASSIGN ...
// Pattern: IDENT ASSIGN <expr>
func singleAssignment(ctx *context) (Stat, error) {

	var e error
	var s SingleAssign

	s.Left, e = expectIdent(ctx.Next())
	if e != nil {
		return s, e
	}

	s.Infix = ctx.Next().Snippet
	s.Right, e = expectExpr(ctx)
	return s, e
}

// Assumes: IDENT DELIM ...
// Pattern: IDENT {DELIM IDENT} ASSIGN <expr> {DELIM <expr>}
func multiAssignment(ctx *context) (Stat, error) {

	var (
		ZERO MultiAssign
		m    MultiAssign
		e    error
	)

	if m.Left, e = multiAssignLeft(ctx); e != nil {
		return ZERO, e
	}

	if ctx.LookAhead().Token != token.ASSIGN {
		return ZERO, newErr("Expected assignment token")
	}
	m.Infix = ctx.Next().Snippet

	if m.Right, e = multiAssignRight(ctx); e != nil {
		return ZERO, e
	}

	return m, nil
}

// Assumes: IDENT DELIM ...
// Pattern: IDENT {DELIM IDENT}
func multiAssignLeft(ctx *context) ([]Expr, error) {

	var (
		r  []Expr
		id Ident
		e  error
	)

	if id, e = expectIdent(ctx.Next()); e != nil {
		return nil, e
	}
	r = append(r, id)

	for ctx.LookAhead().Token == token.DELIM {
		ctx.Next()

		if id, e = expectIdent(ctx.Next()); e != nil {
			return nil, e
		}
		r = append(r, id)
	}

	return r, nil
}

// Pattern: <expr> {DELIM <expr>}
func multiAssignRight(ctx *context) ([]Expr, error) {

	var (
		r  []Expr
		ex Expr
		e  error
	)

	if ex, e = expectExpr(ctx); e != nil {
		return nil, e
	}
	r = append(r, ex)

	for ctx.LookAhead().Token == token.DELIM {
		ctx.Next()

		if ex, e = expectExpr(ctx); e != nil {
			return nil, e
		}
		r = append(r, ex)
	}

	return r, nil
}

// Pattern: BOOL | NUMBER | STRING
func expectExpr(ctx *context) (Expr, error) {

	switch ctx.LookAhead().Token {
	case token.TRUE, token.FALSE:
		return boolLit(ctx.Next()), nil
	case token.NUMBER:
		return numLit(ctx.Next()), nil
	case token.STRING:
		return strLit(ctx.Next()), nil
	}

	return nil, newErr("Unknown expression")
}
