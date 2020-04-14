package recursive

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
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

	for !p.itr.Empty() {
		s := p.statement()
		ss = append(ss, s)
	}

	return ss
}

func (p *parser) statement() Statement {

	s := &Statement{}

	for p.accept(lexeme.LEXEME_ANY) {

		if p.was(lexeme.LEXEME_ID) {
			if p.assignment(s) {
				p.accept(lexeme.LEXEME_ANY)
			}
		}

		for p.factor(s, true) {
			p.expect(lexeme.LEXEME_TERMINATOR)
			p.expect(lexeme.LEXEME_ANY)

			if p.was(lexeme.LEXEME_EOF) {
				break
			}
		}

		if p.was(lexeme.LEXEME_EOF) {
			break
		}
	}

	return *s
}

func (p *parser) assignment(s *Statement) bool {

	// a, b, c
	for p.inspect(lexeme.LEXEME_DELIM) {
		s.IDs = append(s.IDs, p.tk)   // Append first or the previous iterations ID
		p.expect(lexeme.LEXEME_DELIM) // Skip delimitier
		p.expect(lexeme.LEXEME_ID)    // Must be an ID after a delimitier
	}

	// ? :=
	if p.inspect(lexeme.LEXEME_ASSIGN) {
		s.IDs = append(s.IDs, p.tk)    // Append last ID
		p.expect(lexeme.LEXEME_ASSIGN) // Load assignment
		s.Assign = p.tk                // Store assignment
		return true
	}

	return false
}

func (p *parser) factor(s *Statement, loaded bool) bool {

	if !loaded {
		p.accept(lexeme.LEXEME_ANY)
	}

	switch {
	case p.was(lexeme.LEXEME_STRING):
		fallthrough
	case p.was(lexeme.LEXEME_TEMPLATE):
		s.Exprs = append(s.Exprs, Value{p.tk})

	default:
		return false
	}

	return true
}
