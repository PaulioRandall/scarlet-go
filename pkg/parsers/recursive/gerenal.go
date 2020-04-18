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
		p.itr.Next()
		return true
	}

	return false
}

// expect is used when you want to load the next token into p.tk and it must be
// of the specified kind. A panic ensues if your demands are not met.
func (p *parser) expect(tag string, lex token.TokenType) bool {
	if !p.accept(lex) {
		panic(unexpected(tag, p.itr.Peek(), lex))
	}

	return true
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
	return p.itr.Next()
}

// confirm is used when you want to check if p.tk is of a specific token type.
func (p *parser) confirm(lex token.TokenType) bool {
	if lex == token.ANY {
		return p.itr.Past().Type != token.UNDEFINED
	}

	return p.itr.Past().Type == lex
}

// affirm is used when you want to assert that p.tk is of a specific token type
// and if not, panic
func (p *parser) affirm(tag string, lex token.TokenType) bool {
	if p.confirm(lex) {
		return true
	}

	panic(unexpected(tag, p.itr.Past(), lex))
}

// Prior is used when you want the last token that was accepted.
func (p *parser) prior() token.Token {
	return p.itr.Past()
}

// retract is used when you want to backtrack one token.
func (p *parser) retract() {
	p.itr.Back()
}
