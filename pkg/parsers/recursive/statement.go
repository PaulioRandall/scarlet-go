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
	// pattern := assignment | guard | match | loop | expression TERMINATOR

	switch {
	case isIncOrDec(p):
		return parseIncOrDec(p)

	case isAssignment(p):
		return parseAssignment(p)

	case isGuard(p):
		return parseGuard(p)

	case isMatch(p):
		return parseMatch(p)

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

func isIncOrDec(p *pipe) bool {
	return p.matchSequence(IDENTIFIER, INCREMENT) ||
		p.matchSequence(IDENTIFIER, DECREMENT)
}

func parseIncOrDec(p *pipe) Statement {
	// pattern := ID (INCREMENT | DECREMENT)

	inc := Increment{
		ID: p.expect(`parseIncOrDec`, IDENTIFIER),
	}

	if !p.matchAny(INCREMENT, DECREMENT) {
		err.Panic("SANITY CHECK! Expected INCREMENT or DECREMENT", err.At(p.peek()))
	}

	inc.Direction = p.next()

	p.expect(`parseIncOrDec`, TERMINATOR)
	return inc
}
