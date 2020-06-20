package esmerelda

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type TokenStream interface {
	Next() Token
	Peek() Token
	PeekBeyond() Token
}

func ParseStatements(tks []Token) ([]Expression, error) {
	p := newPipeline(tks, nil)
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

	case p.match(TK_EXIT):
		return exit(p)

	case p.match(TK_SPELL):
		return spellCall(p)
	}

	return expression(p)
}

func expectStatement(p *pipeline) (Expression, error) {

	st, e := statement(p)

	if e == nil && st == nil {
		return nil, err.New("Expected statement", err.At(p.any()))
	}

	return st, e
}

func exit(p *pipeline) (Expression, error) {

	tk, e := p.expect(TK_EXIT)
	if e != nil {
		return nil, e
	}

	code, e := expectExpression(p)
	if e != nil {
		return nil, e
	}

	return newExit(tk, code), nil
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

	count, e := countAssignments(p, targets, sources)
	if e != nil {
		return nil, e
	}

	return newAssignmentBlock(targets, sources, count), nil
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

	return expressions(p)
}

func assignmentTargets(p *pipeline) ([]Expression, error) {
	// pattern := assignmentTarget {DELIMITER assignment_target}

	var ats []Expression

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

func assignmentTarget(p *pipeline) (Expression, error) {
	// pattern := identifer | VOID

	if p.match(TK_IDENTIFIER) {
		return assignmentIdentifier(p)
	}

	if p.match(TK_VOID) {
		return newVoid(p.any()), nil
	}

	return nil, err.New("Expected assignment target", err.At(p.any()))
}

func assignmentIdentifier(p *pipeline) (Expression, error) {
	// pattern := IDENTIFIER [list_accessor | function_call]

	tk, e := p.expect(TK_IDENTIFIER)
	if e != nil {
		return nil, e
	}

	var id Expression = newIdentifier(tk)

	for p.match(TK_GUARD_OPEN) {

		id, e = collectionAccessor(p, id)
		if e != nil {
			return nil, e
		}
	}

	return id, nil
}

func countAssignments(p *pipeline, targets, sources []Expression) (int, error) {

	var n int

	for i := 0; i < len(targets) || i < len(sources); i++ {

		if i >= len(targets) {
			line, col := sources[i].Begin()
			return 0, err.New("Too many expressions", err.Pos(line, col))
		}

		if i >= len(sources) {
			return 0, err.New("Expected expression", err.At(p.any()))
		}

		n++
	}

	return n, nil
}
