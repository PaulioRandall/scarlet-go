package parser

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
)

// Pattern: IDENT
func expectIdent(l token.Lexeme) (id tree.Ident, e error) {

	if l.Token != token.IDENT {
		e = errSnip(l.Snippet, "Expected identifier")
		return
	}

	id = tree.Ident{
		Snippet: l.Snippet,
		Val:     l.Val,
	}
	return
}

// Parses: BOOL | NUMBER | STRING
func expectLiteral(ctx *context) (tree.Expr, error) {

	if !ctx.More() {
		return nil, errPos(ctx.End(), "Expected expression but reached EOF")
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

// Parses: <terminator>
func expectTerminator(ctx *context) error {
	if !ctx.More() {
		return errPos(ctx.End(), "Expected terminator but reached EOF")
	}
	if tk := ctx.Next(); !tk.IsTerminator() {
		return errSnip(tk.Snippet, "Expected expression but reached EOF")
	}
	return nil
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

	if ctx.Peek().Token == token.L_PAREN {
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
		return nil, errPos(ctx.End(), "Missing left parenthesis")
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
		return nil, errPos(ctx.End(), "Missing right parenthesis")
	}

	if l := ctx.Next(); l.Token != token.R_PAREN {
		return nil, errSnip(l.Snippet,
			"Expected right parenthesis but got %s", l.Token.String())
	}

	return r, nil
}

// Parses: {<operator> <expr>}
func maybeBinaryExpr(ctx *context, left tree.Expr, leftOpPrec int) (tree.Expr, error) {

	if !ctx.Peek().IsBinaryOperator() {
		return left, nil
	}

	if leftOpPrec >= ctx.Peek().Precedence() {
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
		Snippet: token.SuperSnippet(left.Pos(), right.Pos()),
	}

	return maybeBinaryExpr(ctx, binExpr, leftOpPrec)
}
