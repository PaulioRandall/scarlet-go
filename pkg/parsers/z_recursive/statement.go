package z_recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func parseStatements(p *pipe) []Statement {
	// pattern := {statement}

	var stats []Statement

	for !p.itr.Empty() && !p.match(BLOCK_CLOSE) {
		stat := parseStatement(p)
		stats = append(stats, stat)
	}

	return stats
}

func parseStatement(p *pipe) Statement {
	// pattern := assignment | guard | match | loop | expression TERMINATOR

	switch {
	case isAssignment(p):
		return parseAssignment(p)

		//	case isGuard(p):
		//		return parseGuard(p)

		//	case isMatch(p):
		//	return parseMatch(p)

		//	case isLoop(p):
		//	return parseLoop(p)
	}

	exp := parseExpression(p)

	if exp != nil {
		p.expect(`parseStatement`, TERMINATOR)
		return exp
	}

	panic(unexpected("parseStatement", p.peek(), ANY.String()))
}
