package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"
)

func (p *parser) inspect(lex lexeme.Lexeme) bool {
	if p.itr.Peek().Lexeme == lex {
		return true
	}

	return false
}

func (p *parser) accept(lex lexeme.Lexeme) bool {

	if lex == lexeme.LEXEME_UNDEFINED {
		return false
	}

	if lex == lexeme.LEXEME_ANY || p.inspect(lex) {
		p.tk = p.itr.Next()
		return true
	}

	return false
}

func (p *parser) was(lex lexeme.Lexeme) bool {
	if lex == lexeme.LEXEME_ANY {
		return p.tk.Lexeme != lexeme.LEXEME_UNDEFINED
	}

	return p.tk.Lexeme == lex
}

func (p *parser) expect(lex lexeme.Lexeme) {
	if !p.accept(lex) {
		panic(unexpected("expect", p.itr.Peek(), lex))
	}
}
