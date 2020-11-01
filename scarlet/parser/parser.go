// Package parser converts a stream of Tokens into ASTs.
package parser

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
)

// Parse parses a slice of Tokens into a slice of abstract syntax trees.
func ParseAll(tks []token.Lexeme) ([]tree.Node, error) {
	itr := token.NewLexItr(tks)
	ctx := newCtx(itr, nil)
	return statements(ctx)
}

func newCtx(itr *token.LexItr, parent *context) *context {
	return &context{
		LexItr: itr,
		parent: parent,
	}
}

// Parses: {<assign> | <expr>}
func statements(ctx *context) ([]tree.Node, error) {

	nodes := []tree.Node{}

	for ctx.More() {
		n, e := statement(ctx)
		if e != nil {
			return nil, e
		}
		nodes = append(nodes, n)
		expectTerminator(ctx)
	}

	return nodes, nil
}

// Parses: <assign> | <expr>
func statement(ctx *context) (n tree.Node, e error) {
	switch l := ctx.Peek(); {
	case l.Token == token.IDENT:
		n, e = identLeads(ctx)

	case l.IsLiteral(), l.Token == token.L_PAREN:
		n, e = expectExpr(ctx)

	case l.Token == token.SPELL:
		n, e = spellCall(ctx)

	default:
		e = errSnip(l.Snippet,
			"%s does not lead any known statement", l.Token.String())
	}

	return
}

// Assumes: IDENT ...
func identLeads(ctx *context) (tree.Node, error) {
	switch ctx.Next(); ctx.Peek().Token {
	case token.ASSIGN:
		return singleAssign(ctx)

	case token.DELIM:
		return multiAssign(ctx)

	default:
		ctx.Back()
		return expectExpr(ctx)
	}
}

// Assumes: IDENT ASSIGN ...
// Parses: IDENT ASSIGN <expr>
func singleAssign(ctx *context) (tree.Node, error) {

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

	s.Snippet = token.SuperSnippet(s.Left.Pos(), s.Right.Pos())
	return s, nil
}

// Assumes: IDENT DELIM ...
// Parses: IDENT {DELIM IDENT} ASSIGN <expr> {DELIM <expr>}
func multiAssign(ctx *context) (tree.Node, error) {

	var (
		lSnip token.Snippet
		rSnip token.Snippet
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
	m.Snippet = token.SuperSnippet(lSnip, rSnip)

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
func multiAssignLeft(ctx *context) ([]tree.Assignee, token.Snippet, error) {

	var (
		zero  token.Snippet
		snip  token.Snippet
		l     token.Lexeme
		nodes []tree.Assignee
		id    tree.Ident
		e     error
	)

	l = ctx.Get()
	snip = l.Snippet

	if id, e = expectIdent(l); e != nil {
		return nil, zero, e
	}
	nodes = append(nodes, id)

	for ctx.Peek().Token == token.DELIM {
		ctx.Next()
		l = ctx.Next()

		if id, e = expectIdent(l); e != nil {
			return nil, zero, e
		}
		nodes = append(nodes, id)
	}

	snip = token.SuperSnippet(snip, l.Snippet)
	return nodes, snip, nil
}

// Parses: <expr> {DELIM <expr>}
func multiAssignRight(ctx *context) ([]tree.Expr, token.Snippet, error) {

	var (
		zero  token.Snippet
		snip  token.Snippet
		nodes []tree.Expr
		ex    tree.Expr
		e     error
	)

	if ex, e = expectExpr(ctx); e != nil {
		return nil, zero, e
	}
	nodes = append(nodes, ex)
	snip = ex.Pos()

	for ctx.Peek().Token == token.DELIM {
		ctx.Next()

		if ex, e = expectExpr(ctx); e != nil {
			return nil, zero, e
		}
		nodes = append(nodes, ex)
	}

	snip = token.SuperSnippet(snip, ex.Pos())
	return nodes, snip, nil
}

// Assumes SPELL ...
// Parses: SPELL L_PAREN [<expr> {DELIM <expr>}] R_PAREN
func spellCall(ctx *context) (tree.Node, error) {

	var e error
	sp := ctx.Next()
	n := tree.SpellCall{
		Snippet: sp.Snippet,
		Name:    sp.Val[1:],
	}

	if n.Args, e = expectParams(ctx); e != nil {
		return nil, e
	}

	n.ArgCount = len(n.Args)
	return n, nil
}
