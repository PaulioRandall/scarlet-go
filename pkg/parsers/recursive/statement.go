package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// Expects the following token pattern:
// pattern := {statement}
func parseStatements(p *pipe) []st.Statement {

	var stats []st.Statement

	for !p.itr.Empty() && !p.accept(token.EOF) && !p.match(token.BLOCK_CLOSE) {
		stat := parseStatement(p)
		stats = append(stats, stat)
	}

	return stats
}

// Expects the following token pattern:
// pattern := assignment | guard | match | (expression TERMINATOR)
func parseStatement(p *pipe) st.Statement {

	switch {
	case isAssignment(p):
		return parseAssignment(p)

	case isGuard(p):
		return parseGuard(p)

	case isMatch(p):
		return parseMatch(p)
	}

	exp := parseExpression(p)

	if exp != nil {
		p.expect(`parseStatement`, token.TERMINATOR)
		return exp
	}

	panic(unexpected("parseStatement", p.peek(), token.ANY))
}
