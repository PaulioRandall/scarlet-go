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

func expectStatement(p *parser) (Statement, error) {

	st, e := statement(p)

	if e == nil && st == nil {
		return nil, err.New("Expected statement", err.At(p.any()))
	}

	return st, e
}

func statement(p *parser) (Statement, error) {

	switch {
	case p.match(IDENTIFIER):
		left := p.any()
		return assignOrExpr(p, left)

	case p.match(VOID):
		return assignment(p, p.any())
	}

	return expression(p)
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

	_, e = p.expect(ASSIGN)
	if e != nil {
		return nil, e
	}

	return assignmentExprs(p, tks)
}

func assignmentTokens(p *parser, left Token) ([]Token, error) {

	tks := []Token{left}

	for p.accept(DELIMITER) {

		tk, e := p.expectAnyOf(IDENTIFIER, VOID)
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
			_, e := p.expect(DELIMITER)
			if e != nil {
				return nil, e
			}
		}

		expr, e := expectExpression(p)
		if e != nil {
			return nil, e
		}

		as[i] = p.NewAssignment(tk, expr)
	}

	return p.NewAssignmentBlock(as), nil
}
