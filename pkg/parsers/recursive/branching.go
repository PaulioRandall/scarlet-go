package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isMatch(p *pipe) bool {
	return p.inspect(token.MATCH_OPEN)
}

// Expects the following token pattern:
// pattern := MATCH_OPEN guard {guard} BLOCK_END
// guard := GUARD_OPEN expression GUARD_CLOSE (statement | block)
func parseMatch(p *pipe) st.Match {

	m := st.Match{
		Open:  p.expect(`parseMatch`, token.MATCH_OPEN),
		Cases: parseGuards(p),
	}

	if m.Cases == nil {
		panic(unexpected("parseMatch", p.snoop(), token.GUARD_OPEN))
	}

	m.Close = p.expect(`parseMatch`, token.BLOCK_CLOSE)
	return m
}

// Expects the following token pattern:
// pattern := {guard}
// guard := GUARD_OPEN expression GUARD_CLOSE (statement | block)
func parseGuards(p *pipe) []st.Guard {

	var gs []st.Guard

	for isGuard(p) {
		g := parseGuard(p)
		gs = append(gs, g)
	}

	return gs
}

func isGuard(p *pipe) bool {
	return p.inspect(token.GUARD_OPEN)
}

// Expects the following token pattern:
// pattern := GUARD_OPEN expression GUARD_CLOSE (statement | block)
func parseGuard(p *pipe) st.Guard {

	g := st.Guard{
		Open: p.expect(`parseGuard`, token.GUARD_OPEN),
		Cond: parseExpression(p),
	}

	if g.Cond == nil {
		panic(err("parseGuard", p.snoop(), `Expected expression`))
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
	return p.inspect(token.BLOCK_OPEN)
}

// Expects the following token pattern:
// pattern := BLOCK_OPEN {statement} BLOCK_CLOSE
func parseGuardBlock(p *pipe) st.Block {
	return st.Block{
		Open:  p.expect(`parseGuardBlock`, token.BLOCK_OPEN),
		Stats: statements(p),
		Close: p.expect(`parseGuardBlock`, token.BLOCK_CLOSE),
	}
}

// Expects the following token pattern:
// pattern := statement
func parseGuardStatement(p *pipe) st.Block {
	return st.Block{
		Open:  p.snoop(),
		Stats: []st.Statement{statement(p)},
		Close: p.prior(),
	}
}
