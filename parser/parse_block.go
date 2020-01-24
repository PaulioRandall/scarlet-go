package parser

import (
	"strings"

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
	return ex.TabString(0)
}

// TabString satisfies the Expr interface.
func (ex blockStat) TabString(tabs int) (s string) {

	pre := strings.Repeat("\t", tabs)
	s += pre + "Block (" + ex.opener.String() + ")\n"

	for _, stat := range ex.block {
		s += pre + stat.TabString(1) + "\n"
	}

	s += pre + "(" + ex.closer.String() + ")"
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
		switch tk := p.peek(); tk.Kind {
		case token.END:
			if opener.Kind == token.SOF {
				panic(tk.String() + ": Expected EOF, found a block closing token instead")
			}
			goto BLOCK_PARSED

		case token.EOF:
			if opener.Kind != token.SOF {
				panic(tk.String() + ": Expected a block closing token, found EOF instead")
			}
			goto BLOCK_PARSED

		default:
			s := p.parseStat()
			b.block = append(b.block, s)
		}
	}

BLOCK_PARSED:
	b.closer = p.take()
	return b
}
