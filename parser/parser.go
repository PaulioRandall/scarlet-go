package parser

import (
	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// Parser is parser for a stream of tokens.
type Parser struct {
	in  chan token.Token
	buf *token.Token // buffer, allows for look ahead by 1
}

// New creates a new Parser.
func New(in chan token.Token) *Parser {
	return &Parser{
		in: in,
	}
}

// Parse parses tokens obtained via the input channel into statements until the
// channel is closed. A master statement is returned that represents the block
// of statements.
func (p *Parser) Parse() (_ Stat) {
	return p.parseStats(token.Token{
		Kind: token.KIND_SOF,
	})
}

// peek returns the token on the buffer. If the buffer is empty the the next
// token is read from the input channel into the buffer first. Panics if the
// channel is closed when attempting a read.
func (p *Parser) peek() token.Token {

	if p.buf == nil {
		tk, ok := <-p.in

		if !ok {
			panic(bard.NewHorror(tk, nil,
				"Token input channel closed prematurely",
			))
		}

		p.buf = &tk
	}

	return *p.buf
}

// take removes and returns the token in the buffer. If the buffer is empty then
// a peek operation is performed to fill it first.
func (p *Parser) take() (tk token.Token) {

	if p.buf == nil {
		p.peek()
	}

	tk, p.buf = *p.buf, nil
	return
}

// ensure will panic if the specified token is not one of the specified kinds.
func (p *Parser) ensure(tk token.Token, ks ...token.Kind) {

	var errMsg string
	tkk := tk.Kind

	for _, k := range ks {
		if tkk == k {
			return
		}
	}

	errMsg = "Expected "
	if len(ks) == 1 {
		errMsg += string(ks[0])
	} else {
		errMsg = "either"
		for _, k := range ks {
			errMsg += " " + string(k)
		}
	}

	errMsg += " but was " + string(tk.Kind)
	panic(bard.NewHorror(tk, nil, errMsg))
}

// peekEnsure returns the next token in the input channel but will panic if
// the if the channel is closed or the specified token is not one of the
// specified kinds.
func (p *Parser) peekEnsure(ks ...token.Kind) token.Token {
	tk := p.peek()
	p.ensure(tk, ks...)
	return tk
}

// takeEnsure returns the next token in the input channel but will panic if
// the if the channel is closed or the specified token is not one of the
// specified kinds.
func (p *Parser) takeEnsure(ks ...token.Kind) token.Token {
	tk := p.take()
	p.ensure(tk, ks...)
	return tk
}
