// Package parser converts a stream of Tokens into a parser tree...
package parser

import (
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
	"github.com/PaulioRandall/scarlet-go/token2/tree"
)

// Parse parses a series of Tokens into a series of parse trees.
func Parse(itr TokenItr) ([]tree.Node, error) {
	ctx := newCtx(itr, nil)
	return statements(ctx)
}

func newCtx(itr TokenItr, parent *context) *context {
	return &context{
		TokenItr: itr,
		parent:   parent,
	}
}

// Parses: {<assign> | <expr>}
func statements(ctx *context) ([]tree.Node, error) {

	r := []tree.Node{}

	for ctx.More() {
		n, e := statement(ctx)
		if e != nil {
			return nil, e
		}
		r = append(r, n)
	}

	return r, nil
}

// Parses: <assign> | <expr>
func statement(ctx *context) (n tree.Node, e error) {
	switch l := ctx.LookAhead(); {
	case l.Token == token.IDENT:
		ctx.Next()
		n, e = identLeads(ctx)

	case l.IsLiteral(), l.Token == token.L_PAREN:
		n, e = expectExpr(ctx)

	default:
		e = errSnip(l.Snippet,
			"%s does not lead any known statement", l.Token.String())
	}

	return
}

// identLeads must only be used when the next Token is an IDENT and it begins
// a statement.
func identLeads(ctx *context) (tree.Node, error) {

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
// Parses: IDENT ASSIGN <expr>
func singleAssignment(ctx *context) (tree.Node, error) {

	var e error
	var s tree.SingleAssign

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
// Parses: IDENT {DELIM IDENT} ASSIGN <expr> {DELIM <expr>}
func multiAssignment(ctx *context) (tree.Node, error) {

	var (
		lSnip position.Snippet
		rSnip position.Snippet
		zero  tree.MultiAssign
		m     tree.MultiAssign
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
// Parses: IDENT {DELIM IDENT}
func multiAssignLeft(ctx *context) ([]tree.Assignee, position.Snippet, error) {

	var (
		zero position.Snippet
		snip position.Snippet
		l    lexeme.Lexeme
		r    []tree.Assignee
		id   tree.Ident
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

// Parses: <expr> {DELIM <expr>}
func multiAssignRight(ctx *context) ([]tree.Expr, position.Snippet, error) {

	var (
		zero position.Snippet
		snip position.Snippet
		r    []tree.Expr
		ex   tree.Expr
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

// Parses: BOOL | NUMBER | STRING
func expectLiteral(ctx *context) (tree.Expr, error) {

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

// Parses:  <expr> {<operator> <expr>}
// Parses:  L_PAREN <expr> {<operator> <expr>} R_PAREN
func expectExpr(ctx *context) (tree.Expr, error) {
	return expectExprRight(ctx, 0)
}

// Parses: <term>
func expectTerm(ctx *context) (tree.Expr, error) {
	return expectLiteral(ctx)
}

// Parses:  <expr> {<operator> <expr>}
// Parses:  L_PAREN <expr> {<operator> <expr>} R_PAREN
func expectExprRight(ctx *context, leftOpPrec int) (tree.Expr, error) {

	var left tree.Expr
	var e error

	if ctx.LookAhead().Token == token.L_PAREN {
		left, e = expectExprParen(ctx)
	} else {
		left, e = expectTerm(ctx)
	}

	if e != nil {
		return nil, e
	}

	return maybeBinaryExpr(ctx, left, leftOpPrec)
}

// Parses: L_PAREN <expr> {<operator> <expr>} R_PAREN
func expectExprParen(ctx *context) (tree.Expr, error) {

	if !ctx.More() {
		return nil, errPos(ctx.Snippet().End, "Missing left parenthesis")
	}

	if l := ctx.Next(); l.Token != token.L_PAREN {
		return nil, errSnip(l.Snippet,
			"Expected left parenthesis but got %s", l.Token.String())
	}

	r, e := expectExprRight(ctx, 0)
	if e != nil {
		return nil, e
	}

	if !ctx.More() {
		return nil, errPos(ctx.Snippet().End, "Missing right parenthesis")
	}

	if l := ctx.Next(); l.Token != token.R_PAREN {
		return nil, errSnip(l.Snippet,
			"Expected right parenthesis but got %s", l.Token.String())
	}

	return r, nil
}

// Parses: {<operator> <expr>}
func maybeBinaryExpr(ctx *context, left tree.Expr, leftOpPrec int) (tree.Expr, error) {

	if !ctx.LookAhead().IsBinaryOperator() {
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

	binExpr := tree.BinaryExpr{
		Left:    left,
		Op:      op.Token,
		OpPos:   op.Snippet,
		Right:   right,
		Snippet: position.SuperSnippet(left.Pos(), right.Pos()),
	}

	return maybeBinaryExpr(ctx, binExpr, leftOpPrec)
}
