package parser

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

type (
	// ParseTree is a recursion based tokeniser. It returns an AST and another
	// ParseTree function to obtain the following AST. On error or while
	// obtaining the last AST, ParseTree will be nil.
	ParseTree func() (ast.Tree, ParseTree, error)
)

// New returns a new ParseTree function.
func New(itr LexIterator) ParseTree {
	if itr.More() {
		return nextFunc(itr)
	}
	return nil
}

// ParseAll parses all lexemes into a slice of ASTs.
func ParseAll(itr LexIterator) ([]ast.Tree, error) {

	var (
		r  []ast.Tree
		t  ast.Tree
		pt = New(itr)
		e  error
	)

	for pt != nil {
		if t, pt, e = pt(); e != nil {
			return nil, e
		}
		r = append(r, t)
	}

	return r, nil
}

func nextFunc(itr LexIterator) ParseTree {
	return func() (ast.Tree, ParseTree, error) {
		t, e := parse(itr)
		if e != nil {
			return ast.Tree{}, nil, e
		}
		return t, nextFunc(itr), nil
	}
}

func parse(itr LexIterator) (ast.Tree, error) {
	var t ast.Tree

	n, e := statement(itr)
	if e != nil {
		return t, e
	}

	t = ast.Tree{Root: n}
	return t, nil
}

// STMT
func statement(itr LexIterator) (n ast.Node, e error) {
	switch {
	case itr.MatchPat(token.IDENT, token.DEFINE),
		itr.MatchPat(token.IDENT, token.ASSIGN),
		itr.MatchPat(token.IDENT, token.DELIM):
		return defineOrAssign(itr)
	default:
		return nil, err(itr, "Unknown statement")
	}
}

// IDENT_LIST = IDENT {"," IDENT}
func identList(itr LexIterator) ([]ast.Ident, error) {

	var ids []ast.Ident
	readIdent := func() error {
		if !itr.More() || !itr.Match(token.IDENT) {
			return err(itr, "Expected IDENT")
		}
		id := makeIdent(itr.Read())
		ids = append(ids, id)
		return nil
	}

	if e := readIdent(); e != nil {
		return nil, e
	}
	for itr.Accept(token.DELIM) {
		if e := readIdent(); e != nil {
			return nil, e
		}
	}

	return ids, nil
}

// DEFINE = IDENT_LIST ":=" EXPR {"," EXPR}
// ASSIGN = IDENT_LIST "<-" EXPR {"," EXPR}
func defineOrAssign(itr LexIterator) (ast.Assign, error) {

	var zero ast.Assign

	ids, e := identList(itr)
	if e != nil {
		return zero, e
	}

	if !itr.MatchAny(token.DEFINE, token.ASSIGN) {
		return zero, err(itr, "Expected DEFINE or ASSIGN")
	}

	op := itr.Read()
	exprs, e := expressions(itr)
	if e != nil {
		return zero, e
	}

	return makeAssign(ids, op, exprs), nil
}

// ASSIGN = IDENT_LIST "<-" EXPR {"," EXPR}
func assign(itr LexIterator, ids []ast.Ident) (ast.Assign, error) {
	if !itr.Accept(token.ASSIGN) {
		return ast.Assign{}, err(itr, "Expected ASSIGN")
	}
	return ast.Assign{}, nil
}

// EXPRS {"," EXPRS}
func expressions(itr LexIterator) ([]ast.Expr, error) {

	var (
		r  []ast.Expr
		ex ast.Expr
		e  error
	)

	if ex, e = expression(itr); e != nil {
		return nil, e
	}
	r = append(r, ex)

	for itr.Accept(token.DELIM) {
		if ex, e = expression(itr); e != nil {
			return nil, e
		}
		r = append(r, ex)
	}

	return r, nil
}

// EXPR = LITERAL | IDENT
func expression(itr LexIterator) (ast.Expr, error) {
	switch {
	case !itr.More():
		return nil, err(itr, "Expected EXPR")
	case itr.MatchAny(token.BOOL, token.NUM, token.STR):
		return makeLit(itr.Read()), nil
	case itr.Match(token.IDENT):
		return makeIdent(itr.Read()), nil
	default:
		return nil, err(itr, "Expected EXPR")
	}
}

// LITERAL = BOOL | NUMBER | STRING
func literal(itr LexIterator) (ast.Node, error) {
	if itr.MatchAny(token.BOOL, token.NUM, token.STR) {
		return makeLit(itr.Read()), nil
	}
	return nil, err(itr, "Expected LITERAL")
}
