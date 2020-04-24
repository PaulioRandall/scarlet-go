package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isAssignment(p *pipe) bool {

	if p.match(token.FIX) ||
		p.matchSequence(token.ID, token.DELIM) ||
		p.matchSequence(token.ID, token.ASSIGN) {

		return true
	}

	return false
}

// Expects the following token pattern:
// pattern := [FIX] ID { DELIM ID } ASSIGN expression {expression}
func parseAssignment(p *pipe) st.Assignment {

	fixed := p.accept(token.FIX)
	a := st.Assignment{}

	a.IDs = parseAssignmentIds(p, fixed)
	a.Assign = p.expect(`parseAssignment`, token.ASSIGN)
	a.Exprs = parseExpressions(p)

	if a.Exprs == nil {
		panic(unexpected("parseAssignment", p.peek(), token.ANY))
	}

	return a
}

// Expects the following token pattern:
// pattern := ID { DELIM ID }
func parseAssignmentIds(p *pipe, fixed bool) []st.Identifier {

	var ids []st.Identifier

	for {
		idTk := p.expect(`parseAssignmentIds`, token.ID)
		id := st.Identifier{fixed, idTk}
		ids = append(ids, id)

		if !p.accept(token.DELIM) {
			break
		}
	}

	return ids
}
