package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Parser is parser for a stream of tokens.
type Parser struct {
	in  chan token.Token
	out chan Expr
}

// New creates a new Parser.
func New(in chan token.Token, out chan Expr) *Parser {
	return &Parser{
		in:  in,
		out: out,
	}
}

// checkTokenKind panics if the specified token is not of the specified kind.
func (p *Parser) checkToken(tk token.Token, k token.Kind) {
	if tk.Kind != k {
		panic("Expected token of kind '" + k + "' but was '" + tk.Kind + "'")
	}
}

// NEXT
func (p *Parser) Parse() (_ Expr) {
	return
}
