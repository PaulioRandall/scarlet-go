package recursive

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isBlock(p *pipe) bool {
	return p.match(token.BLOCK_OPEN)
}

func parseBlock(p *pipe) st.Block {
	// pattern := BLOCK_OPEN {statement} BLOCK_CLOSE

	return st.Block{
		Open:  p.expect(`parseBlock`, token.BLOCK_OPEN),
		Stats: parseStatements(p),
		Close: p.expect(`parseBlock`, token.BLOCK_CLOSE),
	}
}

func parseStatBlock(p *pipe) st.Block {
	return st.Block{
		Open:  p.peek(),
		Stats: []st.Statement{parseStatement(p)},
		Close: p.past(),
	}
}
