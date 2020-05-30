package recursive

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type pipe struct {
	tks  []Token
	size int
	pos  int
}

func newPipe(tks []Token) *pipe {
	return &pipe{tks, len(tks), 0}
}

func (p *pipe) hasMore() bool {
	return p.pos < p.size
}

func (p *pipe) match(m Morpheme) bool {

	tk := p._peek()

	if tk == nil {
		return false
	}

	return m == ANY || m == tk.Morpheme()
}

func (p *pipe) any() Token {
	return p._next()
}

func (p *pipe) accept(m Morpheme) bool {

	tk := p._peek()

	if tk == nil {
		return false
	}

	if m == ANY || m == tk.Morpheme() {
		p._next()
		return true
	}

	return false
}

func (p *pipe) expect(exp Morpheme) (Token, error) {

	if p.accept(exp) {
		return p._prev(), nil
	}

	tk := p._peek()

	if tk == nil {
		s := fmt.Sprintf("Expected %s, got UNDEFINED", exp.String())
		return nil, err.New(s, err.After(p._prev()))
	}

	s := fmt.Sprintf(
		"Expected %s, got %s",
		exp.String(), tk.Morpheme().String(),
	)

	return nil, err.New(s, err.At(tk))
}

func (p *pipe) _peek() Token {

	if p.pos >= p.size {
		return nil
	}

	return p.tks[p.pos]
}

func (p *pipe) _next() Token {

	tk := p._peek()

	if tk != nil {
		p.pos++
	}

	return tk
}

func (p *pipe) _prev() Token {

	if p.pos > 0 {
		return p.tks[p.pos-1]
	}

	return nil
}

//****************************************************************************

/*
func (p *pipe) matchAny(ms ...Morpheme) bool {
	for _, t := range ms {
		if p.match(t) {
			return true
		}
	}

	return false
}

func (p *pipe) matchSequence(ms ...Morpheme) bool {

	count := 0

	defer func() { // Undo all calls to p.itr.Next
		for ; count > 0; count-- {
			p.itr.Back()
		}
	}()

	for _, m := range ms {

		if m == ANY || p.match(m) {
			p.itr.Next()
			count++
			continue
		}

		return false
	}

	return true
}


func (p *pipe) expectOneOf(tag string, ms ...Morpheme) Token {
	for _, m := range ms {
		if p.accept(m) {
			return p.itr.Past()
		}
	}

	s := ""
	for i, m := range ms {
		if i != 0 {
			s += " "
		}

		s += m.String()
	}

	err.Panic(errMsg(tag, s, p.peek()), err.At(p.peek()))
	return nil
}
*/
