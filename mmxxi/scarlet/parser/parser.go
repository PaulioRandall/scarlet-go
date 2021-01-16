package parser

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

// ParseTree is a recursion based tokeniser. It returns an AST and another
// ParseTree function to obtain the following AST. On error or while
// obtaining the last AST, ParseTree will be nil.
type ParseTree func() (ast.Tree, ParseTree, error)

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
	if !itr.More() {
		return nil
	}

	return func() (ast.Tree, ParseTree, error) {
		t, e := parseNext(itr)
		if e != nil {
			return ast.Tree{}, nil, e
		}
		return t, nextFunc(itr), nil
	}
}

func parseNext(itr LexIterator) (ast.Tree, error) {
	stmt, e := terminatedStatement(itr)
	if e != nil {
		return ast.Tree{}, e
	}
	return ast.Tree{Root: stmt}, nil
}

// TERMIN_STMT = STMT TERMINATOR
func terminatedStatement(itr LexIterator) (n ast.Stmt, e error) {

	s, e := statement(itr)
	if e != nil {
		return nil, e
	}

	if !itr.Accept(token.TERMINATOR) {
		return nil, err(itr, "Expected TERMINATOR")
	}

	return s, nil
}

// STMT = DEFINE | ASSIGN
func statement(itr LexIterator) (n ast.Stmt, e error) {
	switch {
	case !itr.More():
		return nil, err(itr, "Expected statement")

	case itr.Match(token.IDENT) && itr.InRange(1) && itr.At(1).IsType():
		fallthrough
	case itr.MatchPat(token.IDENT, token.DEFINE),
		itr.MatchPat(token.IDENT, token.ASSIGN),
		itr.MatchPat(token.IDENT, token.DELIM):
		return binding(itr)

	default:
		return nil, err(itr, "Unknown statement type")
	}
}

// DEC_IDENT = IDENT [TYPE]
func decIdent(itr LexIterator) (ast.Ident, error) {

	zero := ast.Ident{}

	if !itr.More() || !itr.Match(token.IDENT) {
		return zero, err(itr, "Expected IDENT")
	}
	v := itr.Read()

	t := ast.T_INFER
	if itr.More() && itr.Peek().IsType() {
		switch lx := itr.Read(); lx.Token {
		case token.T_BOOL:
			t = ast.T_BOOL
		case token.T_NUM:
			t = ast.T_NUM
		case token.T_STR:
			t = ast.T_STR
		default:
			return zero, errLex(lx, "Unknown type")
		}
	}

	return ast.MakeIdent(v, t), nil
}

// IDENT_LIST = DEC_IDENT {"," DEC_IDENT}
func identList(itr LexIterator) ([]ast.Ident, error) {

	var ids []ast.Ident
	readIdent := func() error {
		id, e := decIdent(itr)
		if e != nil {
			return e
		}
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

// DEFINE = IDENT_LIST ":=" EXPR_LIST
// ASSIGN = IDENT_LIST "<-" EXPR_LIST
func binding(itr LexIterator) (ast.Binding, error) {

	var zero ast.Binding

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

	return ast.MakeBinding(ids, op, exprs), nil
}

// EXPR_LIST = EXPR {"," EXPR}
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

// EXPR = IDENT | LITERAL
func expression(itr LexIterator) (ast.Expr, error) {
	switch {
	case !itr.More():
		return nil, err(itr, "Expected EXPR")

	case itr.Match(token.IDENT):
		return ast.MakeIdent(itr.Read(), ast.T_INFER), nil

	case itr.MatchAny(token.BOOL, token.NUM, token.STR):
		return ast.MakeLiteral(itr.Read()), nil

	default:
		return nil, err(itr, "Expected EXPR")
	}
}

// LITERAL = BOOL | NUMBER | STRING
func literal(itr LexIterator) (ast.Node, error) {
	if itr.MatchAny(token.BOOL, token.NUM, token.STR) {
		return ast.MakeLiteral(itr.Read()), nil
	}
	return nil, err(itr, "Expected LITERAL")
}

func err(itr LexIterator, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", itr.Line(), m)
	return errors.New(m)
}

func errNode(n ast.Node, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", n.Snippet().Start.Line, m)
	return errors.New(m)
}

func errLex(lx token.Lexeme, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", lx.Snippet.Start.Line, m)
	return errors.New(m)
}
