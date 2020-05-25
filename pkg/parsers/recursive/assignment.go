package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isAssignment(p *pipe) bool {
	// match := FIX
	// match := ID (DELIM | ASSIGN | GUARD_OPEN)

	return p.match(FIX) ||
		p.matchSequence(IDENTIFIER, DELIMITER) ||
		p.matchSequence(IDENTIFIER, ASSIGN) ||
		p.matchSequence(IDENTIFIER, GUARD_OPEN)
}

func parseAssignment(p *pipe) Assignment {
	// pattern := [FIX] assign_target {assign_target} ASSIGN expression {expression}

	a := Assignment{
		Fixed: p.accept(FIX),
	}

	a.Targets = parseAssignTargets(p)
	a.Assign = p.expect(`parseAssignment`, ASSIGN)

	if isFuncDef(p) {
		a.Exprs = []Expression{parseFuncDef(p)}
	} else {
		a.Exprs = parseExpressions(p)
	}

	if a.Exprs == nil {
		err.Panic(
			errMsg("parseAssignment", `expression`, p.peek()),
			err.At(p.peek()),
		)
	}

	return a
}

func parseAssignTargets(p *pipe) []AssignTarget {
	// pattern := assignTarget { DELIM assignTarget }

	var ats []AssignTarget

	for !p.itr.Empty() {

		at := parseAssignTarget(p)
		ats = append(ats, at)

		if !p.accept(DELIMITER) {
			break
		}
	}

	return ats
}

func parseAssignTarget(p *pipe) AssignTarget {
	// pattern := ID [GUARD_OPEN (NUMBER | ID) GUARD_CLOSE]

	at := AssignTarget{
		ID: p.expect(`parseAssignTarget`, IDENTIFIER),
	}

	if p.accept(GUARD_OPEN) {

		switch {
		case p.matchAny(LIST_START, LIST_END):
			at.Index = ListItemRef{p.next()}

		case p.match(NUMBER):
			at.Index = parseExpression(p)

		case p.match(IDENTIFIER):
			at.Index = parseExpression(p)

		default:
			err.Panic(
				errMsg("parseAssignTarget", `NUMBER or ID`, p.peek()),
				err.At(p.peek()),
			)
		}

		p.expect(`parseAssignTarget`, GUARD_CLOSE)
	}

	return at
}
