package gytha

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

// pipe encapsulates a TokenIterator providing general iterator functionality
// useful when parsing.
type pipe struct {
	itr *TokenIterator
}

func (p *pipe) next() Token {
	return p.itr.Next()
}

func (p *pipe) peek() Token {
	return p.itr.Peek()
}

func (p *pipe) past() Token {
	return p.itr.Past() // Previous token, no iteration
}

func (p *pipe) back() {
	p.itr.Back()
}

func (p *pipe) match(ty TokenType) bool {
	if p.itr.Peek() == nil {
		return false
	}

	o := p.itr.Peek().Type()
	return o == TK_ANY || o == ty
}

func (p *pipe) matchAny(tys ...TokenType) bool {
	for _, ty := range tys {
		if p.match(ty) {
			return true
		}
	}

	return false
}

func (p *pipe) matchSequence(tys ...TokenType) bool {

	count := 0

	defer func() { // Undo all calls to p.itr.Next
		for ; count > 0; count-- {
			p.itr.Back()
		}
	}()

	for _, ty := range tys {

		if ty == TK_ANY || p.match(ty) {
			p.itr.Next()
			count++
			continue
		}

		return false
	}

	return true
}

func (p *pipe) accept(ty TokenType) bool {

	if ty == TK_UNDEFINED {
		return false
	}

	if ty == TK_ANY || p.match(ty) {
		p.itr.Skip()
		return true
	}

	return false
}

func (p *pipe) expect(tag string, ty TokenType) Token {
	if !p.accept(ty) {
		err.Panic(
			errMsg(tag, ty.String(), p.peek()),
			err.At(p.peek()),
		)
	}

	return p.itr.Past()
}

func (p *pipe) expectOneOf(tag string, tys ...TokenType) Token {
	for _, ty := range tys {
		if p.accept(ty) {
			return p.itr.Past()
		}
	}

	s := ""
	for i, ty := range tys {
		if i != 0 {
			s += " "
		}

		s += ty.String()
	}

	err.Panic(errMsg(tag, s, p.peek()), err.At(p.peek()))
	return nil
}
