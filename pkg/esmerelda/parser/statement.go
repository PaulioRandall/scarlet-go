package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

func statements(p *pipeline) ([]Expr, error) {

	r := []Expr{}

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

func statement(p *pipeline) (Expr, error) {
	// pattern := assignment | expression

	switch {
	case p.match(TK_DEFINITION):
		return assignment(p)

	case p.match(TK_IDENTIFIER):
		if p.matchBeyond(TK_ASSIGNMENT) || p.matchBeyond(TK_DELIMITER) {
			return assignment(p)
		}

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

	case p.match(TK_EXIT):
		return exit(p)

	case p.match(TK_SPELL):
		return spellCall(p)
	}

	return expression(p)
}

func expectStatement(p *pipeline) (Expr, error) {

	st, e := statement(p)

	if e == nil && st == nil {
		return nil, err.NewBySnippet("Expected statement", p.any())
	}

	return st, e
}

func exit(p *pipeline) (Expr, error) {

	tk, e := p.expect(TK_EXIT)
	if e != nil {
		return nil, e
	}

	code, e := expectExpr(p)
	if e != nil {
		return nil, e
	}

	return NewExit(tk, code), nil
}

func assignment(p *pipeline) (Expr, error) {
	// pattern := [DEFINITION] assignment_targets ASSIGN assignment_sources

	final := p.accept(TK_DEFINITION)

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

	count, e := countAssignments(p, targets, sources)
	if e != nil {
		return nil, e
	}

	return NewAssignmentBlock(final, targets, sources, count), nil
}

func assignmentSources(p *pipeline) ([]Expr, error) {

	if p.match(TK_FUNCTION) {

		src, e := funcDef(p)
		if e != nil {
			return nil, e
		}

		return []Expr{src}, nil
	}

	if p.match(TK_EXPR_FUNC) {

		src, e := exprFunc(p)
		if e != nil {
			return nil, e
		}

		return []Expr{src}, nil
	}

	return expressions(p)
}

func assignmentTargets(p *pipeline) ([]Expr, error) {
	// pattern := assignmentTarget {DELIMITER assignment_target}

	var ats []Expr

	for {

		at, e := assignmentTarget(p)
		if e != nil {
			return nil, e
		}

		ats = append(ats, at)

		if p.match(TK_ASSIGNMENT) {
			break
		}

		if !p.accept(TK_DELIMITER) {
			break
		}
	}

	return ats, nil
}

func assignmentTarget(p *pipeline) (Expr, error) {
	// pattern := identifer | VOID

	if p.match(TK_IDENTIFIER) {
		return assignmentIdentifier(p)
	}

	if p.match(TK_VOID) {
		return NewVoid(p.any()), nil
	}

	return nil, err.NewBySnippet("Expected assignment target", p.any())
}

func assignmentIdentifier(p *pipeline) (Expr, error) {
	// pattern := IDENTIFIER [list_accessor | function_call]

	tk, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	var id Expr = NewIdentifier(tk)

	for p.match(TK_GUARD_OPEN) {

		id, e = collectionAccessor(p, id)
		if e != nil {
			return nil, e
		}
	}

	return id, nil
}

func countAssignments(p *pipeline, targets, sources []Expr) (int, error) {

	var n int

	for i := 0; i < len(targets) || i < len(sources); i++ {

		if i >= len(targets) {
			return 0, err.NewBySnippet("Too many expressions", sources[i])
		}

		if i >= len(sources) {
			return 0, err.NewBySnippet("Expected expression", p.any())
		}

		n++
	}

	return n, nil
}
