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
	// pattern := assignment | expression

	var (
		left Expression
		e    error
	)

	switch {
	case p.match(IDENTIFIER):
		left, e = identifier(p)
		if e != nil {
			return nil, e
		}

		return assignOrExpr(p, left)

	case p.match(VOID):
		left = p.NewVoid(p.any())
		return assignment(p, left)
	}

	return expression(p)
}

func assignOrExpr(p *parser, left Expression) (Statement, error) {

	if p.match(DELIMITER) || p.match(ASSIGN) {
		return assignment(p, left)
	}

	return left, nil
}

func assignment(p *parser, left Expression) (Statement, error) {
	// pattern := assignment_targets ASSIGN assignment_sources

	tks, e := assignmentTargets(p, left)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(ASSIGN)
	if e != nil {
		return nil, e
	}

	return assignmentSources(p, tks)
}

func assignmentTargets(p *parser, left Expression) ([]Expression, error) {
	// pattern := assignmentTarget {DELIMITER assignment_target}

	ats := []Expression{left}

	for p.accept(DELIMITER) {

		at, e := assignmentTarget(p)
		if e != nil {
			return nil, e
		}

		ats = append(ats, at)
	}

	return ats, nil
}

func assignmentTarget(p *parser) (Expression, error) {
	// pattern := IDENTIFIER | list_accessor | VOID

	switch {
	case p.match(IDENTIFIER):
		return identifier(p)

	case p.match(VOID):
		return p.NewVoid(p.any()), nil
	}

	return nil, err.New("Expected assignment target", err.At(p.any()))
}

func assignmentSources(p *parser, ats []Expression) (Statement, error) {
	// pattern := [expression {DELIMITER expression}]

	as := make([]Assignment, len(ats))

	for i, at := range ats {

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

		as[i] = p.NewAssignment(at, expr)
	}

	return p.NewAssignmentBlock(as), nil
}
