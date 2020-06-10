package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
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
	// pattern := assignment | guard | when | loop | expression TERMINATOR

	switch {
	case isAssignment(p):
		return parseAssignment(p)

	case isGuard(p):
		return parseGuard(p)

	case isWhen(p):
		return parseWhen(p)

	case isLoop(p):
		return parseLoop(p)
	}

	exp := parseExpression(p)

	if exp != nil {
		p.expect(`parseStatement`, TERMINATOR)
		return exp
	}

	err.Panic(
		errMsg("parseStatement", `statement or expression`, p.peek()),
		err.At(p.peek()),
	)

	return nil
}
