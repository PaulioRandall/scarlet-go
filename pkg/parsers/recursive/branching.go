package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isMatch(p *pipe) bool {
	return p.match(token.MATCH)
}

// Expects the following token pattern:
// pattern := MATCH_OPEN guard {guard} BLOCK_END
func parseMatch(p *pipe) st.Match {

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

// Expects the following token pattern:
// pattern := {guard}
func parseGuards(p *pipe) []st.Guard {

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

// Expects the following token pattern:
// pattern := GUARD_OPEN expression GUARD_CLOSE (statement | block)
func parseGuard(p *pipe) st.Guard {

	g := st.Guard{
		Open: p.expect(`parseGuard`, token.GUARD_OPEN),
		Cond: parseExpression(p),
	}

	if g.Cond == nil {
		panic(err("parseGuard", p.peek(), `Expected expression`))
	}

	if !isBoolOperation(g.Cond) {
		panic(err("parseGuard", g.Cond.Token(),
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

	if _, ok := ex.(st.Identifier); ok {
		return true
	}

	if v, ok := ex.(st.Value); ok {
		return v.Source.Type == token.BOOL
	}

	if v, ok := ex.(st.Operation); ok {
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

// Expects the following token pattern:
// pattern := BLOCK_OPEN {statement} BLOCK_CLOSE
func parseGuardBlock(p *pipe) st.Block {
	return st.Block{
		Open:  p.expect(`parseGuardBlock`, token.BLOCK_OPEN),
		Stats: parseStatements(p),
		Close: p.expect(`parseGuardBlock`, token.BLOCK_CLOSE),
	}
}

// Expects the following token pattern:
// pattern := statement
func parseGuardStatement(p *pipe) st.Block {
	return st.Block{
		Open:  p.peek(),
		Stats: []st.Statement{parseStatement(p)},
		Close: p.past(),
	}
}
