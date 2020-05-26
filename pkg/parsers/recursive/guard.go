package recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isGuard(p *pipe) bool {
	return p.match(GUARD_OPEN)
}

func parseGuards(p *pipe) []Guard {
	// pattern := {guard}

	var gs []Guard

	for isGuard(p) {
		g := parseGuard(p)
		gs = append(gs, g)
	}

	return gs
}

func parseGuard(p *pipe) Guard {
	// pattern := GUARD_OPEN expression GUARD_CLOSE (statement | block)

	g := Guard{
		Open:      p.expect(`parseGuard`, GUARD_OPEN),
		Condition: parseExpression(p),
		Close:     p.expect(`parseGuard`, GUARD_CLOSE),
	}

	if isBlock(p) {
		g.Block = parseBlock(p)
	} else {
		g.Block = parseStatBlock(p)
	}

	return g
}

func isBoolOperation(ex Expression) bool {

	switch v := ex.(type) {
	case Identifier:
		return true

	case Value:
		return v.Token().Morpheme() == BOOL

	case Operation:
		return isBoolOperator(v.Operator.Morpheme())
	}

	return false
}

func isBoolOperator(m Morpheme) bool {
	switch m {
	case LESS_THAN,
		LESS_THAN_OR_EQUAL,
		MORE_THAN,
		MORE_THAN_OR_EQUAL,
		EQUAL,
		NOT_EQUAL,
		AND,
		OR:

		return true
	}

	return false
}
