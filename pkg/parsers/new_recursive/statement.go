package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func ParseStatements(tks []Token) ([]Expression, error) {
	p := newPipeline(tks)
	return statements(p)
}

func statements(p *pipeline) ([]Expression, error) {

	r := []Expression{}

	for p.hasMore() {

		st, e := expectStatement(p)
		if e != nil {
			return nil, e
		}

		r = append(r, st)

		_, e = p.expect(TK_TERMINATOR)
		if e != nil {
			return nil, e
		}
	}

	return r, nil
}

func expectStatement(p *pipeline) (Expression, error) {

	st, e := statement(p)

	if e == nil && st == nil {
		return nil, err.New("Expected statement", err.At(p.any()))
	}

	return st, e
}

func statement(p *pipeline) (Expression, error) {
	// pattern := assignment | expression

	switch {
	case p.matchSequence(TK_IDENTIFIER, TK_ASSIGNMENT):
		fallthrough
	case p.matchSequence(TK_IDENTIFIER, TK_DELIMITER):
		return assignment(p)

	case p.match(TK_GUARD_OPEN):
		return guard(p)

	case p.match(TK_WATCH):
		return watch(p)

	case p.match(TK_WHEN):
		return when(p)

	case p.match(TK_LOOP):
		return loop(p)

	case p.match(TK_VOID):
		return assignment(p)
	}

	return operation(p)
}

func assignment(p *pipeline) (Expression, error) {
	// pattern := assignment_targets ASSIGN assignment_sources

	targets, e := assignmentTargets(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_ASSIGNMENT)
	if e != nil {
		return nil, e
	}

	sources, e := assignmentSources(p)
	if e != nil {
		return nil, e
	}

	r, e := createAssignments(p, targets, sources)
	if e != nil {
		return nil, e
	}

	return newAssignmentBlock(r), nil
}

func assignmentSources(p *pipeline) ([]Expression, error) {

	if p.match(TK_FUNCTION) {

		src, e := function(p)
		if e != nil {
			return nil, e
		}

		return []Expression{src}, nil
	}

	if p.match(TK_EXPR_FUNC) {

		src, e := expressionFunction(p)
		if e != nil {
			return nil, e
		}

		return []Expression{src}, nil
	}

	return operations(p)
}

func assignmentTargets(p *pipeline) ([]Expression, error) {
	// pattern := assignmentTarget {DELIMITER assignment_target}

	var ats []Expression

	for first := true; first || p.accept(TK_DELIMITER); first = false {

		at, e := assignmentTarget(p)
		if e != nil {
			return nil, e
		}

		ats = append(ats, at)
	}

	return ats, nil
}

func assignmentTarget(p *pipeline) (Expression, error) {
	// pattern := identifer | VOID

	switch {
	case p.match(TK_IDENTIFIER):
		return identifier(p)

	case p.match(TK_VOID):
		return newVoid(p.any()), nil
	}

	return nil, err.New("Expected assignment target", err.At(p.any()))
}

func createAssignments(p *pipeline, targets, sources []Expression) ([]Assignment, error) {

	var r []Assignment

	for i := 0; i < len(targets) || i < len(sources); i++ {

		if i >= len(targets) {
			line, col := sources[i].Begin()
			return nil, err.New("Too many expressions", err.Pos(line, col))
		}

		if i >= len(sources) {
			return nil, err.New("Expected expression", err.At(p.any()))
		}

		a := newAssignment(targets[i], sources[i])
		r = append(r, a)
	}

	return r, nil
}
