package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Parser is parser for a stream of tokens.
type Parser struct {
	in chan token.Token
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
	return p.parseStats(token.Token{})
}

// checkTokenKind panics if the specified token is not of the specified kind.
func (p *Parser) checkToken(tk token.Token, k token.Kind) {
	if tk.Kind != k {
		panic("Expected token of kind '" + k + "' but was '" + tk.Kind + "'")
	}
}

// parseStat parses the next statement.
func (p *Parser) parseStat(lead token.Token) (s Stat) {

	switch lead.Kind {
	case token.ID:
		s = p.parseAssign(lead)
	default:
		panic("Token of kind '" + lead.Kind + "' is not known to the parser " +
			"or parsing has not been implemented for it yet")
	}

	return
}
