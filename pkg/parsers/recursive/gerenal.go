package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"
)

// inspect is used if you want to confirm the next token is of a specific token
// type.
func (p *parser) inspect(lex lexeme.Lexeme) bool {
	if p.itr.Peek().Lexeme == lex {
		return true
	}

	return false
}

// accept is used when you want to load the next token into p.tk but only if it
// is of the specified kind. LEXEME_ANY may be used if any non-zero token is
// acceptable.
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

// expect is used when you want to load the next token into p.tk and it must be
// of the specified kind. A panic ensues if your demands are not met.
func (p *parser) expect(lex lexeme.Lexeme) {
	if !p.accept(lex) {
		panic(unexpected("expect", p.itr.Peek(), lex))
	}
}

// confirm is used when you want to check if p.tk is of a specific token type.
func (p *parser) confirm(lex lexeme.Lexeme) bool {
	if lex == lexeme.LEXEME_ANY {
		return p.tk.Lexeme != lexeme.LEXEME_UNDEFINED
	}

	return p.tk.Lexeme == lex
}
