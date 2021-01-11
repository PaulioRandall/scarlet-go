package parser

import (
	"errors"
	"fmt"

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
func statement(itr LexIterator) (ast.Node, error) {
	switch {
	case itr.Match(token.IDENT):
		return stmtExpr(itr)
	default:
		return nil, err(itr, "Unknown statement")
	}
}

// DEFINE = IDENT {"," IDENT} ":=" EXPR {"," EXPR}
// ASSIGN = IDENT {"," IDENT} "<-" EXPR {"," EXPR}
// EXPR
func stmtExpr(itr LexIterator) (ast.Node, error) {
	// TODO
	return nil, nil
}

// EXPR = LITERAL | IDENT
func expr(itr LexIterator) (ast.Node, error) {
	// TODO
	return nil, nil
}

// LITERAL = BOOL | NUMBER | STRING
func literal(itr LexIterator) (ast.Node, error) {
	// TODO
	return nil, nil
}

// IDENT
func ident(id token.Lexeme) ast.Ident {
	return ast.Ident{
		Snip: id.Snippet,
		Lex:  id,
	}
}

func err(itr LexIterator, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", itr.Line(), m)
	return errors.New(m)
}
