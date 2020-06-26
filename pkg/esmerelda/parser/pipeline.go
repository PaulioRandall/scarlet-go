package parser

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type pipeline struct {
	stream TokenStream
	prev   Token
}

func newPipeline(stream TokenStream) *pipeline {
	return &pipeline{stream, nil}
}

func (p *pipeline) _peek() Token {
	p._ignoreRedundant()
	return p.stream.Peek()
}

func (p *pipeline) hasMore() bool {
	return p._peek() != nil
}

func (p *pipeline) match(ty TokenType) bool {

	tk := p._peek()
	if tk == nil {
		return false
	}

	return ty == TK_ANY || ty == tk.Type()
}

func (p *pipeline) matchBeyond(ty TokenType) bool {

	tk := p.stream.PeekBeyond()
	if tk == nil {
		return false
	}

	return ty == TK_ANY || ty == tk.Type()
}

func (p *pipeline) peek() Token {
	return p._peek()
}

func (p *pipeline) any() Token {

	if p.accept(TK_ANY) {
		return p.prev
	}

	return nil
}

func (p *pipeline) accept(ty TokenType) bool {

	tk := p._peek()
	if tk == nil {
		return false
	}

	if ty == TK_ANY || ty == tk.Type() {
		p.prev = p.stream.Next()
		return true
	}

	return false
}

func (p *pipeline) expect(exp TokenType) (Token, error) {

	if p.accept(exp) {
		return p.prev, nil
	}

	if p.hasMore() {
		return nil, p._unexpected(p.stream.Peek(), exp.String())
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
		return nil, p._unexpected(p.stream.Peek(), JoinTypes(exp...))
	}

	return nil, p._outOfTokens(p.prev, JoinTypes(exp...))
}

func (p *pipeline) _ignoreRedundant() {

	for next := p.stream.Peek(); next != nil; next = p.stream.Peek() {

		ty := next.Type()

		if ty == TK_NEWLINE {
			ty = TK_TERMINATOR
			line, col := next.Begin()
			next = NewToken(ty, next.Value(), line, col)
		}

		switch {
		case ty == TK_COMMENT:
			p.stream.Next()

		case ty == TK_WHITESPACE:
			p.stream.Next()

		case ty != TK_TERMINATOR:
			return

			// next must be a TERMINATOR
		case p.prev == nil: // Ignore TERMINATORs at start of script
			p.stream.Next()

		case p.prev.Type() == TK_DELIMITER: // Allow "NEWLINE" after delimiter
			p.stream.Next()

		case p.prev.Type() == TK_BLOCK_OPEN: // Allow "NEWLINE" after block start
			p.stream.Next()

		case p.prev.Type() == TK_PAREN_OPEN: // Allow "NEWLINE" after paren start
			p.stream.Next()

		case p.prev.Type() == TK_TERMINATOR: // Ignore successive TERMINATORs
			p.stream.Next()

		default: // TERMINATOR
			return
		}
	}
}

func (p *pipeline) _outOfTokens(prev Token, exp string) error {
	s := fmt.Sprintf("Expected %s; got UNDEFINED", exp)
	return err.NewAfterSnippet(s, prev)
}

func (p *pipeline) _unexpected(next Token, exp string) error {
	s := fmt.Sprintf("Expected %s; got %s", exp, next.Type().String())
	return err.NewBySnippet(s, next)
}
