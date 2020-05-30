package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type parser struct {
	*pipeline
	Factory
}

func ParseStatements(fac Factory, tks []Token) ([]Statement, error) {
	p := &parser{newPipeline(tks), fac}
	return statements(p)
}

func statements(p *parser) ([]Statement, error) {

	r := []Statement{}

	for p.hasMore() {

		st, e := expectStatement(p)
		if e != nil {
			return nil, e
		}

		r = append(r, st)

		_, e = p.expect(TERMINATOR)
		if e != nil {
			return nil, e
		}
	}

	return r, nil
}

func statement(p *parser) (Statement, error) {

	switch {
	case p.match(IDENTIFIER):
		left := p.any()
		return assignOrExpr(p, left)

	case p.match(VOID):
		return p.NewIdentifier(p.any()), nil

	case p.match(BOOL), p.match(NUMBER), p.match(STRING):
		return p.NewLiteral(p.any()), nil
	}

	return expression(p)
}

func expectStatement(p *parser) (Statement, error) {

	st, e := statement(p)

	if e == nil && st == nil {
		return nil, err.New("Expected statement", err.At(p.any()))
	}

	return st, e
}

func assignOrExpr(p *parser, left Token) (Statement, error) {

	if p.match(DELIMITER) || p.match(ASSIGN) {
		return assignment(p, left)
	}

	return p.NewIdentifier(left), nil
}

func assignment(p *parser, left Token) (Statement, error) {

	tks, e := assignmentTokens(p, left)
	if e != nil {
		return nil, e
	}

	return assignmentExprs(p, tks)
}

func assignmentTokens(p *parser, left Token) ([]Token, error) {

	tks := []Token{left}

	for p.match(DELIMITER) {

		tk, e := p.expectAny(IDENTIFIER, VOID)
		if e != nil {
			return nil, e
		}

		tks = append(tks, tk)
	}

	return tks, nil
}

func assignmentExprs(p *parser, tks []Token) (Statement, error) {

	as := make([]Assignment, len(tks))

	for i, tk := range tks {

		if i > 0 {
			p.expect(DELIMITER)
		}

		expr, e := expectExpression(p)
		if e != nil {
			return nil, e
		}

		a := p.NewAssignment(tk, expr)
		as = append(as, a)
	}

	return p.NewAssignmentBlock(as), nil
}
