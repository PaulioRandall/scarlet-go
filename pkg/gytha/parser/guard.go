package parser

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func isGuard(p *pipe) bool {
	return p.match(TK_GUARD_OPEN)
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
		Open:      p.expect(`parseGuard`, TK_GUARD_OPEN),
		Condition: parseExpression(p),
		Close:     p.expect(`parseGuard`, TK_GUARD_CLOSE),
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
		return v.Token().Type() == TK_BOOL

	case Operation:
		return isBoolOperator(v.Operator.Type())
	}

	return false
}

func isBoolOperator(ty TokenType) bool {
	switch ty {
	case TK_LESS_THAN,
		TK_LESS_THAN_OR_EQUAL,
		TK_MORE_THAN,
		TK_MORE_THAN_OR_EQUAL,
		TK_EQUAL,
		TK_NOT_EQUAL,
		TK_AND,
		TK_OR:

		return true
	}

	return false
}
