package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// blockStat represents a block of statements.
type blockStat struct {
	opener token.Token
	closer token.Token
	block  []Stat
}

// Token satisfies the Expr interface.
func (ex blockStat) Token() token.Token {
	return ex.opener
}

// String satisfies the Expr interface.
func (ex blockStat) String() (s string) {

	s += "Block (" + ex.opener.String() + ")\n"

	for _, stat := range ex.block {
		s += "\t" + stat.String() + "\n"
	}

	s += "(" + ex.closer.String() + ")"
	return
}

// Eval satisfies the Expr interface.
func (ex blockStat) Eval(ctx Context) (_ Value) {

	for _, stat := range ex.block {
		stat.Eval(ctx)
	}

	return
}

// parseStats parses a block of statements until the end of file or a block
// closing token is encountered. If the 'opener' is empty then an EOF is
// expected else a block closing token is expected; a panic will ensue
// otherwise. If 'opener' is NOT empty it is assumed that it has been validated
// prior to being passed.
func (p *Parser) parseStats(opener token.Token) Stat {

	b := blockStat{
		opener: opener,
		block:  []Stat{},
	}

	for {
		switch tk := p.take(); tk.Kind {
		case token.END:
			if opener.Kind == token.SOF {
				panic("Expected EOF, found a block closing token instead")
			}

			b.closer = tk
			goto BLOCK_PARSED

		case token.EOF:
			if opener.Kind != token.SOF {
				panic("Expected a block closing token, found EOF instead")
			}

			b.closer = tk
			goto BLOCK_PARSED

		default:
			s := p.parseStat(tk)
			b.block = append(b.block, s)
		}
	}

BLOCK_PARSED:
	return b
}
