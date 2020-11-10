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
	case l.Token == token.IDENT, l.Token == token.VOID:
		n, e = assigneeLeads(ctx)

	case l.IsLiteral(), l.Token == token.L_PAREN:
		n, e = expectExpr(ctx)

	case l.Token == token.SPELL:
		n, e = spellCall(ctx)

	case l.Token == token.L_SQUARE:
		n, e = guard(ctx)

	case l.Token == token.L_CURLY:
		n, e = expectBlock(ctx)

	default:
		e = errSnip(l.Snippet,
			"%s does not lead any known statement", l.Token.String())
	}

	return
}

// Assumes: <assignee> ...
func assigneeLeads(ctx *context) (tree.Node, error) {
	switch ctx.Next(); ctx.Peek().Token {
	case token.ASSIGN:
		return singleAssign(ctx)

	case token.DELIM:
		return multiOrAsymAssign(ctx)

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

	s.Left, e = expectAssignee(ctx.Get())
	if e != nil {
		return s, e
	}

	ctx.Next()
	s.Right, e = expectExpr(ctx)
	if e != nil {
		return s, e
	}

	return s, nil
}

// Assumes: IDENT DELIM ...
// Parses: IDENT {DELIM IDENT} ASSIGN <expr> {DELIM <expr>}
func multiOrAsymAssign(ctx *context) (tree.Node, error) {

	var (
		lSnip token.Snippet
		rSnip token.Snippet
		left  []tree.Assignee
		right []tree.Expr
		e     error
	)

	if left, lSnip, e = multiAssignLeft(ctx); e != nil {
		return nil, e
	}

	op := ctx.Next()
	if op.Token != token.ASSIGN {
		return nil, errSnip(op.Snippet, "Expected assignment symbol")
	}

	if right, rSnip, e = multiAssignRight(ctx); e != nil {
		return nil, e
	}

	snip := token.SuperSnippet(lSnip, rSnip)

	var m tree.Node

	switch lSize, rSize := len(left), len(right); {
	case lSize < rSize:
		return nil, errSnip(snip,
			"Not enough expressions on left or too many on right of assignment")

	case rSize == 1:
		if _, ok := right[0].(tree.SpellCall); !ok {
			return nil, errSnip(snip,
				"Too many expressions on left or not enough on right of assignment")
		}
		m = tree.AsymAssign{
			Left:  left,
			Right: right[0],
		}

	case lSize > rSize:
		return nil, errSnip(snip,
			"Too many expressions on left or not enough on right of assignment")

	default:
		m = tree.MultiAssign{
			Left:  left,
			Right: right,
		}
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
		a     tree.Assignee
		e     error
	)

	l = ctx.Get()
	snip = l.Snippet

	if a, e = expectAssignee(l); e != nil {
		return nil, zero, e
	}
	nodes = append(nodes, a)

	for ctx.Peek().Token == token.DELIM {
		ctx.Next()
		l = ctx.Next()

		if a, e = expectAssignee(l); e != nil {
			return nil, zero, e
		}
		nodes = append(nodes, a)
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

	for ctx.Peek().Token == token.DELIM {
		ctx.Next()

		if ex, e = expectExpr(ctx); e != nil {
			return nil, zero, e
		}
		nodes = append(nodes, ex)
	}

	return nodes, snip, nil
}

// Assumes SPELL ...
// Parses: SPELL L_PAREN [<expr> {DELIM <expr>}] R_PAREN
func spellCall(ctx *context) (tree.Node, error) {

	var e error
	sp := ctx.Next()
	n := tree.SpellCall{
		Name: sp.Val[1:],
	}

	if n.Args, e = expectParams(ctx); e != nil {
		return nil, e
	}

	return n, nil
}

func spellCallExpr(ctx *context) (tree.Expr, error) {
	n, e := spellCall(ctx)
	return n.(tree.Expr), e
}

// Parses: {<assign> | <expr>}
func blockStatements(ctx *context) ([]tree.Node, error) {

	nodes := []tree.Node{}

	for ctx.More() && ctx.Peek().Token != token.R_CURLY {
		n, e := statement(ctx)
		if e != nil {
			return nil, e
		}
		nodes = append(nodes, n)
		expectTerminator(ctx)
	}

	return nodes, nil
}

// guard: [<expr>] L_CURLY {<assign> | <expr>} R_CURLY
func guard(ctx *context) (tree.Guard, error) {

	var e error
	var zero, g tree.Guard

	if !ctx.More() {
		return zero, errPos(ctx.End(), "Missing left square brace")
	}

	if l := ctx.Next(); l.Token != token.L_SQUARE {
		return zero, errSnip(l.Snippet,
			"Expected left square brace but got %s", l.Token.String())
	}

	if g.Cond, e = expectExpr(ctx); e != nil {
		return zero, e
	}

	if l := ctx.Next(); l.Token != token.R_SQUARE {
		return zero, errSnip(l.Snippet,
			"Expected right square brace but got %s", l.Token.String())
	}

	if g.Body, e = expectBlock(ctx); e != nil {
		return zero, e
	}

	return g, nil
}
