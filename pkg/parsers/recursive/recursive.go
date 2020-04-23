package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// ParseAll parses all tokens in tks into Statements.
func ParseAll(tks []token.Token) []st.Statement {
	p := &pipe{token.NewIterator(tks)}
	return statements(p)
}

// statements parses all statements within the parsers iterator.
//
// Preconditions: None
func statements(p *pipe) (ss []st.Statement) {

	for !p.itr.Empty() && !p.accept(token.EOF) && !p.inspect(token.BLOCK_CLOSE) {
		s := statement(p)
		ss = append(ss, s)
	}

	return
}

// statement parses a single statement from the parsers iterator.
//
// Preconditions:
// - Next token is not empty
func statement(p *pipe) (s st.Statement) {

	if isAssignment(p) {
		return parseAssignment(p)
	}

	if isGuard(p) {
		return parseGuard(p)
	}

	if isMatch(p) {
		return parseMatch(p)
	}

	if ex := parseExpression(p); ex != nil {
		p.expect(`statement`, token.TERMINATOR)
		return st.Assignment{
			Exprs: []st.Expression{ex},
		}
	}

	panic(unexpected("statement", p.snoop(), token.ANY))
}
