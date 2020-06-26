package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

func statements(p *pipeline) ([]Expr, error) {

	r := []Expr{}

	for p.hasMore() {

		st, e := statement(p)
		if e != nil {
			return nil, e
		}

		r = append(r, st)
	}

	return r, nil
}

func statement(p *pipeline) (st Expr, e error) {
	// pattern := assignment | expression

	switch {
	case p.match(TK_DEFINITION):
		st, e = assignment(p)

	case p.match(TK_IDENTIFIER):

		p.next()
		if p.match(TK_ASSIGNMENT) || p.match(TK_DELIMITER) {
			p.backup()
			st, e = assignment(p)
			break
		}

		p.backup()

	case p.match(TK_GUARD_OPEN):
		st, e = guard(p)

	case p.match(TK_WATCH):
		st, e = watch(p)

	case p.match(TK_WHEN):
		st, e = when(p)

	case p.match(TK_LOOP):
		st, e = loop(p)

	case p.match(TK_VOID):
		st, e = assignment(p)

	case p.match(TK_EXIT):
		st, e = exit(p)

	case p.match(TK_SPELL):
		st, e = spellCall(p)

	default:
		st, e = expression(p)
	}

	if e != nil {
		return nil, e
	}

	if st == nil {
		return nil, err.NewBySnippet("Expected statement", p.next())
	}

	_, e = p.expect(TK_TERMINATOR)
	if e != nil {
		return nil, e
	}

	return
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

	targets, e := assignTargets(p)
	if e != nil {
		return nil, e
	}

	_, e = p.expect(TK_ASSIGNMENT)
	if e != nil {
		return nil, e
	}

	sources, e := assignSources(p)
	if e != nil {
		return nil, e
	}

	count, e := countAssignments(p, targets, sources)
	if e != nil {
		return nil, e
	}

	return NewAssignBlock(final, targets, sources, count), nil
}

func assignSources(p *pipeline) ([]Expr, error) {

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

func assignTargets(p *pipeline) ([]Expr, error) {
	// pattern := assignmentTarget {DELIMITER assignment_target}

	var ats []Expr

	for {

		at, e := assignTarget(p)
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

func assignTarget(p *pipeline) (Expr, error) {
	// pattern := identifer | VOID

	if p.match(TK_IDENTIFIER) {
		return assignIdentifier(p)
	}

	if p.match(TK_VOID) {
		return NewVoid(p.next()), nil
	}

	return nil, err.NewBySnippet("Expected assignment target", p.next())
}

func assignIdentifier(p *pipeline) (Expr, error) {
	// pattern := IDENTIFIER [list_accessor | function_call]

	tk, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	var id Expr = NewIdentifier(tk)

	for p.match(TK_GUARD_OPEN) {

		id, e = containerItem(p, id)
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
			return 0, err.NewBySnippet("Expected expression", p.next())
		}

		n++
	}

	return n, nil
}
