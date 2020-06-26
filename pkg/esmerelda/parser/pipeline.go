package parser

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type pipeline struct {
	buffer sanitiser
}

func newPipeline(stream TokenStream) *pipeline {
	return &pipeline{
		buffer: newSanitiser(stream),
	}
}

func (p *pipeline) hasMore() bool {
	return !p.buffer.empty()
}

func (p *pipeline) empty() bool {
	return p.buffer.empty()
}

func (p *pipeline) peek() Token {
	return p.buffer.peek()
}

func (p *pipeline) next() Token {
	return p.buffer.next()
}

func (p *pipeline) backup() {
	p.buffer.backup()
}

func (p *pipeline) match(ty TokenType) bool {
	return p._match(ty)
}

func (p *pipeline) accept(ty TokenType) bool {

	if p.buffer.empty() {
		return false
	}

	if ty == TK_ANY || p._match(ty) {
		p.buffer.next()
		return true
	}

	return false
}

func (p *pipeline) expect(ty TokenType) (Token, error) {

	if p.buffer.empty() {
		return nil, p._outOfTokens(p.buffer.past(), ty.String())
	}

	if p._match(ty) {
		return p.buffer.next(), nil
	}

	return nil, p._unexpected(p.buffer.peek(), ty.String())
}

func (p *pipeline) expectAnyOf(tys ...TokenType) (Token, error) {

	if p.buffer.empty() {
		return nil, p._outOfTokens(p.buffer.past(), JoinTypes(tys...))
	}

	for _, ty := range tys {
		if p._match(ty) {
			return p.buffer.next(), nil
		}
	}

	return nil, p._unexpected(p.buffer.peek(), JoinTypes(tys...))
}

func (p *pipeline) _match(ty TokenType) bool {

	tk := p.buffer.peek()
	if tk == nil {
		return false
	}

	return ty == TK_ANY || ty == tk.Type()
}

func (p *pipeline) _outOfTokens(prev Token, exp string) error {
	msg := fmt.Sprintf("Out of tokens, expected %q", exp)
	return err.NewAfterSnippet(msg, prev)
}

func (p *pipeline) _unexpected(next Token, exp string) error {
	msg := fmt.Sprintf("Expected %q; got %q", exp, next.Type().String())
	return err.NewBySnippet(msg, next)
}
