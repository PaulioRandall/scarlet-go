package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

// True if one of the following token patterns match:
// - FIX, ...
// - ID, DELIM, ...
// - ID, ASSIGN, ...
func isAssignment(p *pipe) bool {

	if p.inspect(token.FIX) ||
		p.isSequence(token.ID, token.DELIM) ||
		p.isSequence(token.ID, token.ASSIGN) {

		return true
	}

	return false
}

// Assumes isAssignment returns true.
func parseAssignment(p *pipe) st.Assignment {

	fixed := p.accept(token.FIX)
	a := st.Assignment{}

	a.IDs = parseAssignmentIds(p, fixed)
	a.Assign = p.expect(`parseAssignment`, token.ASSIGN)
	a.Exprs = parseExpressions(p)

	if a.Exprs == nil {
		panic(unexpected("parseAssignment", p.snoop(), token.ANY))
	}

	p.expect(`parseAssignment`, token.TERMINATOR)
	return a
}

// Expects one of the following token patterns:
// - ID, ...
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
