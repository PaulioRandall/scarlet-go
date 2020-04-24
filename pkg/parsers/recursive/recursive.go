package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// ParseAll parses all tokens in tks into Statements.
func ParseAll(tks []token.Token) []st.Statement {
	p := &pipe{token.NewIterator(tks)}
	return parseStatements(p)
}

// Expects the following token pattern:
// pattern := {statement}
func parseStatements(p *pipe) []st.Statement {

	var sts []st.Statement

	for !p.itr.Empty() && !p.accept(token.EOF) && !p.inspect(token.BLOCK_CLOSE) {
		st := parseStatement(p)
		sts = append(sts, st)
	}

	return sts
}

// Expects the following token pattern:
// pattern := assignment | guard | match | expression
func parseStatement(p *pipe) st.Statement {

	switch {
	case isAssignment(p):
		return parseAssignment(p)

	case isGuard(p):
		return parseGuard(p)

	case isMatch(p):
		return parseMatch(p)
	}

	if ex := parseExpression(p); ex != nil {
		p.expect(`parseStatement`, token.TERMINATOR)
		return st.Assignment{
			Exprs: []st.Expression{ex},
		}
	}

	panic(unexpected("parseStatement", p.snoop(), token.ANY))
}
