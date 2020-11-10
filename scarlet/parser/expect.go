package parser

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
)

// Pattern: BOOL
func boolLit(l token.Lexeme) tree.BoolLit {
	return tree.BoolLit{
		Val: l.Token == token.TRUE,
	}
}

// Pattern: NUMBER
func numLit(l token.Lexeme) tree.NumLit {
	v, e := strconv.ParseFloat(l.Val, 64)
	if e != nil {
		panic("SANITY CHECK! Invalid number, should have been detected prior")
	}
	return tree.NumLit{Val: v}
}

// Pattern: STRING
func strLit(l token.Lexeme) tree.StrLit {
	return tree.StrLit{Val: l.Val}
}

// Pattern: IDENT || VOID
func expectAssignee(l token.Lexeme) (a tree.Assignee, e error) {
	switch l.Token {
	case token.IDENT:
		a = tree.Ident{Val: l.Val}
	case token.VOID:
		a = tree.AnonIdent{}
	default:
		e = errSnip(l.Snippet, "Expected identifier")
	}
	return
}

// Pattern: IDENT
func expectIdent(l token.Lexeme) (id tree.Ident, e error) {

	if l.Token != token.IDENT {
		e = errSnip(l.Snippet, "Expected identifier")
		return
	}

	id = tree.Ident{Val: l.Val}
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
	if ctx.Peek().Token == token.R_CURLY {
		return nil
	}
	if tk := ctx.Next(); !tk.IsTerminator() {
		return errSnip(tk.Snippet, "Expected terminator but got %v", tk.Token)
	}
	return nil
}

// Parses:  <expr> {<operator> <expr>}
// Parses:  L_PAREN <expr> {<operator> <expr>} R_PAREN
func expectExpr(ctx *context) (tree.Expr, error) {
	return expectExprRight(ctx, 0)
}

// Parses: <term>
func expectTerm(ctx *context) (ex tree.Expr, e error) {
	switch {
	case !ctx.More():
		e = errPos(ctx.End(), "Expected term")

	case ctx.LexItr.Peek().IsLiteral():
		ex, e = expectLiteral(ctx)

	case ctx.LexItr.Peek().Token == token.IDENT:
		ex, e = expectIdent(ctx.Next())

	case ctx.LexItr.Peek().Token == token.SPELL:
		ex, e = spellCallExpr(ctx)

	default:
		e = errSnip(ctx.LexItr.Peek().Snippet, "Expected term")
	}

	if e != nil {
		return nil, e
	}

	return maybePostUnaryOp(ctx, ex), e
}

func maybePostUnaryOp(ctx *context, left tree.Expr) tree.Expr {
	if !ctx.More() || !ctx.Peek().IsPostUnaryOperator() {
		return left
	}

	switch l := ctx.Next(); l.Token {
	case token.EXIST:
		return tree.UnaryExpr{
			Term: left,
			Op:   tree.OP_EXIST,
		}

	default:
		panic("SANITY CHECK! Unknown post unary operator")
	}
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
		Left:  left,
		Op:    tokenToOperator(op.Token),
		Right: right,
	}

	return maybeBinaryExpr(ctx, binExpr, leftOpPrec)
}

// Parses: L_PAREN [<expr> {DELIM <expr>}] R_PAREN
func expectParams(ctx *context) ([]tree.Expr, error) {

	var (
		l     token.Lexeme
		snip  token.Snippet
		nodes []tree.Expr
		e     error
	)

	if !ctx.More() {
		return nil, errPos(ctx.End(), "Missing parameters")
	}

	if l = ctx.Next(); l.Token != token.L_PAREN {
		return nil, errSnip(l.Snippet,
			"Expected left parenthesis but got %s", l.Token.String())
	}
	snip = l.Snippet

	if ctx.Peek().Token == token.R_PAREN {
		nodes = []tree.Expr{}
	} else {
		if nodes, e = expectParamsSet(ctx); e != nil {
			return nil, e
		}
	}

	if !ctx.More() {
		return nil, errPos(ctx.End(), "Missing right parenthesis")
	}

	if l = ctx.Next(); l.Token != token.R_PAREN {
		return nil, errSnip(l.Snippet,
			"Expected right parenthesis but got %s", l.Token.String())
	}
	snip = token.SuperSnippet(snip, l.Snippet)

	return nodes, nil
}

// Parses:  <expr> {DELIM <expr>}
func expectParamsSet(ctx *context) ([]tree.Expr, error) {

	var (
		nodes = []tree.Expr{}
		ex    tree.Expr
		e     error
	)

	for {
		if ex, e = expectExpr(ctx); e != nil {
			return nil, e
		}
		nodes = append(nodes, ex)

		if ctx.Peek().Token != token.DELIM {
			break
		}
		ctx.Next()
	}

	return nodes, nil
}

// Parsers: L_SQUARE <stmt> R_SQUARE
func expectBlock(ctx *context) (tree.Block, error) {

	var e error
	var zero, b tree.Block

	if !ctx.More() {
		return zero, errPos(ctx.End(), "Missing left curly brace")
	}

	if l := ctx.Next(); l.Token != token.L_CURLY {
		return zero, errSnip(l.Snippet,
			"Expected left curly brace but got %s", l.Token.String())
	}

	if b.Stmts, e = blockStatements(ctx); e != nil {
		return zero, e
	}

	if l := ctx.Next(); l.Token != token.R_CURLY {
		return zero, errSnip(l.Snippet,
			"Expected right curly brace but got %s", l.Token.String())
	}

	return b, nil
}

func tokenToOperator(tk token.Token) tree.Operator {
	switch tk {
	case token.ADD:
		return tree.OP_ADD
	case token.SUB:
		return tree.OP_SUB
	case token.MUL:
		return tree.OP_MUL
	case token.DIV:
		return tree.OP_DIV
	case token.REM:
		return tree.OP_REM

	case token.AND:
		return tree.OP_AND
	case token.OR:
		return tree.OP_OR

	case token.LT:
		return tree.OP_LT
	case token.MT:
		return tree.OP_MT
	case token.LTE:
		return tree.OP_LTE
	case token.MTE:
		return tree.OP_MTE

	case token.EQU:
		return tree.OP_EQU
	case token.NEQ:
		return tree.OP_NEQ

	case token.EXIST:
		return tree.OP_EXIST

	default:
		panic("SANITY CHECK! Unknown operator")
	}
}
