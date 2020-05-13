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

func parseGuards(p *pipe) []st.Guard {
	// pattern := {guard}

	var gs []st.Guard

	for isGuard(p) {
		g := parseGuard(p)
		gs = append(gs, g)
	}

	return gs
}

func isGuard(p *pipe) bool {
	return p.match(token.GUARD_OPEN)
}

func parseGuard(p *pipe) st.Guard {
	// pattern := GUARD_OPEN expression GUARD_CLOSE (statement | block)

	g := st.Guard{
		Open:      p.expect(`parseGuard`, token.GUARD_OPEN),
		Condition: parseExpression(p),
	}

	if g.Condition == nil {
		panic(err("parseGuard", p.peek(), `Expected expression`))
	}

	if !isBoolOperation(g.Condition) {
		panic(err("parseGuard", g.Condition.Token(),
			`Expected operation with bool result`,
		))
	}

	g.Close = p.expect(`parseGuard`, token.GUARD_CLOSE)

	if isGuardBlock(p) {
		g.Block = parseGuardBlock(p)
	} else {
		g.Block = parseGuardStatement(p)
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

func isGuardBlock(p *pipe) bool {
	return p.match(token.BLOCK_OPEN)
}

func parseGuardBlock(p *pipe) st.Block {
	// pattern := BLOCK_OPEN {statement} BLOCK_CLOSE

	return st.Block{
		Open:  p.expect(`parseGuardBlock`, token.BLOCK_OPEN),
		Stats: parseStatements(p),
		Close: p.expect(`parseGuardBlock`, token.BLOCK_CLOSE),
	}
}

func parseGuardStatement(p *pipe) st.Block {
	return st.Block{
		Open:  p.peek(),
		Stats: []st.Statement{parseStatement(p)},
		Close: p.past(),
	}
}
