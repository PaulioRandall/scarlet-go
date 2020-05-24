package recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isMatch(p *pipe) bool {
	return p.match(MATCH)
}

func parseMatch(p *pipe) Match {
	// pattern := MATCH BLOCK_OPEN guard {guard} BLOCK_CLOSE

	m := Match{
		Key:   p.expect(`parseMatch`, MATCH),
		Open:  p.expect(`parseMatch`, BLOCK_OPEN),
		Cases: parseGuards(p),
	}

	if m.Cases == nil {
		panic(unexpected("parseMatch", p.peek(), GUARD_OPEN.String()))
	}

	m.Close = p.expect(`parseMatch`, BLOCK_CLOSE)
	return m
}
