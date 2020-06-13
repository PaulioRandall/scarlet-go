package esmerelda

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

func (p *pipeline) match(ty TokenType) bool {

	tk := p._peek()
	if tk == nil {
		return false
	}

	return ty == TK_ANY || ty == tk.Type()
}

func (p *pipeline) matchSequence(tys ...TokenType) bool {
	for i, ty := range tys {

		tk := p._at(i)
		if tk == nil {
			return false
		}

		if ty != TK_ANY && ty != tk.Type() {
			return false
		}
	}

	return true
}

func (p *pipeline) peek() Token {
	return p._peek()
}

func (p *pipeline) any() Token {
	return p._next()
}

func (p *pipeline) accept(ty TokenType) bool {

	tk := p._peek()
	if tk == nil {
		return false
	}

	if ty == TK_ANY || ty == tk.Type() {
		p._next()
		return true
	}

	return false
}

func (p *pipeline) expect(exp TokenType) (Token, error) {

	if p.accept(exp) {
		return p.prev, nil
	}

	if p.hasMore() {
		return nil, p._unexpected(p._peek(), exp.String())
	}

	return nil, p._outOfTokens(p.prev, exp.String())
}

func (p *pipeline) expectAnyOf(exp ...TokenType) (Token, error) {

	for _, m := range exp {
		if p.accept(m) {
			return p.prev, nil
		}
	}

	if p.hasMore() {
		return nil, p._unexpected(p._peek(), JoinTypes(exp...))
	}

	return nil, p._outOfTokens(p.prev, JoinTypes(exp...))
}

func (p *pipeline) _peek() Token {

	p._ignoreRedundant()
	if p.pos >= p.size {
		return nil
	}

	return p.tks[p.pos]
}

func (p *pipeline) _at(i int) Token {

	p._ignoreRedundant()
	if p.pos+i >= p.size {
		return nil
	}

	return p.tks[p.pos+i]
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

		next := p.tks[p.pos].Type()

		switch {
		case next == TK_COMMENT:
			p.pos++

		case next == TK_WHITESPACE:
			p.pos++

		case next != TK_TERMINATOR:
			return

			// next must be a TERMINATOR
		case p.prev == nil: // Ignore TERMINATORs at start of script
			p.pos++

		case p.prev.Type() == TK_DELIMITER: // Allow "NEWLINE" after delimiter
			p.pos++

		case p.prev.Type() == TK_BLOCK_OPEN: // Allow "NEWLINE" after block start
			p.pos++

		case p.prev.Type() == TK_PAREN_OPEN: // Allow "NEWLINE" after paren start
			p.pos++

		case p.prev.Type() == TK_TERMINATOR: // Ignore successive TERMINATORs
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
	s := fmt.Sprintf("Expected %s; got %s", exp, next.Type().String())
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
