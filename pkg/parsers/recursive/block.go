package recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isBlock(p *pipe) bool {
	return p.match(TK_BLOCK_OPEN)
}

func parseBlock(p *pipe) Block {
	// pattern := BLOCK_OPEN {statement} BLOCK_CLOSE

	return Block{
		Open:  p.expect(`parseBlock`, TK_BLOCK_OPEN),
		Stats: parseStatements(p),
		Close: p.expect(`parseBlock`, TK_BLOCK_CLOSE),
	}
}

func parseStatBlock(p *pipe) Block {
	return Block{
		Open:  p.peek(),
		Stats: []Statement{parseStatement(p)},
		Close: p.past(),
	}
}
