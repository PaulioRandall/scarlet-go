package recursive

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type pipeline struct {
	tks  []Token
	size int
	pos  int
	prev Token
}

func newPipeline(tks []Token) *pipeline {
	return &pipeline{tks, len(tks), 0, nil}
}

func (p *pipeline) hasMore() bool {
	p._ignoreRedundant()
	return p.pos < p.size
}

func (p *pipeline) match(m Morpheme) bool {

	tk := p._peek()
	if tk == nil {
		return false
	}

	return m == ANY || m == tk.Morpheme()
}

func (p *pipeline) any() Token {
	return p._next()
}

func (p *pipeline) accept(m Morpheme) bool {

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

func (p *pipeline) expect(exp Morpheme) (Token, error) {

	if p.accept(exp) {
		return p.prev, nil
	}

	if p.hasMore() {
		return nil, p._unexpected(p._peek(), exp.String())
	}

	return nil, p._outOfTokens(p.prev, exp.String())
}

func (p *pipeline) expectAnyOf(exp ...Morpheme) (Token, error) {

	for _, m := range exp {
		if p.accept(m) {
			return p.prev, nil
		}
	}

	if p.hasMore() {
		return nil, p._unexpected(p._peek(), JoinMorphemes(exp...))
	}

	return nil, p._outOfTokens(p.prev, JoinMorphemes(exp...))
}

func (p *pipeline) _peek() Token {

	p._ignoreRedundant()
	if p.pos >= p.size {
		return nil
	}

	return p.tks[p.pos]
}

func (p *pipeline) _next() Token {

	tk := p._peek()
	if tk == nil {
		return nil
	}

	p.pos++
	p.prev = tk
	return tk
}

func (p *pipeline) _ignoreRedundant() {

	for p.pos < p.size {

		next := p.tks[p.pos].Morpheme()

		switch {
		case next == COMMENT:
			p.pos++

		case next == WHITESPACE:
			p.pos++

		case next != TERMINATOR:
			return

			// next must be a TERMINATOR
		case p.prev == nil: // Ignore TERMINATORs at start of script
			p.pos++

		case p.prev.Morpheme() == TERMINATOR: // Ignore successive TERMINATORs
			p.pos++

		default: // TERMINATOR
			return
		}
	}
}

func (p *pipeline) _outOfTokens(prev Token, exp string) error {
	s := fmt.Sprintf("Expected %s; got UNDEFINED", exp)
	return err.New(s, err.After(prev))
}

func (p *pipeline) _unexpected(next Token, exp string) error {
	s := fmt.Sprintf("Expected %s; got %s", exp, next.Morpheme().String())
	return err.New(s, err.At(next))
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
