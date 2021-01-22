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
	stmt, e := cappedStatement(itr)
	if e != nil {
		return ast.Tree{}, e
	}
	return ast.Tree{Root: stmt}, nil
}

// cappedStatement = statement TERMINATOR
func cappedStatement(itr LexIterator) (ast.Stmt, error) {

	s, e := statement(itr)
	if e != nil {
		return nil, e
	}

	if !itr.Accept(token.TERMINATOR) {
		return nil, err(itr, "Expected TERMINATOR")
	}

	return s, nil
}

// statement = binding
func statement(itr LexIterator) (ast.Stmt, error) {
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

// binding = varList (":=" | "<-") expressions
func binding(itr LexIterator) (ast.Binding, error) {

	var zero ast.Binding

	vars, e := varList(itr)
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

	return ast.MakeBinding(vars, op, exprs), nil
}

// varList = typedVarList {"," typedVarList}
func varList(itr LexIterator) ([]ast.Var, error) {

	var vars []ast.Var

	for {
		typVars, e := typedVarList(itr)
		if e != nil {
			return nil, e
		}

		vars = append(vars, typVars...)

		if !itr.Accept(token.DELIM) {
			break
		}
	}

	return vars, nil
}

// typedVarList = IDENT {DELIM IDENT} [valType]
func typedVarList(itr LexIterator) ([]ast.Var, error) {

	var vars []ast.Var

	for {
		if !itr.Match(token.IDENT) {
			return nil, err(itr, "Expected IDENT")
		}

		lx := itr.Read()
		v := ast.MakeVar(lx, ast.T_UNDEFINED)
		vars = append(vars, v)

		if !itr.Accept(token.DELIM) {
			break
		}
	}

	t, e := valType(itr)
	if e != nil {
		return nil, e
	}

	for i, _ := range vars {
		vars[i].ValType = t
	}

	return vars, nil
}

// valType = T_BOOL | T_NUM | T_STR
func valType(itr LexIterator) (ast.ValType, error) {

	if !itr.More() || !itr.Peek().IsType() {
		return ast.T_INFER, nil
	}

	switch lx := itr.Read(); lx.Token {
	case token.T_BOOL:
		return ast.T_BOOL, nil
	case token.T_NUM:
		return ast.T_NUM, nil
	case token.T_STR:
		return ast.T_STR, nil
	default:
		return ast.T_UNDEFINED, errLex(lx, "Unknown type")
	}
}

// expressions = expression {"," expression}
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

// expression = IDENT | literal
func expression(itr LexIterator) (ast.Expr, error) {
	switch {
	case !itr.More():
		return nil, err(itr, "Expected EXPR")

	case itr.Match(token.IDENT):
		return ast.MakeIdent(itr.Read(), ast.T_RESOLVE), nil

	case itr.MatchAny(token.BOOL, token.NUM, token.STR):
		return ast.MakeLiteral(itr.Read()), nil

	default:
		return nil, err(itr, "Expected EXPR")
	}
}

// literal = BOOL | NUMBER | STRING
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
