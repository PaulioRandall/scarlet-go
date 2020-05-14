package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isMatch(p *pipe) bool {
	return p.match(token.MATCH)
}

func parseMatch(p *pipe) st.Match {
	// pattern := MATCH BLOCK_OPEN guard {guard} BLOCK_CLOSE

	m := st.Match{
		Key:   p.expect(`parseMatch`, token.MATCH),
		Open:  p.expect(`parseMatch`, token.BLOCK_OPEN),
		Cases: parseGuards(p),
	}

	if m.Cases == nil {
		panic(unexpected("parseMatch", p.peek(), token.GUARD_OPEN))
	}

	m.Close = p.expect(`parseMatch`, token.BLOCK_CLOSE)
	return m
}
