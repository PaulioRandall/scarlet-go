package recursive

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isAssignment(p *pipe) bool {
	// match := FIX | ID DELIM | ID ASSIGN

	if p.match(token.FIX) ||
		p.matchSequence(token.ID, token.DELIM) ||
		p.matchSequence(token.ID, token.ASSIGN) {

		return true
	}

	return false
}

func parseAssignment(p *pipe) st.Assignment {
	// pattern := [FIX] ID { DELIM ID } ASSIGN expression {expression}

	a := st.Assignment{
		Fixed: p.accept(token.FIX),
	}

	a.IDs = parseAssignmentIds(p)
	a.Assign = p.expect(`parseAssignment`, token.ASSIGN)

	if isFuncDef(p) {
		a.Exprs = []st.Expression{parseFuncDef(p)}
	} else {
		a.Exprs = parseExpressions(p)
	}

	if a.Exprs == nil {
		panic(unexpected("parseAssignment", p.peek(), token.ANY))
	}

	return a
}

func parseAssignmentIds(p *pipe) []token.Token {
	// pattern := ID { DELIM ID }

	var ids []token.Token

	for !p.itr.Empty() {

		id := p.expect(`parseAssignmentIds`, token.ID)
		ids = append(ids, id)

		if !p.accept(token.DELIM) {
			break
		}
	}

	return ids
}
