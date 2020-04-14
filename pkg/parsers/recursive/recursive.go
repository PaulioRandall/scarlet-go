package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func ParseAll(tks []lexeme.Token) Statements {
	p := parser{
		itr: lexeme.NewIterator(tks),
	}
	return Statements(p.statements())
}

type parser struct {
	itr *lexeme.TokenIterator
	tk  lexeme.Token
}

func (p *parser) statements() []Statement {

	var ss []Statement

	for !p.itr.Empty() && !p.accept(lexeme.LEXEME_EOF) {
		s := p.statement()
		ss = append(ss, s)
	}

	return ss
}

func (p *parser) statement() Statement {

	s := Statement{}

	p.accept(lexeme.LEXEME_ANY)

	if p.was(lexeme.LEXEME_ID) {
		if p.assignment(&s) {
			return s
		}
	}

	if !p.expressions(&s, true) {
		panic(unexpected("statement", p.tk, lexeme.LEXEME_ANY))
	}

	return s
}

func (p *parser) assignment(s *Statement) bool {

	// a, b, c
	for p.inspect(lexeme.LEXEME_DELIM) {
		s.IDs = append(s.IDs, p.tk)   // Append first or the previous iterations ID
		p.expect(lexeme.LEXEME_DELIM) // Skip delimitier
		p.expect(lexeme.LEXEME_ID)    // Must be an ID after a delimitier
	}

	if !p.inspect(lexeme.LEXEME_ASSIGN) {
		return false
	}

	// ? :=
	s.IDs = append(s.IDs, p.tk)    // Append last ID
	p.expect(lexeme.LEXEME_ASSIGN) // Load assignment
	s.Assign = p.tk                // Store assignment

	p.expressions(s, false)
	return true
}

func (p *parser) expressions(s *Statement, loaded bool) bool {

	var found bool

	if !loaded {
		p.accept(lexeme.LEXEME_ANY)
	}

	for p.expression(s, true, true) {
		found = true

		if !p.accept(lexeme.LEXEME_DELIM) {
			break
		}

		p.expect(lexeme.LEXEME_ANY)
	}

	p.expect(lexeme.LEXEME_TERMINATOR)
	return found
}

func (p *parser) expression(s *Statement, loaded, required bool) bool {
	switch {
	case p.factor(s, loaded, required):
	default:
		return false
	}

	return true
}

func (p *parser) factor(s *Statement, loaded, required bool) bool {

	if !loaded {
		p.accept(lexeme.LEXEME_ANY)
	}

	switch {
	case p.was(lexeme.LEXEME_VOID),
		p.was(lexeme.LEXEME_BOOL),
		p.was(lexeme.LEXEME_INT),
		p.was(lexeme.LEXEME_FLOAT),
		p.was(lexeme.LEXEME_STRING),
		p.was(lexeme.LEXEME_TEMPLATE):

		s.Exprs = append(s.Exprs, Value{p.tk})

	default:
		if required {
			panic(unexpected("factor", p.tk, lexeme.LEXEME_ANOTHER))
		}

		return false
	}

	return true
}
