package recursive

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
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
		p.error("expect", p.itr.Peek(), lex)
	}
}

func (p *parser) error(f string, tk lexeme.Token, expected lexeme.Lexeme) {
	msg := "[parser." + f + "] "
	msg += "Expected " + string(expected) + ", got " + tk.String()
	panic(string(msg))
}
