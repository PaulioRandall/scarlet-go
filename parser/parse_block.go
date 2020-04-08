package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// parseStats parses a block of statements until the end of file or a block
// closing token is encountered. If the 'opener' is empty then an EOF is
// expected else a block closing token is expected; a panic will ensue
// otherwise. If 'opener' is NOT empty it is assumed that it has been validated
// prior to being passed.
func (p *Parser) parseStats(opener lexeme.Token) Stat {

	b := blockStat{
		opener: opener,
		block:  []Stat{},
	}

	for {
		switch tk := p.peek(); tk.Lexeme {
		case lexeme.LEXEME_END:
			if opener.Lexeme == lexeme.LEXEME_SOF {
				panic(bard.NewHorror(tk, nil,
					"Expected EOF, found a block closing token instead",
				))
			}
			goto BLOCK_PARSED

		case lexeme.LEXEME_EOF:
			if opener.Lexeme != lexeme.LEXEME_SOF {
				panic(bard.NewHorror(tk, nil,
					"Expected a block closing token, found EOF instead",
				))
			}
			goto BLOCK_PARSED

		default:
			s := p.parseStat(false)
			b.block = append(b.block, s)
		}
	}

BLOCK_PARSED:
	b.closer = p.take()
	return b
}

// blockStat represents a block of statements.
type blockStat struct {
	opener lexeme.Token
	closer lexeme.Token
	block  []Stat
}

// String satisfies the Expr interface.
func (ex blockStat) String() (s string) {

	s = "Block (" + ex.opener.String() + ")"

	for _, stat := range ex.block {
		s += "\n" + stat.String()
	}

	s = strings.ReplaceAll(s, "\n", "\n\t")

	if ex.closer != (lexeme.Token{}) {
		s += "\n(" + ex.closer.String() + ")"
	}

	return
}

// Eval satisfies the Expr interface.
func (ex blockStat) Eval(ctx Context) (_ Value) {

	for _, stat := range ex.block {
		stat.Eval(ctx)
	}

	return
}
