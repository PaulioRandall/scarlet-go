package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func parseStatements(p *pipe) []st.Statement {
	// pattern := {statement}

	var stats []st.Statement

	for !p.itr.Empty() && !p.accept(token.EOF) && !p.match(token.BLOCK_CLOSE) {
		stat := parseStatement(p)
		stats = append(stats, stat)
	}

	return stats
}

func parseStatement(p *pipe) st.Statement {
	// pattern := assignment | guard | match | expression TERMINATOR

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
