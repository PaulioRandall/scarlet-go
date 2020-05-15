package recursive

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isAssignment(p *pipe) bool {
	// match := FIX
	// match := ID (DELIM | ASSIGN | GUARD_OPEN)

	return p.match(token.FIX) ||
		p.matchSequence(token.ID, token.DELIM) ||
		p.matchSequence(token.ID, token.ASSIGN) ||
		p.matchSequence(token.ID, token.GUARD_OPEN)
}

func parseAssignment(p *pipe) st.Assignment {
	// pattern := [FIX] assign_target {assign_target} ASSIGN expression {expression}

	a := st.Assignment{
		Fixed: p.accept(token.FIX),
	}

	a.Targets = parseAssignTargets(p)
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

func parseAssignTargets(p *pipe) []st.AssignTarget {
	// pattern := assignTarget { DELIM assignTarget }

	var ats []st.AssignTarget

	for !p.itr.Empty() {

		at := parseAssignTarget(p)
		ats = append(ats, at)

		if !p.accept(token.DELIM) {
			break
		}
	}

	return ats
}

func parseAssignTarget(p *pipe) st.AssignTarget {
	// pattern := ID [GUARD_OPEN (NUMBER | ID) GUARD_CLOSE]

	at := st.AssignTarget{
		ID: p.expect(`parseAssignTarget`, token.ID),
	}

	if p.accept(token.GUARD_OPEN) {

		switch {
		case p.matchAny(token.PREPEND, token.APPEND):
			at.Index = st.ListItemRef(p.next())
		case p.match(token.NUMBER):
			at.Index = parseExpression(p)
		case p.match(token.ID):
			at.Index = parseExpression(p)
		default:
			panic(unexpected("parseAssignTarget", p.peek(), `token.NUMBER | token.ID`))
		}

		p.expect(`parseAssignTarget`, token.GUARD_CLOSE)
	}

	return at
}
