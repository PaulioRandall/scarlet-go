package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// pipe encapsulates a TokenIterator providing general iterator functionality
// useful when parsing.
type pipe struct {
	itr *token.TokenIterator
}

func (p *pipe) next() token.Token {
	return p.itr.Next()
}

func (p *pipe) peek() token.Token {
	return p.itr.Peek()
}

func (p *pipe) past() token.Token {
	return p.itr.Past() // Previous token, no iteration
}

func (p *pipe) back() {
	p.itr.Back()
}

func (p *pipe) match(typ token.TokenType) bool {
	t := p.itr.Peek().Type
	return t == token.ANY || t == typ
}

func (p *pipe) matchAny(types ...token.TokenType) bool {
	for _, t := range types {
		if p.match(t) {
			return true
		}
	}

	return false
}

func (p *pipe) matchSequence(types ...token.TokenType) bool {

	count := 0

	defer func() { // Undo all calls to p.itr.Next
		for ; count > 0; count-- {
			p.itr.Back()
		}
	}()

	for _, t := range types {

		if t == token.ANY || p.match(t) {
			p.itr.Next()
			count++
			continue
		}

		return false
	}

	return true
}

func (p *pipe) accept(lex token.TokenType) bool {

	if lex == token.UNDEFINED {
		return false
	}

	if lex == token.ANY || p.match(lex) {
		p.itr.Skip()
		return true
	}

	return false
}

func (p *pipe) expect(tag string, lex token.TokenType) token.Token {
	if !p.accept(lex) {
		panic(unexpected(tag, p.itr.Peek(), lex))
	}

	return p.itr.Past()
}
