package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// pipe encapsulates a TokenIterator to provide general iterator functionality
// useful for a parser.
type pipe struct {
	itr *token.TokenIterator
}

// ****************************************************************************
// These functions are for moving forward or backward in the pipe.
// ****************************************************************************

// proceed is used to move forward in the pipe. You are responsible for checking
// the token being returned is what you want.
func (p *pipe) proceed() token.Token {
	return p.itr.Next()
}

// retract is used to move backwards in the pipe.
func (p *pipe) retract() {
	p.itr.Back()
}

// ****************************************************************************
// These functions are for inspecting tokens in the pipe without moving forward
// or backward.
// ****************************************************************************

// snoop is used to obtain the next token without moving forward in the pipe.
func (p *pipe) snoop() token.Token {
	return p.itr.Peek()
}

// inspect is used to confirm the next token is of a specific type.
func (p *pipe) inspect(lex token.TokenType) bool {
	if p.itr.Peek().Type == lex {
		return true
	}

	return false
}

// confirm is used to confirm the previous token is of a specific type.
func (p *pipe) confirm(lex token.TokenType) bool {
	if lex == token.ANY {
		return p.itr.Past().Type != token.UNDEFINED
	}

	return p.itr.Past().Type == lex
}

// affirm is used to assert the previous token is of a specific type and panic
// if not. The tag is used when printing an error.
func (p *pipe) affirm(tag string, lex token.TokenType) bool {
	if p.confirm(lex) {
		return true
	}

	panic(unexpected(tag, p.itr.Past(), lex))
}

// Prior is used when you want the previous token.
func (p *pipe) prior() token.Token {
	return p.itr.Past()
}

// ****************************************************************************
// These functions are for moving along in the pipe if particular conditions
// are met.
// ****************************************************************************

// accept is used to proceed to the next token if it is of the specified type.
// token.ANY may be used if any non-zero token is acceptable.
func (p *pipe) accept(lex token.TokenType) bool {

	if lex == token.UNDEFINED {
		return false
	}

	if lex == token.ANY || p.inspect(lex) {
		p.itr.Next()
		return true
	}

	return false
}

// expect is used to proceed to the next token if it is of the specified type.
// token.ANY may be used if any non-zero token is acceptable. A panic ensues if
// your demands are not met. The tag is used when printing an error.
func (p *pipe) expect(tag string, lex token.TokenType) bool {
	if !p.accept(lex) {
		panic(unexpected(tag, p.itr.Peek(), lex))
	}

	return true
}
