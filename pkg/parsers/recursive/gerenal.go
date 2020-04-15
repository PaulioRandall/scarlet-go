package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// inspect is used if you want to confirm the next token is of a specific token
// type.
func (p *parser) inspect(lex token.TokenType) bool {
	if p.itr.Peek().Type == lex {
		return true
	}

	return false
}

// accept is used when you want to load the next token into p.tk but only if it
// is of the specified kind. ANY may be used if any non-zero token is
// acceptable.
func (p *parser) accept(lex token.TokenType) bool {

	if lex == token.UNDEFINED {
		return false
	}

	if lex == token.ANY || p.inspect(lex) {
		p.tk = p.itr.Next()
		return true
	}

	return false
}

// expect is used when you want to load the next token into p.tk and it must be
// of the specified kind. A panic ensues if your demands are not met.
func (p *parser) expect(lex token.TokenType) {
	if !p.accept(lex) {
		panic(unexpected("expect", p.itr.Peek(), lex))
	}
}

// snoop is used when you want to see the next token without loading it into
// p.tk.
func (p *parser) snoop() token.Token {
	return p.itr.Peek()
}

// proceed is used when you've already checked the next token is of an
// acceptable type and want it loaded into p.tk. It may as well be returned
// since you're likely to use it immediately.
func (p *parser) proceed() token.Token {
	p.tk = p.itr.Next()
	return p.tk
}

// confirm is used when you want to check if p.tk is of a specific token type.
func (p *parser) confirm(lex token.TokenType) bool {
	if lex == token.ANY {
		return p.tk.Type != token.UNDEFINED
	}

	return p.tk.Type == lex
}
