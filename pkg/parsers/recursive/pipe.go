package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// pipe encapsulates a TokenIterator to provide general iterator functionality
// useful for a parser.
type pipe struct {
	itr *token.TokenIterator
}

// next is used to move forward in the pipe. You are responsible for checking
// the token being returned is what you want.
func (p *pipe) next() token.Token {
	return p.itr.Next()
}

// peek is used to obtain the next token without moving forward in the pipe.
func (p *pipe) peek() token.Token {
	return p.itr.Peek()
}

// past is used when you want the previous token.
func (p *pipe) past() token.Token {
	return p.itr.Past()
}

// match is used to confirm the next token is of a specific type.
func (p *pipe) match(typ token.TokenType) bool {
	t := p.itr.Peek().Type
	return t == token.ANY || t == typ
}

// matchAny is used to confirm the next token is any from the slice of types.
func (p *pipe) matchAny(types ...token.TokenType) bool {
	for _, t := range types {
		if p.match(t) {
			return true
		}
	}

	return false
}

// matchSequence is used to determine if the next sequence of tokens matches the
// the input sequence.
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

// accept is used to proceed to the next token if it is of the specified type.
// token.ANY may be used if any non-zero token is acceptable.
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

// expect is used to proceed to the next token if it is of the specified type.
// token.ANY may be used if any non-zero token is acceptable. A panic ensues if
// your demands are not met. The tag is used when printing an error.
func (p *pipe) expect(tag string, lex token.TokenType) token.Token {
	if !p.accept(lex) {
		panic(unexpected(tag, p.itr.Peek(), lex))
	}

	return p.itr.Past()
}
