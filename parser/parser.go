package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Parser is parser for a stream of tokens.
type Parser struct {
	in chan token.Token
	i  int
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
		Kind: token.SOF,
	})
}

// take returns the next token in the input channel or panics if it is closed.
func (p *Parser) take() token.Token {

	tk, ok := <-p.in

	if !ok {
		panic(tk.String() + ": Token input channel closed prematurely")
	}

	return tk
}

// ensure will panic if the specified token is not of the specified kind.
func (p *Parser) ensure(tk token.Token, ks ...token.Kind) {

	tkk := tk.Kind

	for _, k := range ks {
		if tkk == k {
			return
		}
	}

	msg := "Expected any kind from" + "\n"
	for _, k := range ks {
		msg += "\t" + string(k) + ",\n"
	}

	msg += "but was '" + string(tk.Kind) + "'"
	panic(msg)
}

// takeEnsure returns the next token in the input channel but will panic if
// the if the channel is closed or the specified token is not of the specified
// kind.
func (p *Parser) takeEnsure(ks ...token.Kind) token.Token {
	tk := p.take()
	p.ensure(tk, ks...)
	return tk
}
