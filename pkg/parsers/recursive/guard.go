package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isGuard(p *pipe) bool {
	return p.match(token.GUARD_OPEN)
}

func parseGuards(p *pipe) []st.Guard {
	// pattern := {guard}

	var gs []st.Guard

	for isGuard(p) {
		g := parseGuard(p)
		gs = append(gs, g)
	}

	return gs
}

func parseGuard(p *pipe) st.Guard {
	// pattern := GUARD_OPEN expression GUARD_CLOSE (statement | block)

	g := st.Guard{
		Open:      p.expect(`parseGuard`, token.GUARD_OPEN),
		Condition: parseExpression(p),
	}

	if g.Condition == nil {
		panic(err("parseGuard", p.past(), 1, `Expected expression`))
	}

	if !isBoolOperation(g.Condition) {
		panic(err("parseGuard", g.Condition.Token(), 0,
			`Expected operation with bool result`,
		))
	}

	g.Close = p.expect(`parseGuard`, token.GUARD_CLOSE)

	if isBlock(p) {
		g.Block = parseBlock(p)
	} else {
		g.Block = parseStatBlock(p)
	}

	return g
}

func isBoolOperation(ex st.Expression) bool {

	switch v := ex.(type) {
	case st.Identifier:
		return true

	case st.Value:
		return v.Token().Type == token.BOOL

	case st.Operation:
		return isBoolOperator(v.Operator.Type)
	}

	return false
}

func isBoolOperator(typ token.TokenType) bool {
	switch typ {
	case token.LESS_THAN,
		token.LESS_THAN_OR_EQUAL,
		token.MORE_THAN,
		token.MORE_THAN_OR_EQUAL,
		token.EQUAL,
		token.NOT_EQUAL,
		token.AND,
		token.OR:

		return true
	}

	return false
}
